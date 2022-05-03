package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *queryResolver) Mygroupdeal(ctx context.Context) (*model.MyGroupDeal, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	offset, limit := 0, 1000
	onlyLive := true
	_, cnt, err := ioc.Repo.Exhibitions.ListGroupDeals(offset, limit, onlyLive, domain.GROUPDEAL_OPEN)
	if err != nil {
		return nil, err
	}

	return &model.MyGroupDeal{
		User:              mapper.MapUserDaoToUser(user),
		NumParticipates:   99,
		NumLiveGroupdeals: cnt,
	}, nil
}

func (r *queryResolver) Mygroupdeals(ctx context.Context, status model.GroupdealStatus) ([]*model.Exhibition, error) {
	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	offset, limit := 0, 1000
	onlyLive := true

	groupDealStatus := domain.GROUPDEAL_CLOSED
	if status == model.GroupdealStatusOpen {
		groupDealStatus = domain.GROUPDEAL_OPEN
	} else if status == model.GroupdealStatusPending {
		groupDealStatus = domain.GROUPDEAL_PENDING
	}

	exhibitionDaos, _, err := ioc.Repo.Exhibitions.ListGroupDeals(offset, limit, onlyLive, groupDealStatus)
	if err != nil {
		return nil, err
	}

	exhibitions := []*model.Exhibition{}
	for _, exhibitionDao := range exhibitionDaos {
		exhibitions = append(exhibitions, mapper.MapExhibition(exhibitionDao, true))
	}

	// TODO: dev code for giving group users
	user := mapper.MapUserDaoToUser(userDao)
	userGroup := model.UserGroup{
		GroupID: primitive.NewObjectID().Hex(),
		Users:   []*model.User{user},
	}
	for _, exhibition := range exhibitions {
		exhibition.UserGroup = &userGroup
	}

	return exhibitions, nil
}

func (r *queryResolver) Groupdeal(ctx context.Context, id string) (*model.Exhibition, error) {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
	if err != nil {
		return nil, err
	}

	exhibition := mapper.MapExhibition(exhibitionDao, false)

	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return exhibition, nil
	}

	// TODO: dev code for giving group users
	user := mapper.MapUserDaoToUser(userDao)
	userGroup := model.UserGroup{
		GroupID: primitive.NewObjectID().Hex(),
		Users:   []*model.User{user},
	}
	exhibition.UserGroup = &userGroup
	return exhibition, nil
}

func (r *queryResolver) Groupdeals(ctx context.Context, offset int, limit int, status model.GroupdealStatus) ([]*model.Exhibition, error) {
	onlyLive := true

	groupDealStatus := domain.GROUPDEAL_CLOSED
	if status == model.GroupdealStatusOpen {
		groupDealStatus = domain.GROUPDEAL_OPEN
	} else if status == model.GroupdealStatusPending {
		groupDealStatus = domain.GROUPDEAL_PENDING
	}
	exhibitionDaos, _, err := ioc.Repo.Exhibitions.ListGroupDeals(offset, limit, onlyLive, groupDealStatus)
	if err != nil {
		return nil, err
	}

	exhibitions := []*model.Exhibition{}
	for _, exhibitionDao := range exhibitionDaos {
		exhibitions = append(exhibitions, mapper.MapExhibition(exhibitionDao, true))
	}

	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return exhibitions, nil
	}

	// TODO: dev code for giving group users
	user := mapper.MapUserDaoToUser(userDao)
	userGroup := model.UserGroup{
		GroupID: primitive.NewObjectID().Hex(),
		Users:   []*model.User{user},
	}
	for _, exhibition := range exhibitions {
		exhibition.UserGroup = &userGroup
	}

	return exhibitions, nil
}

func (r *queryResolver) CreateGroup(ctx context.Context, exhibitionID string) (*model.Group, error) {
	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	// TODO: CheckWhetherCanCreateGroupOrNot
	newGroup := domain.GroupDAO{
		ID:           primitive.NewObjectID(),
		ExhibitionID: exhibitionID,
		Users:        []*domain.UserDAO{userDao},
	}
	newGroupDao, err := ioc.Repo.Groups.Insert(&newGroup)
	if err != nil {
		return nil, err
	}

	return mapper.MapGroupDaoToGroup(newGroupDao), nil
}

func (r *queryResolver) JoinGroup(ctx context.Context, groupID string) (*model.Group, error) {
	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	// TODO: CheckWhetherCanJoinGroupOrNot
	groupDao, err := ioc.Repo.Groups.Get(groupID)
	if err != nil {
		return nil, err
	}

	groupDao.Users = append(groupDao.Users, userDao)
	// TODO: Add Users for the groupid

	return mapper.MapGroupDaoToGroup(groupDao), nil

}
