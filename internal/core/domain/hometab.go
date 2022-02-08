package domain

import (
	"time"

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
	Path   string
	Params string
}

type HomeTabItemDAO struct {
	ID           primitive.ObjectID `bon:"_id, omitempty"`
	Title        string
	Description  string
	Tags         []string
	BackImageUrl string
	Type         HomeTabItemType
	Products     []*ProductDAO
	Brands       []*BrandDAO
	Exhibitions  []*ExhibitionDAO
	Reference    *ReferenceTarget
	CreatedAt    time.Time
	UpdatedAt    time.Time
	StartedAt    time.Time
	EndedAt      time.Time
}
