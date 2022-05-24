package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	productDao, err := ioc.Repo.Products.Get(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapProduct(productDao), nil
}

func (r *queryResolver) Products(ctx context.Context, input model.ProductsInput) (*model.ProductsOutput, error) {
	priceRanges, priceSorting := mapper.MapProductSortingAndRanges(input.Sorting)

	pdType := domain.Female
	if input.ProductType == model.AlloffProductTypeKids {
		pdType = domain.Kids
	} else if input.ProductType == model.AlloffProductTypeMale {
		pdType = domain.Male
	} else if input.ProductType == model.AlloffProductTypeSports {
		pdType = domain.Sports
	}

	query := product.ProductListInput{
		Offset:        input.Offset,
		Limit:         input.Limit,
		ProductType:   pdType,
		ExhibitionID:  input.ExhibitionID,
		BrandIDs:      input.BrandIds,
		AlloffSizeIDs: input.AlloffSizeIds,
		PriceRanges:   priceRanges,
		PriceSorting:  priceSorting,
	}

	if input.AlloffCategoryID != nil {
		query.AlloffCategoryID = *input.AlloffCategoryID
	}

	pdDaos, cnt, err := product.ListProducts(query)
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for _, productDao := range pdDaos {
		newProd := mapper.MapProduct(productDao)
		products = append(products, newProd)
	}

	return &model.ProductsOutput{
		TotalCount:   cnt,
		Offset:       input.Offset,
		Limit:        input.Limit,
		ExhibitionID: input.ExhibitionID,
		Products:     products,
	}, nil
}
