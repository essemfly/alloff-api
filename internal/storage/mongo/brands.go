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

type brandsRepo struct {
	col *mongo.Collection
}

func (repo *brandsRepo) Get(ID string) (*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	brandObjectId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": brandObjectId}

	var brand *domain.BrandDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&brand); err != nil {
		return nil, err
	}

	return brand, nil
}

func (repo *brandsRepo) GetByKeyname(keyname string) (*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"keyname": keyname}

	var brand *domain.BrandDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&brand); err != nil {
		return nil, err
	}

	return brand, nil
}

func (repo *brandsRepo) List(limit, offset int, filter, sortingOptions interface{}) ([]*domain.BrandDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(sortingOptions)
	options.SetLimit(int64(limit))
	options.SetSkip(int64(offset))

	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}

	var brands []*domain.BrandDAO
	err = cursor.All(ctx, &brands)
	if err != nil {
		return nil, 0, err
	}

	return brands, int(totalCount), nil
}

func (repo *brandsRepo) Upsert(brand *domain.BrandDAO) (*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"keyname": &brand.KeyName}
	brand.ID = ""
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &brand}, opts); err != nil {
		return nil, err
	}

	var updatedBrand *domain.BrandDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedBrand); err != nil {
		return nil, err
	}

	return updatedBrand, nil
}

func MongoBrandsRepo(conn *MongoRepo) repository.BrandsRepository {
	return &brandsRepo{
		col: conn.brandsCol,
	}
}
