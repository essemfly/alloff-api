package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AlloffSizeDAO struct {
	ID         primitive.ObjectID `bson:"_id, omitempty"`
	SizeName   string
	GuideImage string
}
