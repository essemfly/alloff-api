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
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	productBaseUrl := "https://outlet.theory.com"
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

		discountedPriceStr := e.ChildText(".price .sales .value")
		discountedPriceStr = strings.Replace(discountedPriceStr, "$", "", -1)
		discountedPrice, err := strconv.ParseFloat(discountedPriceStr, 32)
		if err != nil {
			log.Println("err", err)
			return
		}
		if discountedPrice == 0 {
			discountedPrice = originalPrice
		} else if originalPrice == 0.0 {
			originalPrice = discountedPrice
		}

		//		productID := e.ChildAttr(".info_area a", "href")
		productId := e.ChildAttr(".link", "href")
		productUrl := productBaseUrl + productId
		productCode := productId[:len(productId)-9]

		productName := e.ChildText(".link")
		images, sizes, colors, inventories, description := getTheoryDetail(productUrl, productCode)
		log.Println(images)
		log.Println(sizes)
		log.Println(colors)
		log.Println(inventories)
		log.Println(description)

		addRequest := &product.ProductCrawlingAddRequest{
			Brand:               brand,
			Source:              source,
			ProductID:           productId,
			ProductName:         productName,
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

// https://outlet.theory.com/on/demandware.store/Sites-theory_outlet-Site/default/Product-Variation?dwvar_L094505R__G8E_color=G8E&dwvar_L094505R__G8E_size=P&pid=L094505R_G8E&quantity=1
func getTheoryDetail(productUrl, productCode string) (imageUrls, sizes, colors []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("outlet.theory.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	description = map[string]string{}

	// images
	c.OnHTML(".js-pdp-vertical-gallery", func(e *colly.HTMLElement) {
		e.ForEach(".js-pdp-primary-image", func(_ int, el *colly.HTMLElement) {
			imageUrls = append(imageUrls, "https:"+el.Attr("src"))
		})
	})

	// sizes / inventories
	c.OnHTML(".js-size-attributes", func(e *colly.HTMLElement) {
		e.ForEach(".swatch-size ", func(_ int, el *colly.HTMLElement) {
			size := el.ChildText("span")
			sizes = append(sizes, size)
			unselectableSize := el.ChildText(".unselectable")
			// 해당 사이즈가 선택 불가능한 사이즈가 아닐 경우 재고가 있는 걸로 판단
			if size != unselectableSize {
				inventories = append(inventories, domain.InventoryDAO{
					Quantity: 10,
					Size:     size,
				})
			}
		})
	})

	// colors : 같은 디자인의 별개 색상 상품을 별개의 상품 id 로 구분하고 있음, 이때 상품 색상은 다 넣어줘야하는지 ?
	c.OnHTML(".attributes", func(e *colly.HTMLElement) {
		color := e.ChildText(".selected-attr")
		colors = append(colors, color)
	})

	// descriptions
	c.OnHTML(".description-and-detail", func(e *colly.HTMLElement) {
		originalDesc := e.ChildText(".description")
		originalDesc = strings.Replace(originalDesc, "Style #:", "", -1)
		originalDesc = strings.Replace(originalDesc, productCode, "", -1)
		originalDesc = strings.Replace(originalDesc, "-", "", -1)
		originalDesc = strings.Replace(originalDesc, "\n\n", "\n", -1)
		originalDesc = strings.Trim(originalDesc, "\n")
		originalDesc = strings.TrimSpace(originalDesc)
		description["설명"] = originalDesc
	})

	err := c.Visit(productUrl)
	if err != nil {
		log.Println("err occured in crawling theory", err)
	}
	return
}
