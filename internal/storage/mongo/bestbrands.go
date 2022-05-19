package mongo

import (
	"context"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type bestBrandsRepo struct {
	col *mongo.Collection
}

func (repo *bestBrandsRepo) Insert(bestBrandDao *domain.BestBrandDAO) (*domain.BestBrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	bestBrandDao.CreatedAt = time.Now()
	_, err := repo.col.InsertOne(ctx, bestBrandDao)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return bestBrandDao, nil
}

func (repo *bestBrandsRepo) GetLatest() (*domain.BestBrandDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.M{"_id": -1})
	options.SetLimit(1)
	items := []*domain.BestBrandDAO{}
	cursor, err := repo.col.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &items)
	if err != nil {
		return nil, err
	}

	if len(items) == 0 {
		return nil, nil
	}

	return items[0], nil
}

func MongoBestBrandsRepo(conn *MongoDB) repository.BestBrandRepository {
	return &bestBrandsRepo{
		col: conn.bestBrandsCol,
	}
}
