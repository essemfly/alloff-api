package main

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/pkg/seeder"
)

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	catDaos := seeder.LoadAlloffCats()

	for _, catDao := range catDaos {
		ioc.Repo.AlloffCategories.Upsert(catDao)
	}

	seeder.MakeClassifyRules()
}
