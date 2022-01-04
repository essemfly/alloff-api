package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

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
