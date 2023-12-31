package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/config"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/crawler/malls"
	"go.mongodb.org/mongo-driver/bson"
)

const numWorkers = 20

func main() {
	cmd.SetBaseConfig()

	crawlModules := []string{
		// "lottefashion",
		// "ssfmall",
		// "sivillage",
		// "kolon",
		// "babathe",
		// "idfmall",
		// "daehyun",
		// "niceclaup",
		// "lacoste",
		// "sisley",
		// "benetton",
		// "theamall",
		// "loungeb",
		// "bylynn",
		// "intrend",
		// "theoutnet",
		// "sandro",
		// "maje",
		// "theory",
		// "claudiePierlot",
		// "flannels",
		//"afound",
		"colognese",
	}

	StartCrawling(crawlModules)
	// product.UpdateManuelProducts()

	crawler.WriteCrawlRecords(crawlModules)
}

func StartCrawling(crawlModules []string) {

	workers := make(chan bool, numWorkers)
	done := make(chan bool, numWorkers)

	for c := 0; c < numWorkers; c++ {
		done <- true
	}

	msg := "======== \n " + "Crawling Started: " + time.Now().String() + " for " + strings.Join(crawlModules[:], ", ")
	log.Println(msg)
	config.WriteSlackMessage(msg)

	for _, module := range crawlModules {
		filter := bson.M{
			"crawlmodulename": module,
		}
		sources, _, err := ioc.Repo.CrawlSources.List(filter)
		if err != nil {
			log.Println(err)
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(sources), func(i, j int) { sources[i], sources[j] = sources[j], sources[i] })

		for _, source := range sources {
			workers <- true
			<-done
			switch source.CrawlModuleName {
			case "lottefashion":
				go malls.CrawlLotteFashion(workers, done, source)
			case "ssfmall":
				go malls.CrawlSSFMall(workers, done, source)
			case "idlook":
				go malls.CrawlIdLook(workers, done, source)
			case "sivillage":
				go malls.CrawlSiVillage(workers, done, source)
			case "kolon":
				go malls.CrawlKolon(workers, done, source)
			case "babathe":
				go malls.CrawlBabathe(workers, done, source)
			case "idfmall":
				go malls.CrawlIDFMall(workers, done, source)
			case "daehyun":
				go malls.CrawlDaehyun(workers, done, source)
			case "niceclaup":
				go malls.CrawlNiceClaup(workers, done, source)
			case "lacoste":
				go malls.CrawlLacoste(workers, done, source)
			case "sisley":
				go malls.CrawlSisley(workers, done, source)
			case "benetton":
				go malls.CrawlBenetton(workers, done, source)
			case "theamall":
				go malls.CrawlTheamall(workers, done, source)
			case "loungeb":
				go malls.CrawlLoungeB(workers, done, source)
			case "bylynn":
				go malls.CrawlBylynn(workers, done, source)
			case "intrend":
				go malls.CrawlIntrend(workers, done, source)
			case "sandro":
				go malls.CrawlSandro(workers, done, source)
			case "maje":
				go malls.CrawlMaje(workers, done, source)
			case "theoutnet":
				go malls.CrawlTheoutnet(workers, done, source)
			case "theory":
				go malls.CrawlTheory(workers, done, source)
			case "claudiePierlot":
				go malls.CrawlClaudiePierlot(workers, done, source)
			case "afound":
				go malls.CrawlAfound(workers, done, source)
			case "flannels":
				go malls.CrawlFlannels(workers, done, source)
			case "colognese":
				go malls.CrawlColognese(workers, done, source)
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

//
