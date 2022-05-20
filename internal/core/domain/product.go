package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductDAO struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	ProductInfo     *ProductMetaInfoDAO
	ProductGroupID  string
	ExhibitionID    string
	AlloffName      string
	ThumbnailImage  string
	OriginalPrice   int
	DiscountedPrice int
	DiscountRate    int
	Inventory       []*InventoryDAO
	IsSoldout       bool
	IsRemoved       bool
	Score           *ProductScoreInfoDAO
	Created         time.Time
	Updated         time.Time
}

func (pd *ProductDAO) UpdatePrice(origPrice, discountedPrice int) bool {
	pd.OriginalPrice = origPrice
	pd.DiscountedPrice = discountedPrice
	if origPrice == 0 {
		pd.OriginalPrice = discountedPrice
	}

	return true
}

func (pd *ProductDAO) UpdateScore(newScore *ProductScoreInfoDAO) {
	pd.Score = newScore
}

func (pd *ProductDAO) CheckSoldout() {
	isSoldout := true
	for _, inv := range pd.Inventory {
		if inv.Quantity > 0 {
			isSoldout = false
			break
		}
	}

	pd.IsSoldout = isSoldout
}

func (pd *ProductDAO) UpdateInventory(newInven []*InventoryDAO) {
	pd.Inventory = newInven

	isSoldout := true
	for _, inv := range newInven {
		if inv.Quantity > 0 {
			isSoldout = false
			break
		}
	}

	pd.IsSoldout = isSoldout
}

func (pd *ProductDAO) MapAlloffInventory() {
	mappingPolicies := pd.ProductInfo.Brand.InventoryMappingPolicies
	productInventory := pd.Inventory
	alloffInventory := []*AlloffInventoryDAO{}

	if mappingPolicies != nil {
		for _, mappingPolicy := range mappingPolicies {
			for _, inv := range productInventory {
				if mappingPolicy.BrandSize == inv.Size {
					alloffInventory = append(alloffInventory, &AlloffInventoryDAO{
						AlloffSize: mappingPolicy.AlloffSize,
						Quantity:   inv.Quantity,
					})
				}
			}
		}
	}

	if len(alloffInventory) > 0 {
		pd.ProductInfo.IsInventoryMapped = true
	}

	pd.ProductInfo.AlloffInventory = alloffInventory
}

// func (pd *ProductDAO) UpdateAlloffCategory(cat *ProductAlloffCategoryDAO) {
// 	pd.ProductInfo.AlloffCategories = cat
// }

func (pd *ProductDAO) UpdateInstruction(instruction *AlloffInstructionDAO) {
	pd.ProductInfo.SalesInstruction = instruction
}

func (pd *ProductDAO) Release(size string, quantity int) error {
	for idx, option := range pd.Inventory {
		if option.Size == size {
			if option.Quantity < quantity {
				return errors.New("insufficient product quantity")
			}
			pd.Inventory[idx].Quantity -= quantity

			return nil
		}
	}
	return errors.New("no matched product size option")
}

func (pd *ProductDAO) Revert(size string, quantity int) error {
	for idx, option := range pd.Inventory {
		if option.Size == size {
			if option.Quantity == 0 {
				pd.IsSoldout = false
			}
			pd.Inventory[idx].Quantity += quantity

			return nil
		}
	}
	return errors.New("no matched product size option")
}

type PriceDAO struct {
	OriginalPrice float32
	CurrencyType  CurrencyType
	CurrentPrice  float32
	History       []*PriceHistoryDAO
}

type LikeProductDAO struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Productid  string
	OldProduct *ProductDAO `bson:"product"`
	Userid     string
	IsPushed   bool
	LastPrice  int
	Removed    bool
	Created    time.Time
	Updated    time.Time
}

type ProductDiffDAO struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	OldPrice   int                `json:"oldprice"`
	NewProduct *ProductDAO        `json:"newproduct"`
	Type       string             `json:"type"`
	IsPushed   bool
}
