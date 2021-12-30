package mongo

import (
	"context"
	"errors"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type brandsRepo struct {
	col *mongo.Collection
}

func (repo *brandsRepo) Get(id string) (*domain.BrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	brandObjectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": brandObjectId}

	var brand *domain.BrandDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&brand); err != nil {
		return nil, errors.New("brand not found error")
	}

	return brand, nil
}

func (repo *brandsRepo) List(alloffCategoryID *string) ([]*domain.BrandDAO, error) {
	return nil, nil
}

func (repo *brandsRepo) Upsert(*domain.BrandDAO) error {
	return nil
}

func MongoBrandsRepo(conn *MongoRepo) repository.BrandsRepository {
	return &brandsRepo{
		col: conn.brandsCol,
	}
}
