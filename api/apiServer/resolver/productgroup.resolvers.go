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

func (r *queryResolver) ProductGroup(ctx context.Context, id string) (*model.ProductGroup, error) {
	pgDao, err := ioc.Repo.ProductGroups.Get(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapProductGroupDao(pgDao), nil
}

func (r *queryResolver) ProductGroups(ctx context.Context) ([]*model.ProductGroup, error) {
	offset, limit := 0, 100
	keyword := ""
	pgDaos, _, err := ioc.Repo.ProductGroups.List(offset, limit, nil, keyword)
	if err != nil {
		return nil, err
	}

	pgs := []*model.ProductGroup{}

	for _, pgDao := range pgDaos {
		pgs = append(pgs, mapper.MapProductGroupDao(pgDao))
	}

	return pgs, nil
}

func (r *queryResolver) Exhibition(ctx context.Context, id string) (*model.Exhibition, error) {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapExhibition(exhibitionDao, false), nil
}

func (r *queryResolver) Exhibitions(ctx context.Context) ([]*model.Exhibition, error) {
	offset, limit := 0, 100 // IGNORRED SINCE ONLY LIVE
	onlyLive := true
	query := ""
	exhibitionDaos, _, err := ioc.Repo.Exhibitions.List(offset, limit, onlyLive, domain.EXHIBITION_NORMAL, query)
	if err != nil {
		return nil, err
	}

	exs := []*model.Exhibition{}

	for _, exhibitionDao := range exhibitionDaos {
		exs = append(exs, mapper.MapExhibition(exhibitionDao, true))
	}

	return exs, nil
}

func (r *queryResolver) Timedeal(ctx context.Context) (*model.Exhibition, error) {
	// For not force update users
	offset, limit := 0, 100
	onlyLive := true
	query := ""
	exhibitionDaos, _, err := ioc.Repo.Exhibitions.List(offset, limit, onlyLive, domain.EXHIBITION_TIMEDEAL, query)
	if err != nil {
		return nil, err
	}
	if len(exhibitionDaos) > 0 {
		return mapper.MapExhibition(exhibitionDaos[0], false), nil
	}
	return nil, nil
}

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
