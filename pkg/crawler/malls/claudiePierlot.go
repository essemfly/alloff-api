package malls

import (
	"encoding/json"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
)

type ImagesURLs struct {
	ImagesURLs string `json:"imagesURL"`
}

func CrawlClaudiePierlot(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("de.claudiepierlot.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.SetRequestTimeout(45 * time.Second)
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

		sizes, inventories, description, colors := getClaudiePierlotDetail(productUrl)

		addRequest := &productinfo.AddMetaInfoRequest{
			AlloffName:           productName,
			ProductID:            productId,
			ProductUrl:           productUrl,
			ProductType:          []domain.AlloffProductType{domain.Female},
			OriginalPrice:        float32(originalPrice),
			DiscountedPrice:      float32(discountedPrice),
			CurrencyType:         domain.CurrencyEUR,
			Brand:                brand,
			Source:               source,
			AlloffCategory:       &domain.AlloffCategoryDAO{},
			Images:               images,
			Colors:               colors,
			Sizes:                sizes,
			Inventory:            inventories,
			Information:          description,
			IsTranslateRequired:  true,
			ModuleName:           source.CrawlModuleName,
			IsRemoved:            false,
			IsSoldout:            false,
			IsForeignDelivery:    true,
			EarliestDeliveryDays: 14,
			LatestDeliveryDays:   21,
			IsRefundPossible:     true,
			RefundFee:            100000,
		}

		totalProducts += 1
		productinfo.ProcessCrawlingInfoRequests(addRequest)

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
	sizes []string, inventories []*domain.InventoryDAO, description map[string]string, colors []string,
) {
	c := colly.NewCollector(
		colly.AllowedDomains("de.claudiepierlot.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	sizes = []string{}
	description = map[string]string{}
	inventories = []*domain.InventoryDAO{}
	colors = []string{}

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

	// 주의사항
	c.OnHTML("div.activate-div-tab-3", func(e *colly.HTMLElement) {
		e.ForEach("ul", func(i int, el *colly.HTMLElement) {
			// i == 0 > 소재
			// i == 1 > 취급 시 주의사항
			if i == 0 {
				desc := ""
				composition := el.Text
				composition = strings.Trim(composition, "\n")
				lineComposition := strings.Split(composition, "\n")
				for _, str := range lineComposition {
					str = "- " + str
					desc += str + "\n"
				}
				desc = strings.TrimRight(desc, "\n")
				description["소재"] = desc
			}

			if i == 1 {
				desc := ""
				care := el.Text
				care = strings.Trim(care, "\n")
				lineCare := strings.Split(care, "\n")
				for _, str := range lineCare {
					if str != "" {
						str = "- " + str
						desc += str + "\n"
					}
				}
				desc = strings.TrimRight(desc, "\n")
				description["취급 시 주의사항"] = desc
			}
		})
	})

	// 사이즈
	c.OnHTML(".siz-list-container .size", func(e *colly.HTMLElement) {
		modelSize := e.ChildText(".modelSize")
		modelSize = strings.TrimSpace(modelSize)
		modelSize = strings.Trim(modelSize, "\n")
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			liClass := el.Attr("class")
			liClassArray := strings.Split(liClass, " ")
			unselectable := false
			if len(liClassArray) > 1 {
				if liClassArray[1] == "unselectable" {
					unselectable = true
				}
			}
			isSize := true
			stock := defaultStock
			if unselectable {
				stock = 0
			}
			size := el.Text
			size = strings.TrimSpace(size)
			size = strings.Trim(size, "\n")
			size = strings.Replace(size, "\n\n\nNachricht sobald verfügbar", "", -1)

			if modelSize == size {
				isSize = false
			}

			// 파싱한 사이즈의 값이 "선택"이 아니고, "모델 스펙"이 아닐때에만 입력
			if size != "Größe" {
				if isSize {
					size = strings.Replace(size, " ", "", -1)

					isDigit := regexp.MustCompile(`^\d*\.?\d+$`)
					if isDigit.MatchString(size) {
						intSize, _ := strconv.Atoi(size)
						// if size system is form of 32, 33, 34 ....
						if intSize > 20 {
							size = "FR" + size
						} else {
							// if size system is form of 0, 1, 2, 3 ....
							size = "EU" + size
						}
					}

					// if size is form of DE32/FR42
					sizeArray := strings.Split(size, "/")
					if len(sizeArray) >= 2 {
						size = sizeArray[0]
					}

					sizes = append(sizes, size)
					inventories = append(inventories, &domain.InventoryDAO{
						Size:     size,
						Quantity: stock,
					})
				}
			}
		})
	})

	// 색상
	c.OnHTML(".color button.current-attribute", func(e *colly.HTMLElement) {
		color := e.Text
		color = strings.Trim(color, "\n")
		color = strings.TrimSpace(color)
		colors = append(colors, color)
	})

	err := c.Visit(productUrl)
	if err != nil {
		log.Println("err occurred in crawling claudie pierlot : ", err)
	}
	return
}
