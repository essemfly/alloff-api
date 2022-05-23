package product

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// For API Servers
type ProductListInput struct {
	Offset           int
	Limit            int
	AlloffCategoryID string
	ExhibitionID     string
	ProductGroupID   string
	Modulename       string
	AlloffSizeIDs    []string
	BrandIDs         []string
	PriceRanges      []PriceRangeType
	PriceSorting     PriceSortingType
}

func (input *ProductListInput) BuildFilter() (bson.M, error) {
	filter := bson.M{"isnotsale": false}

	if input.AlloffCategoryID != "" {
		alloffcat, err := ioc.Repo.AlloffCategories.Get(input.AlloffCategoryID)
		if err == nil {
			if alloffcat.Level == 1 {
				filter["productinfo.alloffcategory.first.keyname"] = alloffcat.KeyName
			} else if alloffcat.Level == 2 {
				filter["productinfo.alloffcategory.second.keyname"] = alloffcat.KeyName
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

	if len(input.BrandIDs) > 0 {
		query := []bson.M{}
		for _, id := range input.BrandIDs {
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				continue
			}
			query = append(query, bson.M{"productinfo.brand._id": oid})
		}
		filter["$or"] = query
	}

	priceQueryRanges := []bson.M{}
	for _, priceRange := range input.PriceRanges {
		if priceRange == "30" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"productinfo.price.discountrate": bson.M{"$lt": 30}},
				bson.M{"productinfo.price.discountrate": bson.M{"$gte": 0}},
			}})
		}
		if priceRange == "50" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"productinfo.price.discountrate": bson.M{"$lt": 50}},
				bson.M{"productinfo.price.discountrate": bson.M{"$gte": 30}},
			}})
		}
		if priceRange == "70" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"productinfo.price.discountrate": bson.M{"$lt": 70}},
				bson.M{"productinfo.price.discountrate": bson.M{"$gte": 50}},
			}})
		}
		if priceRange == "100" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"productinfo.price.discountrate": bson.M{"$gte": 70}},
			}})
		}
	}

	if len(priceQueryRanges) > 0 {
		filter["$or"] = priceQueryRanges
	}

	return filter, nil
}

func (input *ProductListInput) BuildSorting() (bson.D, error) {
	options := bson.D{{Key: "productinfo.issoldout", Value: 1}}
	if input.PriceSorting == PRICE_ASCENDING {
		options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.price.currentprice", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == PRICE_DESCENDING {
		options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.price.currentprice", Value: -1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == DISCOUNTRATE_ASCENDING {
		options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.price.discountrate", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == DISCOUNTRATE_DESCENDING {
		options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.price.discountrate", Value: -1}, {Key: "_id", Value: 1}}
	}

	return options, nil
}

func ListProducts(input ProductListInput) ([]*domain.ProductDAO, int, error) {
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
