package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TopBannerDAO struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	ImageUrl     string
	ExhibitionID string
	Title        string
	SubTitle     string
	IsLive       bool
	Weight       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
