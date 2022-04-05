package malls

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"log"
	"strings"
	"testing"
	"time"
)

func TestCrawlClaudiePierlot(t *testing.T) {
	c := colly.NewCollector(
		colly.AllowedDomains("de.claudiepierlot.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.SetRequestTimeout(45 * time.Second)

	//totalProducts := 0

	//_, err := ioc.Repo.Brands.GetByKeyname("CLAUDIEPIERLOT")
	//if err != nil {
	//	log.Println(err)
	//}

	c.OnHTML(".grid-tile", func(e *colly.HTMLElement) {
		productUrlBase := "https://de.claudiepierlot.com/de_DE/"
		productUrl := e.ChildAttr(".product-tile .product-image .thumb-link", "href")
		productId := productUrl[len(productUrlBase):]
		productName := e.ChildText(".titleProduct .name-link")

		originalPriceStr := e.ChildText(".price-standard")
		discountedPriceStr := e.ChildText(".product-sales-price")
		originalPrice := parseEuro(originalPriceStr)
		discountedPrice := parseEuro(discountedPriceStr)

		originalImageUrls := e.ChildAttr(".product-tile", "data-productmedia")
		imgUrls := ImagesURLs{}
		err := json.Unmarshal([]byte(originalImageUrls), &imgUrls)
		if err != nil {
			log.Println(err)
		}
		images := strings.Split(imgUrls.ImagesURLs, ",")

		log.Println(productName)
		log.Println(images)
		log.Println(originalPrice)
		log.Println(discountedPrice)
		log.Println(productId)

	})

	err := c.Visit("https://de.claudiepierlot.com/de_DE/categories/accessories/?sz=10000&start=0")
	if err != nil {
		log.Println(err)
	}
}

func TestGetClaudiePierlotDetail(t *testing.T) {
	productUrl := "https://de.claudiepierlot.com/de_DE/kategorien/oberteil-outlet/cocteaue20/CFPCM00057.html?dwvar_CFPCM00057_color=K009"
	sz, iv, dc := getClaudiePierlotDetail(productUrl)
	log.Println(dc)
	log.Println(sz)
	log.Println(iv)
}
