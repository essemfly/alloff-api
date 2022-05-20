package mapper

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

func ProductMapper(pd *domain.ProductDAO) *grpcServer.ProductMessage {
	var alloffCategoryName, alloffCategoryID string
	IsClassifiedDone, IsClassifiedTouched := false, false
	if pd.ProductInfo.AlloffCategory != nil {
		IsClassifiedDone, IsClassifiedTouched = pd.ProductInfo.AlloffCategory.Done, pd.ProductInfo.AlloffCategory.Touched
		if pd.ProductInfo.AlloffCategory.Done {
			if pd.ProductInfo.AlloffCategory.Second != nil {
				alloffCategoryName = pd.ProductInfo.AlloffCategory.Second.Name
				alloffCategoryID = pd.ProductInfo.AlloffCategory.Second.ID.Hex()
			} else if pd.ProductInfo.AlloffCategory.First != nil {
				alloffCategoryName = pd.ProductInfo.AlloffCategory.First.Name
				alloffCategoryID = pd.ProductInfo.AlloffCategory.First.ID.Hex()
			}
		}
	}
	totalScore := pd.Score.TotalScore

	images := pd.ProductInfo.Images
	if pd.ProductInfo.IsImageCached {
		images = pd.ProductInfo.CachedImages
	}

	return &grpcServer.ProductMessage{
		AlloffProductId:      pd.ID.Hex(),
		AlloffName:           pd.AlloffName,
		IsForeignDelivery:    pd.ProductInfo.Source.IsForeignDelivery,
		ProductId:            pd.ProductInfo.ProductID,
		ProductUrl:           pd.ProductInfo.ProductUrl,
		OriginalPrice:        int32(pd.OriginalPrice),
		DiscountedPrice:      int32(pd.DiscountedPrice),
		BrandKorName:         pd.ProductInfo.Brand.KorName,
		Inventory:            InventoryMapper(pd),
		Description:          pd.ProductInfo.SalesInstruction.Description.Texts,
		EarliestDeliveryDays: int32(pd.ProductInfo.SalesInstruction.DeliveryDescription.EarliestDeliveryDays),
		LatestDeliveryDays:   int32(pd.ProductInfo.SalesInstruction.DeliveryDescription.LatestDeliveryDays),
		RefundFee:            int32(pd.ProductInfo.SalesInstruction.CancelDescription.RefundFee),
		IsRefundPossible:     pd.ProductInfo.SalesInstruction.CancelDescription.RefundAvailable,
		Images:               images,
		DescriptionImages:    pd.ProductInfo.SalesInstruction.Description.Images,
		CategoryName:         pd.ProductInfo.Category.Name,
		AlloffCategoryName:   alloffCategoryName,
		AlloffCategoryId:     alloffCategoryID,
		IsRemoved:            pd.IsRemoved,
		IsSoldout:            pd.IsSoldout,
		TotalScore:           int32(totalScore),
		ModuleName:           pd.ProductInfo.Source.CrawlModuleName,
		IsClassifiedDone:     IsClassifiedDone,
		IsClassifiedTouched:  IsClassifiedTouched,
		ProductInfos:         pd.ProductInfo.SalesInstruction.Information,
		DescriptionInfos:     pd.ProductInfo.SalesInstruction.Description.Infos,
		ThumbnailImage:       pd.ThumbnailImage,
	}
}

func InventoryMapper(pd *domain.ProductDAO) []*grpcServer.ProductInventoryMessage {
	invMessages := []*grpcServer.ProductInventoryMessage{}
	for _, inv := range pd.Inventory {
		invMessages = append(invMessages, &grpcServer.ProductInventoryMessage{
			Size:     inv.Size,
			Quantity: int32(inv.Quantity),
		})
	}
	return invMessages
}

func ProductSortingAndRangesMapper(sortings []grpcServer.SortingOptions) (priceRanges []product.PriceRangeType, priceSorting product.PriceSortingType) {
	for _, sorting := range sortings {
		if sorting == grpcServer.SortingOptions_PRICE_ASCENDING {
			priceSorting = product.PRICE_ASCENDING
		} else if sorting == grpcServer.SortingOptions_PRICE_DESCENDING {
			priceSorting = product.PRICE_DESCENDING
		} else if sorting == grpcServer.SortingOptions_DISCOUNTRATE_ASCENDING {
			priceSorting = product.DISCOUNTRATE_ASCENDING
		} else if sorting == grpcServer.SortingOptions_DISCOUNTRATE_DESCENDING {
			priceSorting = product.DISCOUNTRATE_DESCENDING
		} else {
			if sorting == grpcServer.SortingOptions_DISCOUNT_0_30 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_30)
			} else if sorting == grpcServer.SortingOptions_DISCOUNT_30_50 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_50)
			} else if sorting == grpcServer.SortingOptions_DISCOUNT_50_70 {
				priceRanges = append(priceRanges, product.PRICE_RANGE_70)
			} else {
				priceRanges = append(priceRanges, product.PRICE_RANGE_100)
			}
		}
	}

	return
}
