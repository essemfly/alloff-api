package product

import (
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

// MakeSnapshot :
// 1. 전체 베스트상품 가져와서 100개 채움 (Elsticsearch 에서 100개 못채우는 경우 랜덤으로 채움)
// 2. alloffLv1 카테고리에 해당하는 베스트 상품 가져와서 100개 채움 (Elasticsearch 에서 100개 못채우는 경우 랜덤으로 채움)
func MakeSnapshot() {
	log.Println("Making snapshot Start!!")

	limit := 1_000
	to := time.Now()
	from := to.Add(-24 * time.Hour)

	// get best 100 products from elastic search
	allBestPds, _ := getTotalBestProducts(limit, from, to)
	totalBest := domain.BestProductDAO{
		ID:               primitive.NewObjectID(),
		AlloffCategoryID: "",
		Products:         allBestPds,
	}
	_, err := ioc.Repo.BestProducts.Insert(&totalBest)
	if err != nil {
		log.Println("err occurred in make snapshot : ", err)
	}

	// get alloffLv1 category ids
	// and make map of string key & []*domain.ProductDAO value
	alloffCatIds := []string{""}
	alloffLev1Cats, _ := ioc.Repo.AlloffCategories.List(&alloffCatIds[0])
	catMap := make(map[string][]*domain.ProductDAO)
	for _, cat := range alloffLev1Cats {
		catMap[cat.ID.Hex()] = []*domain.ProductDAO{}
	}

	// loop catMap, add product to each category and persist it
	for k, _ := range catMap {
		catMap[k], _ = getCatBestProducts(limit, from, to, k)

		catBest := domain.BestProductDAO{
			ID:               primitive.NewObjectID(),
			AlloffCategoryID: k,
			Products:         catMap[k],
		}
		_, err := ioc.Repo.BestProducts.Insert(&catBest)
		if err != nil {
			log.Println("err occured in make snapshot", err)
		}
	}
}

func getTotalBestProducts(limit int, from, to time.Time) ([]*domain.ProductDAO, error) {
	var allBestPds []*domain.ProductDAO
	rankDoc, _ := ioc.Repo.ProductLog.GetRank(limit, from, to)

	ids := rankDoc.GetIds()

	for _, pdId := range ids {
		pd, err := ioc.Repo.Products.Get(pdId)
		if err != nil {
			log.Println("error on get product for product ids : ", err)
			return nil, err
		}
		if !pd.Soldout && !pd.Removed {
			allBestPds = append(allBestPds, pd)
		}
	}

	if len(allBestPds) >= 100 {
		allBestPds = allBestPds[0:100]
	} else {
		limit := 100 - len(allBestPds)
		randomPds := GetBestProductsFromAll(limit)
		log.Printf("product count of totalBest is less than 100 (now %v), add %v of random product\n", len(allBestPds), limit)
		allBestPds = append(allBestPds, randomPds...)
	}
	log.Printf("successfully loaded %v best products for all\n", len(allBestPds))
	return allBestPds, nil
}

func getCatBestProducts(limit int, from, to time.Time, catId string) ([]*domain.ProductDAO, error) {
	var catBestPds []*domain.ProductDAO
	rankDoc, err := ioc.Repo.ProductLog.GetRankByCatId(limit, from, to, catId)
	if err != nil {
		log.Println("error on get rank of products : ", err)
		return nil, err
	}

	ids := rankDoc.GetIds()
	for _, pdId := range ids {
		pd, err := ioc.Repo.Products.Get(pdId)
		if err != nil {
			log.Println("error on get product for product ids : ", err)
			return nil, err
		}
		if !pd.Soldout && !pd.Removed {
			catBestPds = append(catBestPds, pd)
		}
	}

	if len(catBestPds) >= 100 {
		catBestPds = catBestPds[0:100]
	} else {
		limit := 100 - len(catBestPds)
		randomPds := GetAlloffCategoryProducts(catId, limit)
		log.Printf("product count of category %s is less than 100 (now %v), add %v of random product\n", catId, len(catBestPds), limit)
		catBestPds = append(catBestPds, randomPds...)
	}
	log.Printf("successfully loaded %v best products for cat id : %s\n", len(catBestPds), catId)
	return catBestPds, nil
}

func GetAlloffCategoryProducts(alloffCatID string, limit int) []*domain.ProductDAO {
	query := ProductListInput{
		Offset:                    0,
		Limit:                     limit,
		AlloffCategoryID:          alloffCatID,
		IncludeSpecialProductType: NOT_SPECIAL_PRODUCTS,
		IncludeClassifiedType:     NO_MATTER_CLASSIFIED,
		PriceRanges:               []PriceRangeType{PRICE_RANGE_50, PRICE_RANGE_70},
	}

	pds, _, err := Listing(query)
	if err != nil {
		log.Println("err occured in alloff cats product recording")
	}

	return pds
}

func GetBestProductsFromAll(limit int) []*domain.ProductDAO {
	query := ProductListInput{
		Offset:                    0,
		Limit:                     limit,
		IncludeSpecialProductType: NOT_SPECIAL_PRODUCTS,
		IncludeClassifiedType:     NO_MATTER_CLASSIFIED,
		PriceRanges:               []PriceRangeType{PRICE_RANGE_100, PRICE_RANGE_70},
	}

	productDaos, _, err := Listing(query)
	if err != nil {
		log.Println("err occured in products listing")
	}

	return productDaos
}
