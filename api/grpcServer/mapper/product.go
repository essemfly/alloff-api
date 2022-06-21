package mapper

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

func ProductInfoMapper(pdInfo *domain.ProductMetaInfoDAO) *grpcServer.ProductMessage {
	var alloffCategoryName, alloffCategoryID string
	IsClassifiedDone, IsClassifiedTouched := false, false
	if pdInfo.AlloffCategory != nil {
		IsClassifiedDone, IsClassifiedTouched = pdInfo.AlloffCategory.Done, pdInfo.AlloffCategory.Touched
		if pdInfo.AlloffCategory.Done {
			if pdInfo.AlloffCategory.Second != nil {
				alloffCategoryName = pdInfo.AlloffCategory.Second.Name
				alloffCategoryID = pdInfo.AlloffCategory.Second.ID.Hex()
			} else if pdInfo.AlloffCategory.First != nil {
				alloffCategoryName = pdInfo.AlloffCategory.First.Name
				alloffCategoryID = pdInfo.AlloffCategory.First.ID.Hex()
			}
		}
	}

	images := pdInfo.Images
	if pdInfo.IsImageCached {
		images = pdInfo.CachedImages
	}

	isForeginDelivery := true
	if pdInfo.SalesInstruction.DeliveryDescription.DeliveryType == domain.Domestic {
		isForeginDelivery = false
	}

	return &grpcServer.ProductMessage{
		AlloffProductId:      pdInfo.ID.Hex(),
		AlloffName:           pdInfo.AlloffName,
		IsForeignDelivery:    isForeginDelivery,
		ProductId:            pdInfo.ProductID,
		ProductUrl:           pdInfo.ProductUrl,
		OriginalPrice:        int32(pdInfo.Price.OriginalPrice),
		DiscountedPrice:      int32(pdInfo.Price.CurrentPrice),
		SpecialPrice:         int32(pdInfo.Price.CurrentPrice),
		BrandKorName:         pdInfo.Brand.KorName,
		BrandKeyName:         pdInfo.Brand.KeyName,
		Inventory:            InventoryMapper(pdInfo),
		Description:          pdInfo.SalesInstruction.Description.Texts,
		EarliestDeliveryDays: int32(pdInfo.SalesInstruction.DeliveryDescription.EarliestDeliveryDays),
		LatestDeliveryDays:   int32(pdInfo.SalesInstruction.DeliveryDescription.LatestDeliveryDays),
		RefundFee:            int32(pdInfo.SalesInstruction.CancelDescription.RefundFee),
		IsRefundPossible:     pdInfo.SalesInstruction.CancelDescription.RefundAvailable,
		Images:               images,
		DescriptionImages:    pdInfo.SalesInstruction.Description.Images,
		CategoryName:         pdInfo.Category.Name,
		AlloffCategoryName:   alloffCategoryName,
		AlloffCategoryId:     alloffCategoryID,
		IsRemoved:            pdInfo.IsRemoved,
		IsSoldout:            pdInfo.IsSoldout,
		ModuleName:           pdInfo.Source.CrawlModuleName,
		IsClassifiedDone:     IsClassifiedDone,
		IsClassifiedTouched:  IsClassifiedTouched,
		ProductInfos:         pdInfo.SalesInstruction.Information,
		RawProductInfos:      pdInfo.SalesInstruction.RawInformation,
		DescriptionInfos:     pdInfo.SalesInstruction.Description.Infos,
		RawDescriptionInfos:  pdInfo.SalesInstruction.Description.RawInfos,
		ThumbnailImage:       pdInfo.ThumbnailImage,
		ProductTypes:         ProductTypeMapper(pdInfo.ProductType),
		ExhibitionHistory:    ExhibitionHistoryMapper(pdInfo.ExhibitionHistory),
	}
}

func ProductTypeMapper(pdTypes []domain.AlloffProductType) []grpcServer.ProductType {
	pdTypes = removeDuplicateType(pdTypes)
	productTypes := []grpcServer.ProductType{}
	for _, pdtype := range pdTypes {
		if pdtype == domain.Female {
			productTypes = append(productTypes, grpcServer.ProductType_FEMALE)
		} else if pdtype == domain.Male {
			productTypes = append(productTypes, grpcServer.ProductType_MALE)
		} else if pdtype == domain.Kids {
			productTypes = append(productTypes, grpcServer.ProductType_KIDS)
		} else if pdtype == domain.Boy {
			productTypes = append(productTypes, grpcServer.ProductType_BOY)
		} else if pdtype == domain.Girl {
			productTypes = append(productTypes, grpcServer.ProductType_GIRL)
		}
	}
	return productTypes
}

