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

	// (TODO) Be Refactored
	config.InitIamPort(conf)
	config.InitSlack(conf)

	// scripts.AddAlloffCategory()
	// scripts.AddBrandsSeeder()
	// scripts.AddClassifyRules()
	scripts.UpdateBrands()
	scripts.AddProductGroupsSeeder()
	scripts.UpdateHomeitems()
	// scripts.AddProductDiffNotificaton()
	// scripts.ConfirmOrders()
	// scripts.SendNotification()
	// scripts.AddMockOrders()
	// scripts.InsertDiffNotification()
}
