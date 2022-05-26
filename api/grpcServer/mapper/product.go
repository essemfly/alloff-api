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

	return &grpcServer.ProductMessage{
		AlloffProductId:      pdInfo.ID.Hex(),
		AlloffName:           pdInfo.AlloffName,
		IsForeignDelivery:    pdInfo.Source.IsForeignDelivery,
		ProductId:            pdInfo.ProductID,
		ProductUrl:           pdInfo.ProductUrl,
		OriginalPrice:        int32(pdInfo.Price.OriginalPrice),
		DiscountedPrice:      int32(pdInfo.Price.CurrentPrice),
		SpecialPrice:         int32(pdInfo.Price.CurrentPrice),
		BrandKorName:         pdInfo.Brand.KorName,
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
		DescriptionInfos:     pdInfo.SalesInstruction.Description.Infos,
		ThumbnailImage:       pdInfo.ThumbnailImage,
		ProductTypes:         ProductTypeMapper(pdInfo.ProductType),
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
		} else if pdtype == domain.Sports {
			productTypes = append(productTypes, grpcServer.ProductType_SPORTS)
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
				AlloffSize: alloffSizes,
				Quantity:   int32(inv.Quantity),
				Size:       inv.Size,
			})
		} else {
			invMessage = append(invMessage, &grpcServer.ProductInventoryMessage{
				AlloffSize: nil,
				Quantity:   int32(inv.Quantity),
				Size:       inv.Size,
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
