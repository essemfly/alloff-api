package seeder

// import (
// 	"log"
// 	"strconv"
// 	"time"

// 	"math/rand"

// 	"github.com/lessbutter/alloff/internal"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func AddProductGroups() {
// 	loc, _ := time.LoadLocation("Asia/Seoul")
// 	timedealInstrctuion := []string{
// 		"타임딜 상품은 올오프 MD가 팩토리 아울렛 및 현대/롯데/신세계 프리미엄 아울렛에서 직소싱한 상품입니다.",
// 		"타임딜 상품은 오프라인 매장에서 동시에 판매되고 있습니다. 재고가 제한적이라 결제 완료 후 오프라인 매장에서 판매가 완료되어 품절될 수 있습니다.",
// 		"한섬 팩토리 아울렛 타임딜의 경우, 매장에서는 불가능했던 교환/반품이 올오프에서는 가능합니다."}
// 	timedeal0StartTime := time.Date(2021, time.September, 7, 15, 0, 0, 0, loc)
// 	timedeal1StartTime := time.Date(2021, time.September, 9, 15, 0, 0, 0, loc)
// 	timedeal2StartTime := time.Date(2021, time.September, 10, 15, 0, 0, 0, loc)
// 	timedeal3StartTime := time.Date(2021, time.September, 11, 15, 0, 0, 0, loc)

// 	log.Println("ProductGroup seeding start!")
// 	pgid0 := primitive.NewObjectID()
// 	pgid1 := primitive.NewObjectID()
// 	pgid2 := primitive.NewObjectID()
// 	pgid3 := primitive.NewObjectID()

// 	pgs := []*internal.ProductGroupDAO{
// 		{
// 			ID:          pgid0,
// 			Title:       "한섬 팩토리 아울렛\n 가을 준비 Collection",
// 			ShortTitle:  "",
// 			Instruction: timedealInstrctuion,
// 			StartTime:   timedeal0StartTime,
// 			FinishTime:  timedeal0StartTime.Add(5 * time.Hour),
// 			ImgUrl:      "https://alloff.s3.ap-northeast-2.amazonaws.com/promotion/1st_timedeal.jpeg",
// 			Created:     time.Now(),
// 			Hidden:      false,
// 			NumAlarms:   rand.Intn(50) + 10,
// 		},
// 		{
// 			ID:          pgid1,
// 			Title:       "현대 프리미엄 아울렛\n 간절기 아우터 특가",
// 			ShortTitle:  "",
// 			Instruction: timedealInstrctuion,
// 			StartTime:   timedeal1StartTime,
// 			FinishTime:  timedeal1StartTime.Add(5 * time.Hour),
// 			ImgUrl:      "https://alloff.s3.ap-northeast-2.amazonaws.com/promotion/2nd_timedeal.jpeg",
// 			Created:     time.Now(),
// 			Hidden:      false,
// 			NumAlarms:   rand.Intn(50) + 10,
// 		},
// 		{
// 			ID:          pgid2,
// 			Title:       "한섬 팩토리 아울렛\n 역시즌 Collection",
// 			ShortTitle:  "",
// 			Instruction: timedealInstrctuion,
// 			StartTime:   timedeal2StartTime,
// 			FinishTime:  timedeal2StartTime.Add(5 * time.Hour),
// 			ImgUrl:      "https://alloff.s3.ap-northeast-2.amazonaws.com/promotion/3rd_timedeal.jpeg",
// 			Created:     time.Now(),
// 			Hidden:      false,
// 			NumAlarms:   rand.Intn(50) + 10,
// 		},
// 		{
// 			ID:          pgid3,
// 			Title:       "컨템포러리 MD의 SELECT\n 타임/마인/시스템 외",
// 			ShortTitle:  "",
// 			Instruction: timedealInstrctuion,
// 			StartTime:   timedeal3StartTime,
// 			FinishTime:  timedeal3StartTime.Add(5 * time.Hour),
// 			ImgUrl:      "https://alloff.s3.ap-northeast-2.amazonaws.com/promotion/4th_timedeal.jpeg",
// 			Created:     time.Now(),
// 			Hidden:      false,
// 			NumAlarms:   rand.Intn(50) + 10,
// 		},
// 	}

