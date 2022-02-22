package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
)

func (r *queryResolver) ProductGroup(ctx context.Context, id string) (*model.ProductGroup, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapProductGroupDao(pgDao), nil
}

func (r *queryResolver) ProductGroups(ctx context.Context) ([]*model.ProductGroup, error) {
	numPassedPgsToShow := 10
	pgDaos, err := ioc.Repo.ProductGroups.List(numPassedPgsToShow)
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

	return mapper.MapExhibition(exhibitionDao), nil
}

func (r *queryResolver) Exhibitions(ctx context.Context) ([]*model.Exhibition, error) {
	offset, limit := 0, 100 // IGNORRED SINCE ONLY LIVE
	onlyLive := true
	exhibitionDaos, _, err := ioc.Repo.Exhibitions.List(offset, limit, onlyLive)
	if err != nil {
		return nil, err
	}

	exs := []*model.Exhibition{}

	for _, exhibitionDao := range exhibitionDaos {
		exs = append(exs, mapper.MapExhibition(exhibitionDao))
	}

	return exs, nil
}

func (r *queryResolver) Timedeal(ctx context.Context) (*model.ProductGroup, error) {
	// TO BE SPECIFIED WITH PRODUCTGROUP TYPES
	livePgs, err := ioc.Repo.ProductGroups.ListTimedeals(0, 1, true)
	if err != nil {
		return nil, err
	}

	if len(livePgs) > 0 {
		return mapper.MapProductGroupDao(livePgs[0]), nil
	}
	return nil, nil
}
