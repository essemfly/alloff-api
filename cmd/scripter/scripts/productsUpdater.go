package scripts

import (
	"log"
	"time"

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
	productsIsNotSaleCounts := 0
	for _, pd := range pds {
		if pd.ExhibitionFinishTime.Before(time.Now()) {
			pd.IsNotSale = true
			log.Println("Product ID", pd.ID.Hex(), pd.ExhibitionID, pd.ExhibitionFinishTime)
			productsIsNotSaleCounts += 1
			_, err = ioc.Repo.Products.Upsert(pd)
			if err != nil {
				log.Println("err on product upsert", err)
			}
		}
	}
}
