package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func (r *mutationResolver) LikeProduct(ctx context.Context, input *model.LikeProductInput) (bool, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("ERR000:invalid token")
	}

	return ioc.Repo.LikeProducts.Like(user.ID.Hex(), input.ProductID)
}

func (r *queryResolver) Find(ctx context.Context, input model.ProductQueryInput) (*model.ProductsOutput, error) {
	priceRanges, priceSorting := mapper.MapProductSortingAndRanges(input.Sorting)

	query := product.ProductListInput{
		Offset:                    input.Offset,
		Limit:                     input.Limit,
		IncludeSpecialProductType: product.NOT_SPECIAL_PRODUCTS,
		IncludeClassifiedType:     product.NO_MATTER_CLASSIFIED,
		Keyword:                   input.Keyword,
		PriceRanges:               priceRanges,
		PriceSorting:              priceSorting,
	}

	pdDaos, cnt, err := product.Listing(query)
	if err != nil {
		return nil, err
	}

	products := []*model.Product{}
	for _, pd := range pdDaos {
		products = append(products, mapper.MapProductDaoToProduct(pd))
	}

	go ioc.Repo.SearchLog.Index(input.Keyword)

	return &model.ProductsOutput{
		Products:   products,
		Offset:     input.Offset,
		Limit:      input.Limit,
		TotalCount: cnt,
	}, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	pdDao, err := ioc.Repo.Products.Get(id)
	if err != nil {
		return nil, err
	}

	//go elasticsearch.ProductLogRequest(pdDao, domain.PRODUCT_VIEW)
	go ioc.Repo.ProductLog.Index(pdDao, domain.PRODUCT_VIEW)

	return mapper.MapProductDaoToProduct(pdDao), nil
}

func (r *queryResolver) Products(ctx context.Context, input model.ProductsInput) (*model.ProductsOutput, error) {
	priceRanges, priceSorting := mapper.MapProductSortingAndRanges(input.Sorting)

	query := product.ProductListInput{
		Offset:                    input.Offset,
		Limit:                     input.Limit,
		IncludeSpecialProductType: product.NOT_SPECIAL_PRODUCTS,
		IncludeClassifiedType:     product.NO_MATTER_CLASSIFIED,
		PriceRanges:               priceRanges,
		PriceSorting:              priceSorting,
		AlloffClassifier:          mapper.MapAlloffClassifierModelToDAO(input.AlloffClassifier),
		CategoryClassifierName:    input.CategoryClassifier,
	}

	if input.Brand != nil {
		brandDao, err := ioc.Repo.Brands.Get(*input.Brand)
		if err != nil || !brandDao.IsOpenBrand() {
			return &model.ProductsOutput{
				Products:   nil,
				Offset:     input.Offset,
				Limit:      input.Limit,
				TotalCount: 0,
			}, err
		}

		query.BrandID = *input.Brand
		if input.Category != nil {
			if brandDao.UseAlloffCategory {
				query.AlloffCategoryID = *input.Category
			} else {
				query.CategoryID = *input.Category
			}
		}
	}

	if input.ExhibitionID != nil {
		query.IncludeSpecialProductType = product.ALL_PRODUCTS
		query.ExhibitionID = *input.ExhibitionID
	}

	if input.ProductGroupID != nil {
		query.ProductGroupID = *input.ProductGroupID
	}

	pdDaos, cnt, err := product.Listing(query)
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for _, productDao := range pdDaos {
		newProd := mapper.MapProductDaoToProduct(productDao)
		products = append(products, newProd)
	}

	return &model.ProductsOutput{
		Products:   products,
		Offset:     input.Offset,
		Limit:      input.Limit,
		TotalCount: cnt,
	}, nil
}

func (r *queryResolver) AlloffCategoryProducts(ctx context.Context, input model.AlloffCategoryProductsInput) (*model.AlloffCategoryProducts, error) {
	priceRanges, priceSorting := mapper.MapProductSortingAndRanges(input.Sorting)

	query := product.ProductListInput{
		Offset:                    input.Offset,
		Limit:                     input.Limit,
		AlloffCategoryID:          input.AlloffcategoryID,
		IncludeSpecialProductType: product.NOT_SPECIAL_PRODUCTS,
		IncludeClassifiedType:     product.NO_MATTER_CLASSIFIED,
		PriceRanges:               priceRanges,
		PriceSorting:              priceSorting,
	}

	productDaos, totalCount, err := product.Listing(query)
	if err != nil {
		return nil, err
	}

	brandDaos, err := ioc.Repo.Products.ListDistinctBrands(input.AlloffcategoryID)
	if err != nil {
		return nil, err
	}

	alloffCatDao, err := ioc.Repo.AlloffCategories.Get(input.AlloffcategoryID)
	if alloffCatDao == nil {
		return nil, err
	}

	var products []*model.Product
	for _, productDao := range productDaos {
		products = append(products, mapper.MapProductDaoToProduct(productDao))
	}

	var brands []*model.Brand
	includeCategory := false
	for _, brandDao := range brandDaos {
		brands = append(brands, mapper.MapBrandDaoToBrand(brandDao, includeCategory))
	}

	return &model.AlloffCategoryProducts{
		Alloffcategory: mapper.MapAlloffCatDaoToAlloffCat(alloffCatDao),
		Products:       products,
		AllBrands:      brands,
		TotalCount:     totalCount,
		Offset:         input.Offset,
		Limit:          input.Limit,
		SelectedBrands: input.BrandIds,
	}, nil
}

func (r *queryResolver) Likeproducts(ctx context.Context) ([]*model.LikeProductOutput, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	likes, _ := ioc.Repo.LikeProducts.List(user.ID.Hex())

	var products []*model.LikeProductOutput

	for _, like := range likes {
		newProduct, err := ioc.Repo.Products.Get(like.Productid)
		if err != nil {
			log.Println("like products id not found :" + like.Productid)
			continue
		}
		if like.OldProduct == nil {
			return nil, errors.New("old product is missing")
		}

		oldProduct := mapper.MapProductDaoToProduct(like.OldProduct)
		likeProduct := model.LikeProductOutput{
			OldProduct: oldProduct,
			NewProduct: mapper.MapProductDaoToProduct(newProduct),
		}

		products = append(products, &likeProduct)
	}

	return products, nil
}
