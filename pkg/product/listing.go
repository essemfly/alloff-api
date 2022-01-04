package pkg

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
)

// (Future) Mongodb에 종속적인 함수: bson이 사용되었다.
func ProductsListing(offset, limit int, brandID, categoryID, alloffCategoryID string, priceSorting string, priceRanges []string) ([]*domain.ProductDAO, int, error) {
	filter := bson.M{"removed": false}

	if brandID != "" {
		filter["brand._id"] = brandID
		if categoryID != "" {
			filter["category._id"] = categoryID
		}
	} else {
		alloffCat, _ := ioc.Repo.AlloffCategories.Get(alloffCategoryID)
		filter["alloffcategories.done"] = true
		if alloffCat.Level == 1 {
			filter["alloffcategories.first._id"] = alloffCat.ID
		} else if alloffCat.Level == 2 {
			filter["alloffcategories.second._id"] = alloffCat.ID
		}
	}

	priceQueryRanges := []bson.M{}
	for _, priceRange := range priceRanges {
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

	sortingOptions := bson.D{{Key: "soldout", Value: 1}, {Key: "isnewlycrawled", Value: -1}, {Key: "_id", Value: 1}}
	if priceSorting == "ascending" {
		sortingOptions = bson.D{{Key: "soldout", Value: 1}, {Key: "discountedprice", Value: 1}, {Key: "_id", Value: 1}}
	} else if priceSorting == "descending" {
		sortingOptions = bson.D{{Key: "soldout", Value: 1}, {Key: "discountedprice", Value: -1}, {Key: "_id", Value: 1}}
	} else if priceSorting == "discountrateAescending" {
		sortingOptions = bson.D{{Key: "soldout", Value: 1}, {Key: "discountrate", Value: 1}, {Key: "_id", Value: 1}}
	} else if priceSorting == "discountrateDescending" {
		sortingOptions = bson.D{{Key: "soldout", Value: 1}, {Key: "discountrate", Value: -1}, {Key: "_id", Value: 1}}
	}

	products, cnt, err := ioc.Repo.Products.List(offset, limit, filter, sortingOptions)
	if err != nil {
		return nil, cnt, err
	}

	return products, cnt, nil
}
