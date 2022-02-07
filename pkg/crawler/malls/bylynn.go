package malls

import (
	"encoding/json"
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

func CrawlBylynn(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("bylynn.shop"),
	)

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	totalProducts := 0

	c.OnHTML(".productlist5 ul li ", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".cont .price span:first-child")
		originalPrice := utils.ParsePriceString(originalPriceStr)
		secondPriceStr := e.ChildText(".cont .price span:nth-child(2)")
		secondPrice := utils.ParsePriceString(secondPriceStr)

		discountedPriceStr := e.ChildText(".cont .price strong")
		discountedPriceStr = strings.Split(discountedPriceStr, " ")[1]
		discountedPrice := utils.ParsePriceString(discountedPriceStr)

		if secondPrice != 0 {
			if discountedPrice > secondPrice {
				discountedPrice = secondPrice
			}
		}

		if discountedPrice == 0 {
			discountedPrice = originalPrice
		} else if originalPrice == 0 {
			originalPrice = discountedPrice
		}

		urlString := e.ChildAttr(".img", "href")
		productID := strings.Split(urlString, "?")[1]
		productUrl := "https://bylynn.shop/goods/detail.do?" + productID
		mobileUrl := "https://bylynn.shop/mw/goods/detail.do?" + productID

		title, images, sizes, colors, inventories, description := getBylynnDetail(productUrl)

		addRequest := &product.ProductCrawlingAddRequest{
			Brand:         brand,
			Source:        source,
			ProductID:     productID,
			ProductName:   title,
			ProductUrl:    mobileUrl,
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
		product.AddProductInCrawling(addRequest)
	})

	c.OnHTML(".paging", func(e *colly.HTMLElement) {
		lastPageUrl := e.ChildAttr("ul .last a", "href")

		if lastPageUrl == "" {
			lastPageUrl = e.ChildAttr("ul li:last-child a", "href")
		}

		lastPage := strings.Split(strings.Split(lastPageUrl, "page=")[1], "&")[0]
		lastPageNum, _ := strconv.Atoi(lastPage)
		currentPage := e.ChildText("ul li a strong")
		currentPageNum, _ := strconv.Atoi(currentPage)

		if currentPageNum < lastPageNum {
			url := source.CrawlUrl + "&page=" + strconv.Itoa(currentPageNum+1)
			c.Visit(url)
		}

	})
	c.Visit(source.CrawlUrl)

	crawler.PrintCrawlResults(source, totalProducts)

	<-worker
	done <- true
}

type BylynnStock struct {
	STYCD    string `json:"STYCD"`
	SIZECD   string `json:"SIZECD"`
	COLCD    string `json:"COLCD"`
	SIZECDNM string `json:"SIZECDNM"`
	SALECNT  int    `json:"SALECNT"`
	STOCKQTY int    `json:"STOCKQTY"`
}

func getBylynnDetail(productUrl string) (title string, imageUrls []string, sizes, colors []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("bylynn.shop"),
	)

	description = map[string]string{}

	c.OnHTML(".scrollcontent .scroll .state", func(e *colly.HTMLElement) {
		title = e.ChildText(".title")
	})

	c.OnHTML("#detailGoodsImage .photolist", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			if el.Attr("class") != "movie" {
				imageUrl := el.ChildAttr("img", "src")
				if strings.HasPrefix(imageUrl, "//") {
					imageUrl = "https:" + imageUrl
				}
				if len(imageUrl) > 0 {
					imageUrls = append(imageUrls, imageUrl)
				}
			}
		})
	})

	c.OnHTML("#contents100", func(e *colly.HTMLElement) {
		colors := map[string]string{}
		sizeInventories := []domain.InventoryDAO{}

		e.ForEach(".options dl", func(_ int, el *colly.HTMLElement) {
			if el.ChildText("dt") == "COLOR" {
				el.ForEach("ul li", func(_ int, ele *colly.HTMLElement) {
					colorName := ele.ChildAttr("input", "value")
					colors[colorName] = ele.ChildText("label")
				})
			}
		})

		jscodesText := e.ChildText("script:nth-child(4)")
		stockLists := strings.Split(strings.Split(jscodesText, "storeStockList = ")[1], ";")[0]

		stockCrawlResponse := &[]BylynnStock{}
		json.Unmarshal([]byte(stockLists), stockCrawlResponse)
		for _, stock := range *stockCrawlResponse {
			sizes = append(sizes, stock.SIZECDNM)
			if stock.STOCKQTY > 0 {
				sizeInventories = append(inventories, domain.InventoryDAO{
					Quantity: stock.STOCKQTY,
					Size:     stock.SIZECDNM,
				})

			}
		}

		// Colors 별로 inventory가 맞는지는 모르겠다.
		for _, inv := range sizeInventories {
			for _, name := range colors {
				inventories = append(inventories, domain.InventoryDAO{
					Quantity: inv.Quantity,
					Size:     name + " - " + inv.Size,
				})
			}
		}
	})

	c.OnHTML(".tabcontents", func(e *colly.HTMLElement) {
		if e.ChildText("h3") == "PRODUCT INFO" {
			e.ForEach("table tbody tr", func(_ int, el *colly.HTMLElement) {
				if !strings.Contains(el.ChildText("th"), "A/S") {
					description[el.ChildText("th")] = el.ChildText("td")
				}

			})
			return
		}
	})

	c.Visit(productUrl)
	return
}
