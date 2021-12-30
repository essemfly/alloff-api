package mongo

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type brandsRepo struct {
	col *mongo.Collection
}

func (repo *brandsRepo) Get(id string) (*domain.BrandDAO, error) {
	return nil, nil
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
