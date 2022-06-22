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

type topBannersRepo struct {
	col *mongo.Collection
}

func (repo *topBannersRepo) Insert(banner *domain.TopBannerDAO) (*domain.TopBannerDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	banner.CreatedAt = time.Now()
	banner.UpdatedAt = time.Now()
	_, err := repo.col.InsertOne(ctx, banner)

	if err != nil {
		return nil, err
	}

	return banner, nil
}

func (repo *topBannersRepo) Get(itemID string) (*domain.TopBannerDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	itemObjID, _ := primitive.ObjectIDFromHex(itemID)

	item := &domain.TopBannerDAO{}
	if err := repo.col.FindOne(ctx, bson.M{"_id": itemObjID}).Decode(item); err != nil {
		return nil, err
	}

	return item, nil
}

func (repo *topBannersRepo) List(offset, limit int, onlyLive bool) ([]*domain.TopBannerDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{}
	options := options.Find()
	if onlyLive {
		filter["islive"] = true
		options.SetSort(bson.D{{Key: "weight", Value: -1}})
	} else {
		options.SetLimit(int64(limit))
		options.SetSkip(int64(offset))
	}

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cur, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		log.Println("err occured in hometabitem lists", err)
		return nil, 0, err
	}

	var items []*domain.TopBannerDAO
	err = cur.All(ctx, &items)
	if err != nil {
		log.Println("err occured in decoding top banner list", err)
		return nil, 0, err
	}
	return items, int(totalCount), nil
}

func (repo *topBannersRepo) Update(item *domain.TopBannerDAO) (*domain.TopBannerDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if _, err := repo.col.UpdateByID(ctx, item.ID, bson.M{"$set": &item}); err != nil {
		return nil, err
	}

	var updatedItem *domain.TopBannerDAO
	if err := repo.col.FindOne(ctx, bson.M{"_id": item.ID}).Decode(&updatedItem); err != nil {
		return nil, err
	}

	return updatedItem, nil
}

func MongoTopBannersRepo(conn *MongoDB) repository.TopBannersRepository {
	return &topBannersRepo{
		col: conn.topBannersCol,
	}
}
