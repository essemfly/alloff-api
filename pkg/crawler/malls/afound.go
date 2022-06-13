package malls

import (
	"encoding/json"
	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

type ProductInfoJson struct {
	Name         string   `json:"Name"`
	URL          string   `json:"Url"`
	DisplayPrice float64  `json:"DisplayPrice"`
	WasPrice     float64  `json:"WasPrice"`
	Sizes        []string `json:"Sizes"`
}

func CrawlAfound(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.afound.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.SetRequestTimeout(10 * time.Second)

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		config.Logger.Error("error occurred on get brand by key name : ", zap.Error(err))
	}

	totalProducts := 0
	currentPageNum := 1
	c.OnHTML("#afpl-products-list-container > div", func(e *colly.HTMLElement) {
		var productInfoJson ProductInfoJson
		productInfoStr := e.ChildAttr("a", "data-gtm-product")

		err := json.Unmarshal([]byte(productInfoStr), &productInfoJson)
		if err != nil {
			config.Logger.Error("error occurred on unmarshal json to struct : ", zap.Error(err))
		}

		productName := productInfoJson.Name
		productUrl := productInfoJson.URL
		discountedPrice := productInfoJson.DisplayPrice
		originalPrice := productInfoJson.WasPrice
		sizes := productInfoJson.Sizes
		productId := strings.Replace(productUrl, "https://www.afound.com/de-de/produkte/", "", -1)

		inventories := []*domain.InventoryDAO{}
		for _, size := range sizes {
			inventories = append(inventories, &domain.InventoryDAO{
				Size:     size,
				Quantity: 1,
			})
		}

		images, colors, infos, description, notSale := getAfoundDetail(productUrl)

		// 판매하지 않지만 리스트에 남아있는 상품은 패스한다.
		if notSale {
			return
		}

		addRequest := &productinfo.AddMetaInfoRequest{
			AlloffName:          productName,
			ProductID:           productId,
			ProductUrl:          productUrl,
			ProductType:         []domain.AlloffProductType{domain.Female},
			OriginalPrice:       float32(originalPrice),
			DiscountedPrice:     float32(discountedPrice),
			CurrencyType:        domain.CurrencyEUR,
			Brand:               brand,
			Source:              source,
			AlloffCategory:      &domain.AlloffCategoryDAO{},
			Images:              images,
			Colors:              colors,
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
	})

	c.OnHTML("div.seo-pagination.desktop.tiled-list-wrapper.space-below--huge > div > a.tiled-list__item.button-select.last-item", func(e *colly.HTMLElement) {
		lastPageStr := e.Text
		lastPageNum, _ := strconv.Atoi(lastPageStr)
		if lastPageNum > 1 {
			if currentPageNum < lastPageNum {
				currentPageNum += 1
				url := source.CrawlUrl + "&page=" + strconv.Itoa(currentPageNum)
				c.Visit(url)
			}
		}
	})

	err = c.Visit(source.CrawlUrl)
	if err != nil {
		config.Logger.Error("error occurred in crawl afound ", zap.Error(err))
	}

	<-worker
	done <- true
}

func getAfoundDetail(productUrl string) (images, colors []string, infos, description map[string]string, notSale bool) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.afound.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.SetRequestTimeout(10 * time.Second)

	//isDigit := regexp.MustCompile(`^\d*\.?\d+$`)

	infos = map[string]string{
		"소재": "",
		"색상": "",
	}
	description = map[string]string{
		"사이즈 및 핏": "",
		"제품설명":    "",
	}
	images = []string{}
	colors = []string{}
	sizeAndFit := ""

	// 상품 url에 접속했지만 상품 정보가 없는경우 리스트로 리다이렉트된다. 이떄 리스트로 리다이렉트 됐는지를 판단하여 상품 유효성을 검증한다.
	c.OnHTML(".af-product-listing__category-filters__section__header--wrapper", func(e *colly.HTMLElement) {
		if e.Text != "" {
			notSale = true
			return
		}
	})

	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
		"section:nth-child(1) > div > p", func(e *colly.HTMLElement) {
		desc := e.Text
		desc = strings.TrimSpace(desc)
		description["제품설명"] = desc
	})

	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
		"section:nth-child(1) > div > div > dl:nth-child(6) > dd", func(e *colly.HTMLElement) {
		sizeAndFit += e.Text + ", "
	})
	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
		"section:nth-child(1) > div > div > dl:nth-child(5) > dd", func(e *colly.HTMLElement) {
		sizeAndFit += e.Text
	})

	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
		"section:nth-child(1) > div > div > dl:nth-child(1) > dd", func(e *colly.HTMLElement) {
		infos["소재"] = e.Text
	})

	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
		"section:nth-child(1) > div > div > dl:nth-child(2) > dd", func(e *colly.HTMLElement) {
		infos["색상"] = e.Text
		colors = append(colors, e.Text)
	})

	c.OnHTML("div.product-page__grid > section.product-page__grid-col-left", func(e *colly.HTMLElement) {
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			imgUrl := el.ChildAttr("img", "src")
			imgUrl = strings.Split(imgUrl, "?")[0]
			imgUrl += "?preset=product-details-desktop"
			images = append(images, imgUrl)
		})
	})

	c.Visit(productUrl)
	description["사이즈 및 핏"] = sizeAndFit

	return
}
