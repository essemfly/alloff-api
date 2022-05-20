package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	exhibitionService "github.com/lessbutter/alloff-api/pkg/exhibition"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateGroup(ctx context.Context, exhibitionID string) (*model.Group, error) {
	var mutex = &sync.Mutex{}
	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	userID := userDao.ID.Hex()
	params := domain.GroupRequestParams{
		UserID:       &userID,
		ExhibitionID: &exhibitionID,
	}

	userGroupRequest, _ := ioc.Repo.GroupRequest.Get(params)
	if userGroupRequest != nil {
		return nil, fmt.Errorf("ERR502:user already has group")
	}

	exhibition, err := ioc.Repo.Exhibitions.Get(exhibitionID)
	if err != nil {
		return nil, err
	}

	if time.Now().After(exhibition.StartTime) {
		return nil, fmt.Errorf("ERR500:expired create group time")
	}

	mutex.Lock()
	newGroup := domain.GroupDAO{
		ID:               primitive.NewObjectID(),
		ExhibitionID:     exhibitionID,
		NumUsersRequired: exhibition.NumUsersRequired,
		Users:            []*domain.UserDAO{userDao},
	}
	newGroupDao, err := ioc.Repo.Groups.Insert(&newGroup)
	if err != nil {
		return nil, err
	}

	newRequest := &domain.GroupRequestDAO{
		ID:           primitive.NewObjectID(),
		UserID:       userDao.ID.Hex(),
		ExhibitionID: exhibitionID,
		GroupID:      newGroup.ID.Hex(),
		RequestLink:  "",
		Status:       domain.GroupRequestStatusSuccess,
	}
	_, err = ioc.Repo.GroupRequest.Insert(newRequest)
	if err != nil {
		return nil, err
	}

	exhibition.TotalGroups += 1
	_, err = ioc.Repo.Exhibitions.Upsert(exhibition)
	if err != nil {
		return nil, err
	}
	mutex.Unlock()

	return mapper.MapGroupDaoToGroup(newGroupDao), nil
}

func (r *mutationResolver) JoinGroup(ctx context.Context, exhibitionID string, groupID string, requestLink string) (*model.Group, error) {
	var mutex = &sync.Mutex{}
	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	userID := userDao.ID.Hex()
	params := domain.GroupRequestParams{
		UserID:       &userID,
		ExhibitionID: &exhibitionID,
	}

	userGroupRequest, _ := ioc.Repo.GroupRequest.Get(params)
	if userGroupRequest != nil {
		return nil, fmt.Errorf("ERR502:user already has group")
	}

	mutex.Lock()
	groupDao, err := ioc.Repo.Groups.Get(groupID)
	if err != nil {
		return nil, err
	}

	exhibition, err := ioc.Repo.Exhibitions.Get(groupDao.ExhibitionID)
	if err != nil {
		return nil, err
	}

	if time.Now().After(exhibition.StartTime) {
		return nil, fmt.Errorf("ERR504:expired join group time")
	}

	if exhibition.RecruitStartTime.After(userDao.Created) && !exhibition.AllowOldUser {
		return nil, fmt.Errorf("ERR503:not new user")
	}

	if len(groupDao.Users) < exhibition.NumUsersRequired {
		groupDao.Users = append(groupDao.Users, userDao)
		_, err = ioc.Repo.Groups.Update(groupDao)
		if err != nil {
			return nil, err
		}
		newRequest := &domain.GroupRequestDAO{
			ID:           primitive.NewObjectID(),
			UserID:       userDao.ID.Hex(),
			ExhibitionID: groupDao.ExhibitionID,
			GroupID:      groupID,
			RequestLink:  requestLink,
			Status:       domain.GroupRequestStatusSuccess,
		}
		_, err = ioc.Repo.GroupRequest.Insert(newRequest)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("ERR501:group already completed")
	}
	mutex.Unlock()
	exhibitionService.UpdateGroupdealInfo(groupDao, exhibition)
	return mapper.MapGroupDaoToGroup(groupDao), nil
}

// func (r *mutationResolver) AddMockGroupdeal(ctx context.Context, input *model.AddMockGroupdealInput) (*model.Exhibition, error) {
// for prevent query to prod env
// cmd.SetBaseConfig("dev")

// 	recruitStartTime := time.Now().Round(time.Second)
// 	startTime := time.Now().Round(time.Second)
// 	finishTime := time.Now().Round(time.Second)

