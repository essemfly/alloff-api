package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func (r *queryResolver) ProductGroup(ctx context.Context, id string) (*model.ProductGroup, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapProductGroupDao(pgDao), nil
}

// (2022/05/02)이 함수가 현재 Graphql API에서 쓰이는지 궁금함 + 특히 OnlyLive가 필요없지 않나 생각됨
func (r *queryResolver) ProductGroups(ctx context.Context) ([]*model.ProductGroup, error) {
	offset, limit := 0, 100
	keyword := ""
	pgDaos, _, err := ioc.Repo.ProductGroups.List(offset, limit, nil, keyword)
	if err != nil {
		return nil, err
	}

	pgs := []*model.ProductGroup{}

	for _, pgDao := range pgDaos {
		pgs = append(pgs, mapper.MapProductGroupDao(pgDao))
	}

	return pgs, nil
}

func (r *queryResolver) Exhibition(ctx context.Context, id string) (*model.Exhibition, error) {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapExhibition(exhibitionDao, false), nil
}

// 기획전 API
func (r *queryResolver) Exhibitions(ctx context.Context) ([]*model.Exhibition, error) {
	offset, limit := 0, 100 // IGNORRED SINCE ONLY LIVE
	onlyLive := true
	query := ""
	exhibitionDaos, _, err := ioc.Repo.Exhibitions.List(offset, limit, onlyLive, domain.EXHIBITION_NORMAL, query)
	if err != nil {
		return nil, err
	}

	exs := []*model.Exhibition{}

	for _, exhibitionDao := range exhibitionDaos {
		exs = append(exs, mapper.MapExhibition(exhibitionDao, true))
	}

	return exs, nil
}

func (r *queryResolver) Timedeal(ctx context.Context) (*model.Exhibition, error) {
	// For not force update users
	offset, limit := 0, 100
	onlyLive := true
	query := ""
	exhibitionDaos, _, err := ioc.Repo.Exhibitions.List(offset, limit, onlyLive, domain.EXHIBITION_TIMEDEAL, query)
	if err != nil {
		return nil, err
	}
	if len(exhibitionDaos) > 0 {
		return mapper.MapExhibition(exhibitionDaos[0], false), nil
	}
	return nil, nil
}

func (r *queryResolver) Groupdeal(ctx context.Context) (*model.Exhibition, error) {
	// For not force update users
	offset, limit := 0, 100
	onlyLive := true
	query := ""
	exhibitionDaos, _, err := ioc.Repo.Exhibitions.List(offset, limit, onlyLive, domain.EXHIBITION_GROUPDEAL, query)
	if err != nil {
		return nil, err
	}
	if len(exhibitionDaos) > 0 {
		return mapper.MapExhibition(exhibitionDaos[0], false), nil
	}
	return nil, nil
}
