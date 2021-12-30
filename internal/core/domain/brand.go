package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BrandDAO struct {
	ID              string `bson:"_id,omitempty"`
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

type LikeDAO struct {
	Userid  string
	Brands  []*BrandDAO
	Created time.Time
	Updated time.Time
}

type CategoryDAO struct {
	ID string `bson:"_id,omitempty"`
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
	ID           string `bson:"_id, omitempty"`
	Name         string
	KeyName      string
	Level        int
	ParentId     primitive.ObjectID
	CategoryType string `json:"type" bson:"type"`
	ImgURL       string
}
