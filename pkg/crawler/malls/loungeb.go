package malls

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
)

func CrawlLoungeB(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("lounge-b.co.kr"),
	)

	totalProducts := 0
	baseUrl := "https://lounge-b.co.kr/product/list.html"
	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".xans-product .prdList .item", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".description ul li:first-child")
		discountedPriceStr := e.ChildText(".description ul li:nth-child(2)")
		originalPriceStr = strings.Split(originalPriceStr, ": ")[1]
		discountedPriceStr = strings.Split(discountedPriceStr, ": ")[1]
		originalPrice := utils.ParsePriceString(originalPriceStr)
		discountedPrice := utils.ParsePriceString(discountedPriceStr)
		productID := e.Attr("id")
		productUrl := e.ChildAttr(".thumbnail a", "href")
		productUrl = "https://lounge-b.co.kr" + productUrl
		title := e.ChildText(".description .name a")
		title = strings.Split(title, "] ")[1]
		if discountedPrice == 0 {
			discountedPrice = originalPrice
		} else if originalPrice == 0 {
			originalPrice = discountedPrice
		}

		images, colors, sizes, inventories := getLoungebDetail(productUrl)

		if len(images) < 2 {
			return
		}

		mobileUrl := strings.Replace(productUrl, "https://lounge-b.co.kr", "https://m.lounge-b.co.kr", 1)

		addRequest := &productinfo.AddMetaInfoRequest{
			AlloffName:          title,
			ProductID:           productID,
			ProductUrl:          mobileUrl,
			ProductType:         []domain.AlloffProductType{domain.Female},
			OriginalPrice:       float32(originalPrice),
			DiscountedPrice:     float32(discountedPrice),
			CurrencyType:        domain.CurrencyKRW,
			Brand:               brand,
			Source:              source,
			AlloffCategory:      &domain.AlloffCategoryDAO{},
			Images:              images,
			Colors:              colors,
			Sizes:               sizes,
			Inventory:           inventories,
			Information:         nil,
			IsForeignDelivery:   false,
			IsTranslateRequired: false,
			ModuleName:          source.CrawlModuleName,
		}

		totalProducts += 1
		productinfo.ProcessAddProductInfoRequests(addRequest)
	})

	c.OnHTML(".xans-product-normalpaging a:nth-last-child(2)", func(e *colly.HTMLElement) {
		url := baseUrl + e.Attr("href")
		c.Visit(url)
	})

	c.Visit(source.CrawlUrl)

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}

type LoungebParser struct {
	Instances map[string]StockInfo
}

type StockInfo struct {
	StockPrice  string `json:"stock_price"`
	UseStock    bool   `json:"use_stock"`
	StockNumber int    `json:"stock_number"`
	OptionName  string `json:"option_name"`
	OptionValue string `json:"option_value"`
	// Option별로 가격이 다를 가능성이있음
}

func getLoungebDetail(productUrl string) (imageUrls, colors, sizes []string, inventories []*domain.InventoryDAO) {
	colors = nil
	c := colly.NewCollector(
		colly.AllowedDomains("lounge-b.co.kr"),
	)

	c.OnHTML(".listImg", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			imageUrl := el.ChildAttr("img", "src")
			if len(imageUrl) > 0 {
				imageUrls = append(imageUrls, "https:"+imageUrl)
			}
		})
	})

	c.OnHTML("script", func(e *colly.HTMLElement) {
		jscodesText := e.Text
		if strings.Contains(jscodesText, "option_stock_data = '") {
			stocks := strings.Split(strings.Split(jscodesText, "option_stock_data = '")[1], "';")[0]
			stocks = strings.ReplaceAll(stocks, "\\\"", "\"")
			stocks = strings.ReplaceAll(stocks, "\\\\", "\\")
			stockCrawlRseponse := &LoungebParser{}
			json.Unmarshal([]byte(stocks), &stockCrawlRseponse.Instances)
			for _, val := range stockCrawlRseponse.Instances {
				sizes = append(sizes, val.OptionValue)
				if val.StockNumber > 0 {
					inventories = append(inventories, &domain.InventoryDAO{
						Size:     val.OptionValue,
						Quantity: val.StockNumber,
					})
				}
			}
		}
	})

	c.OnHTML("#prdDetail .cont", func(e *colly.HTMLElement) {
		e.ForEach("p img", func(_ int, el *colly.HTMLElement) {
			imageUrl := el.Attr("src")
			if strings.HasPrefix(el.Attr("src"), "//") {
				imageUrl = "https:" + el.Attr("src")
			}
			imageUrls = append(imageUrls, imageUrl)
		})
	})

	c.Visit(productUrl)
	return
}
