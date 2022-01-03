package mongo

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type homeitemRepo struct {
	col *mongo.Collection
}

type featuredRepo struct {
	col *mongo.Collection
}

func (repo *homeitemRepo) Insert(item *domain.HomeItemDAO) error {
	return nil
}

func (repo *homeitemRepo) Update(item *domain.HomeItemDAO) error {
	return nil
}

func (repo *homeitemRepo) List() ([]*domain.HomeItemDAO, error) {
	return nil, nil
}

func (repo *featuredRepo) Insert(item *domain.FeaturedDAO) error {
	return nil
}

func (repo *featuredRepo) List() ([]*domain.FeaturedDAO, error) {
	return nil, nil
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
