package services

import (
	"context"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/broker"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExhibitionService struct {
	grpcServer.ExhibitionServer
}

func (s *ExhibitionService) GetExhibition(ctx context.Context, req *grpcServer.GetExhibitionRequest) (*grpcServer.GetExhibitionResponse, error) {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(req.ExhibitionId)
	if err != nil {
		return nil, err
	}

	return &grpcServer.GetExhibitionResponse{
		Exhibition: mapper.ExhibitionMapper(exhibitionDao),
	}, nil
}

func (s *ExhibitionService) ListExhibitions(ctx context.Context, req *grpcServer.ListExhibitionsRequest) (*grpcServer.ListExhibitionsResponse, error) {
	onlyLive := false
	exhibitionDaos, cnt, err := ioc.Repo.Exhibitions.List(int(req.Offset), int(req.Limit), onlyLive)
	if err != nil {
		return nil, err
	}

	exs := []*grpcServer.ExhibitionMessage{}
	for _, exDao := range exhibitionDaos {
		exs = append(exs, mapper.ExhibitionMapper(exDao))
	}
	return &grpcServer.ListExhibitionsResponse{
		Exhibitions: exs,
		Offset:      req.Offset,
		Limit:       req.Limit,
		TotalCounts: int32(cnt),
	}, nil
}

func (s *ExhibitionService) EditExhibition(ctx context.Context, req *grpcServer.EditExhibitionRequest) (*grpcServer.EditExhibitionResponse, error) {
	exDao, err := ioc.Repo.Exhibitions.Get(req.ExhibitionId)
	if err != nil {
		return nil, err
	}

	layout := "2006-01-02T15:04:05Z07:00"

	if req.BannerImage != nil {
		exDao.BannerImage = *req.BannerImage
	}
	if req.ThumbnailImage != nil {
		exDao.ThumbnailImage = *req.ThumbnailImage
	}
	if req.Title != nil {
		exDao.Title = *req.Title
	}
	if req.Description != nil {
		exDao.Description = *req.Description
	}
	if req.StartTime != nil {
		startTimeObj, _ := time.Parse(layout, *req.StartTime)
		exDao.StartTime = startTimeObj
	}
	if req.FinishTime != nil {
		finishTimeObj, _ := time.Parse(layout, *req.FinishTime)
		exDao.FinishTime = finishTimeObj
	}
	if req.IsLive != nil {
		exDao.IsLive = *req.IsLive
	}
	if req.PgIds != nil && len(req.PgIds) > 0 {
		pgs := []*domain.ProductGroupDAO{}
		for _, pgID := range req.PgIds {
			pg, err := ioc.Repo.ProductGroups.Get(pgID)
			if err != nil {
				log.Println("get product group failed: "+pgID, err)
				continue
			}
			pg.StartTime = exDao.StartTime
			pg.FinishTime = exDao.FinishTime
			newPg, err := ioc.Repo.ProductGroups.Upsert(pg)
			if err != nil {
				log.Println("update product group failed: "+pgID, err)
			}
			pgs = append(pgs, newPg)
		}

		exDao.ProductGroups = pgs
	}

	newExhibitionDao, err := ioc.Repo.Exhibitions.Upsert(exDao)
	if err != nil {
		return nil, err
	}

	broker.ExhibitionSyncer(newExhibitionDao)

	return &grpcServer.EditExhibitionResponse{
		Exhibition: mapper.ExhibitionMapper(newExhibitionDao),
	}, nil
}

func (s *ExhibitionService) CreateExhibition(ctx context.Context, req *grpcServer.CreateExhibitionRequest) (*grpcServer.CreateExhibitionResponse, error) {
	layout := "2006-01-02T15:04:05Z07:00"

	startTimeObj, _ := time.Parse(layout, req.StartTime)
	finishTimeObj, _ := time.Parse(layout, req.FinishTime)

	exDao := &domain.ExhibitionDAO{
		ID:             primitive.NewObjectID(),
		BannerImage:    req.BannerImage,
		ThumbnailImage: req.ThumbnailImage,
		Title:          req.Title,
		SubTitle:       req.Subtitle,
		Description:    req.Description,
		StartTime:      startTimeObj,
		FinishTime:     finishTimeObj,
		IsLive:         false,
	}

	pgs := []*domain.ProductGroupDAO{}
	for _, pgID := range req.PgIds {
		pg, err := ioc.Repo.ProductGroups.Get(pgID)
		if err != nil {
			log.Println("get product group failed: "+pgID, err)
			continue
		}
		pg.StartTime = startTimeObj
		pg.FinishTime = startTimeObj
		pg.GroupType = domain.PRODUCT_GROUP_EXHIBITION
		pg.ExhibitionID = exDao.ID.Hex()
		newPg, err := ioc.Repo.ProductGroups.Upsert(pg)
		if err != nil {
			log.Println("update product group failed: "+pgID, err)
		}
		pgs = append(pgs, newPg)
	}

	exDao.ProductGroups = pgs

	newExDao, err := ioc.Repo.Exhibitions.Upsert(exDao)
	if err != nil {
		log.Println("Exhibition create error", err)
		return nil, err
	}

	return &grpcServer.CreateExhibitionResponse{
		Exhibition: mapper.ExhibitionMapper(newExDao),
	}, nil
}
