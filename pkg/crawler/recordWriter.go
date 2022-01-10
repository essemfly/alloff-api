package crawler

import (
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func PrintCrawlResults(source *domain.CrawlSourceDAO, totalProducts int) {
	msg := source.CrawlModuleName + ": " + source.Category.KeyName + " " + strconv.Itoa(totalProducts) + "products crawled"
	log.Println(msg)
}

func WriteCrawlRecords(brandModules []string) {
	lastRecord, err := ioc.Repo.CrawlRecords.GetLast()
	lastUpdatedDate := lastRecord.Date
	if err != nil {
		lastUpdatedDate = time.Now().Add(-1 * time.Hour)
	}

	numNewProducts := ioc.Repo.Products.CountNewProducts(brandModules)
	numOutProducts := ioc.Repo.Products.MakeOutdateProducts(brandModules, lastUpdatedDate)

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

	msg := "Update Finished: " + lastUpdatedDate.String() + "\n" + "New Products: " + strconv.Itoa(numNewProducts) + "  Out Products: " + strconv.Itoa(numOutProducts)
	log.Println(msg)
}
