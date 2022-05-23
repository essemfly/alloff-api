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

func CrawlNiceClaup(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.niceclaup.co.kr"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	totalProducts := 0
	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".item_list li", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".figcaption .price .origin")
		discountedPriceStr := e.ChildText(".figcaption .price .saled")
		originalPrice := utils.ParsePriceString(originalPriceStr)
		discountedPrice := utils.ParsePriceString(discountedPriceStr)
		if originalPrice == 0 {
			originalPrice = discountedPrice
		}
		productIdRaw := e.ChildAttr("a", "href")
		productID := strings.Split(productIdRaw, "=")[1]
		productUrl := "https://www.niceclaup.co.kr/goods?id=" + productID
		title := e.ChildText(".figcaption .name")

		images, colors, sizes, inventories, description := getNiceClaupDetail(productUrl)

		addRequest := &product.AddMetaInfoRequest{
			AlloffName:      title,
			ProductID:       productID,
			ProductUrl:      productUrl,
			ProductType:     []domain.AlloffProductType{domain.Female},
			OriginalPrice:   float32(originalPrice),
			DiscountedPrice: float32(discountedPrice),
			CurrencyType:    domain.CurrencyKRW,
			Brand:           brand,
			Source:          source,
			// AlloffCategory:  nil,
			Images:              images,
			Colors:              colors,
			Sizes:               sizes,
			Inventory:           inventories,
			Information:         description,
			IsForeignDelivery:   false,
			IsTranslateRequired: false,
			ModuleName:          source.CrawlModuleName,
		}

		totalProducts += 1
		product.ProcessAddProductInfoRequests(addRequest)
	})

	c.OnHTML(".paging", func(e *colly.HTMLElement) {
		currentPageString := e.ChildText(".on")
		if currentPageString != "" {
			currentPage, _ := strconv.Atoi(currentPageString)
			url := source.CrawlUrl + "&page=" + strconv.Itoa(currentPage+1)
			c.Visit(url)
		}
	})

	c.Visit(source.CrawlUrl)

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}

func getNiceClaupDetail(productUrl string) (imageUrls, colors, sizes []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.niceclaup.co.kr"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	colors = nil
	description = map[string]string{}

	c.OnHTML(".item_dt_img", func(e *colly.HTMLElement) {
		e.ForEach("div img", func(_ int, el *colly.HTMLElement) {
			imageUrl := el.Attr("src")
			imageUrls = append(imageUrls, imageUrl)
		})
	})

	c.OnHTML(".item_dt_option_bt > .filter_input", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			size := el.ChildText("label div")
			disabled := el.ChildAttr("label input", "disabled")
			sizes = append(sizes, size)
			if disabled != "disabled" {
				inventories = append(inventories, domain.InventoryDAO{
					Size:     size,
					Quantity: 10,
				})
			}
		})
	})

	c.OnHTML(".item_dt_dinfo .fr", func(e *colly.HTMLElement) {
		e.ForEach("table tbody tr", func(_ int, el *colly.HTMLElement) {
			if !strings.Contains(el.ChildText("th"), "A/S") {
				description[el.ChildText("th")] = el.ChildText("td")
			}
		})
	})

	c.Visit(productUrl)
	return
}
