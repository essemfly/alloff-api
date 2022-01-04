package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductGroupDAO struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Hidden      bool               `json:"hidden"`
	Instruction []string           `json:"instruction"`
	ShortTitle  string             `json:"shorttitle"`
	Title       string             `json:"title"`
	ImgUrl      string             `json:"imgurl"`
	NumAlarms   int
	Products    []*ProductPriorityDAO
	StartTime   time.Time
	FinishTime  time.Time
	Created     time.Time
	Updated     time.Time
}

type ProductPriorityDAO struct {
	Priority int
	Product  *ProductDAO
}

type SpecialExhibitionDAO struct {
	ID primitive.ObjectID `bons:"_id, omitempty"`
}
