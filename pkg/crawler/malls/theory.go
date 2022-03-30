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

		productId := e.ChildAttr(".link", "href")
		productUrl := productBaseUrl + productId
		productCode := productId[:len(productId)-9]    // 뒤에 html 이랑 색상 코드 지우고
		productCode = productCode[len(productCode)-8:] // 앞에 카테고리 분류 지우면 순수 상품코드 추출

		productName := e.ChildText(".link")
		images, sizes, colors, inventories, description := getTheoryDetail(productUrl, productCode)

		addRequest := &product.ProductCrawlingAddRequest{
			Brand:               brand,
			Source:              source,
			ProductID:           productId,
			ProductName:         productName,
			ProductUrl:          productUrl,
			Images:              images,
			Sizes:               sizes,
			Inventories:         inventories,
			Colors:              colors,
			Description:         description,
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
		desc := e.ChildText(".description")
		desc = strings.TrimSpace(desc)
		desc = strings.Replace(desc, "Style #:", "", -1)
		desc = strings.Replace(desc, productCode, "", -1)
		desc = strings.Replace(desc, "-", "", -1)
		desc = strings.Replace(desc, "\n\n", " ", -1)
		desc = strings.Replace(desc, "\n", " ", -1)
		desc = strings.Replace(desc, "  ", " ", -1)
		desc = strings.Replace(desc, "  ", " ", -1)
		desc = strings.Trim(desc, "\n")
		desc = strings.TrimSpace(desc)
		description["설명"] = desc
	})

	// fit
	c.OnHTML(".pdp-fit", func(e *colly.HTMLElement) {
		fit := e.ChildText("ul")
		fit = strings.TrimSpace(fit)
		fit = strings.Replace(fit, "  ", " ", -1)
		description["핏"] = fit
	})

	// composition
	c.OnHTML("div.pdp-composition", func(e *colly.HTMLElement) {
		composition := ""
		e.ForEach(".pdp-details-info", func(_ int, el *colly.HTMLElement) {
			originalComposition := el.Text
			originalComposition = strings.TrimSpace(originalComposition)
			composition += originalComposition + " "
		})
		description["제품 주 소재"] = composition
	})

	c.OnHTML("div.pdp-care", func(e *colly.HTMLElement) {
		care := e.Text
		care = strings.TrimSpace(care)
		description["취급시 주의사항"] = care
	})

	err := c.Visit(productUrl)
	if err != nil {
		log.Println("err occured in crawling theory", err)
	}
	return
}
