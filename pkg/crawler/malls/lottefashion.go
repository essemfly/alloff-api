package malls

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
)

type LotteFashionResponseParse struct {
	Results LotteFashionResponseResult `json:"results"`
}

type LotteFashionResponseResult struct {
	Products  []LotteFashionProduct `json:"products"`
	Page      int                   `json:"page"`
	TotalPage int                   `json:"totalPage"`
	Total     int                   `json:"total"`
}

type LotteFashionProduct struct {
	Name           string                     `json:"name"`
	Images         []LotteFashionImage        `json:"images"`
	Soldout        bool                       `json:"soldout"`
	OriginalPrice  int                        `json:"originalPrice"`
	DiscountRate   int                        `json:"discountRate"`
	SalePrice      int                        `json:"salePrice"`
	Season         string                     `json:"season"`
	BrandEngName   string                     `json:"brandEngName"`
	BrandName      string                     `json:"brandName"`
	ReviewCount    int                        `json:"reviewCount"`
	ReviewSdomain  float32                    `json:"reviewSdomain"`
	Property       LotteFashionProperty       `json:"properties"`
	LotteFashionID string                     `json:"id"`
	CreateYmdt     int                        `json:"createYmdt"`
	Sizes          []LotteFashionProductStock `json:"sizes"`
}

type LotteFashionImage struct {
	Url string `json:"url"`
}

type LotteFashionProductStock struct {
	ItemId string `json:"itemId"`
	Size   string `json:"value"`
	Stock  int    `json:"stock"`
}

type LotteFashionProperty struct {
	StyleYear    string `json:"STYLE_YEAR"`
	Transparency string `json:"TRANSPARENCY"`
	Flexibility  string `json:"FLEXIBILITY"`
	Lining       string `json:"LINING"`
	Touch        string `json:"TOUCH"`
	Weight       string `json:"WEIGHT"`
	Thick        string `json:"THICK"`
	Season       string `json:"SEASON"`
}

func CrawlLotteFashion(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	pageSize := 100
	pageNum := 1
	totalProducts := 0

	jsonStr := BuildLFJsonBody(pageNum, pageSize, source)

	errorMessage := "Crawl Failed: Source " + source.Category.KeyName
	resp, err := utils.RequestRetryer(source.CrawlUrl, utils.REQUEST_POST, utils.GetGeneralHeader(), jsonStr, errorMessage)
	if err != nil {
		<-worker
		done <- true
		return
	}
	defer resp.Body.Close()

	brand, _ := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)

	crawlResponse := &LotteFashionResponseParse{}
	json.NewDecoder(resp.Body).Decode(crawlResponse)
	crawlResults := crawlResponse.Results

	for crawlResults.TotalPage > pageNum {
		pageNum += 1
		jsonStr := BuildLFJsonBody(pageNum, pageSize, source)

		errorMessage = "Crawl Failed: Source Page" + source.Category.KeyName + strconv.Itoa(pageNum)
		resp, err := utils.RequestRetryer(source.CrawlUrl, utils.REQUEST_POST, utils.GetGeneralHeader(), jsonStr, errorMessage)
		if err != nil {
			break
		}

		results := &LotteFashionResponseParse{}
		json.NewDecoder(resp.Body).Decode(results)
		defer resp.Body.Close()

		totalProducts += MapLotteCrawlResultsToModels(results.Results.Products, source, brand)
	}

	crawler.WriteCrawlResults(source, totalProducts)
	<-worker
	done <- true
}

func BuildLFJsonBody(pageNum int, pageSize int, source *domain.CrawlSourceDAO) string {
	return `{"page":` + strconv.Itoa(pageNum) + `, "size":` + strconv.Itoa(pageSize) + `,"tid":[` + source.MainCategoryKey + `], "sendLogYN":"N","brandGroup":["` + source.BrandIdentifier + `"], "pid":[` + source.Category.CatIdentifier + `]}`
}

func MapLotteCrawlResultsToModels(products []LotteFashionProduct, source *domain.CrawlSourceDAO, brand *domain.BrandDAO) int {
	c := colly.NewCollector(
		colly.AllowedDomains("www.lfmall.co.kr"),
	)

	numProducts := 0

	for _, product := range products {
		description := map[string]string{}
		images := []string{}
		for _, img := range product.Images {
			imageUrl := strings.Replace(img.Url, "/320/", "/640/", 1)
			images = append(images, imageUrl)
		}

		colors := []string{}
		newSizes := []string{}
		inventories := []domain.InventoryDAO{}
		for _, sizeOption := range product.Sizes {
			if !utils.ItemExists(newSizes, sizeOption.Size) {
				newSizes = append(newSizes, sizeOption.Size)

				if sizeOption.Stock > 0 {
					inventories = append(inventories, domain.InventoryDAO{
						Size:     sizeOption.Size,
						Quantity: sizeOption.Stock,
					})
				}
			}
		}

		url := "https://www.lfmall.co.kr/productNew.do?cmd=getProductDetail&PROD_CD=" + product.LotteFashionID

		c.OnHTML(".tbl_accordion .table_wrap", func(e *colly.HTMLElement) {
			e.ForEach("table tbody tr td table tbody tr", func(_ int, el *colly.HTMLElement) {
				if !strings.Contains(el.ChildText("th"), "A/S") {
					description[el.ChildText("th")] = el.ChildText("td")
				}
			})
		})

		c.Visit(url)

		addRequest := crawler.ProductsAddRequest{
			Brand:         brand,
			Source:        source,
			ProductID:     product.LotteFashionID,
			ProductName:   product.Name,
			ProductUrl:    url,
			Images:        images,
			Sizes:         newSizes,
			Inventories:   inventories,
			Colors:        colors,
			Description:   description,
			OriginalPrice: float32(product.OriginalPrice),
			SalesPrice:    float32(product.SalePrice),
			CurrencyType:  domain.CurrencyKRW,
		}

		numProducts += 1
		crawler.AddProduct(addRequest)
	}

	return numProducts
}
