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
	INVENTORY_ASCENDING     PriceSortingType = "inventoryAscending"
	INVENTORY_DESCENDING    PriceSortingType = "inventoryDescending"
)

type PriceRangeType string

const (
	PRICE_RANGE_30  PriceRangeType = "30"
	PRICE_RANGE_50  PriceRangeType = "50"
	PRICE_RANGE_70  PriceRangeType = "70"
	PRICE_RANGE_100 PriceRangeType = "100"
)

type ProductDAO struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	ProductInfo          *ProductMetaInfoDAO
	ProductGroupID       string
	ExhibitionID         string
	Weight               int
	IsRemoved            bool
	OnSale               bool
	ExhibitionStartTime  time.Time
	ExhibitionFinishTime time.Time
	Created              time.Time
	Updated              time.Time
}

func (pd *ProductDAO) IsSaleable() bool {
	if time.Now().After(pd.ExhibitionStartTime) && time.Now().Before(pd.ExhibitionFinishTime) {
		return true
	}
	return false
}
