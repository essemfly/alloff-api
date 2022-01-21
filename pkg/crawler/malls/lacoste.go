package malls

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func CrawlLacoste(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.lacoste.com"),
	)

	totalProducts := 0
	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".product-tile", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".ff-semibold .standard-price")
		discountedPriceStr := e.ChildText(".ff-semibold .sales-price")
		originalPrice := utils.ParsePriceString(originalPriceStr)
		discountedPrice := utils.ParsePriceString(discountedPriceStr)
		productID := e.Attr("data-sku")
		productUrl := e.ChildAttr(".js-product-tile-link", "href")
		title := e.ChildText(".fs--small a")

		images, colors, sizes, inventories, description := getLacosteDetail(productUrl)

		addRequest := product.ProductsAddRequest{
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
			CurrencyType:  domain.CurrencyKRW,
		}

		totalProducts += 1
		crawler.AddProduct(addRequest)
	})

	c.OnHTML(".pagination", func(e *colly.HTMLElement) {
		currentPageStr := e.ChildText(".is-active")
		currentPage, _ := strconv.Atoi(currentPageStr)
		url := source.CrawlUrl + "?page=" + strconv.Itoa(currentPage+1)
		c.Visit(url)

	})
	c.Visit(source.CrawlUrl)

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}

func getLacosteDetail(productUrl string) (imageUrls, colors, sizes []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.lacoste.com"),
	)

	description = map[string]string{}

	c.OnHTML(".js-pdp-gallery > .js-slideshow", func(e *colly.HTMLElement) {
		e.ForEach(".slide .slide-content picture", func(_ int, el *colly.HTMLElement) {
			imageUrl := el.ChildAttr("img", "data-src")
			if len(imageUrl) > 0 {
				imageUrls = append(imageUrls, "https:"+imageUrl)
			}
		})
	})

	c.OnHTML(".popin-size-list > .js-pdp-popin-size-list", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			sizeRaw := el.ChildText("button")
			size := strings.Trim(sizeRaw, " ")
			attrs := el.ChildAttr("button", "class")
			sizes = append(sizes, size)
			if !strings.Contains(attrs, "is-disabled") {
				inventories = append(inventories, domain.InventoryDAO{
					Quantity: 10,
					Size:     size,
				})
			}
		})
	})

	// c.OnHTML(".item-flag", func(e *colly.HTMLElement) {
	// 	soldoutFlag := e.Text
	// 	soldoutFlag = strings.Trim(soldoutFlag, " ")
	// 	soldoutFlag = strings.Trim(soldoutFlag, "\n")
	// 	if soldoutFlag == "Sold out" {
	// 		soldout = true
	// 	}
	// })

	c.OnHTML("#ada-pdp-care div ul", func(e *colly.HTMLElement) {
		description["상품상세정보"] = e.Text
	})

	c.Visit(productUrl)
	return
}
