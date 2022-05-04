package homeitem

import (
	"log"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func UpdateHomeItems() {
	log.Println("Running Script: Update HomeItems")
	homeitemDaos, err := ioc.Repo.HomeItems.List()
	if err != nil {
		log.Println("Errors homeitem list: ", err)
	}
	for _, item := range homeitemDaos {
		if item.ItemType == model.HomeItemTypeTimedeal {
			timedeal := domain.PRODUCT_GROUP_TIMEDEAL
			pgs, _, _ := ioc.Repo.ProductGroups.List(0, 10, &timedeal, "")
			if err != nil {
				log.Println("Errors productgroups: ", err)
			}
			item.ProductGroups = pgs
		} else if item.ItemType == model.HomeItemTypeBrand {
			var newBrandItems []*domain.HomeBrandItemDAO
			for _, brandItem := range item.Brands {
				updatedBrand, err := ioc.Repo.Brands.Get(brandItem.Brand.ID.Hex())
				if err != nil {
					log.Println("Erros get brand")
				}
				updatedBrandItem := domain.HomeBrandItemDAO{
					ImgUrl: brandItem.ImgUrl,
					Brand:  updatedBrand,
				}
				newBrandItems = append(newBrandItems, &updatedBrandItem)
			}
			item.Brands = newBrandItems
		} else if item.ItemType == model.HomeItemTypeProduct {
			numProductsToShow := 10
			if item.TargetID != "" {
				priceRanges, priceSorting := mapper.MapProductSortingAndRanges(item.Sorting)

				query := product.ProductListInput{
					Offset:                    0,
					Limit:                     numProductsToShow,
					AlloffCategoryID:          item.TargetID,
					IncludeSpecialProductType: product.NOT_SPECIAL_PRODUCTS,
					IncludeClassifiedType:     product.NO_MATTER_CLASSIFIED,
					PriceRanges:               priceRanges,
					PriceSorting:              priceSorting,
				}
				products, _, err := product.Listing(query)
				if err != nil {
					log.Println("Errors list products")
				}
				item.Products = products
			}
		}
	}
	for _, item := range homeitemDaos {
		err := ioc.Repo.HomeItems.Update(item)
		if err != nil {
			log.Println("Homeitem update", err)
		}
	}

}
