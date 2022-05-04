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
	"github.com/lessbutter/alloff-api/pkg/exhibition"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"sync"
	"time"
)

// TODO 10195번 페이지 영우님에게 확인 내가 신청한것만 뜨는건지, 성공한것만 뜨는건지 (성공한것만 뜨는거면 입장권 쿼리 사용해서 최적화)
func (r *queryResolver) Mygroupdeal(ctx context.Context) (*model.MyGroupDeal, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	userid := user.ID.Hex()
	params := domain.GroupRequestListParams{
		UserID: &userid,
	}
	userGroupRequests, err := ioc.Repo.GroupRequest.List(params, []domain.GroupRequestStatus{})
	if err != nil {
		log.Println("error on get list of group requests")
		return nil, err
	}

	offset, limit := 0, 1000
	onlyLive := true

	liveExhibitions, _, err := ioc.Repo.Exhibitions.ListGroupDeals(offset, limit, onlyLive, domain.GROUPDEAL_OPEN)
	if err != nil {
		return nil, err
	}
	liveGroupdealCnt := exhibition.GetUserParticipatesCount(liveExhibitions, userGroupRequests)

	pendingExhibitions, _, err := ioc.Repo.Exhibitions.ListGroupDeals(offset, limit, onlyLive, domain.GROUPDEAL_PENDING)
	if err != nil {
		return nil, err
	}
	pendingGroupdealCnt := exhibition.GetUserParticipatesCount(pendingExhibitions, userGroupRequests)

	return &model.MyGroupDeal{
		User:              mapper.MapUserDaoToUser(user),
		NumParticipates:   pendingGroupdealCnt,
		NumLiveGroupdeals: liveGroupdealCnt,
	}, nil
}

func (r *queryResolver) Mygroupdeals(ctx context.Context, status model.GroupdealStatus) ([]*model.Exhibition, error) {
	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	userGroupsDao, err := ioc.Repo.Groups.ListByUserId(userDao.ID.Hex())
	if err != nil {
		return nil, err
	}

	exhibitions := []*model.Exhibition{}
	for _, userGroupDao := range userGroupsDao {
		exhibitionDao, err := ioc.Repo.Exhibitions.Get(userGroupDao.ExhibitionID)
		if err != nil {
			return nil, err
		}
		exhibitionModel := mapper.MapExhibition(exhibitionDao, true)
		exhibitionModel.UserGroup = mapper.MapGroupDaoToUserGroup(userGroupDao)
		exhibitions = append(exhibitions, exhibitionModel)
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

func (r *mutationResolver) CreateGroup(ctx context.Context, exhibitionID string) (*model.Group, error) {
	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	exhibition, err := ioc.Repo.Exhibitions.Get(exhibitionID)
	if err != nil {
		return nil, err
	}

	if time.Now().After(exhibition.FinishTime) {
		return nil, fmt.Errorf("ERR500:expired creategroup time")
	}

	newGroup := domain.GroupDAO{
		ID:           primitive.NewObjectID(),
		ExhibitionID: exhibitionID,
		Users:        []*domain.UserDAO{userDao},
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

	return mapper.MapGroupDaoToGroup(newGroupDao), nil
}

func (r *mutationResolver) JoinGroup(ctx context.Context, groupID string, requestLink string) (*model.Group, error) {
	var mutex = &sync.Mutex{}
	userDao := middleware.ForContext(ctx)
	if userDao == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	userID := userDao.ID.Hex()
	params := domain.GroupRequestListParams{
		UserID:  &userID,
		GroupID: &groupID,
	}

	userGroupRequest, err := ioc.Repo.GroupRequest.List(params, []domain.GroupRequestStatus{})
	if len(userGroupRequest) > 0 {
		return nil, fmt.Errorf("ERR502:user already joined to this group")
	}

	mutex.Lock()
	groupDao, err := ioc.Repo.Groups.Get(groupID)
	if err != nil {
		return nil, err
	}

	// TODO 리퀘스트에 스테이터스 필요 없을지도 ..?
	if len(groupDao.Users) < 5 {
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
	return mapper.MapGroupDaoToGroup(groupDao), nil
}
