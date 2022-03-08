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
	layout := "2006-01-02T15:04:05Z07:00"

	startTimeObj, _ := time.Parse(layout, req.StartTime)
	finishTimeObj, _ := time.Parse(layout, req.FinishTime)

	groupType := domain.PRODUCT_GROUP_TIMEDEAL
	if req.GroupType == grpcServer.ProductGroupType_PRODUCT_GROUP_EXHIBITION {
		groupType = domain.PRODUCT_GROUP_EXHIBITION
	}

	shortTitle := ""
	if *req.ShortTitle != "" {
		shortTitle = *req.ShortTitle
	}
	imageUrl := ""
	if *req.ImageUrl != "" {
		imageUrl = *req.ImageUrl
	}

	pgDao := &domain.ProductGroupDAO{
		Title:       req.Title,
		ShortTitle:  shortTitle,
		Instruction: req.Instruction,
		ImgUrl:      imageUrl,
		Products:    []*domain.ProductPriorityDAO{},
		StartTime:   startTimeObj,
		FinishTime:  finishTimeObj,
		GroupType:   domain.ProductGroupType(groupType),
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
	if req.Query == nil || req.Query.GroupType == nil {
		numPassedPgsToShow := 10000 // Dev code 임의로 10000개 잡아둠
		pgDaos, err := ioc.Repo.ProductGroups.List(numPassedPgsToShow)
		if err != nil {
			return nil, err
		}
		pgs := []*grpcServer.ProductGroupMessage{}
		for _, pgDao := range pgDaos {
			pgs = append(pgs, mapper.ProductGroupMapper(pgDao))
		}
		return &grpcServer.ListProductGroupsResponse{
			Pgs:         pgs,
			Offset:      req.Offset,
			Limit:       req.Limit,
			TotalCounts: 0,
		}, nil
	}

	if *req.Query.GroupType.Enum() == grpcServer.ProductGroupType_PRODUCT_GROUP_TIMEDEAL {
		pgDaos, err := ioc.Repo.ProductGroups.ListTimedeals(int(req.Offset), int(req.Limit), false)
		if err != nil {
			return nil, err
		}
		pgs := []*grpcServer.ProductGroupMessage{}
		for _, pgDao := range pgDaos {
			pgs = append(pgs, mapper.ProductGroupMapper(pgDao))
		}
		return &grpcServer.ListProductGroupsResponse{
			Pgs:         pgs,
			Offset:      req.Offset,
			Limit:       req.Limit,
			TotalCounts: 0,
		}, nil
	} else if *req.Query.GroupType.Enum() == grpcServer.ProductGroupType_PRODUCT_GROUP_EXHIBITION {
		pgDaos, cnt, err := ioc.Repo.ProductGroups.ListExhibitionPg(int(req.Offset), int(req.Limit))
		if err != nil {
			return nil, err
		}
		pgs := []*grpcServer.ProductGroupMessage{}
		for _, pgDao := range pgDaos {
			pgs = append(pgs, mapper.ProductGroupMapper(pgDao))
		}
		return &grpcServer.ListProductGroupsResponse{
			Pgs:         pgs,
			Offset:      req.Offset,
			Limit:       req.Limit,
			TotalCounts: int32(cnt),
		}, nil
	} else {
		numPassedPgsToShow := 10000 // Dev code 임의로 10000개 잡아둠
		pgDaos, err := ioc.Repo.ProductGroups.List(numPassedPgsToShow)
		if err != nil {
			return nil, err
		}
		pgs := []*grpcServer.ProductGroupMessage{}
		for _, pgDao := range pgDaos {
			pgs = append(pgs, mapper.ProductGroupMapper(pgDao))
		}
		return &grpcServer.ListProductGroupsResponse{
			Pgs:         pgs,
			Offset:      req.Offset,
			Limit:       req.Limit,
			TotalCounts: 0,
		}, nil
	}
}

func (s *ProductGroupService) EditProductGroup(ctx context.Context, req *grpcServer.EditProductGroupRequest) (*grpcServer.EditProductGroupResponse, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}

	layout := "2006-01-02T15:04:05Z07:00"

	if req.Title != nil {
		pgDao.Title = *req.Title
	}

	if req.ShortTitle != nil {
		pgDao.ShortTitle = *req.ShortTitle
	}

	if req.Instruction != nil {
		pgDao.Instruction = req.Instruction
	}

	if req.ImageUrl != nil {
		pgDao.ImgUrl = *req.ImageUrl
	}

	if req.StartTime != nil {
		startTimeObj, _ := time.Parse(layout, *req.StartTime)
		pgDao.StartTime = startTimeObj
	}

	if req.FinishTime != nil {
		finishTimeObj, _ := time.Parse(layout, *req.FinishTime)
		pgDao.FinishTime = finishTimeObj
	}

	if req.GroupType != nil {
		groupType := domain.PRODUCT_GROUP_TIMEDEAL
		if *req.GroupType == grpcServer.ProductGroupType_PRODUCT_GROUP_EXHIBITION {
			groupType = domain.PRODUCT_GROUP_EXHIBITION
		}
		pgDao.GroupType = domain.ProductGroupType(groupType)
	}

	newPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
	if err != nil {
		return nil, err
	}

	updatedPgDao, err := broker.ProductGroupSyncer(newPgDao)
	if err != nil {
		log.Println("product group syncing error", err)
	}
	return &grpcServer.EditProductGroupResponse{Pg: mapper.ProductGroupMapper(updatedPgDao)}, nil
}

func (s *ProductGroupService) PushProductsInProductGroup(ctx context.Context, req *grpcServer.PushProductsInPgRequest) (*grpcServer.PushProductsInPgResponse, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}

	for _, productPriority := range req.ProductPriority {
		productObjId, _ := primitive.ObjectIDFromHex(productPriority.ProductId)
		pdDao, _ := ioc.Repo.Products.Get(productPriority.ProductId)
		isNewProduct := pgDao.AppendProduct(&domain.ProductPriorityDAO{
			Priority:  int(productPriority.Priority),
			ProductID: productObjId,
			Product:   pdDao,
		})

		if isNewProduct {
			pd, err := ioc.Repo.Products.Get(productPriority.ProductId)
			if err != nil {
				return nil, err
			}
			pd.ProductGroupId = pgDao.ID.Hex()
			_, err = ioc.Repo.Products.Upsert(pd)
			if err != nil {
				return nil, err
			}
		}
	}

	newPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
	if err != nil {
		return nil, err
	}

	updatedPgDao, err := broker.ProductGroupSyncer(newPgDao)
	if err != nil {
		log.Println("product group syncing error", err)
	}
	return &grpcServer.PushProductsInPgResponse{Pg: mapper.ProductGroupMapper(updatedPgDao)}, nil
}

func (s *ProductGroupService) RemoveProductInProductGroup(ctx context.Context, req *grpcServer.RemoveProductInPgRequest) (*grpcServer.RemoveProductInPgResponse, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}

	pgDao.RemoveProduct(req.ProductId)
	newPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
	if err != nil {
		return nil, err
	}

	pd, err := ioc.Repo.Products.Get(req.ProductId)
	if err != nil {
		return nil, err
	}
	pd.ProductGroupId = ""
	_, err = ioc.Repo.Products.Upsert(pd)
	if err != nil {
		return nil, err
	}

	updatedPgDao, err := broker.ProductGroupSyncer(newPgDao)
	if err != nil {
		log.Println("product group syncing error", err)
	}

	return &grpcServer.RemoveProductInPgResponse{Pg: mapper.ProductGroupMapper(updatedPgDao)}, nil
}
