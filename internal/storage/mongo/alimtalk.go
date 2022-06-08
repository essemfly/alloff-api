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

type alimtalkRepo struct {
	col *mongo.Collection
}

func (repo *alimtalkRepo) GetByDetail(userID, templateCode, referenceID string) (*domain.AlimtalkDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	filter := bson.M{
		"templatecode": templateCode,
		"referenceid":  referenceID,
		"userid":       userID,
	}

	var alimtalk *domain.AlimtalkDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&alimtalk); err != nil {
		return nil, err
	}
	return alimtalk, nil
}
func (repo *alimtalkRepo) Insert(alimtalk *domain.AlimtalkDAO) (*domain.AlimtalkDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	oid, err := repo.col.InsertOne(ctx, alimtalk)
	if err != nil {
		return nil, err
	}

	var newAlimtalk domain.AlimtalkDAO
	filter := bson.M{"_id": oid.InsertedID}
	err = repo.col.FindOne(ctx, filter).Decode(&newAlimtalk)
	if err != nil {
		return nil, err
	}

	return &newAlimtalk, nil
}

func (repo *alimtalkRepo) Update(alimtalk *domain.AlimtalkDAO) (*domain.AlimtalkDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	alimtalk.UpdatedAt = time.Now()
	filter := bson.M{"_id": alimtalk.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &alimtalk}, opts); err != nil {
		return nil, err
	}

	return alimtalk, nil
}

func MongoAlimtalksRepo(conn *MongoDB) repository.AlimtalksRepository {
	return &alimtalkRepo{
		col: conn.alimtalkCol,
	}
}
