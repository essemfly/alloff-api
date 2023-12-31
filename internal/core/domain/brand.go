package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryDAO struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// 사용자에게 보여지는 Category name
	Name string
	// Category 식별 identifier key name
	KeyName string
	// Category 식별 identifier
	CatIdentifier string
	// Category가 속해있는 브랜드
	BrandKeyname string
	// Size Guide
	SizeGuide string
}

type ClassifierDAO struct {
	BrandKeyname    string
	CategoryName    string
	AlloffCategory1 *AlloffCategoryDAO
	AlloffCategory2 *AlloffCategoryDAO
	HeuristicRules  map[string]string
}

type AlloffCategoryDAO struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string
	KeyName      string
	Level        int
	ParentId     primitive.ObjectID
	CategoryType string `json:"type" bson:"type"`
	ImgURL       string
}

type BrandDAO struct {
	ID                    primitive.ObjectID `bson:"_id,omitempty"`
	KorName               string
	EngName               string
	KeyName               string
	Description           string
	LogoImgUrl            string               `bson:"logoimgurl,omitempty"`
	BackImgUrl            string               `bson:"backimgurl,omitempty"`
	Category              []*CategoryDAO       `bson:"category,omitempty"`
	AlloffCategory        []*AlloffCategoryDAO `bson:"alloffcategory,omitempty"`
	SizeGuide             []SizeGuideDAO       `bson:"sizeguide,omitempty"` // brand seeder에서 업데이트 시 기존 데이터 유지시키기 위함
	Created               time.Time            `bson:"created,omitempty"`
	Onpopular             bool
	MaxDiscountRate       int
	IsOpen                bool
	IsHide                bool
	InMaintenance         bool
	NumNewProductsIn3days int
	UseAlloffCategory     bool
}

type BrandCountsData struct {
	Brand  *BrandDAO
	Counts int
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
