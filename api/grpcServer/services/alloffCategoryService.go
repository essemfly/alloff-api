package services

import (
	"context"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
)

type AlloffCategoryService struct {
	grpcServer.AlloffCategoryServer
}

func (s *AlloffCategoryService) ListAlloffCategory(ctx context.Context, req *grpcServer.ListAlloffCategoryRequest) (*grpcServer.ListAlloffCategoryResponse, error) {
	alloffCatDaos, _ := ioc.Repo.AlloffCategories.List(&req.ParentId)

	cats := []*grpcServer.AlloffCategoryMessage{}
	for _, catDao := range alloffCatDaos {
		cat := mapper.AlloffCategoryMapper(catDao)
		if cat != nil {
			cats = append(cats, cat)
		}
	}
	return &grpcServer.ListAlloffCategoryResponse{
		Categories: cats,
	}, nil
}
