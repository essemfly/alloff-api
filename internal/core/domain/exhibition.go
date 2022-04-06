package domain

import (
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ExhibitionType string

const (
	EXHIBITION_TIMEDEAL  ExhibitionType = "TIMEDEAL"
	EXHIBITION_NORMAL    ExhibitionType = "NORMAL"
	EXHIBITION_GROUPDEAL ExhibitionType = "GROUPDEAL"
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
	ExhibitionType ExhibitionType
	TargetSales    int
}

// TODO: To be specified with real number
func (exDao *ExhibitionDAO) GetCurrentSales() int {
	return rand.Intn(exDao.TargetSales)
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

func (exDao *ExhibitionDAO) IsSales() bool {
	now := time.Now()
	if now.After(exDao.StartTime) && now.Before(exDao.FinishTime) {
		return true
	}
	return false
}
