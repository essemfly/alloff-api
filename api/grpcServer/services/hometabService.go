package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer"
)

type HomeTabService struct {
	grpcServer.HomeTabItemServer
}

func (s *HomeTabService) GetHomeTabItem(ctx context.Context, req *grpcServer.GetHomeTabItemRequest) (*grpcServer.GetHomeTabItemResponse, error) {
	return nil, nil
}

func (s *HomeTabService) ListHomeTabItems(ctx context.Context, req *grpcServer.ListHomeTabItemsRequest) (*grpcServer.ListHomeTabsItemResponse, error) {
	return nil, nil
}

func (s *HomeTabService) EditHomeTabItem(ctx context.Context, req *grpcServer.EditHomeTabItemRequest) (*grpcServer.EditHomeTabItemResponse, error) {
	return nil, nil
}

func (s *HomeTabService) CreateHomeTabItem(ctx context.Context, req *grpcServer.CreateHomeTabItemRequest) (*grpcServer.CreateHomeTabItemResponse, error) {
	return nil, nil
}
