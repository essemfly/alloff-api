package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func MapProductDaoToProduct(pdDao *domain.ProductDAO) *model.Product {
	inventories := []*model.Inventory{}

	for _, inv := range pdDao.Inventory {
		if inv.Quantity > 0 {
			inventories = append(inventories, &model.Inventory{
				Quantity: inv.Quantity,
				Size:     inv.Size,
			})
		}
	}

	var information []*model.KeyValueInfo
	for k, v := range pdDao.ProductInfo.Information {
		var newInfo model.KeyValueInfo
		newInfo.Key = k
		newInfo.Value = v
		information = append(information, &newInfo)
	}

	deliveryDesc := MapDeliveryDescription(pdDao.SalesInstruction.DeliveryDescription)
	if pdDao.ProductInfo.Source.IsForeignDelivery {
		deliveryDesc.DeliveryType = model.DeliveryTypeForeignDelivery
	} else {
		deliveryDesc.DeliveryType = model.DeliveryTypeDomesticDelivery
	}

	alloffPrice := product.GetCurrentPrice(pdDao)
	alloffPriceDiscountRate := utils.CalculateDiscountRate(pdDao.OriginalPrice, alloffPrice)

	isSoldout := true
	for _, inv := range inventories {
		if inv.Quantity > 0 {
			isSoldout = false
		}
	}

	if isSoldout && !pdDao.Soldout {
		pdDao.Soldout = true
		go ioc.Repo.Products.Upsert(pdDao)
	}

	if pdDao.Soldout {
		isSoldout = true
	}

	images := pdDao.ProductInfo.Images
	if pdDao.IsImageCached {
		images = pdDao.Images
	}

	thumbnailImage := ""
	if len(images) > 0 {
		thumbnailImage = images[0]
		if pdDao.ThumbnailImage != "" {
			thumbnailImage = pdDao.ThumbnailImage
		}
	}

	return &model.Product{
		ID:                  pdDao.ID.Hex(),
		Category:            MapCategoryDaoToCategory(pdDao.ProductInfo.Category),
		Brand:               MapBrandDaoToBrand(pdDao.ProductInfo.Brand, false),
		Name:                pdDao.AlloffName,
		OriginalPrice:       pdDao.OriginalPrice,
		ProductGroupID:      pdDao.ProductGroupId,
		Soldout:             isSoldout,
		Images:              images,
		DiscountedPrice:     alloffPrice,
		DiscountRate:        alloffPriceDiscountRate,
		SpecialPrice:        &alloffPrice,
		SpecialDiscountRate: &alloffPriceDiscountRate,
		ProductURL:          pdDao.ProductInfo.ProductUrl,
		Inventory:           inventories,
		IsUpdated:           pdDao.IsUpdated,
		IsNewProduct:        pdDao.Score.IsNewlyCrawled,
		Removed:             pdDao.Removed,
		Information:         information,
		Description:         MapDescription(pdDao.SalesInstruction.Description),
		CancelDescription:   MapCancelDescription(pdDao.SalesInstruction.CancelDescription),
		DeliveryDescription: deliveryDesc,
		ThumbnailImage:      thumbnailImage,
		ProductClassifier:   MapProductClassifier(pdDao.ProductClassifier),
		AlloffInventory:     MapAlloffInventory(pdDao.AlloffInventory),
	}
}

func MapDescription(desc *domain.ProductDescriptionDAO) *model.ProductDescription {
	var information []*model.KeyValueInfo
	for k, v := range desc.Infos {
		var newInfo model.KeyValueInfo
		newInfo.Key = k
		newInfo.Value = v
		information = append(information, &newInfo)
	}
	return &model.ProductDescription{
		Images: desc.Images,
		Texts:  desc.Texts,
		Infos:  information,
	}
}

