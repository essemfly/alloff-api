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

type cartRepo struct {
	col *mongo.Collection
}

func (repo *cartRepo) Get(cartID string) (*domain.Basket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cartObjectId, _ := primitive.ObjectIDFromHex(cartID)
	filter := bson.M{"_id": cartObjectId}
	var cart *domain.Basket
	if err := repo.col.FindOne(ctx, filter).Decode(&cart); err != nil {
		return nil, err
	}
	return cart, nil
}

func (repo *cartRepo) Upsert(cartDao *domain.Basket) (*domain.Basket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": cartDao.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &cartDao}, opts); err != nil {
		return nil, err
	}

	var updatedCart *domain.Basket
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedCart); err != nil {
		return nil, err
	}

	return updatedCart, nil
}

func MongoCartsRepo(conn *MongoDB) repository.CartsRepository {
	return &cartRepo{
		col: conn.cartCol,
	}
}
