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

type alloffSizeRepo struct {
	col *mongo.Collection
}

func (repo *alloffSizeRepo) Upsert(alloffSize *domain.AlloffSizeDAO) (*domain.AlloffSizeDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"alloffsizename":         alloffSize.AlloffSizeName,
		"alloffcategory.keyname": alloffSize.AlloffCategory.KeyName,
		"producttype":            bson.M{"$elemMatch": bson.M{"$in": alloffSize.ProductType}},
	}

	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &alloffSize}, opts); err != nil {
		return nil, err
	}

	var updatedAlloffSize *domain.AlloffSizeDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedAlloffSize); err != nil {
		return nil, err
	}

	return updatedAlloffSize, nil
}

func (repo *alloffSizeRepo) List(offset, limit int) ([]*domain.AlloffSizeDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	totalCount, _ := repo.col.CountDocuments(ctx, bson.M{})
	cursor, err := repo.col.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, 0, err
	}

	var alloffSizes []*domain.AlloffSizeDAO
	err = cursor.All(ctx, &alloffSizes)
	if err != nil {
		return nil, 0, err
	}

	return alloffSizes, int(totalCount), nil
}

func (repo *alloffSizeRepo) Get(alloffSizeID string) (*domain.AlloffSizeDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(alloffSizeID)
	if err != nil {
		return nil, err
	}

	var alloffSize *domain.AlloffSizeDAO
	filter := bson.M{"_id": oid}
	if err := repo.col.FindOne(ctx, filter).Decode(&alloffSize); err != nil {
		return nil, err
	}

	return alloffSize, nil
}

func MongoAlloffSizeRepo(conn *MongoDB) repository.AlloffSizeRepository {
	return &alloffSizeRepo{
		col: conn.alloffSizeCol,
	}
}
