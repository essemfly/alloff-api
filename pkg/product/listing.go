package product

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Classified string

const (
	CLASSIFIED_DONE      Classified = "CLASSIFIED_DONE"
	NOT_CLASSIFIED       Classified = "NOT_CLASSIFIED"
	NO_MATTER_CLASSIFIED Classified = "NO_MATTER_CLASSIFIED"
)

func ProductsSearchListing(offset, limit int, classifiedType Classified, moduleName, brandID, categoryID, alloffCategoryID, keyword string, priceSorting string, priceRanges []string) ([]*domain.ProductDAO, int, error) {
	filter := bson.M{"removed": false}

	if classifiedType != NO_MATTER_CLASSIFIED {
		if classifiedType == CLASSIFIED_DONE {
			filter["alloffcategories.done"] = true
		} else {
			filter["alloffcategories.done"] = false
		}
	}
	if moduleName != "" {
		filter["productinfo.source.crawlmodulename"] = moduleName
	}
	brandObjID, _ := primitive.ObjectIDFromHex(brandID)
	categoryObjID, _ := primitive.ObjectIDFromHex(categoryID)

	if brandID != "" {
		filter["productinfo.brand._id"] = brandObjID
		if categoryID != "" {
			filter["productinfo.category._id"] = categoryObjID
		}
	}

	if alloffCategoryID != "" {
		alloffcat, err := ioc.Repo.AlloffCategories.Get(alloffCategoryID)
		if err == nil {
			if alloffcat.Level == 1 {
				filter["alloffcategories.first.keyname"] = alloffcat.KeyName
			} else if alloffcat.Level == 2 {
				filter["alloffcategories.second.keyname"] = alloffcat.KeyName
			}
		}
	}

	filter["alloffname"] = primitive.Regex{Pattern: keyword, Options: "i"}

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

	sortingOptions := bson.D{{Key: "soldout", Value: 1}, {Key: "score.isnewlycrawled", Value: -1}, {Key: "_id", Value: 1}, {Key: "score.totalscore", Value: -1}}
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

// (Future) Mongodb에 종속적인 함수: bson이 사용되었다.
func ProductsListing(offset, limit int, brandID, categoryID string, priceSorting string, priceRanges []string) ([]*domain.ProductDAO, int, error) {
	filter := bson.M{"removed": false}

	brandObjID, _ := primitive.ObjectIDFromHex(brandID)
	categoryObjID, _ := primitive.ObjectIDFromHex(categoryID)

	if brandID != "" {
		filter["productinfo.brand._id"] = brandObjID
		if categoryID != "" {
			filter["productinfo.category._id"] = categoryObjID
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

	sortingOptions := bson.D{{Key: "soldout", Value: 1}, {Key: "score.isnewlycrawled", Value: -1}, {Key: "score.totalscore", Value: -1}}
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

// TODO: 위의 함수와 합쳐질 필요가 있다.
func AlloffCategoryProductsListing(offset, limit int, brandKeynames []string, alloffCategoryID string, priceSorting string, priceRanges []string) ([]*domain.ProductDAO, int, error) {
	filter := bson.M{"removed": false, "alloffcategories.done": true}
	if len(brandKeynames) > 0 {
		filter["productinfo.brand.keyname"] = bson.M{"$in": brandKeynames}
	}

	alloffCat, _ := ioc.Repo.AlloffCategories.Get(alloffCategoryID)
	if alloffCat.Level == 1 {
		filter["alloffcategories.first._id"] = alloffCat.ID
	} else if alloffCat.Level == 2 {
		filter["alloffcategories.second._id"] = alloffCat.ID
	}

	priceQueryRanges := []bson.M{}
	for _, priceRange := range priceRanges {
		if priceRange == "30" {
			priceQueryRanges = append(priceQueryRanges, bson.M{"$and": []interface{}{
				bson.M{"discountrate": bson.M{"$lt": 30}},
				bson.M{"discountrate": bson.M{"$gt": 0}},
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

	sortingOptions := bson.D{{Key: "soldout", Value: 1}, {Key: "score.isnewlycrawled", Value: -1}, {Key: "score.totalscore", Value: -1}}
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

func ProductGroupProductsListing(offset, limit int, productGroupId string, sorting model.BrandTimedealSortingType) ([]*domain.ProductDAO, int, error) {
	filter := bson.M{
		"productgroupid": productGroupId,
	}
	sortingOptions := bson.D{{Key: "soldout", Value: 1}, {Key: "score.isnewlycrawled", Value: -1}, {Key: "score.totalscore", Value: -1}}
	switch sorting {
	case model.BrandTimedealSortingTypeDiscountDescending:
		sortingOptions = bson.D{{Key: "soldout", Value: 1}, {Key: "discountrate", Value: -1}, {Key: "score.totalscore", Value: -1}}
	case model.BrandTimedealSortingTypePriceAscending:
		sortingOptions = bson.D{{Key: "soldout", Value: 1}, {Key: "specialprice", Value: 1}, {Key: "score.totalscore", Value: -1}}
	case model.BrandTimedealSortingTypePriceDescending:
		sortingOptions = bson.D{{Key: "soldout", Value: 1}, {Key: "specialprice", Value: 1}, {Key: "score.totalscore", Value: -1}}
	}

	products, cnt, err := ioc.Repo.Products.List(offset, limit, filter, sortingOptions)
	if err != nil {
		log.Println("Error on listing productGroups products : ", err)
		return nil, 0, err
	}
	return products, cnt, nil

}