func MapDeliveryDescription(deliveryDesc *domain.DeliveryDescriptionDAO) *model.DeliveryDescription {
	deliveryType := model.DeliveryTypeDomesticDelivery
	if deliveryDesc.DeliveryType == domain.Foreign {
		deliveryType = model.DeliveryTypeForeignDelivery
	}

	return &model.DeliveryDescription{
		DeliveryType:         deliveryType,
		DeliveryFee:          deliveryDesc.DeliveryFee,
		EarliestDeliveryDays: deliveryDesc.EarliestDeliveryDays,
		LatestDeliveryDays:   deliveryDesc.LatestDeliveryDays,
		Texts:                deliveryDesc.Texts,
	}
}

func MapCancelDescription(cancelDesc *domain.CancelDescriptionDAO) *model.CancelDescription {
	return &model.CancelDescription{
		RefundAvailable: cancelDesc.RefundAvailable,
		ChangeAvailable: cancelDesc.ChangeAvailable,
		ChangeFee:       cancelDesc.ChangeFee,
		RefundFee:       cancelDesc.RefundFee,
	}
}

func MapProductSortingAndRanges(sortings []model.SortingType) (priceRanges []product.PriceRangeType, priceSorting product.PriceSortingType) {
	for _, sorting := range sortings {
		if sorting == model.SortingTypePriceAscending {
			priceSorting = product.PRICE_ASCENDING
		} else if sorting == model.SortingTypePriceDescending {
			priceSorting = product.PRICE_DESCENDING
		} else if sorting == model.SortingTypeDiscountrateAscending {
			priceSorting = product.DISCOUNTRATE_ASCENDING
		} else if sorting == model.SortingTypeDiscountrateDescending {
			priceSorting = product.DISCOUNTRATE_DESCENDING
		} else {
			if sorting == model.SortingTypeDiscount0_30 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_30)
			} else if sorting == model.SortingTypeDiscount30_50 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_50)
			} else if sorting == model.SortingTypeDiscount50_70 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_70)
			} else {
				priceRanges = append(priceRanges, product.PRICE_RANGE_100)
			}
		}
	}

	return
}

func MapAlloffInventory(invs []domain.AlloffInventoryDAO) []*model.AlloffInventory {
	res := []*model.AlloffInventory{}
	for _, inv := range invs {
		invModel := &model.AlloffInventory{
			Quantity:   inv.Quantity,
			AlloffSize: MapAlloffSizeDaoToAlloffSize(&inv.AlloffSize),
		}
		res = append(res, invModel)
	}
	return res
}

func MapAlloffClassifier(classifiers []domain.AlloffClassifier) []model.AlloffClassifier {
	res := []model.AlloffClassifier{}
	for _, classifier := range classifiers {
		switch classifier {
		case domain.Male:
			res = append(res, model.AlloffClassifierMale)
		case domain.Female:
			res = append(res, model.AlloffClassifierFemale)
		case domain.Kids:
			res = append(res, model.AlloffClassifierKids)
		case domain.Sports:
			res = append(res, model.AlloffClassifierSports)
		}
	}
	return res
}

func MapCategoryClassifier(classifier domain.CategoryClassifier) *model.CategoryClassifier {
	return &model.CategoryClassifier{
		KeyName: classifier.KeyName,
		Name:    classifier.Name,
	}
}

func MapProductClassifier(classifier *domain.ProductClassifierDAO) *model.ProductClassifier {
	if classifier == nil {
		return &model.ProductClassifier{}
	}
	return &model.ProductClassifier{
		Classifier: MapAlloffClassifier(classifier.Classifier),
		First:      MapCategoryClassifier(classifier.First),
		Second:     MapCategoryClassifier(classifier.Second),
	}
}

func MapAlloffClassifierModelToDAO(classifiers []model.AlloffClassifier) []domain.AlloffClassifier {
	res := []domain.AlloffClassifier{}
	for _, classifier := range classifiers {
		switch classifier {
		case model.AlloffClassifierMale:
			res = append(res, domain.Male)
		case model.AlloffClassifierFemale:
			res = append(res, domain.Female)
		case model.AlloffClassifierKids:
			res = append(res, domain.Kids)
		case model.AlloffClassifierSports:
			res = append(res, domain.Sports)
		}
	}
	return res
}
