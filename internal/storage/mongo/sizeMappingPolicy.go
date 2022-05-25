package mongo

import (
	"context"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

// 카테고리, 타입, 사이즈명으로 찾는다. 그친구의 AlloffSize 를 찾는다.
func (repo *sizeMappingPolicyRepo) GetByDetail(size string, productTypes []domain.AlloffProductType, alloffCategoryID string) (*domain.SizeMappingPolicyDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	alloffCategoryOid, _ := primitive.ObjectIDFromHex(alloffCategoryID)
	filter := bson.M{
		"sizes":              size,
		"alloffcategory._id": alloffCategoryOid,
		"alloffproducttype":  bson.M{"$all": productTypes},
	}

	var sizeMappingPolicy *domain.SizeMappingPolicyDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&sizeMappingPolicy); err != nil {
		return nil, err
	}

	return sizeMappingPolicy, nil
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

func (repo *sizeMappingPolicyRepo) Update(dao *domain.SizeMappingPolicyDAO) (*domain.SizeMappingPolicyDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := repo.col.UpdateOne(ctx, bson.M{"_id": dao.ID}, bson.M{"$set": &dao})
	if err != nil {
		return nil, err
	}

	return dao, nil
}

func MongoSizeMappingPolicyRepo(conn *MongoDB) repository.SizeMappingPolicyRepository {
	return &sizeMappingPolicyRepo{
		col: conn.sizeMappingPolicyCol,
	}
}
