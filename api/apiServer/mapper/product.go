package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func MapProduct(pdDao *domain.ProductDAO) *model.Product {
	if pdDao.IsNotSale {
		return nil
	}

	deliveryDesc := MapDeliveryDescription(pdDao.ProductInfo.SalesInstruction.DeliveryDescription)
	switch pdDao.ProductInfo.Source.IsForeignDelivery {
	case true:
		deliveryDesc.DeliveryType = model.DeliveryTypeForeignDelivery
	case false:
		deliveryDesc.DeliveryType = model.DeliveryTypeDomesticDelivery
	}

	isSoldOut := true
	for _, inv := range pdDao.ProductInfo.AlloffInventory {
		if inv.Quantity > 0 {
			isSoldOut = false
		}
	}

	if isSoldOut && !pdDao.ProductInfo.IsSoldout {
		pdDao.ProductInfo.IsSoldout = true
		go ioc.Repo.Products.Upsert(pdDao)
	}

	if pdDao.ProductInfo.IsSoldout {
		isSoldOut = true
	}

	thumbnailImage := ""
	if len(pdDao.ProductInfo.Images) > 0 {
		thumbnailImage = pdDao.ProductInfo.Images[0]
		if pdDao.ProductInfo.ThumbnailImage != "" {
			thumbnailImage = pdDao.ProductInfo.ThumbnailImage
		}
	}

	var information []*model.KeyValueInfo
	for k, v := range pdDao.ProductInfo.SalesInstruction.Information {
		var newInfo model.KeyValueInfo
		newInfo.Key = k
		newInfo.Value = v
		information = append(information, &newInfo)
	}

	return &model.Product{
		ID:                  pdDao.ID.Hex(),
		IsNotSale:           pdDao.IsNotSale,
		Brand:               MapBrandDaoToBrand(pdDao.ProductInfo.Brand, false),
		AlloffCategory:      MapAlloffCatDaoToAlloffCat(pdDao.ProductInfo.AlloffCategory.First),
		Name:                pdDao.ProductInfo.AlloffName,
		OriginalPrice:       pdDao.ProductInfo.Price.OriginalPrice,
		DiscountedPrice:     pdDao.ProductInfo.Price.CurrentPrice,
		DiscountRate:        pdDao.ProductInfo.Price.DiscountRate,
		Images:              pdDao.ProductInfo.Images,
		ThumbnailImage:      thumbnailImage,
		Inventory:           MapAlloffInventory(pdDao.ProductInfo.AlloffInventory),
		IsSoldout:           isSoldOut,
		Description:         MapDescription(pdDao.ProductInfo.SalesInstruction.Description),
		DeliveryDescription: deliveryDesc,
		CancelDescription:   MapCancelDescription(pdDao.ProductInfo.SalesInstruction.CancelDescription),
		Information:         information,
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

func MapAlloffInventory(invs []*domain.AlloffInventoryDAO) []*model.AlloffInventory {
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

func MapCancelDescription(cancelDesc *domain.CancelDescriptionDAO) *model.CancelDescription {
	return &model.CancelDescription{
		RefundAvailable: cancelDesc.RefundAvailable,
		ChangeAvailable: cancelDesc.ChangeAvailable,
		ChangeFee:       cancelDesc.ChangeFee,
		RefundFee:       cancelDesc.RefundFee,
	}
}

func MapProductSortingAndRanges(sortings []model.ProductsSortingType) (priceRanges []product.PriceRangeType, priceSorting product.PriceSortingType) {
	for _, sorting := range sortings {
		if sorting == model.ProductsSortingTypePriceAscending {
			priceSorting = product.PRICE_ASCENDING
		} else if sorting == model.ProductsSortingTypePriceDescending {
			priceSorting = product.PRICE_DESCENDING
		} else if sorting == model.ProductsSortingTypeDiscountrateAscending {
			priceSorting = product.DISCOUNTRATE_ASCENDING
		} else if sorting == model.ProductsSortingTypeDiscountrateDescending {
			priceSorting = product.DISCOUNTRATE_DESCENDING
		} else {
			if sorting == model.ProductsSortingTypeDiscount0_30 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_30)
			} else if sorting == model.ProductsSortingTypeDiscount30_50 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_50)
			} else if sorting == model.ProductsSortingTypeDiscount50_70 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_70)
			} else {
				priceRanges = append(priceRanges, product.PRICE_RANGE_100)
			}
		}
	}

	return
}
