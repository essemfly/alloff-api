package mapper

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

func ProductMapper(pd *domain.ProductDAO) *grpcServer.ProductMessage {
	var alloffCategoryName, alloffCategoryID string
	IsClassifiedDone, IsClassifiedTouched := false, false
	if pd.AlloffCategories != nil {
		IsClassifiedDone, IsClassifiedTouched = pd.AlloffCategories.Done, pd.AlloffCategories.Touched
		if pd.AlloffCategories.Done {
			if pd.AlloffCategories.Second != nil {
				alloffCategoryName = pd.AlloffCategories.Second.Name
				alloffCategoryID = pd.AlloffCategories.Second.ID.Hex()
			} else if pd.AlloffCategories.First != nil {
				alloffCategoryName = pd.AlloffCategories.First.Name
				alloffCategoryID = pd.AlloffCategories.First.ID.Hex()
			}
		}
	}
	totalScore := pd.Score.TotalScore

	images := pd.ProductInfo.Images
	if pd.IsImageCached {
		images = pd.Images
	}

	return &grpcServer.ProductMessage{
		AlloffProductId:      pd.ID.Hex(),
		AlloffName:           pd.AlloffName,
		IsForeignDelivery:    pd.ProductInfo.Source.IsForeignDelivery,
		ProductId:            pd.ProductInfo.ProductID,
		ProductUrl:           pd.ProductInfo.ProductUrl,
		OriginalPrice:        int32(pd.OriginalPrice),
		DiscountedPrice:      int32(pd.DiscountedPrice),
		SpecialPrice:         int32(pd.SpecialPrice),
		BrandKorName:         pd.ProductInfo.Brand.KorName,
		Inventory:            InventoryMapper(pd),
		Description:          pd.SalesInstruction.Description.Texts,
		EarliestDeliveryDays: int32(pd.SalesInstruction.DeliveryDescription.EarliestDeliveryDays),
		LatestDeliveryDays:   int32(pd.SalesInstruction.DeliveryDescription.LatestDeliveryDays),
		RefundFee:            int32(pd.SalesInstruction.CancelDescription.RefundFee),
		IsRefundPossible:     pd.SalesInstruction.CancelDescription.RefundAvailable,
		Images:               images,
		DescriptionImages:    pd.SalesInstruction.Description.Images,
		CategoryName:         pd.ProductInfo.Category.Name,
		AlloffCategoryName:   alloffCategoryName,
		AlloffCategoryId:     alloffCategoryID,
		IsRemoved:            pd.Removed,
		IsSoldout:            pd.Soldout,
		TotalScore:           int32(totalScore),
		ModuleName:           pd.ProductInfo.Source.CrawlModuleName,
		IsClassifiedDone:     IsClassifiedDone,
		IsClassifiedTouched:  IsClassifiedTouched,
		ProductInfos:         pd.ProductInfo.Information,
		DescriptionInfos:     pd.SalesInstruction.Description.Infos,
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
