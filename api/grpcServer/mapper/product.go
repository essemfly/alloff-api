package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func ProductMapper(pd *domain.ProductDAO) *grpcServer.ProductMessage {
	return &grpcServer.ProductMessage{
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
		AlloffCategoryName:   pd.AlloffCategories.First.KeyName,
		IsRemoved:            pd.Removed,
		IsSoldout:            pd.Soldout,
		TotalScore:           int32(pd.Score.TotalScore),
	}
}

func InventoryMapper(pd *domain.ProductDAO) []*grpcServer.InventoryMessage {
	invMessages := []*grpcServer.InventoryMessage{}
	for _, inv := range pd.Inventory {
		invMessages = append(invMessages, &grpcServer.InventoryMessage{
			Size:     inv.Size,
			Quantity: int32(inv.Quantity),
		})
	}
	return invMessages
}
