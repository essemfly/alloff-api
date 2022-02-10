package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer"
)

type ExhibitionService struct {
	grpcServer.HomeTabItemServer
}

func (s *ExhibitionService) GetExhibition(ctx context.Context, req *grpcServer.GetExhibitionRequest) (*grpcServer.GetExhibitionResponse, error) {
	return nil, nil
}

func (s *ExhibitionService) ListExhibitions(ctx context.Context, req *grpcServer.ListExhibitionsRequest) (*grpcServer.ListExhibitionsResponse, error) {
	return nil, nil
}

func (s *ExhibitionService) EditExhibition(ctx context.Context, req *grpcServer.EditExhibitionRequest) (*grpcServer.EditExhibitionResponse, error) {
	return nil, nil
}

func (s *ExhibitionService) CreateExhibition(ctx context.Context, req *grpcServer.CreateExhibitionRequest) (*grpcServer.CreateExhibitionResponse, error) {
	return nil, nil
}
