package main

import (
	"log"
	"strings"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"go.mongodb.org/mongo-driver/bson"
)

const numWorkers = 10

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	workers := make(chan bool, numWorkers)
	done := make(chan bool, numWorkers)

	for c := 0; c < numWorkers; c++ {
		done <- true
	}

	crawlModules := []string{
		// "lottefashion",
		// "ssfmall",
		// "idlook",
		// "sivillage",
		// "kolon",
		// "michaa",
		"babathe",
		// "idfmall",
		// "daehyun",
		// "niceclaup",
		// "lacoste",
		// "sisley",
		// "benetton",
		// "theamall",
		// "loungeb",
		// "bylynn",
	}

	msg := "======== \n " + "Crawling Started: " + time.Now().String() + " for " + strings.Join(crawlModules[:], ", ")
	log.Println(msg)

	for _, module := range crawlModules {
		filter := bson.M{
			"crawlmodulename": module,
		}
		sources, _, err := ioc.Repo.CrawlSources.List(filter)
		if err != nil {
			log.Println(err)
		}

		for _, source := range sources {
			workers <- true
			<-done
			switch source.CrawlModuleName {
			// case "lottefashion":
			// 	go crawler.CrawlLotteFashion(workers, done, source)
			// case "ssfmall":
			// 	go crawler.CrawlSSFMall(workers, done, source)
			// case "handsome":
			// 	go crawler.CrawlHandsome(workers, done, source)
			// case "idlook":
			// 	go crawler.CrawlIdLook(workers, done, source)
			// case "sivillage":
			// 	go crawler.CrawlSiVillage(workers, done, source)
			// case "kolon":
			// 	go crawler.CrawlKolon(workers, done, source)
			// case "michaa":
			// 	go crawler.CrawlMichaa(workers, done, source)
			case "babathe":
				go crawler.CrawlBabathe(workers, done, source)
			// case "hfashion":
			// 	go crawler.CrawlHfashion(workers, done, source)
			// case "idfmall":
			// 	go crawler.CrawlIDFMall(workers, done, source)
			// case "daehyun":
			// 	go crawler.CrawlDaehyun(workers, done, source)
			// case "niceclaup":
			// 	go crawler.CrawlNiceClaup(workers, done, source)
			// case "lacoste":
			// 	go crawler.CrawlLacoste(workers, done, source)
			// case "sisley":
			// 	go crawler.CrawlSisley(workers, done, source)
			// case "benetton":
			// 	go crawler.CrawlBenetton(workers, done, source)
			// case "theamall":
			// 	go crawler.CrawlTheamall(workers, done, source)
			// case "loungeb":
			// 	go crawler.CrawlLoungeB(workers, done, source)
			// case "bylynn":
			// 	go crawler.CrawlBylynn(workers, done, source)
			default:
				log.Println("Empty Source")
				<-workers
			}
		}
	}

	for c := 0; c < numWorkers; c++ {
		<-done
	}
}
