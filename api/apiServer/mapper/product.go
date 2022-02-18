package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
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

	alloffPrice := pdDao.DiscountedPrice
	if alloffPrice == 0 {
		alloffPrice = int(pdDao.ProductInfo.Price.OriginalPrice)
	}

	if pdDao.ProductGroupId != "" {
		pgDao, err := ioc.Repo.ProductGroups.Get(pdDao.ProductGroupId)

		if err == nil && pgDao.IsLive() {
			alloffPrice = pdDao.SpecialPrice
		}
	}

	alloffPriceDiscountRate := utils.CalculateDiscountRate(pdDao.ProductInfo.Price.OriginalPrice, float32(alloffPrice))

	isSoldout := true
	if len(inventories) > 0 {
		isSoldout = false
	}

	return &model.Product{
		ID:                  pdDao.ID.Hex(),
		Category:            MapCategoryDaoToCategory(pdDao.ProductInfo.Category),
		Brand:               MapBrandDaoToBrand(pdDao.ProductInfo.Brand, false),
		Name:                pdDao.AlloffName,
		OriginalPrice:       int(pdDao.ProductInfo.Price.OriginalPrice),
		ProductGroupID:      pdDao.ProductGroupId,
		Soldout:             isSoldout,
		Images:              pdDao.ProductInfo.Images,
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
	}
}

func MapDescription(desc *domain.ProductDescriptionDAO) *model.ProductDescription {
	return &model.ProductDescription{
		Images: desc.Images,
		Texts:  desc.Texts,
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
