package mongo

import (
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type productGroupRepo struct {
	col *mongo.Collection
}

func MongoProductGroupsRepo(conn *MongoDB) repository.ProductGroupsRepository {
	return &productGroupRepo{
		col: conn.productGroupCol,
	}
}
