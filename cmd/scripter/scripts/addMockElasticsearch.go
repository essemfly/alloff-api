package scripts

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"math/rand"
	"time"
)

func AddMockBestProducts() {
	pds, cnt, err := ioc.Repo.Products.List(0, 10000, bson.M{}, bson.M{})
	if err != nil {
		log.Panic("error on get product list")
	}
	log.Printf("loaded %v products for mocking elastic search \n", cnt)
	pdCnt := 1
	for _, pd := range pds {
		rand.Seed(time.Now().UnixNano())
		times := rand.Intn(10)
		for i := 1; i <= times; i++ {
			sc, _ := ioc.Repo.ProductLog.Index(pd, domain.PRODUCT_VIEW)
			log.Println(i, "/", times, " -> status code : ", sc, " || progress : ", pdCnt, "/", cnt)
		}
		pdCnt += 1
	}
}

func AddMockBestBrands() {
	bds, cnt, err := ioc.Repo.Brands.List(0, 1000, false, true, bson.M{})
	if err != nil {
		log.Panic("error on get brand list")
	}
	log.Printf("loaded %v brands for mocking elastic search \n", cnt)
	bdCnt := 1
	for _, bd := range bds {
		rand.Seed(time.Now().UnixNano())
		times := rand.Intn(100)
		for i := 1; i <= times; i++ {
			sc, _ := ioc.Repo.BrandLog.Index(bd)
			log.Println(i, "/", times, " -> status code : ", sc, " || progress : ", bdCnt, "/", cnt)
		}
		bdCnt += 1
	}
}
