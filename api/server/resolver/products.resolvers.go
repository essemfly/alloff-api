package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	server "github.com/lessbutter/alloff-api/api/server/model"
)

func (r *mutationResolver) LikeProduct(ctx context.Context, input *server.LikeProductInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id string) (*server.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, input server.ProductsInput) (*server.ProductsOutput, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Likeproducts(ctx context.Context) ([]*server.LikeProductOutput, error) {
	panic(fmt.Errorf("not implemented"))
}
