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
)

func CrawlNiceClaup(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.niceclaup.co.kr"),
	)

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

		addRequest := crawler.ProductsAddRequest{
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

		crawler.AddProduct(addRequest)
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

	<-worker
	done <- true
}

func getNiceClaupDetail(productUrl string) (imageUrls, colors, sizes []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.niceclaup.co.kr"),
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