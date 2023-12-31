package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapModelProductTypeToDomain(modelPdType model.AlloffProductType) domain.AlloffProductType {
	if modelPdType == model.AlloffProductTypeKids {
		return domain.Kids
	} else if modelPdType == model.AlloffProductTypeMale {
		return domain.Male
	} else if modelPdType == model.AlloffProductTypeBoy {
		return domain.Boy
	} else if modelPdType == model.AlloffProductTypeGirl {
		return domain.Girl
	}
	return domain.Female
}

func MapProduct(pdDao *domain.ProductDAO) *model.Product {
	if pdDao.IsRemoved {
		return nil
	}

	deliveryDesc := MapDeliveryDescription(pdDao.ProductInfo.SalesInstruction.DeliveryDescription)

	isSoldOut := true
	for _, inv := range pdDao.ProductInfo.Inventory {
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
		ID:                   pdDao.ID.Hex(),
		IsNotSale:            pdDao.IsRemoved,
		Brand:                MapBrandDaoToBrand(pdDao.ProductInfo.Brand, false),
		AlloffCategory:       MapAlloffCatDaoToAlloffCat(pdDao.ProductInfo.AlloffCategory.First),
		Name:                 pdDao.ProductInfo.AlloffName,
		OriginalPrice:        pdDao.ProductInfo.Price.OriginalPrice,
		DiscountedPrice:      pdDao.ProductInfo.Price.CurrentPrice,
		DiscountRate:         pdDao.ProductInfo.Price.DiscountRate,
		Images:               pdDao.ProductInfo.Images,
		ThumbnailImage:       thumbnailImage,
		Inventory:            MapInventory(pdDao.ProductInfo.Inventory),
		IsSoldout:            isSoldOut,
		Description:          MapDescription(pdDao.ProductInfo.SalesInstruction.Description),
		DeliveryDescription:  deliveryDesc,
		CancelDescription:    MapCancelDescription(pdDao.ProductInfo.SalesInstruction.CancelDescription),
		Information:          information,
		ExhibitionID:         pdDao.ExhibitionID,
		ExhibitionStartTime:  pdDao.ExhibitionStartTime.String(),
		ExhibitionFinishTime: pdDao.ExhibitionFinishTime.String(),
	}
}

func MapInventory(invs []*domain.InventoryDAO) []*model.Inventory {
	res := []*model.Inventory{}
	for _, inv := range invs {
		alloffSizes := []*model.AlloffSize{}
		for _, alloffSize := range inv.AlloffSizes {
			alloffSizes = append(alloffSizes, MapAlloffSizeDaoToAlloffSize(alloffSize))
		}

		invModel := &model.Inventory{
			Quantity:    inv.Quantity,
			Size:        inv.Size,
			AlloffSizes: alloffSizes,
		}
		res = append(res, invModel)
	}
	return res
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

func MapProductSortingAndRanges(sortings []model.ProductsSortingType) (priceRanges []domain.PriceRangeType, priceSorting domain.PriceSortingType) {
	for _, sorting := range sortings {
		if sorting == model.ProductsSortingTypePriceAscending {
			priceSorting = domain.PRICE_ASCENDING
		} else if sorting == model.ProductsSortingTypePriceDescending {
			priceSorting = domain.PRICE_DESCENDING
		} else if sorting == model.ProductsSortingTypeDiscountrateAscending {
			priceSorting = domain.DISCOUNTRATE_ASCENDING
		} else if sorting == model.ProductsSortingTypeDiscountrateDescending {
			priceSorting = domain.DISCOUNTRATE_DESCENDING
		} else if sorting == model.ProductsSortingTypeInventoryAscending {
			priceSorting = domain.INVENTORY_ASCENDING
		} else if sorting == model.ProductsSortingTypeInventoryDescending {
			priceSorting = domain.INVENTORY_DESCENDING
		} else {
			if sorting == model.ProductsSortingTypeDiscount0_30 {
				priceRanges = append(priceRanges, domain.PRICE_RANGE_30)
			} else if sorting == model.ProductsSortingTypeDiscount30_50 {
				priceRanges = append(priceRanges, domain.PRICE_RANGE_50)
			} else if sorting == model.ProductsSortingTypeDiscount50_70 {
				priceRanges = append(priceRanges, domain.PRICE_RANGE_70)
			} else {
				priceRanges = append(priceRanges, domain.PRICE_RANGE_100)
			}
		}
	}

	return
}