// 	for _, pg := range pgs {
// 		_, err := pg.Save()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}
// 	log.Println("ProductGroup seeding end!")
// 	AddAlloffProducts([]primitive.ObjectID{pgid0, pgid1, pgid2, pgid3})
// }

// func AddAlloffProducts(pgids []primitive.ObjectID) {
// 	log.Println("Alloff Products seeding start!")
// 	for _, pgid := range pgids {
// 		addSampleProductToProductGroup(pgid)
// 	}

// 	log.Println("Alloff Products seeding Finish!")
// }

// func addSampleProductToProductGroup(pgId primitive.ObjectID) {
// 	r := rand.New(rand.NewSource(99))
// 	cats, _ := internal.ListAlloffCategoriesAll()
// 	catsLen := len(cats)

// 	totalNumProducts := 10
// 	products, _, _ := internal.ListProducts(0, totalNumProducts, "", "", cats[r.Intn(catsLen)].ID, "", nil)
// 	for i := 0; i < totalNumProducts; i++ {
// 		if i%2 == 0 {
// 			newPd := internal.AlloffProductDAO{
// 				ID:    primitive.NewObjectID(),
// 				Brand: products[i].Brand,
// 				ProductType: []string{
// 					"타임딜",
// 					"1년차 B품/샘플",
// 				},
// 				ProductGroupId: pgId.Hex(),
// 				Description: []string{
// 					"1타임딜 상품" + strconv.Itoa(i) + ", 진짜 좋고 90프로 할인해서 판매하는데 너무 안예쁨",
// 					"설명~",
// 					"설오프라인에서만 구매할 수 있었던 프리미엄 브랜드의 특가 상품, 아무말",
// 					"2020 FW 가죽~",
// 					"현대/롯데/신세계 프리미엄 아울렛, 한섬 팩토리 아울렛에서 소싱한 정품이라구~",
// 				},
// 				Instruction: struct {
// 					Title       string
// 					Thumbnail   string
// 					Description []string
// 					Images      []string
// 				}{
// 					Title: "1년차 B품/샘플 유의사항",
// 					Description: []string{
// 						"본 상품은 미세한 스크래치 및 오염/샘플 등의 사유로 저렴하게 판매하는 1년차 아울렛 상품입니다.",
// 						"한섬 팩토리 아울렛 구매 정책 상 교환 및 반품 불가~",
// 						"아래 상세정보를 통해 자세한 사진과 함께 설명해드릴테니 참고해서 구매해주십쇼~",
// 					},
// 					Thumbnail: "http://placekitten.com/200/300",
// 					Images:    products[i].Images,
// 				},
// 				Faults: []struct {
// 					Image       string
// 					Description string
// 				}{
// 					{Image: products[i].Images[0], Description: "정상 상품은 허리에 벨트가 있으나, 샘플 상품의 경우 벨트가 없습니다."},
// 					{Image: products[i].Images[1], Description: "티셔츠 옆 하단에 작은 구멍이 있습니다."},
// 				},
// 				SizeDescription: []string{
// 					"설명~",
// 					"설오프라인에서만 구매할 수 있었던 프리미엄 브랜드의 특가 상품, 아무말",
// 				},
// 				CancelDescription: []string{
// 					"설명~",
// 					"설오프라인에서만 구매할 수 있었던 프리미엄 브랜드의 특가 상품, 아무말",
// 				},
// 				DeliveryDescription: []string{
// 					"타임딜 마감 후 다음 날 택배를 발송하여 평균 2-3일 이내에 상품을 받아보실 수 있습니다.",
// 					"배송비는 각 상품 또는 유통사에 따라 다를 수 있으며, 주문/결제 페이지에서 정확한 금액을 확인하실 수 있습니다.",
// 				},
// 				Name:            products[i].Name,
// 				OriginalPrice:   products[i].OriginalPrice,
// 				DiscountedPrice: products[i].DiscountedPrice,
// 				DiscountRate:    products[i].DiscountRate,
// 				Images:          products[i].Images,
// 				Removed:         false,
// 				Soldout:         false,
// 				Inventory: []struct {
// 					Size     string
// 					Quantity int
// 				}{

