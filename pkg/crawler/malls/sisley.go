package malls

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
)

type SisleyResponseParser struct {
	Page      string                 `json:"page"`
	TotalPage int                    `json:"totalPage"`
	Products  []SisleyProductWrapper `json:"result"`
}

type SisleyProductWrapper struct {
	ProductIdx      string `json:"pridx"`
	ProductCode     string `json:"productcode"`
	ProductName     string `json:"productname"`
	OriginalPrice   string `json:"consumerprice"`
	DiscountedPrice string `json:"sellprice"`
	ProdCode        string `json:"prodcode"`
}

func CrawlSisley(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	pageNum := 0
	totalProducts := 0
	crawlurl := source.CrawlUrl

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	for {
		dataStr := `{"shopcode":"S","code":"` + source.Category.CatIdentifier + `","page":"` + strconv.Itoa(pageNum) + `","limit":"80","isOutlet":"Y"}`
		errorMessage := "Crawl Failed: Source " + source.Category.KeyName
		resp, err := utils.RequestRetryer(crawlurl, utils.REQUEST_POST, utils.GetGeneralHeader(), dataStr, errorMessage)
		if err != nil {
			break
		}

		defer resp.Body.Close()

		crawlResponse := &SisleyResponseParser{}
		json.NewDecoder(resp.Body).Decode(crawlResponse)
		crawlProducts := crawlResponse.Products

		for _, pd := range crawlProducts {
			originalPriceInt := utils.ParsePriceString(pd.OriginalPrice)
			discountedPriceInt := utils.ParsePriceString(pd.DiscountedPrice)

			if discountedPriceInt == 0 {
				discountedPriceInt = originalPriceInt
			} else if originalPriceInt == 0 {
				originalPriceInt = discountedPriceInt
			}

			productUrl := "http://www.sisleymall.com/product/view?code=" + source.Category.CatIdentifier + "&productcode=" + pd.ProductCode
			images, colors, sizes, inventories, description := CrawlSisleyDetail(productUrl)

			addRequest := &productinfo.AddMetaInfoRequest{
				AlloffName:          pd.ProductName,
				ProductID:           pd.ProductCode,
				ProductUrl:          productUrl,
				ProductType:         []domain.AlloffProductType{domain.Female},
				OriginalPrice:       float32(originalPriceInt),
				DiscountedPrice:     float32(discountedPriceInt),
				CurrencyType:        domain.CurrencyKRW,
				Brand:               brand,
				Source:              source,
				AlloffCategory:      &domain.AlloffCategoryDAO{},
				Images:              images,
				Colors:              colors,
				Sizes:               sizes,
				Inventory:           inventories,
				Information:         description,
				IsForeignDelivery:   false,
				IsTranslateRequired: false,
				ModuleName:          source.CrawlModuleName,
				IsRemoved:           false,
				IsSoldout:           false,
			}

			totalProducts += 1
			productinfo.ProcessCrawlingInfoRequests(addRequest)
		}

		pageInt, _ := strconv.Atoi(crawlResponse.Page)
		if crawlResponse.TotalPage > pageInt {
			pageNum += 1
		} else {
			break
		}
	}

	crawler.PrintCrawlResults(source, totalProducts)
	<-worker
	done <- true
}

func CrawlSisleyDetail(productUrl string) (images, colors, sizes []string, inventories []*domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.sisleymall.com", "www.sisleymall.com:443"),
	)

	description = map[string]string{}
	colors = nil

	c.OnHTML(".prd-pics-middle", func(e *colly.HTMLElement) {
		e.ForEach(".prd-img", func(_ int, el *colly.HTMLElement) {
			imageurl := el.ChildAttr("img", "data-src")
			if imageurl != "" {
				images = append(images, imageurl)
			}
		})
	})

	c.OnHTML(".opt-inner .opt-list", func(e *colly.HTMLElement) {
		e.ForEach("li .chk-size", func(_ int, el *colly.HTMLElement) {
			sizeInClass := el.Attr("class")
			sizes = append(sizes, el.ChildText("label span"))
			if !strings.Contains(sizeInClass, "disabled") {
				inventories = append(inventories, &domain.InventoryDAO{
					Size:     el.ChildText("label span"),
					Quantity: 1,
				})
			}
		})
	})

	c.OnHTML(".layer-prdInfo-list", func(e *colly.HTMLElement) {
		e.ForEach("table tbody tr", func(_ int, el *colly.HTMLElement) {
			if !strings.Contains(el.ChildText("th"), "상품코드") {
				description[el.ChildText("th")] = el.ChildText("td")
			}
		})
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Retry on sisley: " + productUrl)
		time.Sleep(5 * time.Second)
		c.Visit(productUrl)
	})

	c.Visit(productUrl)
	return
}
