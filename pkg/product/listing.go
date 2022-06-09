package product

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// For API Servers
type ProductListInput struct {
	Offset           int
	Limit            int
	ProductType      domain.AlloffProductType
	ProductGroupID   string
	ExhibitionID     string
	AlloffCategoryID string
	BrandIDs         []string
	AlloffSizeIDs    []string
	PriceRanges      []domain.PriceRangeType
	PriceSorting     domain.PriceSortingType
}

func (input *ProductListInput) BuildFilter() (bson.M, error) {
	filter := bson.M{"isnotsale": false}

	if input.ProductType != "" {
		filter["productinfo.producttype"] = input.ProductType
	}

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

	andQuery := []bson.M{}
	if len(input.AlloffSizeIDs) > 0 {
		query := []bson.M{}
		for _, id := range input.AlloffSizeIDs {
			oid, err := primitive.ObjectIDFromHex(id)
			if err != nil {
				continue
			}
			query = append(query, bson.M{"productinfo.inventory.alloffsizes._id": oid})
		}
		andQuery = append(andQuery, bson.M{"$or": query})
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
		andQuery = append(andQuery, bson.M{"$or": query})
	}

	if len(andQuery) > 0 {
		filter["$and"] = andQuery
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
	if input.PriceSorting == domain.PRICE_ASCENDING {
		options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.price.currentprice", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == domain.PRICE_DESCENDING {
		options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.price.currentprice", Value: -1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == domain.DISCOUNTRATE_ASCENDING {
		options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.price.discountrate", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == domain.DISCOUNTRATE_DESCENDING {
		options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.price.discountrate", Value: -1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == domain.INVENTORY_ASCENDING {
		options = bson.D{{Key: "totalQuantity", Value: 1}, {Key: "_id", Value: 1}}
		// options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.inventory.quantity", Value: 1}, {Key: "_id", Value: 1}}
	} else if input.PriceSorting == domain.INVENTORY_DESCENDING {
		options = bson.D{{Key: "totalQuantity", Value: -1}, {Key: "_id", Value: 1}}
		// options = bson.D{{Key: "productinfo.issoldout", Value: 1}, {Key: "productinfo.inventory.quantity", Value: -1}, {Key: "_id", Value: 1}}
	}

	return options, nil
}

func ListProducts(input ProductListInput) ([]*domain.ProductDAO, int, error) {
	filter, err := input.BuildFilter()
	if err != nil {
		config.Logger.Error("Error in getting products filter ", zap.Error(err))
		return nil, 0, err
	}

	sortingOptions, err := input.BuildSorting()
	if err != nil {
		config.Logger.Error("Error in getting products sorting options ", zap.Error(err))
		return nil, 0, err
	}

	pipelines := []interface{}{
		bson.M{"$match": filter},
		bson.M{"$addFields": bson.M{
			"totalquantity": bson.M{
				"$add": []bson.M{
					{"$sum": "$productinfo.inventory.quantity"},
				},
			},
		}},
		bson.M{"$sort": sortingOptions},
		bson.M{"$limit": input.Limit + input.Offset},
		bson.M{"$skip": input.Offset},
	}
	products, cnt, err := ioc.Repo.Products.Aggregate(filter, pipelines)
	if err != nil {
		return nil, cnt, err
	}

	return products, cnt, nil
}
