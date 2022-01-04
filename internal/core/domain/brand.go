package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandDAO struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	KorName         string
	EngName         string
	KeyName         string
	Description     string
	LogoImgUrl      string
	Category        []*CategoryDAO
	SizeGuide       []SizeGuideDAO
	Created         time.Time
	Modulename      string
	Onpopular       bool
	MaxDiscountRate int
	IsOpen          bool
	IsHide          bool
	InMaintenance   bool
	NumNewProducts  int
}

type SizeGuideDAO struct {
	Label  string
	ImgUrl string
}

type LikeBrandDAO struct {
	Userid  string
	Brands  []*BrandDAO
	Created time.Time
	Updated time.Time
}
