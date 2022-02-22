package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/product"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) HomeTabItems(ctx context.Context, onlyLive bool, offset *int, limit *int) ([]*model.HomeTabItem, error) {
	if onlyLive {
		tabItems, _, err := ioc.Repo.HomeTabItems.List(0, 100, onlyLive)
		if err != nil {
			return nil, err
		}

		retItems := []*model.HomeTabItem{}
		for _, item := range tabItems {
			retItems = append(retItems, mapper.MapHomeTabItem(item))
		}
		return retItems, nil
	}

	offsetParam := 0
	limitParam := 100
	if offset != nil {
		offsetParam = *offset
	}
	if limit != nil {
		limitParam = *limit
	}
	tabItems, _, err := ioc.Repo.HomeTabItems.List(offsetParam, limitParam, onlyLive)
	if err != nil {
		return nil, err
	}

	retItems := []*model.HomeTabItem{}
	for _, item := range tabItems {
		retItems = append(retItems, mapper.MapHomeTabItem(item))
	}
	return retItems, nil
}

func (r *queryResolver) BestProducts(ctx context.Context, offset int, limit int, alloffCategoryID string, brief bool) ([]*model.Product, error) {
	bestproductDao, err := ioc.Repo.BestProducts.GetLatest(alloffCategoryID)
	if err != nil {
		log.Println("Err occured in get latest best products", err)
		return nil, err
	}

	return mapper.MapBestProducts(bestproductDao, brief, offset, limit), nil
}

func (r *queryResolver) BestBrands(ctx context.Context, offset int, limit int) ([]*model.Brand, error) {
	brandDaos, _, err := ioc.Repo.Brands.List(offset, limit, bson.M{"onpopular": true}, nil)
	if err != nil {
		return nil, err
	}

	brands := []*model.Brand{}
	for _, brandDao := range brandDaos {
		brands = append(brands, mapper.MapBrandDaoToBrand(brandDao, false))
	}

	return brands, nil
}

func (r *queryResolver) BargainProducts(ctx context.Context, offset int, limit int, alloffCategoryID string, brief bool) ([]*model.Product, error) {
	if alloffCategoryID == "" {
		productDaos, _, err := product.ProductsListing(offset, limit, "", "", "", []string{"100"})
		if err != nil {
			return nil, err
		}
		pds := []*model.Product{}
		for _, productDao := range productDaos {
			pds = append(pds, mapper.MapProductDaoToProduct(productDao))
		}
		return pds, nil
	}

	productDaos, _, err := product.AlloffCategoryProductsListing(offset, limit, nil, alloffCategoryID, "", []string{"100"})
	if err != nil {
		return nil, err
	}

	pds := []*model.Product{}
	for _, productDao := range productDaos {
		pds = append(pds, mapper.MapProductDaoToProduct(productDao))
	}
	return pds, nil
}

func (r *queryResolver) TopBanners(ctx context.Context) ([]*model.TopBanner, error) {
	offset, limit := 0, 100
	onlyLive := true
	bannerDaos, _, err := ioc.Repo.TopBanners.List(offset, limit, onlyLive)
	if err != nil {
		return nil, err
	}

	banners := []*model.TopBanner{}
	for _, bannerDao := range bannerDaos {
		banners = append(banners, mapper.MapTopBanner(bannerDao))
	}

	return banners, nil
}
