package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HomeTabItemType string

const (
	HOMETAB_ITEM_BRANDS           = "HOMETAB_ITEM_BRANDS"
	HOMETAB_ITEM_BRAND_EXHIBITION = "HOMETAB_ITEM_BRAND_EXHIBITION"
	HOMETAB_ITEM_EXHIBITIONS      = "HOMETAB_ITEM_EXHIBITIONS"
	HOMETAB_ITEM_EXHIBITION       = "HOMETAB_ITEM_EXHIBITION"
	HOMETAB_ITEM_PRODUCTS         = "HOMETAB_ITEM_PRODUCTS"
)

type ReferenceTarget struct {
	Path    string
	Params  string
	Options []model.SortingType
}

type HomeTabItemDAO struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Title        string
	Description  string
	Tags         []string
	BackImageUrl string
	Type         HomeTabItemType
	Weight       int
	Products     []*ProductDAO
	Brands       []*BrandDAO
	Exhibitions  []*ExhibitionDAO
	Reference    *ReferenceTarget
	CreatedAt    time.Time
	UpdatedAt    time.Time
	StartedAt    time.Time
	FinishedAt   time.Time
}

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
