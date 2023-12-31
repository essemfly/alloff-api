package malls

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.uber.org/zap"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func CrawlIntrend(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("it.intrend.it"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	currentPageNum := 0
	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	totalProducts := 0

	c.OnHTML(".js-product-list .js-pagination-container .js-product-card", func(e *colly.HTMLElement) {
		originalPriceStr := e.ChildText(".full-price")
		originalPrice := 0.0
		if originalPriceStr != "" {
			originalPriceStr = strings.Split(originalPriceStr, " ")[1]
			originalPriceStr = strings.Replace(originalPriceStr, ".", "", -1)
			originalPriceStr = strings.Replace(originalPriceStr, ",", ".", -1)
			originalPrice, err = strconv.ParseFloat(originalPriceStr, 32)
			if err != nil {
				msg := fmt.Sprintf("err on parse original price %v : ", originalPrice)
				config.Logger.Error(msg, zap.Error(err))
				return
			}
		}

		discountedPriceStr := e.ChildText(".price")
		discountedPriceStr = strings.Split(discountedPriceStr, " ")[1]
		discountedPriceStr = strings.Replace(discountedPriceStr, ".", "", -1)
		discountedPriceStr = strings.Replace(discountedPriceStr, ",", ".", -1)
		discountedPrice, err := strconv.ParseFloat(discountedPriceStr, 32)
		if err != nil {
			msg := fmt.Sprintf("err on parse discount price %v : ", originalPrice)
			config.Logger.Error(msg, zap.Error(err))
			return
		}

		if discountedPrice == 0 {
			discountedPrice = originalPrice
		} else if originalPrice == 0.0 {
			originalPrice = float64(genOriginalPrice(float32(discountedPrice)))
		}

		productID := e.Attr("data-product-id")
		productUrl := "https://it.intrend.it" + e.ChildAttr(".js-anchor", "href")

		title, composition, productColor, images, sizes, colors, inventories, description := getIntrendDetail(productUrl)
		if len(sizes) == 0 {
			inventories = append(inventories, &domain.InventoryDAO{
				Size:     "normal",
				Quantity: 1,
			})
		}

		infos := map[string]string{
			"소재": composition,
			"색상": productColor,
		}

		// forbidden 403 case
		if title == "" {
			msg := fmt.Sprintf("not allowed access by intrend server on : %s\n", source.CrawlUrl)
			config.Logger.Error(msg)
			return
		}

		addRequest := &productinfo.AddMetaInfoRequest{
			AlloffName:           title,
			ProductID:            productID,
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
			DescriptionInfos:     infos,
			Sizes:                sizes,
			Inventory:            inventories,
			Information:          description,
			DescriptionImages:    images,
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

	c.OnHTML(".js-pager .container-fluid ul", func(e *colly.HTMLElement) {
		lastPageStr := e.ChildAttr("li:last-child a", "data-page")
		lastPageNum, _ := strconv.Atoi(lastPageStr)
		if currentPageNum < lastPageNum {
			currentPageNum += 1
			url := source.CrawlUrl + "?page=" + strconv.Itoa(currentPageNum)
			c.Visit(url)
		}
	})
	err = c.Visit(source.CrawlUrl)
	if err != nil {
		config.Logger.Error("error occurred in crawl intrend ", zap.Error(err))
	}

	crawler.PrintCrawlResults(source, totalProducts)

	<-worker
	done <- true
}

type IntrendStock struct {
	STYCD    string `json:"STYCD"`
	SIZECD   string `json:"SIZECD"`
	COLCD    string `json:"COLCD"`
	SIZECDNM string `json:"SIZECDNM"`
	SALECNT  int    `json:"SALECNT"`
	STOCKQTY int    `json:"STOCKQTY"`
}

func getIntrendDetail(productUrl string) (title, composition, productColor string, imageUrls []string, sizes, colors []string, inventories []*domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("it.intrend.it"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	isDigit := regexp.MustCompile(`^\d*\.?\d+$`)

	description = map[string]string{}

	c.OnHTML(".pp_mod-prod-desc-head", func(e *colly.HTMLElement) {
		title = e.ChildText(".title")
	})

	c.OnHTML(".product-gallery", func(e *colly.HTMLElement) {
		e.ForEach(".js-item-image .item img", func(_ int, el *colly.HTMLElement) {
			imageUrls = append(imageUrls, strings.Split(el.Attr("src"), "#")[0])
		})
	})

	c.OnHTML(".sizes .sizes-select-wrapper .sizes-select-list", func(e *colly.HTMLElement) {
		e.ForEach(".list-inline li", func(_ int, el *colly.HTMLElement) {
			size := el.ChildText("span .value")
			if isDigit.MatchString(size) {
				size = "IT" + size
			}
			sizes = append(sizes, size)
			if el.Attr("class") != "li-disabled" {
				inventories = append(inventories, &domain.InventoryDAO{
					Quantity: defaultStock,
					Size:     size,
				})
			}
		})
	})

	c.OnHTML(".swatches", func(e *colly.HTMLElement) {
		e.ForEach(".swatch", func(_ int, el *colly.HTMLElement) {
			color := el.ChildAttr("img", "title")
			colors = append(colors, color)
		})
	})

	c.OnHTML(".swatches .title", func(e *colly.HTMLElement) {
		productColor = e.ChildText(".value")
	})

	c.OnHTML("#description .details-tab-content", func(e *colly.HTMLElement) {
		description["설명"] = e.ChildText("p")
	})

	c.OnHTML("#composition .details-tab-content", func(e *colly.HTMLElement) {
		texts := ""
		e.ForEach("ul li", func(idx int, el *colly.HTMLElement) {
			texts += el.Text
		})
		composition = texts
	})

	c.OnHTML("#fitting .details-tab-content", func(e *colly.HTMLElement) {
		texts := ""
		e.ForEach("ul li", func(idx int, el *colly.HTMLElement) {
			texts += el.Text
		})
		description["모델"] = texts
	})

	c.Visit(productUrl)
	return
}

func genOriginalPrice(discountedPrice float32) float32 {
	originalPrice := discountedPrice
	if discountedPrice <= 20 {
		disRate := genRandRate(78, 80)
		originalPrice = discountedPrice * disRate
	} else if 20 < discountedPrice && discountedPrice <= 50 {
		disRate := genRandRate(70, 78)
		originalPrice = discountedPrice * disRate
	} else if 50 < discountedPrice && discountedPrice <= 70 {
		disRate := genRandRate(65, 75)
		originalPrice = discountedPrice * disRate
	} else if 70 < discountedPrice && discountedPrice <= 100 {
		disRate := genRandRate(55, 72)
		originalPrice = discountedPrice * disRate
	} else if 100 < discountedPrice && discountedPrice <= 300 {
		disRate := genRandRate(45, 60)
		originalPrice = discountedPrice * disRate
	} else if 300 < discountedPrice && discountedPrice <= 400 {
		disRate := genRandRate(40, 55)
		originalPrice = discountedPrice * disRate
	} else if 400 < discountedPrice {
		disRate := genRandRate(30, 45)
		originalPrice = discountedPrice * disRate
	}
	return originalPrice
}

func genRandRate(min, max int) float32 {
	rand.Seed(time.Now().UnixNano())
	rng := max - min + 1
	randFloat := (float32(rand.Intn(rng)) + float32(min) + 100.00) / 100.00
	return randFloat
}
