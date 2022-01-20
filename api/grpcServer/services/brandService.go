package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer"
)

type BrandService struct {
	grpcServer.BrandServer
}

func (s *BrandService) CreateBrand(ctx context.Context, req *grpcServer.CreateBrandRequest) (*grpcServer.CreateBrandResponse, error) {
	return nil, nil
}

func (s *BrandService) ListBrand(ctx context.Context, req *grpcServer.ListBrandRequest) (*grpcServer.ListBrandResponse, error) {
	return nil, nil
}
