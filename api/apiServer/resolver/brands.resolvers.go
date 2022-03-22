package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
)

func (r *mutationResolver) LikeBrand(ctx context.Context, input *model.LikeBrandInput) (bool, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("ERR000:invalid token")
	}

	return ioc.Repo.LikeBrands.Like(user.ID.Hex(), input.BrandID)
}

func (r *queryResolver) Brand(ctx context.Context, input *model.BrandInput) (*model.Brand, error) {
	brandDao, err := ioc.Repo.Brands.Get(input.BrandID)
	includeCategory := true
	if err != nil {
		return nil, err
	}

	return mapper.MapBrandDaoToBrand(brandDao, includeCategory), nil
}

func (r *queryResolver) Brands(ctx context.Context, input *model.BrandsInput) ([]*model.Brand, error) {
	if input.OnlyLikes != nil {
		if *input.OnlyLikes {
			user := middleware.ForContext(ctx)
			if user == nil {
				return nil, fmt.Errorf("ERR000:invalid token")
			}

			likeDao, err := ioc.Repo.LikeBrands.List(user.ID.Hex())
			if err != nil {
				return nil, err
			}

			if likeDao == nil {
				return nil, nil
			}

			likeBrands := []*model.Brand{}
			for _, likebrand := range likeDao.Brands {
				if likebrand != nil {
					likeBrands = append(likeBrands, mapper.MapBrandDaoToBrand(likebrand, false))
				}
			}

			return likeBrands, nil
		}
	}

	// Temp code for limits
	offset, limit := 0, 1000
	brandDaos, _, err := ioc.Repo.Brands.List(offset, limit, false, nil)
	if err != nil {
		return nil, err
	}

	brands := []*model.Brand{}
	for _, brand := range brandDaos {
		brands = append(brands, mapper.MapBrandDaoToBrand(brand, false))
	}

	return brands, nil
}
