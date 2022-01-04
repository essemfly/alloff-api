package mongo

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type homeitemRepo struct {
	col *mongo.Collection
}

type featuredRepo struct {
	col *mongo.Collection
}

func (repo *homeitemRepo) Insert(item *domain.HomeItemDAO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.col.InsertOne(ctx, item)
	if err != nil {
		return err
	}
	return nil
}

func (repo *homeitemRepo) Update(item *domain.HomeItemDAO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(false)

	filter := bson.M{"title": item.Title}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &item}, opts); err != nil {
		return err
	}
	return nil
}

func (repo *homeitemRepo) List() ([]*domain.HomeItemDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{}

	options := options.Find()
	options.SetSort(bson.D{{Key: "priority", Value: -1}})

	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	var homeitems []*domain.HomeItemDAO
	err = cursor.All(ctx, &homeitems)
	if err != nil {
		return nil, err
	}

	return homeitems, nil
}

func (repo *featuredRepo) Insert(item *domain.FeaturedDAO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.col.InsertOne(ctx, item)
	if err != nil {
		return err
	}

	return nil
}

func (repo *featuredRepo) List() ([]*domain.FeaturedDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{
		"enddate": bson.M{
			"$gte": primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	options := options.Find()
	options.SetSort(bson.D{{Key: "order", Value: 1}})

	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}

	var featureds []*domain.FeaturedDAO
	err = cursor.All(ctx, &featureds)
	if err != nil {
		return nil, err
	}

	return featureds, nil
}

func MongoHomeItemsRepo(conn *MongoDB) repository.HomeItemsRepository {
	return &homeitemRepo{
		col: conn.homeitemCol,
	}
}

func MongoFeaturedsRepo(conn *MongoDB) repository.FeaturedsRepository {
	return &featuredRepo{
		col: conn.featuredCol,
	}
}
