package broker

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"go.mongodb.org/mongo-driver/bson"
)

func BrandSyncer(brandKeyname string) {
	offset, limit := 0, 20000

	newBrand, _ := ioc.Repo.Brands.GetByKeyname(brandKeyname)

	filter := bson.M{
		"productinfo.brand.keyname": brandKeyname,
	}

	pds, _, err := ioc.Repo.Products.List(offset, limit, filter, nil)
	if err != nil {
		log.Println("err", err)
	}

	for idx, pd := range pds {
		if idx%100 == 0 {
			log.Println("IDX", idx)
		}
		pd.ProductInfo.Brand = newBrand
		_, err := ioc.Repo.Products.Upsert(pd)
		if err != nil {
			log.Println("err occured", err)
		}
	}

}
