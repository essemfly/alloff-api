package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
)

func (r *queryResolver) Exhibition(ctx context.Context, id string) (*model.Exhibition, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Exhibitions(ctx context.Context, input model.ExhibitionsInput) (*model.ExhibitionsOutput, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SetAlarm(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
