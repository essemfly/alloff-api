package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandDAO struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	KorName               string
	EngName               string
	KeyName               string
	Description           string
	LogoImgUrl            string
	BackImgUrl            string
	Category              []*CategoryDAO
	SizeGuide             []SizeGuideDAO
	Created               time.Time
	Onpopular             bool
	MaxDiscountRate       int
	IsOpen                bool
	IsHide                bool
	InMaintenance         bool
	NumNewProductsIn3days int
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

func (brand *BrandDAO) IsOpenBrand() bool {
	return brand.IsOpen && !brand.IsHide && !brand.InMaintenance
}
