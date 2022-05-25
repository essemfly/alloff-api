package product

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryClassifiedType string

const (
	CLASSIFIED_DONE      CategoryClassifiedType = "CLASSIFIED_DONE"
	NOT_CLASSIFIED       CategoryClassifiedType = "NOT_CLASSIFIED"
	NO_MATTER_CLASSIFIED CategoryClassifiedType = "NO_MATTER_CLASSIFIED"
)

type PriceSortingType string

const (
	PRICE_ASCENDING         PriceSortingType = "ascending"
	PRICE_DESCENDING        PriceSortingType = "descending"
	DISCOUNTRATE_ASCENDING  PriceSortingType = "discountrateAscending"
	DISCOUNTRATE_DESCENDING PriceSortingType = "discountrateDescending"
)

type PriceRangeType string

const (
	PRICE_RANGE_30  PriceRangeType = "30"
	PRICE_RANGE_50  PriceRangeType = "50"
	PRICE_RANGE_70  PriceRangeType = "70"
	PRICE_RANGE_100 PriceRangeType = "100"
)

// For Backoffice Servers
type ProductInfoListInput struct {
	Offset                 int
	Limit                  int
	BrandID                string
	AlloffCategoryID       string
	AlloffSizeIDs          []string
	CategoryClassifierName string
	Modulename             string
	Keyword                string
	IncludeClassifiedType  CategoryClassifiedType
	PriceRanges            []PriceRangeType
	PriceSorting           PriceSortingType
	OnlyCategoryClassified bool
}

func (input *ProductInfoListInput) BuildFilter() (bson.M, error) {
	filter := bson.M{}
	if input.BrandID != "" {
		brandObjID, _ := primitive.ObjectIDFromHex(input.BrandID)
		filter["brand._id"] = brandObjID
	}

	if input.AlloffCategoryID != "" {
		alloffcat, err := ioc.Repo.AlloffCategories.Get(input.AlloffCategoryID)
		if err == nil {
			if alloffcat.Level == 1 {
				filter["alloffcategory.first.keyname"] = alloffcat.KeyName
			} else if alloffcat.Level == 2 {
				filter["alloffcategory.second.keyname"] = alloffcat.KeyName
			}
		}
	}

	if input.Modulename != "" {
		filter["source.crawlmodulename"] = input.Modulename
	}

	if input.Keyword != "" {
		filter["alloffname"] = primitive.Regex{Pattern: input.Keyword, Options: "i"}
	}

	if input.IncludeClassifiedType != NO_MATTER_CLASSIFIED {
		if input.IncludeClassifiedType == CLASSIFIED_DONE {
			filter["alloffcategory.done"] = true
		} else {
			filter["alloffcategory.done"] = false
		}
	}

	if input.OnlyCategoryClassified {
		filter["iscategoryclassified"] = true
	}

	if len(input.AlloffSizeIDs) > 0 {
		query := []bson.M{}
		for _, id := range input.AlloffSizeIDs {
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				continue
			}
			query = append(query, bson.M{"alloffinventory.alloffsize._id": oid})
		}
		filter["$or"] = query
	}

	priceQueryRanges := []bson.M{}
	for _, priceRange := range input.PriceRanges {
		if priceRange == "30" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"price.discountrate": bson.M{"$lt": 30}},
				bson.M{"price.discountrate": bson.M{"$gte": 0}},
			}})
		}
		if priceRange == "50" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"price.discountrate": bson.M{"$lt": 50}},
				bson.M{"price.discountrate": bson.M{"$gte": 30}},
			}})
		}
		if priceRange == "70" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"price.discountrate": bson.M{"$lt": 70}},
				bson.M{"price.discountrate": bson.M{"$gte": 50}},
			}})
		}
		if priceRange == "100" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"price.discountrate": bson.M{"$gte": 70}},
			}})
		}
	}

	if len(priceQueryRanges) > 0 {
		filter["$or"] = priceQueryRanges
	}

	return filter, nil
}

func (input *ProductInfoListInput) BuildSorting() (bson.D, error) {
	options := bson.D{{Key: "issoldout", Value: 1}}
	if input.PriceSorting == PRICE_ASCENDING {
		options = bson.D{{Key: "issoldout", Value: 1}, {Key: "price.currentprice", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == PRICE_DESCENDING {
		options = bson.D{{Key: "issoldout", Value: 1}, {Key: "price.currentprice", Value: -1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == DISCOUNTRATE_ASCENDING {
		options = bson.D{{Key: "issoldout", Value: 1}, {Key: "price.discountrate", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == DISCOUNTRATE_DESCENDING {
		options = bson.D{{Key: "issoldout", Value: 1}, {Key: "price.discountrate", Value: -1}, {Key: "_id", Value: 1}}
	}

	return options, nil
}

func ListProductInfos(input ProductInfoListInput) ([]*domain.ProductMetaInfoDAO, int, error) {
	filter, err := input.BuildFilter()
	if err != nil {
		log.Println("Error in getting products filter ", err)
		return nil, 0, err
	}
	sortingOptions, err := input.BuildSorting()
	if err != nil {
		log.Println("Error in getting products sorting ", err)
		return nil, 0, err
	}

	products, cnt, err := ioc.Repo.ProductMetaInfos.List(input.Offset, input.Limit, filter, sortingOptions)
	if err != nil {
		return nil, cnt, err
	}

	return products, cnt, nil
}