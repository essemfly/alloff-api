package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/alimtalk"
	"github.com/lessbutter/alloff-api/pkg/product"
	"go.uber.org/zap"
)

func (r *mutationResolver) SetAlarm(ctx context.Context, id string) (bool, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("ERR000:invalid token")
	}

	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
	if err != nil {
		return false, err
	}

	alimtalkRegisterd, err := alimtalk.ChangeExhibitionNotifyStatus(user, exhibitionDao)
	if err != nil {
		return false, err
	}

	// 기존의 알림톡을 취소하는 경우
	if alimtalkRegisterd == nil {
		return false, nil
	}
	return true, nil
}

func (r *queryResolver) Exhibition(ctx context.Context, id string) (*model.Exhibition, error) {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapExhibition(exhibitionDao, false), nil
}

func (r *queryResolver) Exhibitions(ctx context.Context, input model.ExhibitionInput) (*model.ExhibitionOutput, error) {
	offset, limit := 0, 100
	query := ""

	var exhibitionStatus domain.ExhibitionStatus
	switch input.Status {
	case model.ExhibitionStatusNotOpen:
		exhibitionStatus = domain.EXHIBITION_NOTOPEN
	case model.ExhibitionStatusLive:
		exhibitionStatus = domain.EXHIBITION_LIVE
	case model.ExhibitionStatusClosed:
		exhibitionStatus = domain.EXHIBITION_CLOSED
	}

	// only live
	liveDaos, liveCnt, err := ioc.Repo.Exhibitions.List(offset, limit, false, domain.EXHIBITION_LIVE, domain.EXHIBITION_TIMEDEAL, query)
	if err != nil {
		return nil, err
	}
	// not open
	notOpenDaos, notOpenCnt, err := ioc.Repo.Exhibitions.List(offset, limit, false, domain.EXHIBITION_NOTOPEN, domain.EXHIBITION_TIMEDEAL, query)
	if err != nil {
		return nil, err
	}

	exs := []*model.Exhibition{}

	if exhibitionStatus == domain.EXHIBITION_LIVE {
		for _, exhibitionDao := range liveDaos {
			exs = append(exs, mapper.MapExhibition(exhibitionDao, true))
		}
	} else {
		for _, exhibitionDao := range notOpenDaos {
			exs = append(exs, mapper.MapExhibition(exhibitionDao, true))
		}
	}

	return &model.ExhibitionOutput{
		Exhibitions:   exs,
		Status:        input.Status,
		LiveCounts:    liveCnt,
		NotOpenCounts: notOpenCnt,
	}, nil
}

func (r *queryResolver) ExhibitionInfo(ctx context.Context, input model.MetaInfoInput) (*model.MetaInfoOutput, error) {
	pdType := domain.Female
	if input.ProductType == model.AlloffProductTypeKids {
		pdType = domain.Kids
	} else if input.ProductType == model.AlloffProductTypeMale {
		pdType = domain.Male
	} else if input.ProductType == model.AlloffProductTypeSports {
		pdType = domain.Sports
	}

	brandIds := []string{}
	if len(input.BrandIds) > 0 {
		brandIds = input.BrandIds
	}
	alloffcatID := ""
	if input.AlloffCategoryID != nil {
		alloffcatID = *input.AlloffCategoryID
	}

	query := product.ProductListInput{
		ProductType:      pdType,
		ExhibitionID:     input.ExhibitionID,
		BrandIDs:         brandIds,
		AlloffCategoryID: alloffcatID,
	}

	filter, err := query.BuildFilter()
	if err != nil {
		config.Logger.Error("build filter err on exhibition info", zap.Error(err))
	}

	//brandData, alloffcatData, sizeData := ioc.Repo.Products.ListDistinctInfos(filter)
	brandData, alloffcatData, sizeData := ioc.Repo.Products.ListInfos(filter)

	brandOutputs := []*model.BrandOutput{}
	for _, brandCount := range brandData {
		brandOutputs = append(brandOutputs, &model.BrandOutput{
			Brand:       mapper.MapBrandDaoToBrand(brandCount.Brand, false),
			NumProducts: brandCount.Counts,
		})
	}
	alloffcats := []*model.AlloffCategory{}
	for _, cat := range alloffcatData {
		alloffcats = append(alloffcats, mapper.MapAlloffCatDaoToAlloffCat(cat))
	}
	sizes := []*model.AlloffSize{}
	for _, size := range sizeData {
		sizes = append(sizes, mapper.MapAlloffSizeDaoToAlloffSize(size))
	}

	return &model.MetaInfoOutput{
		Brands:           brandOutputs,
		AlloffCategories: alloffcats,
		AlloffSizes:      sizes,
	}, nil
}
