package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func ProductMapper(pd *domain.ProductDAO) *grpcServer.ProductMessage {
	alloffCategoryName := ""
	totalScore := 0

	return &grpcServer.ProductMessage{
		AlloffProductId:      pd.ID.Hex(),
		AlloffName:           pd.AlloffName,
		IsForeignDelivery:    pd.ProductInfo.Source.IsForeignDelivery,
		ProductId:            pd.ProductInfo.ProductID,
		OriginalPrice:        int32(pd.ProductInfo.Price.OriginalPrice),
		DiscountedPrice:      int32(pd.DiscountedPrice),
		SpecialPrice:         int32(pd.SpecialPrice),
		BrandKorName:         pd.ProductInfo.Brand.KorName,
		Inventory:            InventoryMapper(pd),
		Description:          pd.SalesInstruction.Description.Texts,
		EarliestDeliveryDays: int32(pd.SalesInstruction.DeliveryDescription.LatestDeliveryDays),
		LatestDeliveryDays:   int32(pd.SalesInstruction.DeliveryDescription.EarliestDeliveryDays),
		RefundFee:            int32(pd.SalesInstruction.CancelDescription.RefundFee),
		IsRefundPossible:     pd.SalesInstruction.CancelDescription.RefundAvailable,
		Images:               pd.ProductInfo.Images,
		DescriptionImages:    pd.SalesInstruction.Description.Images,
		CategoryName:         pd.ProductInfo.Category.Name,
		AlloffCategoryName:   alloffCategoryName,
		IsRemoved:            pd.Removed,
		IsSoldout:            pd.Soldout,
		TotalScore:           int32(totalScore),
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
