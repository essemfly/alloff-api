package mongo

import (
	"context"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type categoryRepo struct {
	col *mongo.Collection
}

func (repo *categoryRepo) List(brandKeyname string) ([]*domain.CategoryDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	sortingOptions := bson.D{{Key: "catidentifier", Value: 1}, {Key: "_id", Value: 1}}
	options := options.Find()
	options.SetSort(sortingOptions)

	cursor, err := repo.col.Find(ctx, bson.M{"brandkeyname": brandKeyname}, options)
	if err != nil {
		return nil, err
	}

	var cats []*domain.CategoryDAO
	err = cursor.All(ctx, &cats)
	if err != nil {
		return nil, err
	}

	return cats, nil
}

func (repo *categoryRepo) Upsert(category *domain.CategoryDAO) (*domain.CategoryDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"keyname": &category.KeyName}
	category.ID = ""
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &category}, opts); err != nil {
		return nil, err
	}

	var updatedCat *domain.CategoryDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedCat); err != nil {
		return nil, err
	}
	return updatedCat, nil
}

func MongoCategoriesRepo(conn *MongoDB) repository.CategoriesRepository {
	return &categoryRepo{
		col: conn.categoryCol,
	}
}
