package services

import (
	"context"
	"log"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
)

type ProductService struct {
	grpcServer.ProductServer
}

func (s *ProductService) GetProduct(ctx context.Context, req *grpcServer.GetProductRequest) (*grpcServer.GetProductResponse, error) {
	return nil, nil
}

func (s *ProductService) PutProduct(ctx context.Context, req *grpcServer.PutProductRequest) (*grpcServer.PutProductResponse, error) {
	return nil, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *grpcServer.ListProductsRequest) (*grpcServer.ListProductsResponse, error) {
	log.Println("Request", req)
	brandID := "61db9501dc6a9bb988410a35"
	if req.Query.BrandId != nil {
		brandID = *req.Query.BrandId
	}

	products, cnt, err := product.ProductsListing(int(req.Offset), int(req.Limit), brandID, "", "", nil)
	if err != nil {
		return nil, err
	}

	pds := []*grpcServer.ProductMessage{}

	for _, pd := range products {
		pds = append(pds, ProductMapper(pd))
	}

	ret := &grpcServer.ListProductsResponse{
		Offset:      req.Offset,
		Limit:       req.Limit,
		TotalCounts: int32(cnt),
		Products:    pds,
	}

	return ret, nil
}

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
