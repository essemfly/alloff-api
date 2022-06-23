package services

import (
	"context"
	"log"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TopBannerService struct {
	grpcServer.TopBannerServer
}

func (s *TopBannerService) GetTopBanner(ctx context.Context, req *grpcServer.GetTopBannerRequest) (*grpcServer.GetTopBannerResponse, error) {
	bannerDao, err := ioc.Repo.TopBanners.Get(req.BannerId)
	if err != nil {
		return nil, err
	}

	return &grpcServer.GetTopBannerResponse{
		Banner: mapper.TopbannerMapper(bannerDao),
	}, nil
}

func (s *TopBannerService) ListTopBanners(ctx context.Context, req *grpcServer.ListTopBannersRequest) (*grpcServer.ListTopBannersResponse, error) {
	onlyLive := false
	bannerDaos, cnt, err := ioc.Repo.TopBanners.List(int(req.Offset), int(req.Limit), onlyLive)
	if err != nil {
		return nil, err
	}

	banners := []*grpcServer.TopBannerMessage{}
	for _, bannerDao := range bannerDaos {
		banners = append(banners, mapper.TopbannerMapper(bannerDao))
	}

	return &grpcServer.ListTopBannersResponse{
		Banners:     banners,
		Offset:      req.Offset,
		Limit:       req.Limit,
		TotalCounts: int32(cnt),
	}, nil
}

func (s *TopBannerService) EditTopBanner(ctx context.Context, req *grpcServer.EditTopBannerRequest) (*grpcServer.EditTopBannerResponse, error) {
	bannerDao, err := ioc.Repo.TopBanners.Get(req.BannerId)
	if err != nil {
		return nil, err
	}

	if req.BannerImage != nil {
		bannerDao.ImageUrl = *req.BannerImage
	}
	if req.ExhibitionId != nil {
		bannerDao.ExhibitionID = *req.ExhibitionId
	}
	if req.Title != nil {
		bannerDao.Title = *req.Title
	}
	if req.Subtitle != nil {
		bannerDao.SubTitle = *req.Subtitle
	}
	if req.IsLive != nil {
		bannerDao.IsLive = *req.IsLive
	}
	if req.Weight != nil {
		bannerDao.Weight = int(*req.Weight)
	}

	newBannerDao, err := ioc.Repo.TopBanners.Update(bannerDao)
	if err != nil {
		return nil, err
	}

	return &grpcServer.EditTopBannerResponse{
		Banner: mapper.TopbannerMapper(newBannerDao),
	}, nil
}

func (s *TopBannerService) CreateTopBanner(ctx context.Context, req *grpcServer.CreateTopBannerRequest) (*grpcServer.CreateTopBannerResponse, error) {

	bannerDao := &domain.TopBannerDAO{
		ID:           primitive.NewObjectID(),
		ImageUrl:     req.BannerImage,
		ExhibitionID: req.ExhibitionId,
		Title:        req.Title,
		SubTitle:     req.Subtitle,
		IsLive:       req.IsLive,
		Weight:       int(req.Weight),
	}

	newBannerDao, err := ioc.Repo.TopBanners.Insert(bannerDao)
	if err != nil {
		log.Println("TopBanner create error", err)
		return nil, err
	}

	return &grpcServer.CreateTopBannerResponse{
		Banner: mapper.TopbannerMapper(newBannerDao),
	}, nil
}
