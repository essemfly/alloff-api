package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExhibitionDAO struct {
	ID             primitive.ObjectID `bson:"_id, omitempty"`
	BannerImage    string
	ThumbnailImage string
	Title          string
	SubTitle       string
	Description    string
	StartTime      time.Time
	FinishTime     time.Time
	ProductGroups  []*ProductGroupDAO
	IsLive         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (exDao *ExhibitionDAO) ListCheifProducts() []*ProductDAO {
	numProductsToShow := 10
	products := []*ProductDAO{}
	if len(exDao.ProductGroups) > 0 {
		if len(exDao.ProductGroups[0].Products) > 0 {
			for idx, productPriority := range exDao.ProductGroups[0].Products {
				if idx >= numProductsToShow {
					break
				}
				products = append(products, productPriority.Product)
			}
		}
	}
	return products
}
