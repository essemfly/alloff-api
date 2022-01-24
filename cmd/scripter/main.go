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

	// scripts.AddAlloffCategory()
	// scripts.AddBrandsSeeder()
	// scripts.AddClassifyRules()
	// scripts.UpdateBrands()
	// scripts.AddProductGroupsSeeder()
	// scripts.UpdateHomeitems()
	// scripts.AddProductDiffNotificaton()
	// scripts.ConfirmOrders()
	// scripts.SendNotification()
	// seeder.AddProductGroups()
	scripts.AddMockOrders()
	// scripts.AddMockHomeItems()
	// scripts.InsertDiffNotification()
	// scripts.TestAlimtalk()
}
