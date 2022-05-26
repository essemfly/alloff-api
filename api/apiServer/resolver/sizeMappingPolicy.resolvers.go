package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/config/ioc"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
)

func (r *queryResolver) SizeMappingPolicies(ctx context.Context) ([]*model.SizeMappingPolicy, error) {
	res := []*model.SizeMappingPolicy{}
	policies, err := ioc.Repo.SizeMappingPolicy.List()
	if err != nil {
		return nil, err
	}

	for _, policy := range policies {
		res = append(res, mapper.MapSizeMappingPolicy(policy))
	}

	return res, nil
}
