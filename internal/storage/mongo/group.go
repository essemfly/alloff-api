package mongo

import (
	"context"
	"fmt"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (repo *groupRepo) Get(ID string) (*domain.GroupDAO, error) {
	return nil, fmt.Errorf("TOBO IMPLEMENTED")
}

func (repo *groupRepo) List(exhibitionID string) ([]*domain.GroupDAO, error) {
	return nil, fmt.Errorf("TOBO IMPLEMENTED")
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

func (repo *groupRequestRepo) List(userID, exhibitionID string, status domain.GroupRequestStatus) ([]*domain.GroupRequestDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{
		"userid":       userID,
		"exhibitionid": exhibitionID,
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

func (repo *groupRequestRepo) Update(dao *domain.GroupRequestDAO) (*domain.GroupRequestDAO, error) {
	return nil, fmt.Errorf("TOBO IMPLEMENTED")
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
