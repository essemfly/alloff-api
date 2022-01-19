package grpcServer

import "context"

type ProductService struct {
	ProductServer
}

func (s *ProductService) GetProduct(ctx context.Context, req *GetProdcutRequest) (*GetProductResponse, error) {
	return nil, nil
}

func (s *ProductService) PutProduct(ctx context.Context, req *PutProductRequest) (*PutProductResponse, error) {
	return nil, nil
}

func (s *ProductService) ListProducts(ctx context.Context, req *ListProductsRequest) (*ListProductsResponse, error) {
	return nil, nil
}
