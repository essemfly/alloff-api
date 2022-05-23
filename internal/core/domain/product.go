package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlloffSizeDAO struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	AlloffCategory *AlloffCategoryDAO
	SizeName       string
	GuideImage     string
}

type AlloffInventoryDAO struct {
	AlloffSize AlloffSizeDAO
	Quantity   int
}

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
	Inventory       []*AlloffInventoryDAO
	IsSoldout       bool
	IsRemoved       bool
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

func (pd *ProductDAO) UpdateInventory(newInven []*AlloffInventoryDAO) {
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

	for _, mappingPolicy := range mappingPolicies {
		for _, inv := range productInventory {
			if mappingPolicy.BrandSize == inv.AlloffSize.SizeName {
				alloffInventory = append(alloffInventory, &AlloffInventoryDAO{
					AlloffSize: mappingPolicy.AlloffSize,
					Quantity:   inv.Quantity,
				})
			}
		}
	}

	if len(alloffInventory) > 0 {
		pd.ProductInfo.IsInventoryMapped = true
	}

	pd.ProductInfo.AlloffInventory = alloffInventory
}

func (pd *ProductDAO) Release(size string, quantity int) error {
	for idx, option := range pd.Inventory {
		if option.AlloffSize.SizeName == size {
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
		if option.AlloffSize.SizeName == size {
			if option.Quantity == 0 {
				pd.IsSoldout = false
			}
			pd.Inventory[idx].Quantity += quantity

			return nil
		}
	}
	return errors.New("no matched product size option")
}
