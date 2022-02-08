package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
)

func (r *queryResolver) HomeTabItems(ctx context.Context) ([]*model.HomeTabItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BestProducts(ctx context.Context, offset int, limit int, alloffCategoryID string, brief bool) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BestBrands(ctx context.Context, offset int, limit int) ([]*model.Brand, error) {
	panic(fmt.Errorf("not implemented"))
}
