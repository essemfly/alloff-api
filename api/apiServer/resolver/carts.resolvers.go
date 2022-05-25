package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
)

func (r *mutationResolver) AddCartItem(ctx context.Context, input *model.AddCartItemInput) (*model.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RemoveCartItem(ctx context.Context, productID string) (*model.Cart, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCart(ctx context.Context, id string) (*model.Cart, error) {
	if id == "" {
	}
	panic(fmt.Errorf("not implemented"))
}
