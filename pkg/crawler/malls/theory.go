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
			originalPriceStr = strings.Replace(originalPriceStr, ",", "", -1)
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
		desc := ""
		// 내용 파싱하여 \n 으로 나누기
		originalDesc := e.ChildText(".description")
		originalDesc = strings.TrimSpace(originalDesc)
		originalDesc = strings.Replace(originalDesc, "Style #:", "", -1)
		originalDesc = strings.Replace(originalDesc, productCode, "", -1)
		originalDesc = strings.Replace(originalDesc, "\n\n", "\n", -1)
		originalDesc = strings.Replace(originalDesc, "  ", " ", -1)
		originalDesc = strings.Replace(originalDesc, "  ", " ", -1)
		originalDesc = strings.Trim(originalDesc, "\n")
		originalDesc = strings.TrimSpace(originalDesc)
		// \n 으로 나눈것 split 하여 line by line 으로 가공
		descSlice := strings.Split(originalDesc, "\n")
		for _, str := range descSlice {
			// https://outlet.theory.com/piazza-jkt-2/L091101R_Q1G.html 같이 description 에 공백 한글자가 포함되어 들어올 수 있음
			if str != " " {
				str = strings.TrimSpace(str)
				str = strings.Trim(str, "\n")
				str += "\n"
				desc += str
			}
		}
		desc = strings.TrimRight(desc, "\n")
		description["설명"] = desc
	})

	// fit
	c.OnHTML(".pdp-fit", func(e *colly.HTMLElement) {
		fit := ""
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			lineText := el.Text
			lineText = "- " + lineText
			lineText += "\n"
			fit += lineText
		})
		fit = strings.TrimRight(fit, "\n")
		description["핏"] = fit
	})

	// composition
	c.OnHTML("div.pdp-composition", func(e *colly.HTMLElement) {
		composition := ""
		e.ForEach(".pdp-details-info", func(_ int, el *colly.HTMLElement) {
			originalComposition := el.Text
			originalComposition = strings.TrimSpace(originalComposition)
			originalComposition = "- " + originalComposition
			composition += originalComposition + "\n"
		})
		composition = strings.TrimRight(composition, "\n")
		description["소재"] = composition
	})

	c.OnHTML("div.pdp-care", func(e *colly.HTMLElement) {
		care := ""

		// 내용 파싱하여 마침표 기준으로 나누기
		originalCare := e.Text
		originalCare = strings.TrimSpace(originalCare)
		// 마침표로 나눈것 split 하여 line by line 으로 가공
		careSlice := strings.Split(originalCare, ".")
		for id, str := range careSlice {
			str = strings.TrimSpace(str)
			// 마침표로 구분한 careSlice 의 마지막 요소가 불필요하게 인입되기 때문에, 마지막 요소를 지워주는 작업이 필요함
			if str != "Imported" && id != len(careSlice)-1 {
				str = "- " + str
				care += str + "\n"
			}
		}

		care = strings.TrimRight(care, "\n")
		description["취급 시 주의사항"] = care
	})

	err := c.Visit(productUrl)
	if err != nil {
		log.Println("err occured in crawling theory", err)
	}
	return
}
