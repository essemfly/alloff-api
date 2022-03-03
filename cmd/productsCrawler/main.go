package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/brand"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/crawler/malls"
	"go.mongodb.org/mongo-driver/bson"
)

const numWorkers = 20

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "local"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	cmd.SetBaseConfig(Env)

	crawlModules := []string{
		// "lottefashion",
		// "ssfmall",
		// // "idlook",
		// "sivillage",
		// "kolon",
		// "babathe",
		// "idfmall",
		"daehyun",
		// "niceclaup",
		// "lacoste",
		// "sisley",
		// "benetton",
		// "theamall",
		// "loungeb",
		// "bylynn",
		"intrend",
		"theoutnet",
	}

	StartCrawling(crawlModules)

	brand.UpdateBrandCategory()
	brand.UpdateBrandDiscountRate()
	// product.MakeSnapshot()
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
			case "theoutnet":
				go malls.CrawlTheoutnet(workers, done, source)
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
