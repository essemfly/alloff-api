package crawler

import (
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func PrintCrawlResults(source *domain.CrawlSourceDAO, totalProducts int) {
	if totalProducts > 0 {
		msg := source.CrawlModuleName + ": " + source.Category.KeyName + " " + strconv.Itoa(totalProducts) + "products crawled"
		log.Println(msg)
	}
}

func WriteCrawlRecords(brandModules []string) {
	lastRecord, err := ioc.Repo.CrawlRecords.GetLast()
	lastUpdatedDate := time.Now().Add(7 * -24 * time.Hour)
	if err == nil {
		lastUpdatedDate = lastRecord.Date
	}

	numNewProducts := ioc.Repo.ProductMetaInfos.CountNewProducts(brandModules, lastRecord.Date)
	numOutProducts := ioc.Repo.ProductMetaInfos.MakeOutdatedProducts(brandModules, lastRecord.Date)

	newRecord := &domain.CrawlRecordDAO{
		Date:          time.Now(),
		CrawledBrands: brandModules,
		NewProducts:   numNewProducts,
		OldProducts:   numOutProducts,
	}

	err = ioc.Repo.CrawlRecords.Insert(newRecord)
	if err != nil {
		log.Println("Error occured in inserting crawling Record", err)
	}

	msg := "Update Finished: " + lastUpdatedDate.Format("2006-01-02 15:04:05") + "\n" + "New Products: " + strconv.Itoa(numNewProducts) + "  Out Products: " + strconv.Itoa(numOutProducts)
	log.Println(msg)
	config.WriteSlackMessage(msg)
}
