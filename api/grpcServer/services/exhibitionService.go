package services

import (
	"context"
	"github.com/lessbutter/alloff-api/pkg/exhibition"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/broker"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
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
		Exhibition: mapper.ExhibitionMapper(exhibitionDao, false),
	}, nil
}

func (s *ExhibitionService) ListExhibitions(ctx context.Context, req *grpcServer.ListExhibitionsRequest) (*grpcServer.ListExhibitionsResponse, error) {
	onlyLive := false
	if req.IsLive {
		onlyLive = true
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

	exhibitionDaos, cnt, err := ioc.Repo.Exhibitions.List(int(req.Offset), int(req.Limit), onlyLive, groupType, query)
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
	if req.RecruitStartTime != nil {
		recruitStartTimeObj, _ := time.Parse(layout, *req.RecruitStartTime)
		exDao.RecruitStartTime = recruitStartTimeObj
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

	if req.NumUsersRequired != nil {
		exDao.NumUsersRequired = int(*req.NumUsersRequired)
	}
	if req.AllowOldUser != nil {
		exDao.AllowOldUser = *req.AllowOldUser
	}

	banners := []domain.ExhibitionBanner{}
	if req.Banners != nil {
		for _, banner := range req.Banners {
			bannerDao := domain.ExhibitionBanner{
				ImgUrl:         banner.ImgUrl,
				Title:          banner.Title,
				Subtitle:       banner.Subtitle,
				ProductGroupId: banner.ProductGroupId,
			}
			banners = append(banners, bannerDao)
		}
		exDao.Banners = banners
	}

	pgType := domain.PRODUCT_GROUP_EXHIBITION
	if exDao.ExhibitionType == domain.EXHIBITION_TIMEDEAL {
		pgType = domain.PRODUCT_GROUP_TIMEDEAL
	} else if exDao.ExhibitionType == domain.EXHIBITION_GROUPDEAL {
		pgType = domain.PRODUCT_GROUP_GROUPDEAL
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
			if pg.Brand != nil {
				pgType = domain.PRODUCT_GROUP_BRAND_TIMEDEAL
			}
			pg.GroupType = pgType
			pg.ExhibitionID = exDao.ID.Hex()
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

	go broker.ExhibitionSyncer(newExhibitionDao)
	go exhibition.UpdateCheapestPrice(newExhibitionDao)

	return &grpcServer.EditExhibitionResponse{
		Exhibition: mapper.ExhibitionMapper(newExhibitionDao, false),
	}, nil
}

func (s *ExhibitionService) CreateExhibition(ctx context.Context, req *grpcServer.CreateExhibitionRequest) (*grpcServer.CreateExhibitionResponse, error) {
	layout := "2006-01-02T15:04:05Z07:00"

	recruitStartTimeObj, _ := time.Parse(layout, req.RecruitStartTime)
	startTimeObj, _ := time.Parse(layout, req.StartTime)
	finishTimeObj, _ := time.Parse(layout, req.FinishTime)

	exhibitionGroupType := domain.EXHIBITION_NORMAL
	if req.ExhibitionType == grpcServer.ExhibitionType_EXHIBITION_TIMEDEAL {
		exhibitionGroupType = domain.EXHIBITION_TIMEDEAL
	} else if req.ExhibitionType == grpcServer.ExhibitionType_EXHIBITION_GROUPDEAL {
		exhibitionGroupType = domain.EXHIBITION_GROUPDEAL
	}

	banners := []domain.ExhibitionBanner{}
	if req.Banners != nil {
		for _, banner := range req.Banners {
			bannerDao := domain.ExhibitionBanner{
				ImgUrl:         banner.ImgUrl,
				Title:          banner.Title,
				Subtitle:       banner.Subtitle,
				ProductGroupId: banner.ProductGroupId,
			}
			banners = append(banners, bannerDao)
		}
	}

	numUsersRequired := 0
	if req.NumUsersRequired != nil {
		numUsersRequired = int(*req.NumUsersRequired)
	}

	exDao := &domain.ExhibitionDAO{
		ID:               primitive.NewObjectID(),
		BannerImage:      req.BannerImage,
		ThumbnailImage:   req.ThumbnailImage,
		Title:            req.Title,
		SubTitle:         req.Subtitle,
		Description:      req.Description,
		RecruitStartTime: recruitStartTimeObj,
		StartTime:        startTimeObj,
		FinishTime:       finishTimeObj,
		IsLive:           false,
		ExhibitionType:   exhibitionGroupType,
		TargetSales:      int(req.TargetSales),
		Banners:          banners,
		CreatedAt:        time.Now(),
		NumUsersRequired: numUsersRequired,
		AllowOldUser:     req.AllowOldUser,
	}

	pgs := []*domain.ProductGroupDAO{}
	for _, pgID := range req.PgIds {
		pg, err := ioc.Repo.ProductGroups.Get(pgID)
		if err != nil {
			log.Println("get product group failed: "+pgID, err)
			continue
		}
		productGroupType := pg.GroupType
		pg.StartTime = startTimeObj
		pg.FinishTime = startTimeObj
		if pg.Brand != nil {
			productGroupType = domain.PRODUCT_GROUP_BRAND_TIMEDEAL
		}
		pg.GroupType = productGroupType
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

	go exhibition.UpdateCheapestPrice(newExDao)

	return &grpcServer.CreateExhibitionResponse{
		Exhibition: mapper.ExhibitionMapper(newExDao, false),
	}, nil
}
