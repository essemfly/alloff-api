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
	brandDaos, totalCount, err := ioc.Repo.Brands.List(offset, limit, false, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Total # of brands : " + strconv.Itoa(totalCount))

	for _, brandDao := range brandDaos {
		brandDao.UseAlloffCategory = false
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
		_, pdCount, _ := product.ProductsListing(0, 1, brandDao.ID.Hex(), catDao.ID.Hex(), "", "", nil)
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
		_, pdCount, _ := product.AlloffCategoryProductsListing(0, 1, []string{brandDao.KeyName}, alloffcat.ID.Hex(), "", nil)
		if pdCount == 0 {
			continue
		}
		brandAlloffCats = append(brandAlloffCats, alloffcat)
	}

	return brandAlloffCats
}
