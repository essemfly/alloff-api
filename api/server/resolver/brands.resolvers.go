package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/server/model"
)

func (r *mutationResolver) LikeBrand(ctx context.Context, input *model.LikeBrandInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Brand(ctx context.Context, input *model.BrandInput) (*model.Brand, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Brands(ctx context.Context, input *model.BrandsInput) ([]*model.Brand, error) {
	panic(fmt.Errorf("not implemented"))
}