// 					{
// 						Size:     "S",
// 						Quantity: r.Intn(3),
// 					},
// 					{
// 						Size:     "M",
// 						Quantity: r.Intn(3),
// 					},
// 					{
// 						Size:     "L",
// 						Quantity: r.Intn(1),
// 					},
// 				},
// 			}
// 			_, err := newPd.Save()
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		} else {
// 			newPd := internal.AlloffProductDAO{
// 				ID:    primitive.NewObjectID(),
// 				Brand: products[i].Brand,
// 				ProductType: []string{
// 					"타임딜",
// 					"아울렛 특가 상품",
// 				},
// 				ProductGroupId: pgId.Hex(),
// 				Description: []string{
// 					"1타임딜 상품" + strconv.Itoa(i) + ", 진짜 좋고 90프로 할인해서 판매하는데 너무 안예쁨",
// 					"설명~",
// 					"설오프라인에서만 구매할 수 있었던 프리미엄 브랜드의 특가 상품, 아무말",
// 					"2020 FW 가죽~",
// 					"현대/롯데/신세계 프리미엄 아울렛, 한섬 팩토리 아울렛에서 소싱한 정품이라구~",
// 				},
// 				Instruction: struct {
// 					Title       string
// 					Thumbnail   string
// 					Description []string
// 					Images      []string
// 				}{
// 					Title: "아울렛 특가 상품 유의사항",
// 					Description: []string{
// 						"현대/롯데/신세계 프리미엄 아울렛, 한섬 팩토리 아울렛 구매 정책 상, 반품/교환 기간이 짧습니다. 따라서 기간 내 반품/교환이 어려워~",
// 					},
// 					Thumbnail: "http://placekitten.com/200/300",
// 					Images:    products[1].Images,
// 				},
// 				Faults: nil,
// 				SizeDescription: []string{
// 					"설명~",
// 					"설오프라인에서만 구매할 수 있었던 프리미엄 브랜드의 특가 상품, 아무말",
// 				},
// 				CancelDescription: []string{
// 					"설명~",
// 					"설오프라인에서만 구매할 수 있었던 프리미엄 브랜드의 특가 상품, 아무말",
// 				},
// 				DeliveryDescription: []string{
// 					"타임딜 마감 후 다음 날 택배를 발송하여 평균 2-3일 이내에 상품을 받아보실 수 있습니다.",
// 					"배송비는 각 상품 또는 유통사에 따라 다를 수 있으며, 주문/결제 페이지에서 정확한 금액을 확인하실 수 있습니다.",
// 				},
// 				Name:            products[i].Name,
// 				OriginalPrice:   products[i].OriginalPrice,
// 				DiscountedPrice: products[i].DiscountedPrice,
// 				DiscountRate:    products[i].DiscountRate,
// 				Images:          products[i].Images,
// 				Removed:         false,
// 				Soldout:         false,
// 				Inventory: []struct {
// 					Size     string
// 					Quantity int
// 				}{

// 					{
// 						Size:     "S",
// 						Quantity: r.Intn(3),
// 					},
// 					{
// 						Size:     "M",
// 						Quantity: r.Intn(3),
// 					},
// 					{
// 						Size:     "L",
// 						Quantity: r.Intn(1),
// 					},
// 				},
// 			}
// 			_, err := newPd.Save()
// 			if err != nil {
// 				log.Println(err)
// 			}
// 		}

// 	}
// }
