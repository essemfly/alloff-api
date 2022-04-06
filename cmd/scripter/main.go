package main

import (
	"fmt"

	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/internal/pkg/broker"
)

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "prod"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	cmd.SetBaseConfig(Env)

	broker.HomeTabSyncer()
	// notis, err := ioc.Repo.Notifications.List(0, 100, []domain.NotificationType{domain.NOTIFICATION_EXHIBITION_OPEN_NOTIFICATION}, true)
	// if err != nil {
	// 	log.Println("err on noti", err)
	// }

	// for _, noti := range notis {
	// 	noti.Message = "(광고) 단 3일! 최대 72% SALE"
	// 	_, err := ioc.Repo.Notifications.Update(noti)
	// 	if err != nil {
	// 		log.Println("err", err)
	// 	}
	// }
	// ex, err := ioc.Repo.Exhibitions.Get("621f74eaac1cb3aaaeadbebf")
	// if err != nil {
	// 	log.Println("er", err)
	// }
	// ex.BannerImage = "https://alloff.s3.ap-northeast-2.amazonaws.com/images/3WF7Zr_1646228680_%E1%84%80%E1%85%B5%E1%84%92%E1%85%AC%E1%86%A8%E1%84%8C%E1%85%A5%E1%86%AB+%E1%84%89%E1%85%A1%E1%86%BC%E1%84%89%E1%85%A6+30.png"
	// _, err = ioc.Repo.Exhibitions.Upsert(ex)
	// if err != nil {
	// 	log.Println("err2", err)
	// }
	// filter := bson.M{"istranslaterequired": true}
	// offset, limit := 0, 1000
	// pds, cnt, err := ioc.Repo.Products.List(offset, limit, filter, nil)
	// if err != nil {
	// 	log.Println("err", err)
	// }
	// log.Println("cnt", cnt)
	// for _, pd := range pds {
	// 	_, err = product.TranslateProductInfo(pd)
	// 	if err != nil {
	// 		log.Println("err occured ", err)
	// 	}
	// }
}
