package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
)

func (r *queryResolver) Featureds(ctx context.Context) ([]*model.FeaturedItem, error) {
	featuredDaos, err := ioc.Repo.Featureds.List()
	if err != nil {
		return nil, err
	}

	items := []*model.FeaturedItem{}
	for _, itemDao := range featuredDaos {
		items = append(items, mapper.MapFeatured(itemDao))
	}

	return items, nil
}

func (r *queryResolver) Homeitems(ctx context.Context) ([]*model.HomeItem, error) {
	homeitemDaos, err := ioc.Repo.HomeItems.List()
	if err != nil {
		return nil, err
	}

	items := []*model.HomeItem{}
	for _, itemDao := range homeitemDaos {
		items = append(items, mapper.MapHomeitem(itemDao))
	}

	return items, nil
}
