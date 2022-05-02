package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func (r *queryResolver) HomeTabItems(ctx context.Context, onlyLive bool, offset *int, limit *int) ([]*model.HomeTabItem, error) {
	offsetParam, limitParam := 0, 100
	if offset != nil {
		offsetParam = *offset
	}
	if limit != nil {
		limitParam = *limit
	}

	if onlyLive {
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

func (r *queryResolver) BestProducts(ctx context.Context, offset int, limit int, alloffCategoryID string, brief bool) (*model.ProductsResult, error) {
	bestproductDao, err := ioc.Repo.BestProducts.GetLatest(alloffCategoryID)
	if err != nil {
		log.Println("Err occured in get latest best products", err)
		return nil, err
	}

	lastUpdatedKST := bestproductDao.CreatedAt.Add(9 * time.Hour).Format("2006-01-02 15:04:05")
	flooredLastUpdated := lastUpdatedKST[0:len(lastUpdatedKST)-5] + "00:00"

	return &model.ProductsResult{
		Products:    mapper.MapBestProducts(bestproductDao, brief, offset, limit),
		LastUpdated: flooredLastUpdated,
	}, nil
}

func (r *queryResolver) BestBrands(ctx context.Context, offset int, limit int) (*model.BrandsResult, error) {
	bestBrandsDao, err := ioc.Repo.BestBrands.GetLatest()
	if err != nil {
		log.Println("Err occurred in get latest best products : ", err)
		return nil, err
	}

	brands := []*model.Brand{}
	for _, brandDao := range bestBrandsDao.Brands {
		brands = append(brands, mapper.MapBrandDaoToBrand(brandDao, false))
	}

	lastUpdatedKST := bestBrandsDao.CreatedAt.Add(9 * time.Hour).Format("2006-01-02 15:04:05")
	flooredLastUpdated := lastUpdatedKST[0:len(lastUpdatedKST)-5] + "00:00"

	return &model.BrandsResult{
		Brands:      brands,
		LastUpdated: flooredLastUpdated,
	}, nil
}

func (r *queryResolver) BargainProducts(ctx context.Context, offset int, limit int, alloffCategoryID string, brief bool) ([]*model.Product, error) {
	if alloffCategoryID == "" {
		productDaos, _, err := product.ProductsListing(offset, limit, "", "", "", "", []string{"100"})
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
