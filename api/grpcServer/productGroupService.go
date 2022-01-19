package grpcServer

import (
	"context"
)

type ProductGroupService struct {
	ProductGroupServer
}

func (s *ProductGroupService) GetProductGroup(ctx context.Context, req *GetProductGroupRequest) (*GetProductGroupResponse, error) {
	return nil, nil
}

func (s *ProductGroupService) PostProductGroup(ctx context.Context, req *PostProductGroupRequest) (*PostProductGroupResponse, error) {
	return nil, nil
}

func (s *ProductGroupService) ListProductGroups(ctx context.Context, req *ListProductGroupsRequest) (*ListProductGroupsResponse, error) {
	return nil, nil
}

func (s *ProductGroupService) PushProducts(ctx context.Context, req *PushProductsRequest) (*PushProductsResponse, error) {
	return nil, nil
}
