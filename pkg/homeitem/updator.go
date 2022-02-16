package homeitem

import (
	"log"

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
			numPassedPgsToShow := 15
			pgs, err := ioc.Repo.ProductGroups.List(numPassedPgsToShow)
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
				var sortingOptions []string
				for _, sortingType := range item.Sorting {
					if sortingType == model.SortingTypeDiscount70_100 {
						sortingOptions = append(sortingOptions, "100")
					} else if sortingType == model.SortingTypeDiscount50_70 {
						sortingOptions = append(sortingOptions, "70")
					} else if sortingType == model.SortingTypeDiscount30_50 {
						sortingOptions = append(sortingOptions, "30")
					} else if sortingType == model.SortingTypeDiscount0_30 {
						sortingOptions = append(sortingOptions, "0")
					}
				}
				products, _, err := product.AlloffCategoryProductsListing(0, numProductsToShow, nil, item.TargetID, "", sortingOptions)
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
