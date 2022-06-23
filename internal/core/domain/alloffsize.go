package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AlloffSizeDAO struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	AlloffCategory *AlloffCategoryDAO
	AlloffSizeName string
	Sizes          []string
	ProductType    []AlloffProductType
}
