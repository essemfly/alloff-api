package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
)

func (r *queryResolver) ProductGroup(ctx context.Context, id string) (*model.ProductGroup, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(id)
	if err != nil {
		return nil, err
	}

	return pgDao.ToDTO(), nil
}

func (r *queryResolver) ProductGroups(ctx context.Context) ([]*model.ProductGroup, error) {
	pgDaos, err := ioc.Repo.ProductGroups.List()
	if err != nil {
		return nil, err
	}

	pgs := []*model.ProductGroup{}

	for _, pgDao := range pgDaos {
		pgs = append(pgs, pgDao.ToDTO())
	}

	return pgs, nil
}

func (r *queryResolver) Exhibition(ctx context.Context, id string) (*model.Exhibition, error) {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
	if err != nil {
		return nil, err
	}

	return exhibitionDao.ToDTO(), nil
}

func (r *queryResolver) Exhibitions(ctx context.Context) ([]*model.Exhibition, error) {
	exhibitionDaos, err := ioc.Repo.Exhibitions.List()
	if err != nil {
		return nil, err
	}

	exs := []*model.Exhibition{}

	for _, exexhibitionDao := range exhibitionDaos {
		exs = append(exs, exexhibitionDao.ToDTO())
	}

	return exs, nil
}
