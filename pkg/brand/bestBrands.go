package brand

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"log"
	"time"
)

func MakeSnapshot() {
	var res []*domain.BrandDAO
	limit := 50
	to := time.Now()
	from := to.Add(-24 * time.Hour)

	rankDoc, _ := ioc.Repo.BrandLog.GetRank(limit, from, to)
	ids := rankDoc.GetIds()

	for _, brandId := range ids {
		brand, err := ioc.Repo.Brands.Get(brandId)
		if err != nil {
			log.Println("error on getting brand : ", err)
		}
		if brand.IsOpen && !brand.IsHide {
			res = append(res, brand)
		}
	}

	if len(res) > 30 {
		res = res[0:30]
	}

	bestBrands := &domain.BestBrandDAO{
		Brands: res,
	}

	_, err := ioc.Repo.BestBrands.Insert(bestBrands)
	if err != nil {
		log.Println("err occurred in make snapshot")
	}
}
