package malls

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

const (
	sandroAllowedDomain           = "de.sandro-paris.com"
	sandroRecommendingDescription = "passt perfekt zum"
	sandroReferenceDescription    = "Référence"
	sandroInlineSizeDescription   = "Größenentsprechung"
	sandroModelDescription        = "Das Model"
	sandroDefaultcolor            = "__default__"
)

func CrawlSandro(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := getCollyCollector(sandroAllowedDomain)
	totalProducts := 0

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".product-info", func(e *colly.HTMLElement) {
		colorMap := map[string]string{sandroDefaultcolor: sandroDefaultcolor}
		e.ForEach("ul.swatch-list li", func(_ int, li *colly.HTMLElement) {
			colorId := li.ChildAttr("img", "data-colorid")
			colorName := li.ChildAttr("img", "data-colorname")
			colorMap[colorId] = colorName
		})
		nameUrl := e.ChildAttr(".name-link", "href")
		urlNodes := strings.Split(nameUrl, "/")
		productId := strings.Split(urlNodes[len(urlNodes)-1], ".html")[0]
		colorMapIsEmpty := len(colorMap) == 1
		for colorId, colorName := range colorMap {
			if !colorMapIsEmpty && colorId == sandroDefaultcolor {
				continue
			}
			totalProducts++
			var productDetailUrl string
			var productUrl string
			if colorId == majeDeafultColor || colorId == "" {
				// 색상이 없을 땐 원본 URL 그대로 사용
				productDetailUrl = nameUrl
				productUrl = nameUrl
			} else {
				// 색상이 있을 땐 파싱된 swatch 사용하여 URL 재구성
				productDetailUrl = getSandroDetailUrl(productId, colorId)
				productUrl = getSandroProductUrl(productId, colorId)
			}
			productName, images, sizes, inventories, description, originalPrice, salesPrice := getSandroDetail(productDetailUrl)

			productIdForDb := productId
			productNameForDb := productName
			if colorId != majeDeafultColor && colorId != "" {
				productIdForDb += "-" + colorName
				productNameForDb += " - " + colorName
			}
			addRequest := &product.ProductCrawlingAddRequest{
				Brand:               brand,
				Images:              images,
				Sizes:               sizes,
				Inventories:         inventories,
				Description:         description,
				OriginalPrice:       originalPrice,
				SalesPrice:          salesPrice,
				CurrencyType:        domain.CurrencyEUR,
				Source:              source,
				ProductID:           productIdForDb,
				ProductName:         productNameForDb,
				ProductUrl:          productUrl,
				IsTranslateRequired: true,
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
	if colorId == sandroDefaultcolor || colorId == "" {
		return fmt.Sprintf("https://de.sandro-paris.com/on/demandware.store/Sites-Sandro-DE-Site/de_DE/Product-Variation?pid=%s&Quantity=1", productId)
	}
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
	c := getCollyCollector(sandroAllowedDomain)

	// 상품명
	c.OnHTML("h1.prod-title", func(h1 *colly.HTMLElement) {
		productName = h1.Text
	})

	alreadyCrawledSize := false
	// 사이즈 & 재고
	c.OnHTML("ul.size", func(e *colly.HTMLElement) {
		if alreadyCrawledSize {
			return
		}
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
		alreadyCrawledSize = true
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
			if strings.Contains(text, sandroRecommendingDescription) {
				// "이런 상품과 함께 입으면 좋습니다"는 설명은 쓰지 않음
				return
			}
			if strings.Contains(text, sandroReferenceDescription) {
				// 레퍼런스 넘버는 쓰지 않음
				return
			}
			if strings.Contains(text, sandroInlineSizeDescription) {
				// 설명 내부의 사이즈 정보는 쓰지 않음
				return
			}
			key := descriptionKey
			if strings.Contains(text, sandroModelDescription) {
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
	priceParsed := false
	c.OnHTML("div.product-price", func(div *colly.HTMLElement) {
		if priceParsed {
			// Only parse the first-appearing price information
			return
		}

		salesPrice = parseEuro(div.ChildText("span.price-sales"))
		originalPrice = parseEuro(div.ChildText("span.price-standard"))

		priceParsed = true
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
