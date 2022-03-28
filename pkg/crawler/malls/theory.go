package malls

import (
	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
	"log"
	"strconv"
	"strings"
)

func CrawlTheory(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("outlet.theory.com"),
	)

	productBaseUrl := "https://outlet.theory.com/"
	totalProducts := 0

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".product-grid-tile--small", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".price .list .value")
		originalPrice := 0.0
		if originalPriceStr != "" {
			originalPriceStr = strings.Replace(originalPriceStr, "Comparable Value:", "", -1)
			originalPriceStr = strings.Replace(originalPriceStr, "Price reduced from", "", -1)
			originalPriceStr = strings.Replace(originalPriceStr, "$", "", -1)
			originalPriceStr = strings.Replace(originalPriceStr, "\n", "", -1)
			originalPriceStr = strings.Trim(originalPriceStr, " ")
			originalPrice, err = strconv.ParseFloat(originalPriceStr, 32)
			if err != nil {
				log.Println("err", err)
				return
			}
		}
		log.Println("original price : ", originalPrice)

		discountedPriceStr := e.ChildText(".price .sales .value")
		discountedPriceStr = strings.Replace(discountedPriceStr, "$", "", -1)
		discountedPrice, err := strconv.ParseFloat(discountedPriceStr, 32)
		if err != nil {
			log.Println("err", err)
			return
		}
		log.Println("discounted price : ", discountedPrice)
		if discountedPrice == 0 {
			discountedPrice = originalPrice
		} else if originalPrice == 0.0 {
			originalPrice = discountedPrice
		}

		//		productID := e.ChildAttr(".info_area a", "href")
		productId := e.ChildAttr(".link", "href")
		productUrl := productBaseUrl + productId

		addRequest := &product.ProductCrawlingAddRequest{
			Brand:               brand,
			Source:              source,
			ProductID:           productId,
			ProductName:         "",
			ProductUrl:          productUrl,
			Images:              nil,
			Sizes:               nil,
			Inventories:         nil,
			Colors:              nil,
			Description:         nil,
			OriginalPrice:       float32(originalPrice),
			SalesPrice:          float32(discountedPrice),
			CurrencyType:        domain.CurrencyUSD,
			IsTranslateRequired: true,
		}
		totalProducts += 1
		product.AddProductInCrawling(addRequest)
	})

	err = c.Visit(source.CrawlUrl)
	if err != nil {
		log.Println("err occurred in crawling theory")
	}

	crawler.PrintCrawlResults(source, totalProducts)

	<-worker
	done <- true
}
