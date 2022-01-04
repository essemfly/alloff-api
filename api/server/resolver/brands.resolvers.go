package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/lessbutter/alloff-api/api/middleware"
	"github.com/lessbutter/alloff-api/api/server/model"
	"github.com/lessbutter/alloff-api/config/ioc"
)

func (r *mutationResolver) LikeBrand(ctx context.Context, input *model.LikeBrandInput) (bool, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return false, errors.New("invalid token")
	}

	return ioc.Repo.LikeBrands.Like(user.ID.Hex(), input.BrandID)
}

func (r *queryResolver) Brand(ctx context.Context, input *model.BrandInput) (*model.Brand, error) {
	brandDao, err := ioc.Repo.Brands.Get(input.BrandID)
	includeCategory := true
	if err != nil {
		return nil, err
	}

	return brandDao.ToDTO(includeCategory), nil
}

func (r *queryResolver) Brands(ctx context.Context, input *model.BrandsInput) ([]*model.Brand, error) {
	if input.OnlyLikes != nil {
		if *input.OnlyLikes {
			user := middleware.ForContext(ctx)
			if user == nil {
				return nil, errors.New("invalid token")
			}

			likeDao, err := ioc.Repo.LikeBrands.List(user.ID.Hex())
			if err != nil {
				return nil, err
			}

			likeBrands := []*model.Brand{}
			for _, likebrand := range likeDao.Brands {
				likeBrands = append(likeBrands, likebrand.ToDTO(false))
			}

			return likeBrands, nil
		}
	}

	// Temp code for limits
	offset, limit := 0, 1000
	brandDaos, _, err := ioc.Repo.Brands.List(offset, limit, nil, nil)
	if err != nil {
		return nil, err
	}

	brands := []*model.Brand{}
	for _, brand := range brandDaos {
		brands = append(brands, brand.ToDTO(false))
	}

	return brands, nil
}
