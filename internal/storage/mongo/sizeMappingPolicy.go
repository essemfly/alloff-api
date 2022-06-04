package mongo

import (
	"context"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type sizeMappingPolicyRepo struct {
	col *mongo.Collection
}

func (repo *sizeMappingPolicyRepo) Get(ID string) (*domain.SizeMappingPolicyDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": oid}

	var sizeMappingPolicy *domain.SizeMappingPolicyDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&sizeMappingPolicy); err != nil {
		return nil, err
	}

	return sizeMappingPolicy, nil
}

func (repo *sizeMappingPolicyRepo) ListByDetail(size string, productTypes []domain.AlloffProductType, alloffCategpryID string) ([]*domain.SizeMappingPolicyDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	alloffCategpryOid, _ := primitive.ObjectIDFromHex(alloffCategpryID)
	filter := bson.M{
		"sizes":              size,
		"alloffcategory._id": alloffCategpryOid,
		"alloffproducttype":  bson.M{"$all": productTypes},
	}

	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var sizeMappingPolicies []*domain.SizeMappingPolicyDAO
	err = cursor.All(ctx, &sizeMappingPolicies)
	if err != nil {
		return nil, err
	}
	return sizeMappingPolicies, nil
}

func (repo *sizeMappingPolicyRepo) List() ([]*domain.SizeMappingPolicyDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cursor, err := repo.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var sizeMappingPolicies []*domain.SizeMappingPolicyDAO
	err = cursor.All(ctx, &sizeMappingPolicies)
	if err != nil {
		return nil, err
	}

	return sizeMappingPolicies, nil
}

func (repo *sizeMappingPolicyRepo) Insert(dao *domain.SizeMappingPolicyDAO) (*domain.SizeMappingPolicyDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	oid, err := repo.col.InsertOne(ctx, dao)
	if err != nil {
		return nil, err
	}

	var newDao *domain.SizeMappingPolicyDAO
	filter := bson.M{"_id": oid.InsertedID}
	err = repo.col.FindOne(ctx, filter).Decode(&newDao)
	if err != nil {
		return nil, err
	}

	return newDao, nil
}

func (repo *sizeMappingPolicyRepo) Upsert(dao *domain.SizeMappingPolicyDAO) (*domain.SizeMappingPolicyDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{
		"alloffsize.alloffsizename": dao.AlloffSize.AlloffSizeName,
		"alloffcategory.keyname":    dao.AlloffCategory.KeyName,
		"alloffproducttype":         bson.M{"$elemMatch": bson.M{"$in": dao.AlloffProductType}},
	}

	log.Println(dao)
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": dao}, opts); err != nil {
		return nil, err
	}

	var updatedSizeMappingPolicy *domain.SizeMappingPolicyDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedSizeMappingPolicy); err != nil {
		return nil, err
	}

	return updatedSizeMappingPolicy, nil
}

func MongoSizeMappingPolicyRepo(conn *MongoDB) repository.SizeMappingPolicyRepository {
	return &sizeMappingPolicyRepo{
		col: conn.sizeMappingPolicyCol,
	}
}
