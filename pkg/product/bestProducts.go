package product

import (
	"github.com/lessbutter/alloff-api/internal/core/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MakeSnapshot() {
	log.Println("Making snapshot Start!!")

	var allBestPds []*domain.ProductDAO

	limit := 10_000
	to := time.Now()
	from := to.Add(-24 * time.Hour)

	rankDoc, _ := ioc.Repo.ProductLog.GetRank(limit, from, to)
	ids := getProductIdsFromDocuments(rankDoc)
	pds, err := ioc.Repo.Products.ListByIds(ids)
	if err != nil {
		log.Println("err occurred in make snapshot : ", err)
	}

	if len(pds) == 0 {
		log.Println("no data for make snapshot") // 이떄 sentry에 연락하고싶다..?!
		allBestPds = []*domain.ProductDAO{}
	} else {
		if len(pds) > 99 {
			allBestPds = pds[0:99]
		} else {
			allBestPds = pds
		}
	}

	totalBest := domain.BestProductDAO{
		ID:               primitive.NewObjectID(),
		AlloffCategoryID: "",
		Products:         allBestPds,
	}
	log.Println(len(totalBest.Products))
	//alloffCatIDs := []string{""}
	//alloffLev1Cats, _ := ioc.Repo.AlloffCategories.List(&alloffCatIDs[0])
	//
	//products :=  ()
	//snapshot := domain.BestProductDAO{
	//	ID:               primitive.NewObjectID(),
	//	AlloffCategoryID: "",
	//	Products:         products,
	//}
	//_, err := ioc.Repo.BestProducts.Insert(&snapshot)
	//if err != nil {
	//	log.Println("err occured in make snapshot", err)
	//}
	//
	//for _, cat := range alloffLev1Cats {
	//	products := GetAlloffCategoryProducts(cat.ID.Hex())
	//	snapshot := domain.BestProductDAO{
	//		ID:               primitive.NewObjectID(),
	//		AlloffCategoryID: cat.ID.Hex(),
	//		Products:         products,
	//	}
	//	_, err := ioc.Repo.BestProducts.Insert(&snapshot)
	//	if err != nil {
	//		log.Println("err occured in make snapshot", err)
	//	}
	//}
}

func getProductIdsFromDocuments(doc *dto.DocumentCountDTO) (ids []string) {
	// buckets : Elasticsearch의 Aggregation Query의 결과가 담기는 리스트
	buckets := doc.Aggregations.GroupByState.Buckets
	for _, bucket := range buckets {
		ids = append(ids, bucket.Key)
	}
	return
}

func GetAlloffCategoryProducts(alloffCatID string) []*domain.ProductDAO {
	pds, _, err := AlloffCategoryProductsListing(0, 100, nil, alloffCatID, "", []string{"70", "50"})
	if err != nil {
		log.Println("err occured in alloff cats product recording")
	}

	return pds
}

func GetBestProductsFromAll() []*domain.ProductDAO {
	productDaos, _, err := ProductsListing(0, 100, "", "", "", "", []string{"70", "100"})
	if err != nil {
		log.Println("err occured in products listing")
	}

	return productDaos
}
