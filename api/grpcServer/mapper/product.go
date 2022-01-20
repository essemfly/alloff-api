package mapper

import (
	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func ProductMapper(pd *domain.ProductDAO) *grpcServer.ProductMessage {
	return &grpcServer.ProductMessage{
		ProductId:       pd.ProductInfo.ProductID,
		AlloffName:      pd.AlloffName,
		DiscountedPrice: int32(pd.DiscountedPrice),
		DiscountRate:    int32(pd.DiscountRate),
		SpecialPrice:    int32(pd.SpecialPrice),
		BrandKorName:    pd.ProductInfo.Brand.KorName,
		CategoryName:    pd.ProductInfo.Category.Name,
		IsRemoved:       pd.Removed,
		IsSoldout:       pd.Soldout,
		Inventory:       InventoryMapper(pd),
		TotalScore:      int32(pd.Score.TotalScore),
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
