package grpcServer

import "context"

type NotiService struct {
	NotificationServer
}

func (s *NotiService) CreateNoti(ctx context.Context, req *CreateNotiRequest) (*CreateNotiResponse, error) {
	return nil, nil
}

func (s *NotiService) ListNoti(ctx context.Context, req *ListNotiRequest) (*ListNotiResponse, error) {
	return nil, nil
}
