package malls

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

type KolonResponseParser struct {
	Images           []KolonImage         `json:"images"`
	Soldout          string               `json:"soldOutYn"`
	Options          []KolonProductOption `json:"variantOptions"`
	Fabric           string               `json:"fabric"`
	ProduceYearMonth string               `json:"produceYearMonth"`
	ManufacturerName string               `json:"manufacturerName"`
	OriginCountry    string               `json:"originCountry"`
}

type KolonImage struct {
	ImageLS struct {
		Index int    `json:"index"`
		Url   string `json:"url"`
	}
	ImageLM struct {
		Index int    `json:"index"`
		Url   string `json:"url"`
	}
	ImageLZ struct {
		Index int    `json:"index"`
		Url   string `json:"url"`
	}
}

type KolonProductOption struct {
	Size  string `json:"legacySize"`
	Stock struct {
		Level string `json:"stockLevelStatus"`
	} `json:"stock"`
}

func CrawlKolon(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.kolonmall.com"),
	)

	currentPage := 1

	totalProducts := 0
	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	c.OnHTML(".itmtgi-0", func(e *colly.HTMLElement) {
		productID := e.ChildAttr(".itmtgi-2 .itmtgi-4", "href")
		productUrl := "https://www.kolonmall.com/" + productID
		productName := e.ChildText(".itmtgi-2 .itmtgi-4 .tdtd57-2")
		discountedPriceAll := e.ChildText(".itmtgi-2 .itmtgi-4 .tdtd57-3")
		discountedPriceSplit := strings.Split(discountedPriceAll, "원")
		discountedPriceStr := discountedPriceSplit[0]
		originalPriceStr := e.ChildText(".itmtgi-2 .itmtgi-4 .tdtd57-3 .real-price")
		originalPrice := utils.ParsePriceString(originalPriceStr)
		discountedPrice := utils.ParsePriceString(discountedPriceStr)

		if originalPrice == 0 {
			originalPrice = discountedPrice
		}

		images := []string{}
		sizes := []string{}
		colors := []string{}
		inventories := []*domain.InventoryDAO{}
		description := map[string]string{}

		errorMessage := "Crawl Failed: Product Detail" + source.Category.KeyName + " - " + productUrl
		resp, err := utils.RequestRetryer(productUrl, utils.REQUEST_GET, utils.GetGeneralHeader(), "", errorMessage)
		if err != nil {
			return
		}
		defer resp.Body.Close()

		doc, _ := goquery.NewDocumentFromReader(resp.Body)
		doc.Find("script").Each(func(i int, s *goquery.Selection) {
			tmp := s.Text()
			divider := "var productOrig = "
			if strings.Contains(tmp, "var productOrig = ") {
				productOrigPlainList := strings.Split(tmp, divider)
				productOrig := strings.Split(productOrigPlainList[1], "};")[0] + "}"
				result := &KolonResponseParser{}
				json.Unmarshal([]byte(productOrig), &result)
				if result != nil {
					description["소재"] = result.Fabric
					description["제조연월"] = result.ProduceYearMonth
					description["제조국"] = result.OriginCountry
					description["제조자"] = result.ManufacturerName
					for _, img := range result.Images {
						images = append(images, img.ImageLM.Url)
					}

					for _, size := range result.Options {
						sizes = append(sizes, size.Size)
						if size.Stock.Level != "outOfStock" {
							inventories = append(inventories, &domain.InventoryDAO{
								Size:     size.Size,
								Quantity: 10,
							})
						}
					}
				} else {
					log.Println("ERROR CASE", source.CrawlUrl)
				}
			}
		})

		addRequest := &product.AddMetaInfoRequest{
			AlloffName:      productName,
			ProductID:       productID,
			ProductUrl:      productUrl,
			ProductType:     []domain.AlloffProductType{domain.Female},
			OriginalPrice:   float32(originalPrice),
			DiscountedPrice: float32(discountedPrice),
			CurrencyType:    domain.CurrencyKRW,
			Brand:           brand,
			Source:          source,
			// AlloffCategory:  nil,
			Images:              images,
			Colors:              colors,
			Sizes:               sizes,
			Inventory:           inventories,
			Information:         description,
			IsForeignDelivery:   false,
			IsTranslateRequired: false,
			ModuleName:          source.CrawlModuleName,
		}

		totalProducts += 1
		product.ProcessAddProductInfoRequests(addRequest)

	})

	c.OnHTML(".product-page-arrow:last-child", func(e *colly.HTMLElement) {
		currentPage += 1
		c.Visit(source.CrawlUrl + "&page=" + strconv.Itoa(currentPage))
	})

	c.Visit(source.CrawlUrl)

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}
