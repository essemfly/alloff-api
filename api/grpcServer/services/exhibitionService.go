package services

import (
	"context"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/exhibition"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
	"go.uber.org/zap"
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
		Exhibition: mapper.ExhibitionMapper(exhibitionDao, false),
	}, nil
}

func (s *ExhibitionService) ListExhibitions(ctx context.Context, req *grpcServer.ListExhibitionsRequest) (*grpcServer.ListExhibitionsResponse, error) {
	onlyLive := true
	if !req.IsLive {
		onlyLive = false
	}

	groupType := domain.EXHIBITION_NORMAL
	if req.GroupType == grpcServer.ExhibitionType_EXHIBITION_GROUPDEAL {
		groupType = domain.EXHIBITION_GROUPDEAL
	} else if req.GroupType == grpcServer.ExhibitionType_EXHIBITION_TIMEDEAL {
		groupType = domain.EXHIBITION_TIMEDEAL
	}

	query := ""
	if req.Query != nil {
		query = *req.Query
	}

	exhibitionDaos, cnt, err := ioc.Repo.Exhibitions.List(int(req.Offset), int(req.Limit), onlyLive, domain.EXHIBITION_STATUS_ALL, groupType, query)
	if err != nil {
		return nil, err
	}

	exs := []*grpcServer.ExhibitionMessage{}
	for _, exDao := range exhibitionDaos {
		exs = append(exs, mapper.ExhibitionMapper(exDao, true))
	}
	return &grpcServer.ListExhibitionsResponse{
		Exhibitions: exs,
		Offset:      req.Offset,
		Limit:       req.Limit,
		TotalCounts: int32(cnt),
		IsLive:      onlyLive,
		Query:       query,
		GroupType:   mapper.ExhibitionGroupTypeMapper(groupType),
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
	if req.Subtitle != nil {
		exDao.SubTitle = *req.Subtitle
	}
	if req.Description != nil {
		exDao.Description = *req.Description
	}
	if len(req.Tags) > 0 {
		exDao.Tags = req.Tags
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
		for _, existPg := range exDao.ProductGroups {
			removed := true
			for _, newPgId := range req.PgIds {
				if existPg.ID.Hex() == newPgId {
					removed = false
					break
				}
			}
			if removed {
				err := exhibition.ProductGroupSyncer(existPg)
				if err != nil {
					config.Logger.Error("err found on removing pg", zap.Error(err))
				}
			}
		}

		for _, pgID := range req.PgIds {
			pg, err := ioc.Repo.ProductGroups.Get(pgID)
			if err != nil {
				log.Println("get product group failed: "+pgID, err)
				continue
			}
			pgs = append(pgs, pg)
		}
		exDao.ProductGroups = pgs
	}

	newExhibitionDao, err := exhibition.ExhibitionSyncer(exDao)
	if err != nil {
		return nil, err
	}

	return &grpcServer.EditExhibitionResponse{
		Exhibition: mapper.ExhibitionMapper(newExhibitionDao, false),
	}, nil
}

func (s *ExhibitionService) CreateExhibition(ctx context.Context, req *grpcServer.CreateExhibitionRequest) (*grpcServer.CreateExhibitionResponse, error) {
	layout := "2006-01-02T15:04:05Z07:00"

	startTimeObj, _ := time.Parse(layout, req.StartTime)
	finishTimeObj, _ := time.Parse(layout, req.FinishTime)

	exhibitionGroupType := domain.EXHIBITION_NORMAL
	if req.ExhibitionType == grpcServer.ExhibitionType_EXHIBITION_TIMEDEAL {
		exhibitionGroupType = domain.EXHIBITION_TIMEDEAL
	} else if req.ExhibitionType == grpcServer.ExhibitionType_EXHIBITION_GROUPDEAL {
		exhibitionGroupType = domain.EXHIBITION_GROUPDEAL
	}

	exhibitionReq := &exhibition.ExhibitionRequest{
		BannerImage:     req.BannerImage,
		ThumbnailImage:  req.ThumbnailImage,
		Title:           req.Title,
		SubTitle:        req.Subtitle,
		Description:     req.Description,
		Tags:            req.Tags,
		ProductGroupIDs: req.PgIds,
		ExhibitionType:  exhibitionGroupType,
		StartTime:       startTimeObj,
		FinishTime:      finishTimeObj,
	}

	exDao, err := exhibition.AddExhibition(exhibitionReq)
	if err != nil {
		return nil, err
	}

	exDao, err = exhibition.ExhibitionSyncer(exDao)
	if err != nil {
		return nil, err
	}

	return &grpcServer.CreateExhibitionResponse{
		Exhibition: mapper.ExhibitionMapper(exDao, false),
	}, nil
}
