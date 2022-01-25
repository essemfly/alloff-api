package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductGroupDAO struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Title       string             `json:"title"`
	ShortTitle  string             `json:"shorttitle"`
	Instruction []string           `json:"instruction"`
	ImgUrl      string             `json:"imgurl"`
	NumAlarms   int
	Products    []*ProductPriorityDAO
	StartTime   time.Time
	FinishTime  time.Time
	Created     time.Time
	Updated     time.Time
}

type ProductPriorityDAO struct {
	Priority  int
	ProductID primitive.ObjectID
}

type ExhibitionDAO struct {
	ID             primitive.ObjectID `bons:"_id, omitempty"`
	BannerImage    string
	ThumbnailImage string
	Title          string
	ShortTitle     string
	ProductGroups  []*ProductGroupDAO
}
