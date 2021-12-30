package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	server1 "github.com/lessbutter/alloff-api/api/server"
	server "github.com/lessbutter/alloff-api/api/server/model"
)

func (r *queryResolver) Alloffcategories(ctx context.Context, input *server.AlloffCategoryInput) ([]*server.AlloffCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Alloffcategory(ctx context.Context, input *server.AlloffCategoryID) (*server.AlloffCategory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AlloffcategoryProducts(ctx context.Context, input server.CategoryProductsInput) (*server.AlloffCategoryProducts, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns server1.QueryResolver implementation.
func (r *Resolver) Query() server1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
