package services

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/api/grpcServer"
	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductGroupService struct {
	grpcServer.ProductGroupServer
}

func (s *ProductGroupService) GetProductGroup(ctx context.Context, req *grpcServer.GetProductGroupRequest) (*grpcServer.GetProductGroupResponse, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}
	return &grpcServer.GetProductGroupResponse{
		Pg: mapper.ProductGroupMapper(pgDao),
	}, nil
}

func (s *ProductGroupService) CreateProductGroup(ctx context.Context, req *grpcServer.CreateProductGroupRequest) (*grpcServer.CreateProductGroupResponse, error) {
	layout := "2006-01-02T15:04:05.000Z"

	startTimeObj, _ := time.Parse(layout, req.StartTime)
	finishTimeObj, _ := time.Parse(layout, req.FinishTime)

	pgDao := &domain.ProductGroupDAO{
		Title:       req.Title,
		ShortTitle:  req.ShortTitle,
		Instruction: req.Instruction,
		ImgUrl:      req.ImageUrl,
		Products:    []*domain.ProductPriorityDAO{},
		StartTime:   startTimeObj,
		FinishTime:  finishTimeObj,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	newPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
	if err != nil {
		return nil, err
	}

	return &grpcServer.CreateProductGroupResponse{
		Pg: mapper.ProductGroupMapper(newPgDao),
	}, nil
}

func (s *ProductGroupService) ListProductGroups(ctx context.Context, req *grpcServer.ListProductGroupsRequest) (*grpcServer.ListProductGroupsResponse, error) {
	pgDaos, err := ioc.Repo.ProductGroups.List()
	if err != nil {
		return nil, err
	}

	pgs := []*grpcServer.ProductGroupMessage{}
	for _, pgDao := range pgDaos {
		pgs = append(pgs, mapper.ProductGroupMapper(pgDao))
	}
	return &grpcServer.ListProductGroupsResponse{
		Pgs: pgs,
	}, nil
}

func (s *ProductGroupService) PushProducts(ctx context.Context, req *grpcServer.PushProductsRequest) (*grpcServer.PushProductsResponse, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}

	pds := pgDao.Products
	for idx, productID := range req.ProductId {
		pd, err := ioc.Repo.Products.Get(productID)
		if err != nil {
			return nil, err
		}

		pd.ProductGroupId = &pgDao.ID
		_, err = ioc.Repo.Products.Upsert(pd)
		if err != nil {
			return nil, err
		}

		productObjID, _ := primitive.ObjectIDFromHex(productID)
		pds = append(pds, &domain.ProductPriorityDAO{
			ProductID: productObjID,
			Priority:  idx,
		})
	}
	pgDao.Products = pds
	newPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
	if err != nil {
		return nil, err
	}

	return &grpcServer.PushProductsResponse{Pg: mapper.ProductGroupMapper(newPgDao)}, nil
}
