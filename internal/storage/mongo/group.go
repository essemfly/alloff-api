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

func (repo *groupRepo) List(exhibitionID string) ([]*domain.GroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"exhibitionid": exhibitionID}

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

func (repo *groupRepo) Update(groupDao *domain.GroupDAO) (*domain.GroupDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if _, err := repo.col.UpdateByID(ctx, groupDao.ID, bson.M{"$set": &groupDao}); err != nil {
		return nil, err
	}

	var updatedGroup *domain.GroupDAO
	if err := repo.col.FindOne(ctx, bson.M{"_id": groupDao.ID}).Decode(&updatedGroup); err != nil {
		return nil, err
	}

	return updatedGroup, nil
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

func (repo *groupRequestRepo) List(userID, exhibitionID string, status []domain.GroupRequestStatus) ([]*domain.GroupRequestDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{}
	if len(status) == 0 {
		filter = bson.M{
			"userid":       userID,
			"exhibitionid": exhibitionID,
		}
	} else {
		statusFilter := []interface{}{}
		for _, st := range status {
			statusFilter = append(statusFilter, bson.M{"status": st})
		}
		filter = bson.M{
			"userid":       userID,
			"exhibitionid": exhibitionID,
			"$or":          statusFilter,
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
