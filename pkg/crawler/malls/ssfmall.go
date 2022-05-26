package malls

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
)

type SSFResponseParser struct {
	Products    []SSFProductWrapper `json:"godList"`
	TotalRow    int                 `json:"totalRow"`
	CurrentPage int                 `json:"currentPage"`
	TotalPage   int                 `json:"totalPage"`
}

type SSFProductWrapper struct {
	God           SSFProduct   `json:"god"`
	DspGodPrc     SSFDspGodPrc `json:"dspGodPrc"`
	FrontImageUrl string       `json:"imgFrontUrl"`
	BackImageUrl  string       `json:"imgBackUrl"`
	SeasonGrpname string       `json:"seasonGrpNm"`
	ProdYear      string       `json:"prodYear"`
	ProductId     string       `json:"godNo"`
}

type SSFProduct struct {
	ProductId string `json:"godNo"`
	Name      string `json:"godNm"`
	Soldout   string `json:"godSaleSectCd"`
}

type SSFDspGodPrc struct {
	OriginalPrice   int `json:"rtlPrc"`
	DiscountedPrice int `json:"lastSalePrc"`
	DiscountRate    int `json:"godDcRt"`
}

func CrawlSSFMall(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	pageNum := 1
	url := source.CrawlUrl + "&currentPage=" + strconv.Itoa(pageNum)

	errorMessage := "Crawl Failed: Source " + source.Category.KeyName
	resp, err := utils.RequestRetryer(url, utils.REQUEST_GET, utils.GetSSFHeaders(), "", errorMessage)
	if err != nil {
		log.Println("SSFmall fail on", source)
		<-worker
		done <- true
		return
	}

	defer resp.Body.Close()

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println("brand repo key error", err)
	}

	crawlResponse := &SSFResponseParser{}
	err = json.NewDecoder(resp.Body).Decode(crawlResponse)
	if err != nil {
		log.Println("ssf decode err", err)
	}

	totalProducts := MapSSFCrawlResultsToModels(crawlResponse.Products, source, brand)

	for crawlResponse.TotalPage > crawlResponse.CurrentPage {
		pageNum += 1
		url := source.CrawlUrl + "&currentPage=" + strconv.Itoa(pageNum)
		errorMessage = "Crawl Failed: Source Page" + source.Category.KeyName + strconv.Itoa(pageNum)
		resp, err := utils.RequestRetryer(url, utils.REQUEST_GET, utils.GetSSFHeaders(), "", errorMessage)
		if err != nil {
			log.Println("SSF Request error", err)
			break
		}

		crawlResponse = &SSFResponseParser{}
		err = json.NewDecoder(resp.Body).Decode(crawlResponse)
		if err != nil {
			log.Println("SSF REsponse decode error", err)

		}
		totalProducts += MapSSFCrawlResultsToModels(crawlResponse.Products, source, brand)
	}

	crawler.PrintCrawlResults(source, totalProducts)

	<-worker
	done <- true
}

func MapSSFCrawlResultsToModels(products []SSFProductWrapper, source *domain.CrawlSourceDAO, brand *domain.BrandDAO) int {
	numProducts := 0
	for _, pd := range products {
		productUrl := "https://www.ssfshop.com/" + source.BrandKeyname + "/" + pd.ProductId + "/good"
		images, colors, sizes, inventories, description := getSSFDetailInfo(productUrl)

		addRequest := &productinfo.AddMetaInfoRequest{
			AlloffName:          pd.God.Name,
			ProductID:           pd.ProductId,
			ProductUrl:          productUrl,
			ProductType:         []domain.AlloffProductType{domain.Female},
			OriginalPrice:       float32(pd.DspGodPrc.OriginalPrice),
			DiscountedPrice:     float32(pd.DspGodPrc.DiscountedPrice),
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

		numProducts += 1
		productinfo.ProcessCrawlingInfoRequests(addRequest)
	}
	return numProducts
}

func getSSFDetailInfo(productUrl string) (imageUrls, colors, sizes []string, inventories []*domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.ssfshop.com"),
	)

	description = map[string]string{}
	colors = nil

	c.OnHTML(".detail .summary .gallery ul", func(e *colly.HTMLElement) {
		e.ForEach("li a", func(_ int, el *colly.HTMLElement) {
			imageurl := el.ChildAttr("img", "src")
			if imageurl != "" {
				imageUrls = append(imageUrls, imageurl)
			}
		})
	})

	c.OnHTML(".size_wrap ul", func(e *colly.HTMLElement) {
		e.ForEach("input", func(_ int, el *colly.HTMLElement) {
			if el.Attr("name") == "sizeItmNo" {
				size := el.Attr("sizeitmnm")
				sizes = append(sizes, size)
				stockInStr := el.Attr("onlineusefulinvqty")
				stock, _ := strconv.Atoi(stockInStr)
				inventories = append(inventories, &domain.InventoryDAO{
					Size:     size,
					Quantity: stock,
				})
			}
		})
	})

	c.OnHTML(".essential .data_essential table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			if !strings.Contains(el.ChildText("th"), "AS") {
				description[el.ChildText("th")] = el.ChildText("td")
			}
		})
	})

	c.Visit(productUrl)
	return
}
