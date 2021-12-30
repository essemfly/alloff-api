package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	server "github.com/lessbutter/alloff-api/api/server/model"
)

func (r *queryResolver) Featureds(ctx context.Context) ([]*server.FeaturedItem, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Curations(ctx context.Context) ([]*server.Curation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Homeitems(ctx context.Context) ([]*server.HomeItem, error) {
	panic(fmt.Errorf("not implemented"))
}