// 	switch input.GroupdealStatus {
// 	case model.GroupdealStatusPending:
// 		startTime = startTime.Add(72 * time.Hour)
// 		finishTime = startTime.Add(72 * time.Hour)
// 	case model.GroupdealStatusOpen:
// 		recruitStartTime = recruitStartTime.Add(-72 * time.Hour)
// 		finishTime = finishTime.Add(72 * time.Hour)
// 	case model.GroupdealStatusClosed:
// 		recruitStartTime = recruitStartTime.Add(-144 * time.Hour)
// 		startTime = startTime.Add(-72 * time.Hour)
// 	}

// 	newExhibition := &domain.ExhibitionDAO{
// 		ID:               primitive.NewObjectID(),
// 		BannerImage:      input.BannerImage,
// 		ThumbnailImage:   input.ThumbnailImage,
// 		Title:            input.Title,
// 		SubTitle:         input.Subtitle,
// 		Description:      input.Description,
// 		StartTime:        startTime,
// 		FinishTime:       finishTime,
// 		RecruitStartTime: recruitStartTime,
// 		IsLive:           true,
// 		CreatedAt:        time.Now(),
// 		UpdatedAt:        time.Now(),
// 		ExhibitionType:   domain.EXHIBITION_GROUPDEAL,
// 		NumUsersRequired: input.NumUsersRequired,
// 		AllowOldUser:     input.AllowOldUser,
// 	}

// 	newExhibition.ProductGroups = scripts.AddGroupdealProductGroups(newExhibition)
// 	_, err := ioc.Repo.Exhibitions.Upsert(newExhibition)
// 	if err != nil {
// 		return nil, err
// 	}
// 	exhibitionService.UpdateCheapestPrice(newExhibition)
// 	exhibitionService.UpdateGroupdealInfo(&domain.GroupDAO{}, newExhibition)
// 	return mapper.MapExhibition(newExhibition, false), nil
// }

// func (r *mutationResolver) AddMockGroup(ctx context.Context, exhibitionID string, isCompleted bool) (*model.Group, error) {
// 	// for prevent query to prod env
// 	cmd.SetBaseConfig("dev")

// 	exhibition, err := ioc.Repo.Exhibitions.Get(exhibitionID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	exhibition.TotalGroups += 1
// 	_, err = ioc.Repo.Exhibitions.Upsert(exhibition)
// 	if err != nil {
// 		log.Println("error occurred on update total groups : ", err)
// 		return nil, err
// 	}

// 	group, err := scripts.AddMockGroup(exhibition, isCompleted)
// 	if err != nil {
// 		return nil, err
// 	}

// 	exhibitionService.UpdateGroupdealInfo(group, exhibition)
// 	return mapper.MapGroupDaoToGroup(group), nil
// }

// func (r *mutationResolver) PushMockUserToGroup(ctx context.Context, groupID string) (*model.Group, error) {
// 	// for prevent query to prod env
// 	cmd.SetBaseConfig("dev")

// 	groupDao, err := ioc.Repo.Groups.Get(groupID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	exhibition, err := ioc.Repo.Exhibitions.Get(groupDao.ExhibitionID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	groupDao, err = scripts.PushMockUserToGroup(groupDao)
// 	if err != nil {
// 		return nil, err
// 	}
// 	exhibitionService.UpdateGroupdealInfo(groupDao, exhibition)
// 	return mapper.MapGroupDaoToGroup(groupDao), nil
// }

// func (r *queryResolver) Mygroupdeal(ctx context.Context) (*model.MyGroupDeal, error) {
// 	user := middleware.ForContext(ctx)
// 	if user == nil {
// 		return nil, fmt.Errorf("ERR000:invalid token")
// 	}

// 	userid := user.ID.Hex()
// 	params := domain.GroupRequestParams{
// 		UserID: &userid,
// 	}
// 	userGroupRequests, err := ioc.Repo.GroupRequest.List(params, []domain.GroupRequestStatus{})
// 	if err != nil {
// 		log.Println("error on get list of group requests")
// 		return nil, err
// 	}
// 	userTickets, err := ioc.Repo.GroupdealTickets.ListByDetail(user.ID.Hex(), "")
// 	if err != nil {
// 		log.Println("error on get list of tickets")
// 		return nil, err
// 	}

// 	offset, limit := 0, 1000
// 	onlyLive := true

// 	liveExhibitions, _, err := ioc.Repo.Exhibitions.ListGroupDeals(offset, limit, onlyLive, domain.GROUPDEAL_OPEN)
// 	if err != nil {
// 		return nil, err
// 	}
// 	liveGroupdealCnt := exhibitionService.GetUserPurchasableGroupCount(liveExhibitions, userTickets)

// 	pendingExhibitions, _, err := ioc.Repo.Exhibitions.ListGroupDeals(offset, limit, onlyLive, domain.GROUPDEAL_PENDING)
// 	if err != nil {
// 		return nil, err
// 	}
// 	pendingGroupdealCnt := exhibitionService.GetUserPendingGroupCount(pendingExhibitions, userGroupRequests)

