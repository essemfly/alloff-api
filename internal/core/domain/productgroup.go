package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductGroupType string

const (
	PRODUCT_GROUP_TIMEDEAL       ProductGroupType = "PRODUCT_GROUP_TIMEDEAL"
	PRODUCT_GROUP_GROUPDEAL      ProductGroupType = "PRODUCT_GROUP_GROUPDEAL"
	PRODUCT_GROUP_EXHIBITION     ProductGroupType = "PRODUCT_GROUP_EXHIBITION"
	PRODUCT_GROUP_BRAND_TIMEDEAL ProductGroupType = "PRODUCT_GROUP_BRAND_TIMEDEAL"
)

type ProductGroupDAO struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	Title        string             `json:"title"`
	ShortTitle   string             `json:"shorttitle"`
	Instruction  []string           `json:"instruction"`
	ImgUrl       string             `json:"imgurl"`
	NumAlarms    int
	Products     []*ProductPriorityDAO
	GroupType    ProductGroupType
	StartTime    time.Time
	FinishTime   time.Time
	Created      time.Time
	Updated      time.Time
	ExhibitionID string
	Brand        *BrandDAO
}

type ProductPriorityDAO struct {
	Priority  int
	Product   *ProductDAO
	ProductID primitive.ObjectID
}

func (pgDao *ProductGroupDAO) AppendProduct(priorityDao *ProductPriorityDAO) bool {
	newPds := []*ProductPriorityDAO{}

	isNewProduct := true
	for _, alreadyInProduct := range pgDao.Products {
		if alreadyInProduct.ProductID.Hex() == priorityDao.ProductID.Hex() {
			isNewProduct = false
			if alreadyInProduct.Priority != priorityDao.Priority {
				newPds = append(newPds, priorityDao)
			}
		} else {
			newPds = append(newPds, alreadyInProduct)
		}
	}

	pgDao.Products = newPds
	if isNewProduct {
		pgDao.Products = append(pgDao.Products, priorityDao)
		return true
	}
	return false
}

func (pgDao *ProductGroupDAO) IsLive() bool {
	now := time.Now()
	if now.After(pgDao.StartTime) && now.Before(pgDao.FinishTime) {
		return true
	}
	return false
}
