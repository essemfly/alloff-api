package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlloffSizeDAO struct {
	ID               primitive.ObjectID `bson:"_id, omitempty"`
	AlloffCategory   *AlloffCategoryDAO
	AlloffSizeName   string
	OriginalSizeName string
}

type AlloffInventoryDAO struct {
	AlloffSize AlloffSizeDAO
	Quantity   int
}

type ProductDAO struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ProductInfo    *ProductMetaInfoDAO
	ProductGroupID string
	ExhibitionID   string
	Weight         int
	IsNotSale      bool // Exhibition이 종료되거나 혹은 backoffice에서 삭제했을때
	Created        time.Time
	Updated        time.Time
}
