package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
)

func (r *queryResolver) TopBanners(ctx context.Context) ([]*model.TopBanner, error) {
	offset, limit := 0, 100
	onlyLive := true
	bannerDaos, _, err := ioc.Repo.TopBanners.List(offset, limit, onlyLive)
	if err != nil {
		return nil, err
	}

	banners := []*model.TopBanner{}
	for _, bannerDao := range bannerDaos {
		banners = append(banners, mapper.MapTopBanner(bannerDao))
	}

	return banners, nil
}
