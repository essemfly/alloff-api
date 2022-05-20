package scripts

// import (
// 	"log"
// 	"time"

// 	"github.com/lessbutter/alloff-api/config/ioc"
// 	"github.com/lessbutter/alloff-api/internal/core/domain"
// 	"github.com/lessbutter/alloff-api/internal/utils"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func AddMockTimedeal() {
// 	log.Println("ADD MOCK TIMEDEAL2.0 START ********")
// 	exId := primitive.NewObjectID()
// 	pgs := addTimedealProductGroups(exId.Hex())

// 	totalProducts := 0

// 	var banners []domain.ExhibitionBanner
// 	for _, pg := range pgs {
// 		bn := domain.ExhibitionBanner{
// 			ImgUrl:         "https://picsum.photos/seed/" + utils.CreateShortUUID() + "/400/200",
// 			Title:          "타임딜 2.0 배너 타이틀 " + pg.Title,
// 			Subtitle:       "타임딜 2.0 배너 타이틀 " + pg.Title,
// 			ProductGroupId: pg.ID.Hex(),
// 		}
// 		banners = append(banners, bn)
// 		totalProducts += len(pg.Products)
// 	}

// 	startTime := time.Now()
// 	exDao := domain.ExhibitionDAO{
// 		ID:             exId,
// 		Title:          "타임딜 2.0 석민테스트_19:55",
// 		SubTitle:       "타임딜 2.0 석민테스트_19:55 서브타이틀",
// 		Description:    "타임딜 2.0 석민테스트_19:55 디스크립션!!",
// 		StartTime:      startTime,
// 		FinishTime:     startTime.Add(365 * 24 * time.Hour),
// 		ProductGroups:  pgs,
// 		IsLive:         true,
// 		CreatedAt:      time.Now(),
// 		ExhibitionType: domain.EXHIBITION_TIMEDEAL,
// 		Banners:        banners,
// 	}
// 	_, err := ioc.Repo.Exhibitions.Upsert(&exDao)
// 	if err != nil {
// 		log.Println("Error on Upsert exhibition : ", err)
// 	}
// 	log.Println("ADD MOCK TIMEDEAL2.0 END ********")
// }
