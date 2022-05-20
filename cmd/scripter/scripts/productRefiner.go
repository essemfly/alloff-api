package scripts

// import (
// 	"log"

// 	"github.com/lessbutter/alloff-api/config/ioc"
// 	"go.mongodb.org/mongo-driver/bson"
// )

// func RemoveOver100DiscountRate() {
// 	offset, limit := 0, 1000
// 	filter := bson.M{
// 		"discountrate": bson.M{"$gte": 100},
// 	}
// 	pds, cnt, err := ioc.Repo.Products.List(offset, limit, filter, nil)
// 	if err != nil {
// 		log.Println("err in listing products", err)
// 	}
// 	log.Println("CNT", cnt)
// 	for _, pd := range pds {
// 		pd.Removed = true
// 		_, err := ioc.Repo.Products.Upsert(pd)
// 		if err != nil {
// 			log.Println("fail on upsert product", pd)
// 		}
// 	}
// }

// func FixEmptyAlloffname() {
// 	offset, limit := 0, 1000
// 	filter := bson.M{
// 		"alloffname": "",
// 	}
// 	pds, cnt, err := ioc.Repo.Products.List(offset, limit, filter, nil)
// 	if err != nil {
// 		log.Println("ERR", err)
// 	}

// 	for idx, pd := range pds {
// 		if idx%100 == 0 {
// 			log.Println(pd.ID.Hex(), pd.ProductInfo.Brand.KeyName, pd.ProductInfo.Source.CrawlModuleName)
// 		}
// 		pd.AlloffName = pd.ProductInfo.OriginalName
// 		pd.IsTranslateRequired = true

// 		_, err := ioc.Repo.Products.Upsert(pd)
// 		if err != nil {
// 			log.Println("ERR", err)
// 		}
// 	}

// 	log.Println("CNT", cnt)
// }
