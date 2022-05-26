package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type SizeMappingPolicyDAO struct {
	ID                primitive.ObjectID `bson:"_id, omitempty"`
	AlloffSize        *AlloffSizeDAO
	AlloffCategory    *AlloffCategoryDAO
	Sizes             []string
	AlloffProductType []AlloffProductType
}
