package malls

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
	"log"
	"strings"
)

type ImagesURLs struct {
	ImagesURLs string `json:"imagesURL"`
}

func CrawlClaudiePierlot(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("de.claudiepierlot.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	totalProducts := 0

	//totalProducts := 0

	brand, err := ioc.Repo.Brands.GetByKeyname(source.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".grid-tile", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".price-standard")
		discountedPriceStr := e.ChildText(".product-sales-price")
		originalPrice := parseEuro(originalPriceStr)
		discountedPrice := parseEuro(discountedPriceStr)

		productUrlBase := "https://de.claudiepierlot.com/de_DE/"
		productUrl := e.ChildAttr(".product-tile .product-image .thumb-link", "href")
		productId := productUrl[len(productUrlBase):]
		productName := e.ChildText(".titleProduct .name-link")

		originalImageUrls := e.ChildAttr(".product-tile", "data-productmedia")
		imgUrlsStr := ImagesURLs{}
		err := json.Unmarshal([]byte(originalImageUrls), &imgUrlsStr)
		if err != nil {
			log.Println(err)
		}
		images := strings.Split(imgUrlsStr.ImagesURLs, ",")

		sizes, inventories, description := getClaudiePierlotDetail(productUrl)

		log.Println(images)
		log.Println(productId)
		log.Println(originalPrice)
		log.Println(discountedPrice)
		log.Println(sizes)
		log.Println(inventories)
		log.Println(description)

		addRequest := &product.ProductCrawlingAddRequest{
			Brand:               brand,
			Images:              images,
			Sizes:               sizes,
			Inventories:         inventories,
			Description:         description,
			OriginalPrice:       originalPrice,
			SalesPrice:          discountedPrice,
			CurrencyType:        domain.CurrencyEUR,
			Source:              source,
			ProductID:           productId,
			ProductName:         productName,
			ProductUrl:          productUrl,
			IsTranslateRequired: true,
		}
		totalProducts += 1
		product.AddProductInCrawling(addRequest)
	})

	err = c.Visit(source.CrawlUrl)
	if err != nil {
		log.Println("err occurred in crawling claudie pierlot : ", err)
	}

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}

func getClaudiePierlotDetail(productUrl string) (
	sizes []string, inventories []domain.InventoryDAO, description map[string]string,
) {
	c := colly.NewCollector(
		colly.AllowedDomains("de.claudiepierlot.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	sizes = []string{}
	description = map[string]string{}
	inventories = []domain.InventoryDAO{}

	// 설명
	c.OnHTML(".description", func(e *colly.HTMLElement) {
		desc := ""
		originalDesc := e.ChildText("p")
		originalDesc = originalDesc[:len(originalDesc)-22] // Referenz: 상품코드 제거용
		originalDesc = strings.TrimSpace(originalDesc)
		descSlice := strings.Split(originalDesc, ".")
		for _, str := range descSlice {
			if str != "" {
				str = "- " + str
				desc += str + "\n"
			}
		}
		desc = strings.TrimRight(desc, "\n")
		description["설명"] = desc
	})

	// 사이즈
	c.OnHTML(".siz-list-container .size", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			stock := 10
			outOfStock := e.ChildText(".unclickable .sizeDisplayValue")
			size := el.Text
			size = strings.TrimSpace(size)
			size = strings.Trim(size, "\n")

			if strings.Contains(outOfStock, size) {
				stock = 0
			}

			if size != "Größe" {
				sizes = append(sizes, size)
				inventories = append(inventories, domain.InventoryDAO{
					Size:     size,
					Quantity: stock,
				})
			}
		})
	})

	err := c.Visit(productUrl)
	if err != nil {
		log.Println("err occurred in crawling claudie pierlot : ", err)
	}
	return
}
