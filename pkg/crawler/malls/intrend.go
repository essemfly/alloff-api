package malls

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func CrawlIntrend(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("it.intrend.it"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	currentPageNum := 0
	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	totalProducts := 0

	c.OnHTML(".js-product-list .js-pagination-container .js-product-card", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".full-price")
		originalPrice := 0.0
		if originalPriceStr != "" {
			originalPriceStr = strings.Split(originalPriceStr, " ")[1]
			originalPriceStr = strings.Replace(originalPriceStr, ",", ".", -1)
			originalPrice, err = strconv.ParseFloat(originalPriceStr, 32)
			if err != nil {
				log.Println("err ", err)
				return
			}
		}

		discountedPriceStr := e.ChildText(".price")
		discountedPriceStr = strings.Split(discountedPriceStr, " ")[1]
		discountedPriceStr = strings.Replace(discountedPriceStr, ",", ".", -1)
		discountedPrice, err := strconv.ParseFloat(discountedPriceStr, 32)
		if err != nil {
			log.Println("err ", err)
			return
		}

		if discountedPrice == 0 {
			discountedPrice = originalPrice
		} else if originalPrice == 0.0 {
			originalPrice = discountedPrice
		}

		productID := e.Attr("data-product-id")
		productUrl := "https://it.intrend.it" + e.ChildAttr(".js-anchor", "href")

		title, images, sizes, colors, inventories, description := getIntrendDetail(productUrl)

		if len(sizes) == 0 {
			inventories = append(inventories, domain.InventoryDAO{
				Size:     "normal",
				Quantity: 10,
			})
		}
		addRequest := &product.ProductCrawlingAddRequest{
			Brand:         brand,
			Source:        source,
			ProductID:     productID,
			ProductName:   title,
			ProductUrl:    productUrl,
			Images:        images,
			Sizes:         sizes,
			Inventories:   inventories,
			Colors:        colors,
			Description:   description,
			OriginalPrice: float32(originalPrice),
			SalesPrice:    float32(discountedPrice),
			CurrencyType:  domain.CurrencyEUR,
		}

		totalProducts += 1
		product.AddProductInCrawling(addRequest)
	})

	c.OnHTML(".js-pager .container-fluid ul", func(e *colly.HTMLElement) {
		lastPageStr := e.ChildAttr("li:last-child a", "data-page")
		lastPageNum, _ := strconv.Atoi(lastPageStr)
		if currentPageNum < lastPageNum {
			currentPageNum += 1
			url := source.CrawlUrl + "?page=" + strconv.Itoa(currentPageNum)
			c.Visit(url)
		}
	})
	err = c.Visit(source.CrawlUrl)
	if err != nil {
		log.Println("err occured in crawling intrend", err)
	}

	crawler.PrintCrawlResults(source, totalProducts)

	<-worker
	done <- true
}

type IntrendStock struct {
	STYCD    string `json:"STYCD"`
	SIZECD   string `json:"SIZECD"`
	COLCD    string `json:"COLCD"`
	SIZECDNM string `json:"SIZECDNM"`
	SALECNT  int    `json:"SALECNT"`
	STOCKQTY int    `json:"STOCKQTY"`
}

func getIntrendDetail(productUrl string) (title string, imageUrls []string, sizes, colors []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("it.intrend.it"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	description = map[string]string{}

	c.OnHTML(".pp_mod-prod-desc-head", func(e *colly.HTMLElement) {
		title = e.ChildText(".title")
	})

	c.OnHTML(".product-gallery", func(e *colly.HTMLElement) {
		e.ForEach(".js-item-image .item img", func(_ int, el *colly.HTMLElement) {
			imageUrls = append(imageUrls, strings.Split(el.Attr("src"), "#")[0])
		})
	})

	c.OnHTML(".sizes .sizes-select-wrapper .sizes-select-list", func(e *colly.HTMLElement) {
		e.ForEach(".list-inline li", func(_ int, el *colly.HTMLElement) {
			size := el.ChildText("span .value")
			sizes = append(sizes, size)
			if el.Attr("class") != "li-disabled" {
				inventories = append(inventories, domain.InventoryDAO{
					Quantity: 10,
					Size:     size,
				})
			}
		})
	})

	c.OnHTML(".swatches", func(e *colly.HTMLElement) {
		e.ForEach(".swatch", func(_ int, el *colly.HTMLElement) {
			color := el.ChildAttr("img", "title")
			colors = append(colors, color)
		})
	})

	c.OnHTML("#description .details-tab-content", func(e *colly.HTMLElement) {
		description["설명"] = e.ChildText("p")
	})

	c.OnHTML("#composition .details-tab-content", func(e *colly.HTMLElement) {
		texts := ""
		e.ForEach("ul li", func(idx int, el *colly.HTMLElement) {
			texts += el.Text
		})
		description["소재"] = texts
	})

	c.OnHTML("#fitting .details-tab-content", func(e *colly.HTMLElement) {
		texts := ""
		e.ForEach("ul li", func(idx int, el *colly.HTMLElement) {
			texts += el.Text
		})
		description["모델"] = texts
	})

	c.Visit(productUrl)
	return
}
