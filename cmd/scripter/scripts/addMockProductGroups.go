package scripts

import (
	"log"
	"math/rand"
	"time"

	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func addTiemdealProductGroups(exhibitionId string) []*domain.ProductGroupDAO {
	loc, _ := time.LoadLocation("Asia/Seoul")
	timedealInstrctuion := []string{
		"타임딜 상품은 올오프 MD가 팩토리 아울렛 및 현대/롯데/신세계 프리미엄 아울렛에서 직소싱한 상품입니다.",
		"타임딜 상품은 오프라인 매장에서 동시에 판매되고 있습니다. 재고가 제한적이라 결제 완료 후 오프라인 매장에서 판매가 완료되어 품절될 수 있습니다.",
		"한섬 팩토리 아울렛 타임딜의 경우, 매장에서는 불가능했던 교환/반품이 올오프에서는 가능합니다."}
	timedealStartTime := time.Now().In(loc)

	log.Println("Timedeal ProductGroup seeding start!")

	brands, _, err := ioc.Repo.Brands.List(0, 10, true, nil)
	if err != nil {
		log.Println("Error on listing brands")
	}

	var pgs []*domain.ProductGroupDAO
	for _, brand := range brands {
		pgid := primitive.NewObjectID()
		var backImgUrl string
		if brand.BackImgUrl != "" {
			backImgUrl = brand.BackImgUrl
		} else {
			backImgUrl = "https://picsum.photos/seed/" + utils.CreateShortUUID() + "/375/215"
		}

		pg := domain.ProductGroupDAO{
			ID:          pgid,
			Title:       brand.KorName + " 타임딜 2.0 테스트 타이틀 타이틀",
			ShortTitle:  brand.KorName + " 타임딜 2.0 숏타이틀",
			Instruction: timedealInstrctuion,
			ImgUrl:      backImgUrl,
			GroupType:   domain.PRODUCT_GROUP_BRAND_TIMEDEAL,
			StartTime:   timedealStartTime,
			FinishTime:  timedealStartTime.Add(365 * 24 * time.Hour),
			Created:     time.Now(),
			NumAlarms:   rand.Intn(50) + 10,
			Brand:       brand,
		}

		filter := bson.M{
			"productinfo.brand.keyname": brand.KeyName,
		}
		products, _, err := ioc.Repo.Products.List(0, 10, filter, nil)
		if err != nil {
			log.Println("Error on listing products : ", err)
		}

		var pdPriorities []*domain.ProductPriorityDAO
		for i, pd := range products {
			pd.SpecialPrice = pd.DiscountedPrice // Mocking Data 에서만 이렇게 사용
			pdPriorities = append(pdPriorities, &domain.ProductPriorityDAO{
				Priority:  i,
				Product:   pd,
				ProductID: pd.ID,
			})
			pd.ProductGroupId = pg.ID.Hex()
			_, err = ioc.Repo.Products.Upsert(pd)
			if err != nil {
				log.Println("Error on upsert product : ", err)
			}
		}

		pg.Products = pdPriorities
		pg.ExhibitionID = exhibitionId
		_, err = ioc.Repo.ProductGroups.Upsert(&pg)
		if err != nil {
			log.Println("Error on upsert productGroups : ", err)
		}
		pgs = append(pgs, &pg)
	}
	log.Println("Timedeal2.0 ProductGroup seeding end!")
	return pgs
}

func AddProductGroups() {
	loc, _ := time.LoadLocation("Asia/Seoul")
	timedealInstrctuion := []string{
		"타임딜 상품은 올오프 MD가 팩토리 아울렛 및 현대/롯데/신세계 프리미엄 아울렛에서 직소싱한 상품입니다.",
		"타임딜 상품은 오프라인 매장에서 동시에 판매되고 있습니다. 재고가 제한적이라 결제 완료 후 오프라인 매장에서 판매가 완료되어 품절될 수 있습니다.",
		"한섬 팩토리 아울렛 타임딜의 경우, 매장에서는 불가능했던 교환/반품이 올오프에서는 가능합니다."}
	timedeal0StartTime := time.Now().In(loc)
	timedeal1StartTime := time.Now().Add(2 * time.Hour).In(loc)
	timedeal2StartTime := time.Now().Add(12 * time.Hour).In(loc)
	timedeal3StartTime := time.Now().Add(-12 * time.Hour).In(loc)

	log.Println("ProductGroup seeding start!")
	pgid0 := primitive.NewObjectID()
	pgid1 := primitive.NewObjectID()
	pgid2 := primitive.NewObjectID()
	pgid3 := primitive.NewObjectID()

	pgs := []*domain.ProductGroupDAO{
		{
			ID:          pgid0,
			Title:       "한섬 팩토리 아울렛\n 가을 준비 Collection",
			ShortTitle:  "",
			Instruction: timedealInstrctuion,
			StartTime:   timedeal0StartTime,
			FinishTime:  timedeal0StartTime.Add(10 * time.Hour),
			ImgUrl:      "https://alloff.s3.ap-northeast-2.amazonaws.com/promotion/1st_timedeal.jpeg",
			Created:     time.Now(),
			NumAlarms:   rand.Intn(50) + 10,
		},
		{
			ID:          pgid1,
			Title:       "현대 프리미엄 아울렛\n 간절기 아우터 특가",
			ShortTitle:  "",
			Instruction: timedealInstrctuion,
			StartTime:   timedeal1StartTime,
			FinishTime:  timedeal1StartTime.Add(10 * time.Hour),
			ImgUrl:      "https://alloff.s3.ap-northeast-2.amazonaws.com/promotion/2nd_timedeal.jpeg",
			Created:     time.Now(),
			NumAlarms:   rand.Intn(50) + 10,
		},
		{
			ID:          pgid2,
			Title:       "한섬 팩토리 아울렛\n 역시즌 Collection",
			ShortTitle:  "",
			Instruction: timedealInstrctuion,
			StartTime:   timedeal2StartTime,
			FinishTime:  timedeal2StartTime.Add(10 * time.Hour),
			ImgUrl:      "https://alloff.s3.ap-northeast-2.amazonaws.com/promotion/3rd_timedeal.jpeg",
			Created:     time.Now(),
			NumAlarms:   rand.Intn(50) + 10,
		},
		{
			ID:          pgid3,
			Title:       "컨템포러리 MD의 SELECT\n 타임/마인/시스템 외",
			ShortTitle:  "",
			Instruction: timedealInstrctuion,
			StartTime:   timedeal3StartTime,
			FinishTime:  timedeal3StartTime.Add(10 * time.Hour),
			ImgUrl:      "https://alloff.s3.ap-northeast-2.amazonaws.com/promotion/4th_timedeal.jpeg",
			Created:     time.Now(),
			NumAlarms:   rand.Intn(50) + 10,
		},
	}

	blouseOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbff6")
	outerOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbffc")
	pantsOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cc003")
	onepieceOID, _ := primitive.ObjectIDFromHex("60feb9f98adeef23689cbffa")

	alloffcats := []primitive.ObjectID{
		blouseOID, outerOID, pantsOID, onepieceOID,
	}

	for idx, pg := range pgs {
		totalNumProducts := 10
		products, _, err := product.AlloffCategoryProductsListing(0, totalNumProducts, nil, alloffcats[idx].Hex(), "", nil)
		if err != nil {
			log.Println("add sample product error", err)
		}
		log.Println("#products", len(products))

		pdpriorities := []*domain.ProductPriorityDAO{}
		for idx, pd := range products {
			pd.SpecialPrice = pd.DiscountedPrice // For test code
			pdpriorities = append(pdpriorities, &domain.ProductPriorityDAO{
				Priority:  idx,
				ProductID: pd.ID,
			})
			pd.ProductGroupId = pg.ID.Hex()
			_, err = ioc.Repo.Products.Upsert(pd)
			if err != nil {
				log.Println("product upsert failed")
			}
		}

		pg.Products = pdpriorities
		_, err = ioc.Repo.ProductGroups.Upsert(pg)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("ProductGroup seeding end!")
}

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