func ProductTypeReverseMapper(ptypes []grpcServer.ProductType) []domain.AlloffProductType {
	productTypes := []domain.AlloffProductType{}
	for _, reqPdType := range ptypes {
		if reqPdType == grpcServer.ProductType_FEMALE {
			productTypes = append(productTypes, domain.Female)
		} else if reqPdType == grpcServer.ProductType_MALE {
			productTypes = append(productTypes, domain.Male)
		} else if reqPdType == grpcServer.ProductType_KIDS {
			productTypes = append(productTypes, domain.Kids)
		} else if reqPdType == grpcServer.ProductType_BOY {
			productTypes = append(productTypes, domain.Boy)
		} else if reqPdType == grpcServer.ProductType_GIRL {
			productTypes = append(productTypes, domain.Girl)
		}
	}
	return productTypes
}

func removeDuplicateType(strSlice []domain.AlloffProductType) []domain.AlloffProductType {
	allKeys := make(map[domain.AlloffProductType]bool)
	list := []domain.AlloffProductType{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func InventoryMapper(pd *domain.ProductMetaInfoDAO) []*grpcServer.ProductInventoryMessage {
	invMessage := []*grpcServer.ProductInventoryMessage{}

	for _, inv := range pd.Inventory {
		if inv.AlloffSizes != nil {
			alloffSizes := []*grpcServer.AlloffSizeMessage{}
			for _, alloffSize := range inv.AlloffSizes {
				alloffSizes = append(alloffSizes, AlloffSizeMapper(alloffSize))
			}

			invMessage = append(invMessage, &grpcServer.ProductInventoryMessage{
				AlloffSizes: alloffSizes,
				Quantity:    int32(inv.Quantity),
				Size:        inv.Size,
			})
		} else {
			invMessage = append(invMessage, &grpcServer.ProductInventoryMessage{
				AlloffSizes: nil,
				Quantity:    int32(inv.Quantity),
				Size:        inv.Size,
			})
		}

	}
	return invMessage
}

func ProductSortingAndRangesMapper(sortings []grpcServer.SortingOptions) (priceRanges []domain.PriceRangeType, priceSorting domain.PriceSortingType) {
	for _, sorting := range sortings {
		if sorting == grpcServer.SortingOptions_PRICE_ASCENDING {
			priceSorting = domain.PRICE_ASCENDING
		} else if sorting == grpcServer.SortingOptions_PRICE_DESCENDING {
			priceSorting = domain.PRICE_DESCENDING
		} else if sorting == grpcServer.SortingOptions_DISCOUNTRATE_ASCENDING {
			priceSorting = domain.DISCOUNTRATE_ASCENDING
		} else if sorting == grpcServer.SortingOptions_DISCOUNTRATE_DESCENDING {
			priceSorting = domain.DISCOUNTRATE_DESCENDING
		} else if sorting == grpcServer.SortingOptions_CREATED_DESCENDING {
			priceSorting = domain.CREATED_DESCENDING
		} else {
			if sorting == grpcServer.SortingOptions_DISCOUNT_0_30 {
				priceRanges = append(priceRanges, domain.PRICE_RANGE_30)
			} else if sorting == grpcServer.SortingOptions_DISCOUNT_30_50 {
				priceRanges = append(priceRanges, domain.PRICE_RANGE_50)
			} else if sorting == grpcServer.SortingOptions_DISCOUNT_50_70 {
				priceRanges = append(priceRanges, domain.PRICE_RANGE_70)
			} else {
				priceRanges = append(priceRanges, domain.PRICE_RANGE_100)
			}
		}
	}

	return
}

func ExhibitionHistoryMapper(exhibitionHistory *domain.ExhibitionHistoryDAO) map[string]string {
	// for check nil case
	if exhibitionHistory == nil {
		return map[string]string{}
	}
	startTime := exhibitionHistory.StartTime.Format("2006-01-02T15:04:05Z07:00")
	finishTime := exhibitionHistory.FinishTime.Format("2006-01-02T15:04:05Z07:00")
	return map[string]string{
		"exhibition_id": exhibitionHistory.ExhibitionID,
		"title":         exhibitionHistory.Title,
		"start_time":    startTime,
		"finish_time":   finishTime,
	}
}
