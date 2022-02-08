package mongo

import (
	"context"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type hometabitemRepo struct {
	col *mongo.Collection
}

func (repo *hometabitemRepo) Insert(item *domain.HomeTabItemDAO) (*domain.HomeTabItemDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	_, err := repo.col.InsertOne(ctx, item)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (repo *hometabitemRepo) Get(itemID string) (*domain.HomeTabItemDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	itemObjID, _ := primitive.ObjectIDFromHex(itemID)

	item := &domain.HomeTabItemDAO{}
	if err := repo.col.FindOne(ctx, bson.M{"_id": itemObjID}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (repo *hometabitemRepo) List(offset, limit int, onlyLive bool) ([]*domain.HomeTabItemDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := primitive.NewDateTimeFromTime(time.Now())
	filter := bson.M{}
	options := options.Find()
	if onlyLive {
		filter["finishedat"] = bson.M{"$gte": now}
		options.SetSort(bson.D{{Key: "weight", Value: -1}})
	} else {
		options.SetLimit(int64(limit))
		options.SetSkip(int64(offset))
	}

	cur, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		log.Println("err occured in hometabitem lists", err)
		return nil, err
	}

	var items []*domain.HomeTabItemDAO
	err = cur.All(ctx, &items)
	if err != nil {
		log.Println("err occured in decoding", err)
		return nil, err
	}

	return items, nil
}

func (repo *hometabitemRepo) Update(item *domain.HomeTabItemDAO) (*domain.HomeTabItemDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if _, err := repo.col.UpdateByID(ctx, item.ID, bson.M{"$set": &item}); err != nil {
		return nil, err
	}

	var updatedItem *domain.HomeTabItemDAO
	if err := repo.col.FindOne(ctx, bson.M{"_id": item.ID}).Decode(&updatedItem); err != nil {
		return nil, err
	}

	return updatedItem, nil
}

func MongoHometabItemsRepo(conn *MongoDB) repository.HomeTabItemsRepository {
	return &hometabitemRepo{
		col: conn.hometabItemsCol,
	}
}
