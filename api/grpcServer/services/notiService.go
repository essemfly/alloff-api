package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer"
)

type NotiService struct {
	grpcServer.NotificationServer
}

func (s *NotiService) CreateNoti(ctx context.Context, req *grpcServer.CreateNotiRequest) (*grpcServer.CreateNotiResponse, error) {
	return nil, nil
}

func (s *NotiService) ListNoti(ctx context.Context, req *grpcServer.ListNotiRequest) (*grpcServer.ListNotiResponse, error) {
	return nil, nil
}
