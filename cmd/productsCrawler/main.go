package main

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
)

const numWorkers = 1

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	for i := 0; i < numWorkers; i++ {
		brand, err := ioc.Repo.Brands.Get("60b60075881252a2bd16b848")
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(brand.KorName)
	}
}
