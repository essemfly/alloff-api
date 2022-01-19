package grpcServer

import "context"

type BrandService struct {
	BrandServer
}

func (s *BrandService) CreateBrand(ctx context.Context, req *CreateBrandRequest) (*CreateBrandResponse, error) {
	return nil, nil
}

func (s *BrandService) ListBrand(ctx context.Context, req *ListBrandRequest) (*ListBrandResponse, error) {
	return nil, nil
}
