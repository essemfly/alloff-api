package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer"
)

type ProductGroupService struct {
	grpcServer.ProductGroupServer
}

func (s *ProductGroupService) GetProductGroup(ctx context.Context, req *grpcServer.GetProductGroupRequest) (*grpcServer.GetProductGroupResponse, error) {
	return nil, nil
}

func (s *ProductGroupService) CreateProductGroup(ctx context.Context, req *grpcServer.CreateProductGroupRequest) (*grpcServer.CreateProductGroupResponse, error) {
	return nil, nil
}

func (s *ProductGroupService) ListProductGroups(ctx context.Context, req *grpcServer.ListProductGroupsRequest) (*grpcServer.ListProductGroupsResponse, error) {
	return nil, nil
}

func (s *ProductGroupService) PushProducts(ctx context.Context, req *grpcServer.PushProductsRequest) (*grpcServer.PushProductsResponse, error) {
	return nil, nil
}
