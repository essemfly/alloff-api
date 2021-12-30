package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	server "github.com/lessbutter/alloff-api/api/server/model"
)

func (r *mutationResolver) AddAlloffProduct(ctx context.Context, input *server.AlloffProductInput) (*server.AlloffProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AlarmProductGroup(ctx context.Context, groupID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Alloffproduct(ctx context.Context, id string) (*server.AlloffProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProductGroup(ctx context.Context, id string) (*server.ProductGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProductGroups(ctx context.Context) ([]*server.ProductGroup, error) {
	panic(fmt.Errorf("not implemented"))
}
