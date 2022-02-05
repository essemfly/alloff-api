package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
)

// 기존 ProductGroup에 Product들이 통째로 들어가 있지 않아서, 그것을 넣어주기 위한 1회성 코드
func AddProductInPg() {
	numPassedPgsToShow := 10000
	pgs, _ := ioc.Repo.ProductGroups.List(numPassedPgsToShow)
	for _, pg := range pgs {
		for _, productPriority := range pg.Products {
			product, err := ioc.Repo.Products.Get(productPriority.ProductID.Hex())
			if err != nil {
				log.Println("err occured in pd find", err)
			}
			productPriority.Product = product
		}
		_, err := ioc.Repo.ProductGroups.Upsert(pg)
		if err != nil {
			log.Println("err occured in pg upsert", err)
		}
	}

}
