package main

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/pkg/seeder"
)

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	seeder.AddProductGroups()
}
