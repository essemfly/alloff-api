package main

import (
	"log"

	"github.com/lessbutter/alloff-api/cmd/scripter/scripts"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
)

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	scripts.AddAlloffCategory()
	scripts.AddBrandsSeeder()
	scripts.AddClassifyRules()
	scripts.UpdateBrands()
	scripts.UpdateHomeitems()
	scripts.AddProductDiffNotificaton()
	scripts.AddProductGroupsSeeder()
	scripts.ConfirmOrders()
	scripts.SendNotification()
}
