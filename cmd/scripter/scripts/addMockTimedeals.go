package scripts

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddMockTimedeal() {
	log.Println("ADD MOCK TIMEDEAL2.0 START ********")
	exId := primitive.NewObjectID()
	pgs := addTiemdealProductGroups(exId.Hex())

	totalProducts := 0

	var banners []domain.ExhibitionBanner
	for _, pg := range pgs {
		bn := domain.ExhibitionBanner{
			ImgUrl:         "https://picsum.photos/seed/" + utils.CreateShortUUID() + "/400/200",
			Title:          "타임딜 2.0 배너 타이틀 " + pg.Title,
			Subtitle:       "타임딜 2.0 배너 타이틀 " + pg.Title,
			ProductGroupId: pg.ID.Hex(),
		}
		banners = append(banners, bn)
		totalProducts += len(pg.Products)
	}

	startTime := time.Now()
	exDao := domain.ExhibitionDAO{
		ID:             exId,
		Title:          "타임딜 2.0 테스트 타임딜 아마 백오피스에서 쓸듯",
		SubTitle:       "타임딜 2.0 테스트 타임딜 서브타이틀 아마 백오피스에서 쓸듯",
		Description:    "타임딜 2.0 테스트 타임딜 문구 아마 백오피스에서 쓸듯",
		StartTime:      startTime,
		FinishTime:     startTime.Add(365 * 24 * time.Hour),
		ProductGroups:  pgs,
		IsLive:         true,
		CreatedAt:      time.Now(),
		ExhibitionType: domain.EXHIBITION_TIMEDEAL,
		Banners:        banners,
	}
	_, err := ioc.Repo.Exhibitions.Upsert(&exDao)
	if err != nil {
		log.Println("Error on Upsert exhibition : ", err)
	}
	log.Println("ADD MOCK TIMEDEAL2.0 END ********")
}
