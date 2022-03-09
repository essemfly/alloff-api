package malls

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

const (
	allowedDomain           = "de.sandro-paris.com"
	userAgent               = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"
	defaultStock            = 100
	recommendingDescription = "passt perfekt zum"
	referenceDescription    = "Référence"
	inlineSizeDescription   = "Größenentsprechung"
	modelDescription        = "Das Model"
)

func CrawlSandro(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := getCollyCollector()
	totalProducts := 0

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".product-info", func(e *colly.HTMLElement) {
		colorMap := map[string]string{}
		e.ForEach("ul.swatch-list li", func(_ int, li *colly.HTMLElement) {
			colorId := li.ChildAttr("img", "data-colorid")
			colorName := li.ChildAttr("img", "data-colorname")
			colorMap[colorId] = colorName
		})
		nameUrl := e.ChildAttr(".name-link", "href")
		urlNodes := strings.Split(nameUrl, "/")
		productId := strings.Split(urlNodes[len(urlNodes)-1], ".html")[0]
		for colorId, colorName := range colorMap {
			totalProducts++
			productDetailUrl := getSandroDetailUrl(productId, colorId)
			productName, images, sizes, inventories, description, originalPrice, salesPrice := getSandroDetail(productDetailUrl)
			addRequest := &product.ProductCrawlingAddRequest{
				Brand:         brand,
				Images:        images,
				Sizes:         sizes,
				Inventories:   inventories,
				Description:   description,
				OriginalPrice: originalPrice,
				SalesPrice:    salesPrice,
				CurrencyType:  domain.CurrencyEUR,
				Source:        source,
				ProductID:     productId,
				ProductName:   productName + " - " + colorName,
				ProductUrl:    getSandroProductUrl(productId, colorId),
			}
			totalProducts += 1
			product.AddProductInCrawling(addRequest)
		}
	})

	err = c.Visit(source.CrawlUrl)
	if err != nil {
		log.Println("err occured in crawling sandro", err)
	}

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}

func getSandroDetailUrl(productId string, colorId string) string {
	return getSandroProductUrl(productId, colorId) + "&format=ajax"
}

func getSandroProductUrl(productId string, colorId string) string {
	return fmt.Sprintf("https://de.sandro-paris.com/on/demandware.store/Sites-Sandro-DE-Site/de_DE/Product-Variation?pid=%s&dwvar_%s_color=%s&Quantity=1", productId, productId, colorId)
}

func getSandroDetail(productUrl string) (
	productName string,
	images []string,
	sizes []string,
	inventories []domain.InventoryDAO,
	description map[string]string,
	originalPrice float32,
	salesPrice float32,
) {
	c := getCollyCollector()

	// 상품명
	c.OnHTML("h1.prod-title", func(h1 *colly.HTMLElement) {
		productName = h1.Text
	})

	// 사이즈 & 재고
	c.OnHTML(".siz-list-container", func(e *colly.HTMLElement) {
		e.ForEach("li.emptyswatch", func(_ int, li *colly.HTMLElement) {
			outOfStock := strings.Contains(li.Attr("class"), "notinstock")
			size := li.ChildText("span.sizeDisplayValue")
			stock := defaultStock
			if outOfStock {
				stock = 0
			}
			sizes = append(sizes, size)
			inventories = append(inventories, domain.InventoryDAO{
				Size:     size,
				Quantity: stock,
			})
		})
	})

	// 설명
	description = map[string]string{}
	c.OnHTML(".titleDescPr.toggleMe", func(h2 *colly.HTMLElement) {
		descriptionType := strings.ToLower(strings.TrimSpace(h2.Text))
		descriptionKey := ""
		switch descriptionType {
		case "produktbeschreibung":
			descriptionKey = "설명"
		case "zusammensetzung":
			descriptionKey = "소재"
		}
		if descriptionKey == "" {
			// 다른 설명란의 정보는 쓰지 않음
			return
		}

		html, _ := h2.DOM.Next().Find(".detaildesc").Html()
		nodes := strings.Split(html, "<br/>")
		for _, node := range nodes {
			text := strings.Replace(node, `<br \=""/>`, "", -1)
			text = strings.Replace(text, `•`, "", -1)
			text = strings.TrimSpace(text)
			if strings.Contains(text, recommendingDescription) {
				// "이런 상품과 함께 입으면 좋습니다"는 설명은 쓰지 않음
				return
			}
			if strings.Contains(text, referenceDescription) {
				// 레퍼런스 넘버는 쓰지 않음
				return
			}
			if strings.Contains(text, inlineSizeDescription) {
				// 설명 내부의 사이즈 정보는 쓰지 않음
				return
			}
			key := descriptionKey
			if strings.Contains(text, modelDescription) {
				key = "모델"
			}

			if val, exists := description[key]; exists {
				description[key] = val + "\n" + text
			} else {
				description[key] = text
			}
		}
	})

	// 가격
	c.OnHTML("span.price-sales", func(span *colly.HTMLElement) {
		// 가격을 찾을 수 없으면 panic (MustCompile)
		re := regexp.MustCompile("[0-9]+")
		_salesPrice, _ := strconv.ParseFloat(re.FindAllString(span.Text, -1)[0], 32)
		salesPrice = float32(_salesPrice)
	})
	c.OnHTML("span.price-standard", func(span *colly.HTMLElement) {
		// 가격을 찾을 수 없으면 panic (MustCompile)
		re := regexp.MustCompile("[0-9]+")
		_originalPrice, _ := strconv.ParseFloat(re.FindAllString(span.Text, -1)[0], 32)
		originalPrice = float32(_originalPrice)
	})

	// 이미지
	c.OnHTML("div.image-container", func(container *colly.HTMLElement) {
		src, exists := container.DOM.Find("source").First().Attr("data-srcset")
		if !exists {
			return
		}
		imageUrl := strings.Split(src, "?")[0]
		images = append(images, imageUrl)
	})

	c.Visit(productUrl)
	return
}

func getCollyCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomain),
		colly.UserAgent(userAgent),
	)
	return c
}
