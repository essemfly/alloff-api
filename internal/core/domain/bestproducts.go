package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BestProductDAO struct {
	ID               primitive.ObjectID `bson:"_id, omitempty"`
	AlloffCategoryID string
	Products         []*ProductDAO
	CreatedAt        time.Time
}
