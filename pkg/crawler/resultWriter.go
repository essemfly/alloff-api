package crawler

import (
	"log"
	"strconv"

	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func WriteCrawlResults(source *domain.CrawlSourceDAO, totalProducts int) {
	msg := source.CrawlModuleName + ": " + source.Category.KeyName + " " + strconv.Itoa(totalProducts) + "products crawled"
	log.Println(msg)

}
