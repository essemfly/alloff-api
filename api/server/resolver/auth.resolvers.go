package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	server1 "github.com/lessbutter/alloff-api/api/server"
	server "github.com/lessbutter/alloff-api/api/server/model"
)

func (r *mutationResolver) RegisterNotification(ctx context.Context, deviceID string, allowNotification bool, userID *string) (*server.Device, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input server.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUserInfo(ctx context.Context, input server.UserInfoInput) (*server.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input server.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input server.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context) (*server.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns server1.MutationResolver implementation.
func (r *Resolver) Mutation() server1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
