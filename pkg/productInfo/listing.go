package productinfo

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// For Backoffice Servers
type ProductInfoListInput struct {
	Offset                 int
	Limit                  int
	BrandID                string
	AlloffCategoryID       string
	AlloffSizeIDs          []string
	ProductTypes           []domain.AlloffProductType
	CategoryClassifierName string
	Modulename             string
	Keyword                string
	ProductUrl             string
	IncludeClassifiedType  domain.CategoryClassifiedType
	PriceRanges            []domain.PriceRangeType
	PriceSorting           domain.PriceSortingType
	OnlyCategoryClassified bool
}

func (input *ProductInfoListInput) BuildFilter() (bson.M, error) {
	filter := bson.M{"isremoved": false}
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

	if input.ProductUrl != "" {
		filter["producturl"] = primitive.Regex{Pattern: input.ProductUrl, Options: "i"}
	}

	if len(input.ProductTypes) > 0 {
		filter["producttype"] = bson.M{"$all": input.ProductTypes}
	}

	if input.IncludeClassifiedType != "" && input.IncludeClassifiedType != domain.NO_MATTER_CLASSIFIED {
		if input.IncludeClassifiedType == domain.CLASSIFIED_DONE {
			filter["alloffcategory.done"] = true
		} else {
			filter["alloffcategory.done"] = false
		}
	}

	if input.OnlyCategoryClassified {
		filter["iscategoryclassified"] = true
	}

	if len(input.AlloffSizeIDs) > 0 {
		alloffSizeIds := []primitive.ObjectID{}
		for _, id := range input.AlloffSizeIDs {
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				continue
			}
			alloffSizeIds = append(alloffSizeIds, oid)
		}
		filter["inventory.alloffsizes._id"] = bson.M{"$all": alloffSizeIds}
	}

	andQuries := []bson.M{}
	keyWordQueries := []bson.M{}
	if input.Keyword != "" {
		keyWordQueries = append(keyWordQueries, bson.M{"originalname": primitive.Regex{Pattern: input.Keyword, Options: "i"}})
		keyWordQueries = append(keyWordQueries, bson.M{"alloffname": primitive.Regex{Pattern: input.Keyword, Options: "i"}})
		keyWordQueries = append(keyWordQueries, bson.M{"source.crawlmodulename": primitive.Regex{Pattern: input.Keyword, Options: "i"}})
	}
	if len(keyWordQueries) > 0 {
		andQuries = append(andQuries, bson.M{"$or": keyWordQueries})
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
		andQuries = append(andQuries, bson.M{"$or": priceQueryRanges})
	}

	if len(andQuries) > 0 {
		filter["$and"] = andQuries
	}

	return filter, nil
}

func (input *ProductInfoListInput) BuildSorting() (bson.D, error) {
	options := bson.D{{Key: "issoldout", Value: 1}, {Key: "_id", Value: -1}}
	if input.PriceSorting == domain.PRICE_ASCENDING {
		options = bson.D{{Key: "issoldout", Value: 1}, {Key: "price.currentprice", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == domain.PRICE_DESCENDING {
		options = bson.D{{Key: "issoldout", Value: 1}, {Key: "price.currentprice", Value: -1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == domain.DISCOUNTRATE_ASCENDING {
		options = bson.D{{Key: "issoldout", Value: 1}, {Key: "price.discountrate", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == domain.DISCOUNTRATE_DESCENDING {
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
