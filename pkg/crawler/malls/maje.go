package malls

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

const (
	majeAllowedDomain    = "de.maje.com"
	majeReferenceRegex   = `(?m)Ref\s?:\s?.*`
	majeModelDescription = "trägt Größe"
	majeDeafultColor     = "__default__"
)

func CrawlMaje(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := getCollyCollector(majeAllowedDomain)
	totalProducts := 0

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".infosProduct", func(e *colly.HTMLElement) {
		colorMap := map[string]string{majeDeafultColor: majeDeafultColor}
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
			if !colorMapIsEmpty && colorId == majeDeafultColor {
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
				productDetailUrl = getMajeDetailUrl(productId, colorId)
				productUrl = getMajeProductUrl(productId, colorId)
			}
			productName, images, sizes, inventories, description, originalPrice, salesPrice := getMajeDetail(productDetailUrl)

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
		log.Println("err occured in crawling maje", err)
	}

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}

func getMajeDetailUrl(productId string, colorId string) string {
	return getMajeProductUrl(productId, colorId) + "&ContentTarget=swiper&format=ajax"
}

func getMajeProductUrl(productId string, colorId string) string {
	if colorId == majeDeafultColor || colorId == "" {
		return fmt.Sprintf("https://de.maje.com/on/demandware.store/Sites-Maje-DE-Site/de/Product-Variation?pid=%s&Quantity=1", productId)
	}
	return fmt.Sprintf("https://de.maje.com/on/demandware.store/Sites-Maje-DE-Site/de/Product-Variation?pid=%s&dwvar_%s_color=%s&Quantity=1", productId, productId, colorId)
}

func getMajeDetail(productUrl string) (
	productName string,
	images []string,
	sizes []string,
	inventories []domain.InventoryDAO,
	description map[string]string,
	originalPrice float32,
	salesPrice float32,
) {
	c := getCollyCollector(majeAllowedDomain)

	// 상품명
	c.OnHTML("h1.productName", func(h1 *colly.HTMLElement) {
		productName = strings.TrimSpace(h1.Text)
	})

	alreadyCrawledSize := false
	// 사이즈 & 재고
	c.OnHTML("ul.size", func(e *colly.HTMLElement) {
		if alreadyCrawledSize {
			return
		}
		e.ForEach("li.emptyswatch", func(_ int, li *colly.HTMLElement) {
			outOfStock := strings.Contains(li.Attr("class"), "unselectable")
			size := li.ChildText("div.defaultSize")
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

	description = map[string]string{}

	// 설명 1
	alreadyCrawledDescription1 := false
	c.OnHTML("ul.product-short-info", func(ul *colly.HTMLElement) {
		if alreadyCrawledDescription1 {
			// This info appears twice. Crawl only once.
			return
		}
		ul.ForEach("li", func(_ int, li *colly.HTMLElement) {
			text := strings.Replace(li.Text, `•`, "", -1)
			text = strings.TrimSpace(text)

			key := "설명"
			if strings.Contains(text, majeModelDescription) {
				key = "모델"
			}

			if val, exists := description[key]; exists {
				description[key] = val + "\n" + text
			} else {
				description[key] = text
			}
		})
		alreadyCrawledDescription1 = true
	})

	// 설명 2
	c.OnHTML("div.wrapper-tabs-product ul li h2", func(h2 *colly.HTMLElement) {
		descriptionType := strings.TrimSpace(h2.Text)
		descriptionKey := ""
		switch descriptionType {
		case "Beschreibung":
			descriptionKey = "설명"
		case "Hauptstoff & Pflege":
			descriptionKey = "소재"
		}
		if descriptionKey == "" {
			// 다른 설명란의 정보는 쓰지 않음
			return
		}

		h2.DOM.Next().Each(func(i int, s *goquery.Selection) {
			itemprop, _ := s.Attr("itemprop")
			if descriptionKey == "설명" && itemprop != "description" {
				// Use description div only for 설명.
				return
			}

			key := descriptionKey
			re := regexp.MustCompile(majeReferenceRegex)
			text := re.ReplaceAllString(strings.TrimSpace(s.Text()), "")

			textNodes := strings.Split(text, "\n")
			trimRe := regexp.MustCompile(`\s*`)
			trimmedNodes := utils.Map(textNodes, func(s string) string {
				noWhiteSpaces := trimRe.ReplaceAllString(s, "")
				if noWhiteSpaces == "" {
					return ""
				}
				return strings.TrimSpace(s)
			})
			joinnableNodes := []string{}
			for _, trimmedNode := range trimmedNodes {
				if trimmedNode != "" {
					joinnableNodes = append(joinnableNodes, trimmedNode)
				}
			}

			text = strings.Join(joinnableNodes, "\n")

			if val, exists := description[key]; exists {
				description[key] = val + "\n" + text
			} else {
				description[key] = text
			}
		})

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
	c.OnHTML("ul.swiper-wrapper li", func(container *colly.HTMLElement) {
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
