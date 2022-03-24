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
	priceSorting := ""
	var priceRange []string
	for _, sorting := range input.Sorting {
		if sorting == model.SortingTypePriceAscending {
			priceSorting = "ascending"
		} else if sorting == model.SortingTypePriceDescending {
			priceSorting = "descending"
		} else if sorting == model.SortingTypeDiscountrateAscending {
			priceSorting = "discountrateAescending"
		} else if sorting == model.SortingTypeDiscountrateDescending {
			priceSorting = "discountrateDescending"
		} else {
			if sorting == model.SortingTypeDiscount0_30 {
				priceRange = append(priceRange, "30")
			} else if sorting == model.SortingTypeDiscount30_50 {
				priceRange = append(priceRange, "50")
			} else if sorting == model.SortingTypeDiscount50_70 {
				priceRange = append(priceRange, "70")
			} else {
				priceRange = append(priceRange, "100")
			}
		}
	}

	pdDaos, cnt, err := product.ProductsSearchListing(input.Offset, input.Limit, false, "", "", "", "", input.Keyword, priceSorting, priceRange)
	if err != nil {
		return nil, err
	}

	products := []*model.Product{}
	for _, pd := range pdDaos {
		products = append(products, mapper.MapProductDaoToProduct(pd))
	}

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

	return mapper.MapProductDaoToProduct(pdDao), nil
}

func (r *queryResolver) Products(ctx context.Context, input model.ProductsInput) (*model.ProductsOutput, error) {
	var productDaos []*domain.ProductDAO

	priceSorting := ""
	var priceRange []string
	for _, sorting := range input.Sorting {
		if sorting == model.SortingTypePriceAscending {
			priceSorting = "ascending"
		} else if sorting == model.SortingTypePriceDescending {
			priceSorting = "descending"
		} else if sorting == model.SortingTypeDiscountrateAscending {
			priceSorting = "discountrateAescending"
		} else if sorting == model.SortingTypeDiscountrateDescending {
			priceSorting = "discountrateDescending"
		} else {
			if sorting == model.SortingTypeDiscount0_30 {
				priceRange = append(priceRange, "30")
			} else if sorting == model.SortingTypeDiscount30_50 {
				priceRange = append(priceRange, "50")
			} else if sorting == model.SortingTypeDiscount50_70 {
				priceRange = append(priceRange, "70")
			} else {
				priceRange = append(priceRange, "100")
			}
		}
	}

	totalCount := 0

	brandDao, err := ioc.Repo.Brands.Get(*input.Brand)
	if err != nil || !brandDao.IsOpenBrand() {
		return &model.ProductsOutput{
			Products:   nil,
			Offset:     input.Offset,
			Limit:      input.Limit,
			TotalCount: totalCount,
		}, err
	}

	if input.Category == nil {
		productDaos, totalCount, _ = product.ProductsListing(input.Offset, input.Limit, *input.Brand, "", priceSorting, priceRange)
	} else {
		if brandDao.UseAlloffCategory {
			productDaos, totalCount, _ = product.AlloffCategoryProductsListing(input.Offset, input.Limit, []string{brandDao.KeyName}, *input.Category, priceSorting, priceRange)
		} else {
			productDaos, totalCount, _ = product.ProductsListing(input.Offset, input.Limit, *input.Brand, *input.Category, priceSorting, priceRange)
		}
	}

	var products []*model.Product
	for _, productDao := range productDaos {
		newProd := mapper.MapProductDaoToProduct(productDao)
		products = append(products, newProd)
	}

	return &model.ProductsOutput{
		Products:   products,
		Offset:     input.Offset,
		Limit:      input.Limit,
		TotalCount: totalCount,
	}, nil
}

func (r *queryResolver) AlloffCategoryProducts(ctx context.Context, input model.AlloffCategoryProductsInput) (*model.AlloffCategoryProducts, error) {
	priceSorting := ""
	var priceRange []string
	for _, sorting := range input.Sorting {
		if sorting == model.SortingTypePriceAscending {
			priceSorting = "ascending"
		} else if sorting == model.SortingTypePriceDescending {
			priceSorting = "descending"
		} else {
			if sorting == model.SortingTypeDiscount0_30 {
				priceRange = append(priceRange, "30")
			} else if sorting == model.SortingTypeDiscount30_50 {
				priceRange = append(priceRange, "50")
			} else if sorting == model.SortingTypeDiscount50_70 {
				priceRange = append(priceRange, "70")
			} else {
				priceRange = append(priceRange, "100")
			}
		}
	}

	totalCount := 0

	productDaos, totalCount, err := product.AlloffCategoryProductsListing(input.Offset, input.Limit, input.BrandIds, input.AlloffcategoryID, priceSorting, priceRange)
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
