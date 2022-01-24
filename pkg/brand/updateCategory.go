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
	brandDaos, totalCount, err := ioc.Repo.Brands.List(offset, limit, nil, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Total # of brands : " + strconv.Itoa(totalCount))

	for _, brandDao := range brandDaos {
		brandCats := []*domain.CategoryDAO{}
		catDaos, err := ioc.Repo.Categories.List(brandDao.KeyName)
		if err != nil {
			log.Println(err)
		}

		for _, catDao := range catDaos {
			_, pdCount, _ := product.ProductsListing(0, 1, brandDao.ID.Hex(), catDao.ID.Hex(), "", nil)
			if pdCount == 0 {
				continue
			}
			brandCats = append(brandCats, catDao)
		}

		brandDao.Category = brandCats
		_, err = ioc.Repo.Brands.Upsert(brandDao)
		if err != nil {
			log.Println(err)
		}
	}
}
