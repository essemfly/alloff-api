package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/server/model"
)

func (r *mutationResolver) AddAlloffProduct(ctx context.Context, input *model.AlloffProductInput) (*model.AlloffProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AlarmProductGroup(ctx context.Context, groupID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Alloffproduct(ctx context.Context, id string) (*model.AlloffProduct, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProductGroup(ctx context.Context, id string) (*model.ProductGroup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProductGroups(ctx context.Context) ([]*model.ProductGroup, error) {
	panic(fmt.Errorf("not implemented"))
}
