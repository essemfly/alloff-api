package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
)

func (r *mutationResolver) CreateGroup(ctx context.Context, exhibitionID string) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) JoinGroup(ctx context.Context, exhibitionID string, groupID string, requestLink string) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddMockGroupdeal(ctx context.Context, input *model.AddMockGroupdealInput) (*model.Exhibition, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AddMockGroup(ctx context.Context, exhibitionID string, isCompleted bool) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PushMockUserToGroup(ctx context.Context, groupID string) (*model.Group, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Mygroupdeal(ctx context.Context) (*model.MyGroupDeal, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Mygroupdeals(ctx context.Context, status model.GroupdealStatus) ([]*model.Exhibition, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Groupdeal(ctx context.Context, id string) (*model.Exhibition, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Groupdeals(ctx context.Context, offset int, limit int, status model.GroupdealStatus) ([]*model.Exhibition, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) CheckTicket(ctx context.Context, exhibitionID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
