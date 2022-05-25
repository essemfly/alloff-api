package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type SizeMappingPolicyDAO struct {
	ID                primitive.ObjectID
	AlloffSize        *AlloffSizeDAO
	AlloffCategory    *AlloffCategoryDAO
	Sizes             []string
	AlloffProductType []AlloffProductType
}
