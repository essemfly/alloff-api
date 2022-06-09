package services

import (
	"context"

	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

type HealthService struct {
	grpcServer.HealthServer
}

func (s *HealthService) Check(ctx context.Context, req *grpcServer.HealthCheckRequest) (*grpcServer.HealthCheckResponse, error) {
	return &grpcServer.HealthCheckResponse{
		Status: grpcServer.HealthCheckResponse_SERVING,
	}, nil
}
