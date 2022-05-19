package product

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpecialProductType string

const (
	SPECIAL_PRODUCTS     SpecialProductType = "SPECIAL_PRODUCTS"
	NOT_SPECIAL_PRODUCTS SpecialProductType = "NOT_SPECIAL_PRODUCTS"
	ALL_PRODUCTS         SpecialProductType = "ALL_PRODUCTS"
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

type ProductListInput struct {
	Offset                    int
	Limit                     int
	BrandID                   string
	CategoryID                string
	AlloffCategoryID          string
	ExhibitionID              string
	ProductGroupID            string
	Modulename                string
	Keyword                   string
	IncludeClassifiedType     CategoryClassifiedType
	IncludeSpecialProductType SpecialProductType
	PriceRanges               []PriceRangeType
	PriceSorting              PriceSortingType
}

func (input *ProductListInput) BuildFilter() (bson.M, error) {
	filter := bson.M{"removed": false}
	if input.BrandID != "" {
		brandObjID, _ := primitive.ObjectIDFromHex(input.BrandID)
		filter["productinfo.brand._id"] = brandObjID
	}

	if input.CategoryID != "" {
		categoryObjID, _ := primitive.ObjectIDFromHex(input.CategoryID)
		filter["productinfo.category._id"] = categoryObjID
	}

	if input.AlloffCategoryID != "" {
		alloffcat, err := ioc.Repo.AlloffCategories.Get(input.AlloffCategoryID)
		if err == nil {
			if alloffcat.Level == 1 {
				filter["alloffcategories.first.keyname"] = alloffcat.KeyName
			} else if alloffcat.Level == 2 {
				filter["alloffcategories.second.keyname"] = alloffcat.KeyName
			}
		}
	}

	if input.ExhibitionID != "" {
		filter["exhibitionid"] = input.ExhibitionID

	}
	if input.ProductGroupID != "" {
		filter["productgroupid"] = input.ProductGroupID
	}

	if input.Modulename != "" {
		filter["productinfo.source.crawlmodulename"] = input.Modulename
	}

	if input.Keyword != "" {
		filter["alloffname"] = primitive.Regex{Pattern: input.Keyword, Options: "i"}
	}

	if input.IncludeClassifiedType != NO_MATTER_CLASSIFIED {
		if input.IncludeClassifiedType == CLASSIFIED_DONE {
			filter["alloffcategories.done"] = true
		} else {
			filter["alloffcategories.done"] = false
		}
	}

	if input.IncludeSpecialProductType != ALL_PRODUCTS {
		if input.IncludeSpecialProductType == SPECIAL_PRODUCTS {
			filter["isspecial"] = true
		} else {
			filter["isspecial"] = false
		}
	}

	priceQueryRanges := []bson.M{}
	for _, priceRange := range input.PriceRanges {
		if priceRange == "30" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"discountrate": bson.M{"$lt": 30}},
				bson.M{"discountrate": bson.M{"$gte": 0}},
			}})
		}
		if priceRange == "50" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"discountrate": bson.M{"$lt": 50}},
				bson.M{"discountrate": bson.M{"$gte": 30}},
			}})
		}
		if priceRange == "70" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"discountrate": bson.M{"$lt": 70}},
				bson.M{"discountrate": bson.M{"$gte": 50}},
			}})
		}
		if priceRange == "100" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"discountrate": bson.M{"$gte": 70}},
			}})
		}
	}

	if len(priceQueryRanges) > 0 {
		filter["$or"] = priceQueryRanges
	}

	return filter, nil
}

func (input *ProductListInput) BuildSorting() (bson.D, error) {
	options := bson.D{{Key: "soldout", Value: 1}, {Key: "score.isnewlycrawled", Value: -1}, {Key: "_id", Value: 1}, {Key: "score.totalscore", Value: -1}}
	if input.PriceSorting == PRICE_ASCENDING {
		options = bson.D{{Key: "soldout", Value: 1}, {Key: "discountedprice", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == PRICE_DESCENDING {
		options = bson.D{{Key: "soldout", Value: 1}, {Key: "discountedprice", Value: -1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == DISCOUNTRATE_ASCENDING {
		options = bson.D{{Key: "soldout", Value: 1}, {Key: "discountrate", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == DISCOUNTRATE_DESCENDING {
		options = bson.D{{Key: "soldout", Value: 1}, {Key: "discountrate", Value: -1}, {Key: "_id", Value: 1}}
	}

	return options, nil
}

func Listing(input ProductListInput) ([]*domain.ProductDAO, int, error) {
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

	products, cnt, err := ioc.Repo.Products.List(input.Offset, input.Limit, filter, sortingOptions)
	if err != nil {
		return nil, cnt, err
	}

	return products, cnt, nil
}
