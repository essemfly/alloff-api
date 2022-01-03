package mongo

import (
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type alimtalkRepo struct {
	col *mongo.Collection
}

func MongoAlimtalksRepo(conn *MongoDB) repository.AlimtalksRepository {
	return &alimtalkRepo{
		col: conn.alimtalkCol,
	}
}
