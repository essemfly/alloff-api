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

type bestProductsRepo struct {
	col *mongo.Collection
}

func (repo *bestProductsRepo) Insert(bestProductDao *domain.BestProductDAO) (*domain.BestProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	bestProductDao.CreatedAt = time.Now()
	_, err := repo.col.InsertOne(ctx, bestProductDao)

	if err != nil {
		return nil, err
	}

	return bestProductDao, nil
}

func (repo *bestProductsRepo) GetLatest(alloffCategoryID string) (*domain.BestProductDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSort(bson.M{"_id": -1})
	options.SetLimit(1)
	items := []*domain.BestProductDAO{}
	filter := bson.M{"alloffcategoryid": alloffCategoryID}
	cursor, err := repo.col.Find(ctx, filter, options)
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

func MongoBestProductsRepo(conn *MongoDB) repository.BestProductsRepository {
	return &bestProductsRepo{
		col: conn.bestProductsCol,
	}
}
