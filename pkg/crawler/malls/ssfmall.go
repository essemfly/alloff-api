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
	pdPackage "github.com/lessbutter/alloff-api/pkg/product"
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
	for _, product := range products {
		// isSoldout := false
		// if product.God.Soldout == "SLDOUT" {
		// 	isSoldout = true
		// }

		productUrl := "https://www.ssfshop.com/" + source.BrandKeyname + "/" + product.ProductId + "/good"
		images, colors, sizes, inventories, description := getSSFDetailInfo(productUrl)

		addRequest := pdPackage.ProductsAddRequest{
			Brand:         brand,
			Source:        source,
			ProductID:     product.ProductId,
			ProductName:   product.God.Name,
			ProductUrl:    productUrl,
			Images:        images,
			Sizes:         sizes,
			Inventories:   inventories,
			Colors:        colors,
			Description:   description,
			OriginalPrice: float32(product.DspGodPrc.OriginalPrice),
			SalesPrice:    float32(product.DspGodPrc.DiscountedPrice),
			CurrencyType:  domain.CurrencyKRW,
		}

		numProducts += 1
		crawler.AddProduct(addRequest)

		// newProduct := domain.ProductDAO{
		// 	ProductId:       product.ProductId,
		// 	Category:        &source.Category,
		// 	Brand:           brand,
		// 	Name:            product.God.Name,
		// 	OriginalPrice:   product.DspGodPrc.OriginalPrice,
		// 	DiscountedPrice: product.DspGodPrc.DiscountedPrice,
		// 	DiscountRate:    utils.CalculateDiscountRate(float32(product.DspGodPrc.OriginalPrice), float32(product.DspGodPrc.DiscountedPrice)),
		// 	Soldout:         isSoldout,
		// 	Images:          images,
		// 	Created:         time.Now(),
		// 	Updated:         time.Now(),
		// 	IsNewlyCrawled:  true,
		// 	ProductUrl:      productUrl,
		// 	Removed:         false,
		// 	Description:     description,
		// 	SizeAvailable:   sizes,
		// 	IsImageCached:   false,
		// }

	}
	return numProducts
}

func getSSFDetailInfo(productUrl string) (imageUrls, colors, sizes []string, inventories []domain.InventoryDAO, description map[string]string) {
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

	c.OnHTML("#prdInfoOptionArea #optSelectGnrlDiv .option ul", func(e *colly.HTMLElement) {
		e.ForEach("li a em", func(_ int, el *colly.HTMLElement) {
			size := el.Text
			size = strings.Replace(size, "\n", "", -1)
			size = strings.Replace(size, "\t", "", -1)
			sizes = append(sizes, size)
			if !strings.Contains(size, "품절") {
				inventories = append(inventories, domain.InventoryDAO{
					Size:     size,
					Quantity: 10,
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
