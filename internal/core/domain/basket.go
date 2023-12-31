package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Basket struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Items        []*BasketItem
	ProductPrice int
}

type BasketItem struct {
	Product      *ProductDAO
	ExhibitionID string
	Size         string
	Quantity     int
	Errors       []string
}