// 	return &model.MyGroupDeal{
// 		User:              mapper.MapUserDaoToUser(user),
// 		NumParticipates:   pendingGroupdealCnt,
// 		NumLiveGroupdeals: liveGroupdealCnt,
// 	}, nil
// }

// func (r *queryResolver) Mygroupdeals(ctx context.Context, status model.GroupdealStatus) ([]*model.Exhibition, error) {
// 	userDao := middleware.ForContext(ctx)
// 	if userDao == nil {
// 		return nil, fmt.Errorf("ERR000:invalid token")
// 	}

// 	userGroupsDao, err := ioc.Repo.Groups.ListByUserId(userDao.ID.Hex())
// 	if err != nil {
// 		return nil, err
// 	}

// 	exhibitions := []*model.Exhibition{}
// 	for _, userGroupDao := range userGroupsDao {
// 		exhibitionDao, err := ioc.Repo.Exhibitions.Get(userGroupDao.ExhibitionID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		exhibitionModel := mapper.MapExhibition(exhibitionDao, true)
// 		exhibitionModel.UserGroup = mapper.MapGroupDaoToUserGroup(userGroupDao, userDao)
// 		exhibitions = append(exhibitions, exhibitionModel)
// 	}
// 	return exhibitions, nil
// }

// func (r *queryResolver) Groupdeal(ctx context.Context, id string) (*model.Exhibition, error) {
// 	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	exhibition := mapper.MapExhibition(exhibitionDao, false)

// 	userDao := middleware.ForContext(ctx)
// 	if userDao == nil {
// 		return exhibition, nil
// 	}

// 	userGroup := &model.UserGroup{}
// 	userGroupDao, err := ioc.Repo.Groups.GetByDetail(userDao.ID.Hex(), id)
// 	if err != nil {
// 		userGroup = nil
// 	}
// 	userGroup = mapper.MapGroupDaoToUserGroup(userGroupDao, userDao)

// 	latestPurchase := []*model.OrderItem{}
// 	orderDaos, err := ioc.Repo.OrderItems.ListByExhibitionID(id)
// 	if err != nil {
// 		latestPurchase = nil
// 	}
// 	for _, orderDao := range orderDaos {
// 		latestPurchase = append(latestPurchase, mapper.MapOrderItem(orderDao))
// 	}

// 	totalTickets, err := ioc.Repo.GroupdealTickets.ListByDetail("", id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	exhibition.UserGroup = userGroup
// 	exhibition.LatestPurchase = latestPurchase
// 	exhibition.TotalGroupdealTickets = len(totalTickets)

// 	return exhibition, nil
// }

// func (r *queryResolver) Groupdeals(ctx context.Context, offset int, limit int, status model.GroupdealStatus) ([]*model.Exhibition, error) {
// 	onlyLive := true

// 	groupDealStatus := domain.GROUPDEAL_CLOSED
// 	if status == model.GroupdealStatusOpen {
// 		groupDealStatus = domain.GROUPDEAL_OPEN
// 	} else if status == model.GroupdealStatusPending {
// 		groupDealStatus = domain.GROUPDEAL_PENDING
// 	}
// 	exhibitionDaos, _, err := ioc.Repo.Exhibitions.ListGroupDeals(offset, limit, onlyLive, groupDealStatus)
// 	if err != nil {
// 		return nil, err
// 	}

// 	userDao := middleware.ForContext(ctx)

// 	exhibitions := []*model.Exhibition{}
// 	for _, exhibitionDao := range exhibitionDaos {
// 		userGroup := &model.UserGroup{}
// 		if userDao != nil {
// 			userGroupDao, _ := ioc.Repo.Groups.GetByDetail(userDao.ID.Hex(), exhibitionDao.ID.Hex())
// 			if userGroupDao != nil {
// 				userGroup = mapper.MapGroupDaoToUserGroup(userGroupDao, userDao)
// 			}
// 		}
// 		exhibitionModel := mapper.MapExhibition(exhibitionDao, true)
// 		exhibitionModel.UserGroup = userGroup
// 		exhibitions = append(exhibitions, exhibitionModel)
// 	}
// 	return exhibitions, nil
// }

// func (r *queryResolver) CheckTicket(ctx context.Context, exhibitionID string) (bool, error) {
// 	userDao := middleware.ForContext(ctx)
// 	if userDao == nil {
// 		return false, nil
// 	}

// 	_, err := ioc.Repo.GroupdealTickets.GetByDetail(userDao.ID.Hex(), exhibitionID)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }
