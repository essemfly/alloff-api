package services

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/api/grpcServer/mapper"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	grpcServer "github.com/lessbutter/alloff-grpc-protos/gen/goalloff"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type ProductGroupService struct {
	grpcServer.ProductGroupServer
}

func (s *ProductGroupService) GetProductGroup(ctx context.Context, req *grpcServer.GetProductGroupRequest) (*grpcServer.ProductGroupMessage, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}
	productListInput := product.ProductListInput{
		ProductGroupID: pgDao.ID.Hex(),
	}

	pds, _, err := product.ListProducts(productListInput)
	if err != nil {
		config.Logger.Error("error occured on listing products on pg mapper", zap.Error(err))
		return nil, err
	}

	return mapper.ProductGroupMapper(pgDao, pds), nil
}

func (s *ProductGroupService) CreateProductGroup(ctx context.Context, req *grpcServer.CreateProductGroupRequest) (*grpcServer.ProductGroupMessage, error) {
	layout := "2006-01-02T15:04:05Z07:00"

	startTimeObj, _ := time.Parse(layout, req.StartTime)
	finishTimeObj, _ := time.Parse(layout, req.FinishTime)

	groupType := domain.PRODUCT_GROUP_EXHIBITION
	if req.GroupType == grpcServer.ProductGroupType_PRODUCT_GROUP_EXHIBITION {
		groupType = domain.PRODUCT_GROUP_EXHIBITION
	} else if req.GroupType == grpcServer.ProductGroupType_PRODUCT_GROUP_GROUPDEAL {
		groupType = domain.PRODUCT_GROUP_GROUPDEAL
	} else if req.GroupType == grpcServer.ProductGroupType_PRODUCT_GROUP_TIMEDEAL {
		groupType = domain.PRODUCT_GROUP_TIMEDEAL
	} else if req.GroupType == grpcServer.ProductGroupType_PRODUCT_GROUP_BRAND_TIMEDEAL {
		groupType = domain.PRODUCT_GROUP_BRAND_TIMEDEAL
	}

	shortTitle := ""
	if *req.ShortTitle != "" {
		shortTitle = *req.ShortTitle
	}
	imageUrl := ""
	if *req.ImageUrl != "" {
		imageUrl = *req.ImageUrl
	}

	var brand *domain.BrandDAO
	if req.BrandId != nil {
		brandDao, err := ioc.Repo.Brands.Get(*req.BrandId)
		brand = brandDao
		if err != nil {
			config.Logger.Error("Error on get brand for create product group : "+*req.BrandId, zap.Error(err))
			return nil, err
		}
	} else {
		brand = nil
	}

	pgDao := &domain.ProductGroupDAO{
		Title:       req.Title,
		ShortTitle:  shortTitle,
		Instruction: req.Instruction,
		ImgUrl:      imageUrl,
		Products:    []*domain.ProductPriorityDAO{},
		StartTime:   startTimeObj,
		FinishTime:  finishTimeObj,
		GroupType:   groupType,
		Created:     time.Now(),
		Updated:     time.Now(),
		Brand:       brand,
	}

	newPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
	if err != nil {
		return nil, err
	}

	return mapper.ProductGroupMapper(newPgDao, nil), nil
}

func (s *ProductGroupService) ListProductGroups(ctx context.Context, req *grpcServer.ListProductGroupsRequest) (*grpcServer.ListProductGroupsResponse, error) {
	keyword := ""
	if *req.Query.SearchQuery != "" {
		keyword = *req.Query.SearchQuery
	}
	var groupType domain.ProductGroupType
	if req.Query.GroupType != nil {
		if *req.Query.GroupType.Enum() == grpcServer.ProductGroupType_PRODUCT_GROUP_TIMEDEAL {
			groupType = domain.PRODUCT_GROUP_BRAND_TIMEDEAL
		} else if *req.Query.GroupType.Enum() == grpcServer.ProductGroupType_PRODUCT_GROUP_GROUPDEAL {
			groupType = domain.PRODUCT_GROUP_GROUPDEAL
		} else if *req.Query.GroupType.Enum() == grpcServer.ProductGroupType_PRODUCT_GROUP_EXHIBITION {
			groupType = domain.PRODUCT_GROUP_EXHIBITION
		}
	}

	pgDaos, cnt, err := ioc.Repo.ProductGroups.List(int(req.Offset), int(req.Limit), &groupType, keyword)
	if err != nil {
		return nil, err
	}
	pgs := []*grpcServer.ProductGroupMessage{}
	for _, pgDao := range pgDaos {
		pgs = append(pgs, mapper.ProductGroupMapper(pgDao, nil))
	}
	return &grpcServer.ListProductGroupsResponse{
		Pgs:         pgs,
		Offset:      req.Offset,
		Limit:       req.Limit,
		TotalCounts: int32(cnt),
	}, nil

}

func (s *ProductGroupService) EditProductGroup(ctx context.Context, req *grpcServer.EditProductGroupRequest) (*grpcServer.ProductGroupMessage, error) {
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
		} else if *req.GroupType == grpcServer.ProductGroupType_PRODUCT_GROUP_GROUPDEAL {
			groupType = domain.PRODUCT_GROUP_GROUPDEAL
		}
		pgDao.GroupType = groupType
	}

	if req.BrandId != nil {
		brand, err := ioc.Repo.Brands.GetByKeyname(*req.BrandId)
		if err != nil {
			config.Logger.Error("Error on get brand for edit product group : "+*req.BrandId, zap.Error(err))
			return nil, err
		}
		pgDao.Brand = brand
	}

	newPgDao, err := ioc.Repo.ProductGroups.Upsert(pgDao)
	if err != nil {
		return nil, err
	}

	productListInput := product.ProductListInput{
		ProductGroupID: newPgDao.ID.Hex(),
	}

	pds, _, err := product.ListProducts(productListInput)
	if err != nil {
		config.Logger.Error("error occured on listing products on pg mapper", zap.Error(err))
		return nil, err
	}

	return mapper.ProductGroupMapper(newPgDao, pds), nil
}

func (s *ProductGroupService) PushProductsInProductGroup(ctx context.Context, req *grpcServer.ProductsInPgRequest) (*grpcServer.ProductGroupMessage, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}

	for _, productPriority := range req.ProductPriorities {
		pdInfoDao, err := ioc.Repo.ProductMetaInfos.Get(productPriority.ProductId)
		if err != nil {
			config.Logger.Error("err occured in pdinfo not found: "+productPriority.ProductId, zap.Error(err))
			continue
		}

		newPdDao := &domain.ProductDAO{
			ID:                   primitive.NewObjectID(),
			ProductInfo:          pdInfoDao,
			ProductGroupID:       pgDao.ID.Hex(),
			ExhibitionID:         pgDao.ExhibitionID,
			Weight:               int(productPriority.Priority),
			IsNotSale:            false,
			ExhibitionStartTime:  pgDao.StartTime,
			ExhibitionFinishTime: pgDao.FinishTime,
			Created:              time.Now(),
			Updated:              time.Now(),
		}

		_, err = ioc.Repo.Products.Insert(newPdDao)
		if err != nil {
			config.Logger.Error("err occured on products insert on pg : "+productPriority.ProductId, zap.Error(err))
		}
	}
	exDao, err := ioc.Repo.Exhibitions.Get(pgDao.ExhibitionID)
	if err != nil {
		return nil, err
	}
	go exhibition.ExhibitionSyncer(exDao)

	// TODO edit offset & limit should be fixed normally
	productListInput := product.ProductListInput{
		ProductGroupID: pgDao.ID.Hex(),
		Offset:         0,
		Limit:          10000,
	}

	pds, _, err := product.ListProducts(productListInput)
	if err != nil {
		config.Logger.Error("error occured on listing products on pg mapper", zap.Error(err))
		return nil, err
	}

	return mapper.ProductGroupMapper(pgDao, pds), nil
}

func (s *ProductGroupService) UpdateProductsInProductGroup(ctx context.Context, req *grpcServer.ProductsInPgRequest) (*grpcServer.ProductGroupMessage, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}

	for _, productPriority := range req.ProductPriorities {
		pdDao, err := ioc.Repo.Products.GetByMetaID(productPriority.ProductId, pgDao.ExhibitionID)
		if err != nil {
			config.Logger.Error("err occured in pd not found: "+productPriority.ProductId, zap.Error(err))
			continue
		}

		pdDao.Weight = int(productPriority.Priority)
		pdDao.Updated = time.Now()

		_, err = ioc.Repo.Products.Upsert(pdDao)
		if err != nil {
			config.Logger.Error("err occured on products insert on pg : "+productPriority.ProductId, zap.Error(err))
		}
	}

	productListInput := product.ProductListInput{
		ProductGroupID: pgDao.ID.Hex(),
	}

	pdDaos, _, err := product.ListProducts(productListInput)
	if err != nil {
		config.Logger.Error("error occured on listing products on pg mapper", zap.Error(err))
		return nil, err
	}

	return mapper.ProductGroupMapper(pgDao, pdDaos), nil
}

func (s *ProductGroupService) RemoveProductInProductGroup(ctx context.Context, req *grpcServer.RemoveProductInPgRequest) (*grpcServer.ProductGroupMessage, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(req.ProductGroupId)
	if err != nil {
		return nil, err
	}

	pd, err := ioc.Repo.Products.GetByMetaID(req.ProductId, pgDao.ExhibitionID)
	if err != nil {
		return nil, err
	}

	pd.IsNotSale = true
	_, err = ioc.Repo.Products.Upsert(pd)
	if err != nil {
		return nil, err
	}

	productListInput := product.ProductListInput{
		ProductGroupID: pgDao.ID.Hex(),
	}

	pdDaos, _, err := product.ListProducts(productListInput)
	if err != nil {
		config.Logger.Error("error occured on listing products on pg mapper", zap.Error(err))
		return nil, err
	}

	return mapper.ProductGroupMapper(pgDao, pdDaos), nil
}
