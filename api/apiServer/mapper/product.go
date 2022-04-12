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
