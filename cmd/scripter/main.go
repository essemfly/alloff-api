package main

import (
	"log"

	"github.com/lessbutter/alloff-api/cmd/scripter/scripts"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/internal/storage/postgres"
)

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	pgconn := postgres.NewPostgresDB(conf)
	pgconn.RegisterRepos()

	// (TODO) Be Refactored
	config.InitIamPort(conf)
	config.InitSlack(conf)
	config.InitNotification(conf)

	// scripts.AddAlloffCategory()
	// scripts.AddBrandsSeeder()
	// scripts.AddClassifyRules()
	// scripts.UpdateBrands()
	// scripts.UpdateHomeitems()
	// scripts.AddProductDiffNotificaton()
	// scripts.ConfirmOrders()
	// scripts.SendNotification()
	// scripts.AddProductGroups()
	// scripts.AddMockOrders()
	// scripts.AddMockHomeItems()
	// scripts.InsertDiffNotification()
	// scripts.TestAlimtalk()
	title := "MaxMara(Intrend) 시즌오프:짠:"
	message := "(광고) 품절임박! 시즌오프 세일:폭탄:"
	timedealID := "61a9d41dee2c39d09bebeedb"
	scripts.AddTimedealNotification(timedealID, title, message)
}
