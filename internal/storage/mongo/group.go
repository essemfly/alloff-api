package mongo

import (
	"context"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type groupRepo struct {
	col *mongo.Collection
}

type groupRequestRepo struct {
	col *mongo.Collection
}

type groupdealTicketRepo struct {
	col *mongo.Collection
}

func (repo *groupRepo) Insert(groupDao *domain.GroupDAO) (*domain.GroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := repo.col.InsertOne(ctx, groupDao)
	if err != nil {
		return nil, err
	}
	return groupDao, nil
}

func (repo *groupRepo) Get(groupID string) (*domain.GroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	groupObjectId, _ := primitive.ObjectIDFromHex(groupID)
	filter := bson.M{"_id": groupObjectId}

	group := &domain.GroupDAO{}
	if err := repo.col.FindOne(ctx, filter).Decode(group); err != nil {
		return nil, err
	}

	return group, nil
}

func (repo *groupRepo) Update(groupDao *domain.GroupDAO) (*domain.GroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := repo.col.UpdateOne(ctx, bson.M{"_id": groupDao.ID}, bson.M{"$set": &groupDao})
	if err != nil {
		return nil, err
	}

	return groupDao, nil
}

func (repo *groupRepo) GetByDetail(userId, exhibitionId string) (*domain.GroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	objectUserId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{
		"exhibitionid": exhibitionId,
		"users._id":    objectUserId,
	}

	group := &domain.GroupDAO{}
	if err := repo.col.FindOne(ctx, filter).Decode(group); err != nil {
		return nil, err
	}

	return group, nil
}

func (repo *groupRepo) ListByUserId(userId string) ([]*domain.GroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	objectUserId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.M{
		"users._id": objectUserId,
	}

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var groups []*domain.GroupDAO
	err = cursor.All(ctx, &groups)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (repo *groupRequestRepo) Insert(groupRequest *domain.GroupRequestDAO) (*domain.GroupRequestDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	groupRequest.CreatedAt = time.Now()

	_, err := repo.col.InsertOne(ctx, groupRequest)
	if err != nil {
		return nil, err
	}

	return groupRequest, nil
}

func (repo *groupRequestRepo) List(params domain.GroupRequestParams, status []domain.GroupRequestStatus) ([]*domain.GroupRequestDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{}
	if len(status) == 0 {
		if params.UserID != nil {
			filter["userid"] = &params.UserID
		}
		if params.GroupID != nil {
			filter["groupid"] = &params.GroupID
		}
		if params.ExhibitionID != nil {
			filter["exhibitionid"] = &params.ExhibitionID
		}
	} else {
		statusFilter := []interface{}{}
		for _, st := range status {
			statusFilter = append(statusFilter, bson.M{"status": st})
		}
		filter["$or"] = statusFilter
		if params.UserID != nil {
			filter["userid"] = &params.UserID
		}
		if params.GroupID != nil {
			filter["groupid"] = &params.GroupID
		}
		if params.ExhibitionID != nil {
			filter["exhibitionid"] = &params.ExhibitionID
		}
	}

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var groupRequests []*domain.GroupRequestDAO
	err = cursor.All(ctx, &groupRequests)
	if err != nil {
		return nil, err
	}

	return groupRequests, nil
}

func (repo *groupRequestRepo) ListByGroupID(groupID string, status []domain.GroupRequestStatus) ([]*domain.GroupRequestDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{}
	if len(status) == 0 {
		filter = bson.M{
			"groupid": groupID,
		}
	} else {
		statusFilter := []interface{}{}
		for _, st := range status {
			statusFilter = append(statusFilter, bson.M{"status": st})
		}
		filter = bson.M{
			"groupid": groupID,
			"$or":     statusFilter,
		}
	}

	sortingOptions := bson.D{{Key: "createdat", Value: 1}}
	options := options.Find()
	options.SetSort(sortingOptions)

	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	var groupRequests []*domain.GroupRequestDAO
	err = cursor.All(ctx, &groupRequests)
	if err != nil {
		return nil, err
	}

	return groupRequests, nil
}

func (repo *groupRequestRepo) Update(groupRequestDao *domain.GroupRequestDAO) (*domain.GroupRequestDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if _, err := repo.col.UpdateByID(ctx, groupRequestDao.ID, bson.M{"$set": &groupRequestDao}); err != nil {
		return nil, err
	}

	var updatedGroupRequest *domain.GroupRequestDAO
	if err := repo.col.FindOne(ctx, bson.M{"_id": groupRequestDao.ID}).Decode(&updatedGroupRequest); err != nil {
		return nil, err
	}

	return updatedGroupRequest, nil
}

func (repo *groupRequestRepo) Get(params domain.GroupRequestParams) (*domain.GroupRequestDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{}
	if params.UserID != nil {
		filter["userid"] = &params.UserID
	}
	if params.GroupID != nil {
		filter["groupid"] = &params.GroupID
	}
	if params.ExhibitionID != nil {
		filter["exhibitionid"] = &params.ExhibitionID
	}

	groupRequest := &domain.GroupRequestDAO{}
	if err := repo.col.FindOne(ctx, filter).Decode(groupRequest); err != nil {
		return nil, err
	}

	return groupRequest, nil
}

func (repo *groupdealTicketRepo) GetByDetail(exhibitionID, userID string) (*domain.GroupdealTicketDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{
		"exhibitionid": exhibitionID,
		"userid":       userID,
	}

	var groupdealTicket *domain.GroupdealTicketDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&groupdealTicket); err != nil {
		return nil, err
	}
	return groupdealTicket, nil
}

func (repo *groupdealTicketRepo) Insert(groupdealTicketDao *domain.GroupdealTicketDAO) (*domain.GroupdealTicketDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	groupdealTicketDao.CreatedAt = time.Now()

	_, err := repo.col.InsertOne(ctx, groupdealTicketDao)
	if err != nil {
		return nil, err
	}

	return groupdealTicketDao, nil
}

func (repo *groupdealTicketRepo) ListByDetail(userID, exhibitionID string) ([]*domain.GroupdealTicketDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{}
	if userID != "" {
		filter["userid"] = userID
	}
	if exhibitionID != "" {
		filter["exhibitionid"] = exhibitionID
	}

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	groupdealTickets := []*domain.GroupdealTicketDAO{}
	err = cursor.All(ctx, &groupdealTickets)
	if err != nil {
		return nil, err
	}

	return groupdealTickets, nil
}

func MongoGroupsRepo(conn *MongoDB) repository.GroupRepository {
	return &groupRepo{
		col: conn.groupCol,
	}
}

func MongoGroupRequestsRepo(conn *MongoDB) repository.GroupRequestRepository {
	return &groupRequestRepo{
		col: conn.groupRequestCol,
	}
}

func MongoGroupdealTicketRepo(conn *MongoDB) repository.GroupdealTicketRepository {
	return &groupdealTicketRepo{
		col: conn.groupdealTicketCol,
	}
}
