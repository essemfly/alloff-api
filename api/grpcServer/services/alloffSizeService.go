package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlloffSizeService struct {
	grpcServer.AlloffSizeServer
}

func (s *AlloffSizeService) ListAlloffSize(ctx context.Context, req *grpcServer.ListAlloffSizeRequest) (*grpcServer.ListAlloffSizeResponse, error) {
	alloffSizeDaos, _, _ := ioc.Repo.AlloffSizes.List(0, 10000)

	sizes := []*grpcServer.AlloffSizeMessage{}
	for _, alloffSizeDao := range alloffSizeDaos {
		alloffSize := mapper.AlloffSizeMapper(alloffSizeDao)
		if alloffSize != nil {
			sizes = append(sizes, alloffSize)
		}
	}
	return &grpcServer.ListAlloffSizeResponse{
		AlloffSizes: sizes,
	}, nil
}

func (s *AlloffSizeService) EditAlloffSize(ctx context.Context, req *grpcServer.EditAlloffSizeRequest) (*grpcServer.AlloffSizeMessage, error) {
	alloffSizeDao, err := ioc.Repo.AlloffSizes.Get(req.AlloffSizeId)
	if err != nil {
		return nil, err
	}

	if req.SizeName != nil {
		alloffSizeDao.AlloffSizeName = *req.SizeName
	}

	newAlloffSizeDao, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
	if err != nil {
		return nil, err
	}

	return mapper.AlloffSizeMapper(newAlloffSizeDao), nil
}

func (s *AlloffSizeService) CreateAlloffSize(ctx context.Context, req *grpcServer.CreateAlloffSizeRequest) (*grpcServer.AlloffSizeMessage, error) {
	alloffSizeDao := &domain.AlloffSizeDAO{
		ID:             primitive.NewObjectID(),
		AlloffSizeName: req.SizeName,
		AlloffCategory: nil,
	}

	newAlloffSize, err := ioc.Repo.AlloffSizes.Upsert(alloffSizeDao)
	if err != nil {
		return nil, err
	}

	return mapper.AlloffSizeMapper(newAlloffSize), nil
}

func (s *AlloffSizeService) GetAlloffSize(ctx context.Context, req *grpcServer.GetAlloffSizeRequest) (*grpcServer.AlloffSizeMessage, error) {
	alloffSizeDao, err := ioc.Repo.AlloffSizes.Get(req.AlloffSizeId)
	if err != nil {
		return nil, err
	}

	return mapper.AlloffSizeMapper(alloffSizeDao), nil
}
