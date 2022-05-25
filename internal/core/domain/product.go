package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryClassifiedType string

const (
	CLASSIFIED_DONE      CategoryClassifiedType = "CLASSIFIED_DONE"
	NOT_CLASSIFIED       CategoryClassifiedType = "NOT_CLASSIFIED"
	NO_MATTER_CLASSIFIED CategoryClassifiedType = "NO_MATTER_CLASSIFIED"
)

type PriceSortingType string

const (
	PRICE_ASCENDING         PriceSortingType = "ascending"
	PRICE_DESCENDING        PriceSortingType = "descending"
	DISCOUNTRATE_ASCENDING  PriceSortingType = "discountrateAscending"
	DISCOUNTRATE_DESCENDING PriceSortingType = "discountrateDescending"
)

type PriceRangeType string

const (
	PRICE_RANGE_30  PriceRangeType = "30"
	PRICE_RANGE_50  PriceRangeType = "50"
	PRICE_RANGE_70  PriceRangeType = "70"
	PRICE_RANGE_100 PriceRangeType = "100"
)

type AlloffSizeDAO struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	AlloffCategory *AlloffCategoryDAO
	AlloffSizeName string
}

type AlloffInventoryDAO struct {
	AlloffSize AlloffSizeDAO
	Quantity   int
}

type ProductDAO struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	ProductInfo          *ProductMetaInfoDAO
	ProductGroupID       string
	ExhibitionID         string
	Weight               int
	IsNotSale            bool // Exhibition이 종료되거나 혹은 backoffice에서 삭제했을때
	ExhibitionStartTime  time.Time
	ExhibitionFinishTime time.Time
	Created              time.Time
	Updated              time.Time
}
