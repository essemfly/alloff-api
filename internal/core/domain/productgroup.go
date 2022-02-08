package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductGroupDAO struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Title       string             `json:"title"`
	ShortTitle  string             `json:"shorttitle"`
	Instruction []string           `json:"instruction"`
	ImgUrl      string             `json:"imgurl"`
	NumAlarms   int
	Products    []*ProductPriorityDAO
	StartTime   time.Time
	FinishTime  time.Time
	Created     time.Time
	Updated     time.Time
}

type ProductPriorityDAO struct {
	Priority  int
	Product   *ProductDAO
	ProductID primitive.ObjectID
}

type ExhibitionDAO struct {
	ID             primitive.ObjectID `bons:"_id, omitempty"`
	BannerImage    string
	ThumbnailImage string
	Title          string
	Description    string
	ProductGroups  []*ProductGroupDAO
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

func (pgDao *ProductGroupDAO) RemoveProduct(productID string) {
	newPds := []*ProductPriorityDAO{}
	for _, pd := range pgDao.Products {
		if pd.ProductID.Hex() != productID {
			newPds = append(newPds, pd)
		}
	}
	pgDao.Products = newPds
}
