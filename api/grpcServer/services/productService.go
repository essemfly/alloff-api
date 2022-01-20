package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/product"
)

type ProductService struct {
	grpcServer.ProductServer
}

func (s *ProductService) GetProduct(ctx context.Context, req *grpcServer.GetProductRequest) (*grpcServer.GetProductResponse, error) {
	pdDao, err := ioc.Repo.Products.Get(req.ProductId)
	if err != nil {
		return nil, err
	}

	return &grpcServer.GetProductResponse{
		Product: mapper.ProductMapper(pdDao),
	}, nil
}

func (s *ProductService) PutProduct(ctx context.Context, req *grpcServer.PutProductRequest) (*grpcServer.PutProductResponse, error) {
	pdDao, err := ioc.Repo.Products.Get(req.ProductId)
	if err != nil {
		return nil, err
	}

	pdDao.SpecialPrice = int(req.SpecialPrice)
	newPdDao, err := ioc.Repo.Products.Upsert(pdDao)
	if err != nil {
		return nil, err
	}
	return &grpcServer.PutProductResponse{
		Product: mapper.ProductMapper(newPdDao),
	}, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *grpcServer.ListProductsRequest) (*grpcServer.ListProductsResponse, error) {
	brandID := ""
	if req.Query.BrandId != nil {
		brandID = *req.Query.BrandId
	}
	categoryID := ""
	if req.Query.CategoryId != nil {
		categoryID = *req.Query.CategoryId
	}

	products, cnt, err := product.ProductsListing(int(req.Offset), int(req.Limit), brandID, categoryID, "", nil)
	if err != nil {
		return nil, err
	}

	pds := []*grpcServer.ProductMessage{}

	for _, pd := range products {
		pds = append(pds, mapper.ProductMapper(pd))
	}

	ret := &grpcServer.ListProductsResponse{
		Offset:      req.Offset,
		Limit:       req.Limit,
		TotalCounts: int32(cnt),
		Products:    pds,
	}

	return ret, nil
}
