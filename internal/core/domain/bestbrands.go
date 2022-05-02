package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type BestBrandDAO struct {
	ID        primitive.ObjectID `bson:"_id, omitempty"`
	Brands    []*BrandDAO
	CreatedAt time.Time
}
