package brand

import (
	"log"
	"strconv"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func UpdateBrandCategory() {
	log.Println("Running Script: Update Brand Category")
	offset, limit := 0, 1000
	brandDaos, totalCount, err := ioc.Repo.Brands.List(offset, limit, false, false, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Total # of brands : " + strconv.Itoa(totalCount))

	for _, brandDao := range brandDaos {
		brandDao.AlloffCategory = ListBrandAlloffCategories(brandDao)
		brandDao.Category = ListBrandCategories(brandDao)
		_, err := ioc.Repo.Brands.Upsert(brandDao)
		if err != nil {
			log.Println("err occured on brand udpate", err)
		}
	}
}

func ListBrandCategories(brandDao *domain.BrandDAO) []*domain.CategoryDAO {
	brandCats := []*domain.CategoryDAO{}
	catDaos, err := ioc.Repo.Categories.List(brandDao.KeyName)
	if err != nil {
		log.Println(err)
	}

	for _, catDao := range catDaos {
		query := product.ProductListInput{
			Offset:                    0,
			Limit:                     1,
			BrandID:                   brandDao.ID.Hex(),
			CategoryID:                catDao.ID.Hex(),
			IncludeSpecialProductType: product.NOT_SPECIAL_PRODUCTS,
			IncludeClassifiedType:     product.NO_MATTER_CLASSIFIED,
		}
		_, pdCount, _ := product.Listing(query)
		if pdCount == 0 {
			continue
		}
		brandCats = append(brandCats, catDao)
	}

	return brandCats
}

func ListBrandAlloffCategories(brandDao *domain.BrandDAO) []*domain.AlloffCategoryDAO {
	brandAlloffCats := []*domain.AlloffCategoryDAO{}
	parentID := ""
	alloffcatDaos, err := ioc.Repo.AlloffCategories.List(&parentID)
	if err != nil {
		log.Println(err)
	}

	for _, alloffcat := range alloffcatDaos {
		query := product.ProductListInput{
			Offset:                    0,
			Limit:                     1,
			BrandID:                   brandDao.ID.Hex(),
			AlloffCategoryID:          alloffcat.ID.Hex(),
			IncludeSpecialProductType: product.NOT_SPECIAL_PRODUCTS,
			IncludeClassifiedType:     product.NO_MATTER_CLASSIFIED,
		}
		_, pdCount, _ := product.Listing(query)
		if pdCount == 0 {
			continue
		}
		brandAlloffCats = append(brandAlloffCats, alloffcat)
	}

	return brandAlloffCats
}
