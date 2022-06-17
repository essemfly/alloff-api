package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func ProductInSaleUpdater() {
	productListInput := product.ProductListInput{
		Offset: 0,
		Limit:  100000,
	}
	pds, cnt, err := product.ListProducts(productListInput)
	if err != nil {
		log.Println("Err in listing products", err)
	}
	log.Println("Total products", cnt)
	for _, pd := range pds {
		if pd.IsSaleable() {
			pd.IsNotSale = false
		} else {
			pd.IsNotSale = true
		}

		_, err = ioc.Repo.Products.Upsert(pd)
		if err != nil {
			log.Println("err on product upsert", err)
		}
	}
}
