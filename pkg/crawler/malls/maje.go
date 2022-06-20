package malls

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
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

	c.OnHTML("#kr_maje_bandeaux_pages_marronnier_fin_operations > h2.sub-title.sub-title--desktop", func(e *colly.HTMLElement) {
		if e.Text == "VERKÄUFE KOMMEN BALD AUF MAJE.COM" || e.Text == "DIE LAST CHANCE WIRD BALD AUF MAJE.COM ZURÜCK SEIN" {
			log.Println("no product detected")
			return
		}
	})

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
			productName, productColor, composition, images, sizes, inventories, description, originalPrice, salesPrice := getMajeDetail(productDetailUrl)

			infos := map[string]string{
				"소재": composition,
				"색상": productColor,
			}
			productIdForDb := productId
			productNameForDb := productName
			if colorId != majeDeafultColor && colorId != "" {
				productIdForDb += "-" + colorName
				productNameForDb += " - " + colorName
			}

			addRequest := &productinfo.AddMetaInfoRequest{
				AlloffName:          productNameForDb,
				ProductID:           productIdForDb,
				ProductUrl:          productUrl,
				ProductType:         []domain.AlloffProductType{domain.Female},
				OriginalPrice:       originalPrice,
				DiscountedPrice:     salesPrice,
				CurrencyType:        domain.CurrencyEUR,
				Brand:               brand,
				Source:              source,
				AlloffCategory:      &domain.AlloffCategoryDAO{},
				Images:              images,
				Colors:              nil,
				DescriptionInfos:    infos,
				Sizes:               sizes,
				Inventory:           inventories,
				Information:         description,
				DescriptionImages:   images,
				IsForeignDelivery:   true,
				IsTranslateRequired: true,
				ModuleName:          source.CrawlModuleName,
				IsRemoved:           false,
				IsSoldout:           false,
			}

			totalProducts += 1
			productinfo.ProcessCrawlingInfoRequests(addRequest)
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

// 소재는 상품정보 제공 고시가 아니라 상품 설명으로 이동 220530
func getMajeDetail(productUrl string) (
	productName, productColor, composition string,
	images []string,
	sizes []string,
	inventories []*domain.InventoryDAO,
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

			stock := defaultStock
			if outOfStock {
				stock = 0
			}
			sizes = append(sizes, size)
			inventories = append(inventories, &domain.InventoryDAO{
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

			if key == "소재" {
				composition = text
			} else {
				if val, exists := description[key]; exists {
					description[key] = val + "\n" + text
				} else {
					description[key] = text
				}
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

	c.OnHTML("div.product-image-container ul.swiper-wrapper li", func(container *colly.HTMLElement) {
		src, exists := container.DOM.Find("source").First().Attr("data-srcset")
		if !exists {
			return
		}
		imageUrl := strings.Split(src, "?")[0]
		if !strings.Contains(imageUrl, "VIDEO") {
			images = append(images, imageUrl)
		}
	})

	// 색상
	c.OnHTML("div.product-variations div.value ul.swatches", func(e *colly.HTMLElement) {
		e.ForEach("li.selected", func(_ int, li *colly.HTMLElement) {
			productColor = li.ChildText("span")
		})
	})

	c.Visit(productUrl)
	return
}
