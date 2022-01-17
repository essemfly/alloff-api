package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/server/model"
)

func (r *queryResolver) Version(ctx context.Context) (*model.AppVersion, error) {
	return &model.AppVersion{
		LatestVersion: "0.5.5",
		MinVersion:    "0.5.5",
		Message:       nil,
		IsMaintenance: false,
	}, nil
}
