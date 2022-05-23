package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func MapProduct(pdDao *domain.ProductDAO) *model.Product {
	deliveryDesc := MapDeliveryDescription(pdDao.ProductInfo.SalesInstruction.DeliveryDescription)
	switch pdDao.ProductInfo.Source.IsForeignDelivery {
	case true:
		deliveryDesc.DeliveryType = model.DeliveryTypeForeignDelivery
	case false:
		deliveryDesc.DeliveryType = model.DeliveryTypeDomesticDelivery
	}

	alloffPrice := product.GetCurrentPrice(pdDao)
	alloffPriceDiscountRate := utils.CalculateDiscountRate(pdDao.OriginalPrice, alloffPrice)

	isSoldOut := true
	for _, inv := range pdDao.Inventory {
		if inv.Quantity > 0 {
			isSoldOut = false
		}
	}

	if isSoldOut && !pdDao.IsSoldout {
		pdDao.IsSoldout = true
		go ioc.Repo.Products.Upsert(pdDao)
	}

	if pdDao.IsSoldout {
		isSoldOut = true
	}

	thumbnailImage := ""
	if len(pdDao.ProductInfo.Images) > 0 {
		thumbnailImage = pdDao.ProductInfo.Images[0]
		if pdDao.ThumbnailImage != "" {
			thumbnailImage = pdDao.ThumbnailImage
		}
	}

	return &model.Product{
		ID:                  pdDao.ID.Hex(),
		Brand:               MapBrandDaoToBrand(pdDao.ProductInfo.Brand, false),
		AlloffCategory:      MapAlloffCatDaoToAlloffCat(pdDao.ProductInfo.AlloffCategory.First),
		Name:                pdDao.AlloffName,
		OriginalPrice:       pdDao.OriginalPrice,
		DiscountedPrice:     alloffPrice,
		DiscountRate:        alloffPriceDiscountRate,
		Images:              pdDao.ProductInfo.Images,
		ThumbnailImage:      thumbnailImage,
		Inventory:           MapAlloffInventory(pdDao.Inventory),
		IsSoldout:           isSoldOut,
		Description:         MapDescription(pdDao.ProductInfo.SalesInstruction.Description),
		DeliveryDescription: deliveryDesc,
		CancelDescription:   MapCancelDescription(pdDao.ProductInfo.SalesInstruction.CancelDescription),
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
