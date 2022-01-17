package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/pkg/notification"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertDiffNotification() {
	filter := bson.M{"ispushed": false, "type": "price"}
	diffs, err := ioc.Repo.ProductDiffs.List(filter)
	if err != nil {
		log.Println(err)
		return
	}

	for _, diff := range diffs {
		err = notification.InsertDiffNotification(diff.NewProduct, diff.OldPrice)
		if err != nil {
			log.Println("error occured on insert diff ID", diff.ID, err)
		}
		diff.IsPushed = true
		_, err = ioc.Repo.ProductDiffs.Update(diff)
		if err != nil {
			log.Println("error occured on update diff pushed", diff.ID, err)
		}
	}
	log.Println("Insert Diff notifications finished")
}
