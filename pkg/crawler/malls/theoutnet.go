package malls

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	"github.com/lessbutter/alloff-api/pkg/product"
)

type TheOutnetResponseParser struct {
	PageNumber    int                        `json:"pageNumber"`
	PageSize      int                        `json:"pageSize"`
	Products      []TheOutnetResponseProduct `json:"products"`
	TotalProducts int                        `json:"recordSetTotal"`
}

type TheOutnetResponseProduct struct {
	Name      string `json:"nameEN"`
	ProductID string `json:"productId"`
	Colors    []struct {
		Label string `json:"labelEN"`
	} `json:"productColours"`
	Seo struct {
		SeoUrl string `json:"seoURLKeyword"`
	}
	Images []TheOutnetResponseImage `json:"images"`
	Price  struct {
		RdDiscount struct {
			Amount  int
			Divisor int
		} `json:"rdDiscount"`
		SellingPrice struct {
			Amount  int
			Divisor int
		} `json:"sellingPrice"`
		WasPrice struct {
			Amount  int
			Divisor int
		} `json:"wasPrice"`
		RdSellingPrice struct {
			Amount  int
			Divisor int
		} `json:"rdSellingPrice"`
		Discount struct {
			Amount  int
			Divisor int
		} `json:"discount"`
		Currency struct {
			Symbol string
			Label  string
		} `json:"currency"`
		RdWasPrice struct {
			Amount  int
			Divisor int
		} `json:"rdWasPrice"`
	} `json:"price"`
	Attributes []struct {
		Identifier string
		Values     []struct {
			Identifier string
			Label      string
		}
	} `json:"attributes"`
}

type TheOutnetResponseImage struct {
	View string
	Size struct {
		Width  int
		Height int
	}
	Url string
}

func CrawlTheoutnet(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	pageNum, pageSize := 1, 96

	brand, err := ioc.Repo.Brands.GetByKeyname(source.Category.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	url, err := url.Parse(source.CrawlUrl)
	if err != nil {
		log.Fatal(err)
	}

	values := url.Query()
	values.Add("attrs", "true")
	values.Add("locale", "en_US")
	values.Add("pageSize", strconv.Itoa(pageSize))
	values.Add("pageNumber", strconv.Itoa(pageNum))
	values.Add("category", "/designers/"+source.BrandIdentifier+source.Category.CatIdentifier)
	url.RawQuery = values.Encode()

	errorMessage := "Crawl Failed: Source " + source.Category.KeyName
	resp, err := utils.RequestRetryer(url.String(), utils.REQUEST_GET, utils.GetTheoutnetHeaders(), "", errorMessage)
	if err != nil {
		log.Println("Theoutnet fail on", source, err)
		<-worker
		done <- true
		return
	}

	defer resp.Body.Close()

	outnetResponse := &TheOutnetResponseParser{}

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		log.Println("gzip err", err)
	}

	err = json.NewDecoder(reader).Decode(outnetResponse)
	if err != nil {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Println("body", string(b))
		log.Println("the outnet decode err", err)
	}

	totalProductsNum := MapTheoutnetListProducts(outnetResponse.Products, source, brand)

	for pageNum*pageSize < outnetResponse.TotalProducts {
		pageNum += 1
		values.Set("pageNumber", strconv.Itoa(pageNum))
		url.RawQuery = values.Encode()

		errorMessage := "Crawl Failed: Source " + source.Category.KeyName
		resp, err := utils.RequestRetryer(url.String(), utils.REQUEST_GET, utils.GetTheoutnetHeaders(), "", errorMessage)
		if err != nil {
			log.Println("Theoutnet fail on ", source, strconv.Itoa(pageNum), err)
		}
		outnetResponse = &TheOutnetResponseParser{}
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			log.Println("gzip err", err)
		}
		err = json.NewDecoder(reader).Decode(outnetResponse)
		if err != nil {
			log.Println("the outnet decode err", err)
		}
		MapTheoutnetListProducts(outnetResponse.Products, source, brand)
	}

	crawler.PrintCrawlResults(source, totalProductsNum)

	<-worker
	done <- true
}

func MapTheoutnetListProducts(pds []TheOutnetResponseProduct, source *domain.CrawlSourceDAO, brand *domain.BrandDAO) int {
	numProducts := 0
	urlPrefix := "https://www.theoutnet.com/en-de/shop/product"

	for _, pd := range pds {
		sizes, inventories, description := CrawlTheoutnetDetail(pd.Seo.SeoUrl)

		images := []string{}
		for _, imgResponse := range pd.Images {
			if imgResponse.Size.Width == 580 {
				if strings.HasPrefix(imgResponse.Url, "//") {
					images = append(images, "https:"+imgResponse.Url)
				} else {
					images = append(images, imgResponse.Url)
				}
			}
		}

		colors := []string{}
		for _, colorResp := range pd.Colors {
			colors = append(colors, colorResp.Label)
		}

		addRequest := &product.ProductCrawlingAddRequest{
			Brand:         brand,
			Source:        source,
			ProductID:     pd.ProductID,
			ProductName:   pd.Name,
			ProductUrl:    urlPrefix + pd.Seo.SeoUrl,
			Images:        images,
			Sizes:         sizes,
			Inventories:   inventories,
			Colors:        colors,
			Description:   description,
			OriginalPrice: float32(pd.Price.WasPrice.Amount) / float32(pd.Price.WasPrice.Divisor),
			SalesPrice:    float32(pd.Price.SellingPrice.Amount) / float32(pd.Price.SellingPrice.Divisor),
			CurrencyType:  domain.CurrencyEUR,
		}

		numProducts += 1
		product.AddProductInCrawling(addRequest)
	}
	return numProducts
}

func CrawlTheoutnetDetail(productUrl string) (sizes []string, inventories []domain.InventoryDAO, description map[string]string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.theoutnet.com", "www.theoutnet.com:443"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	urlPrefix := "https://www.theoutnet.com/en-de/shop/product"
	description = map[string]string{}

	c.OnHTML(".multipleSizes", func(e *colly.HTMLElement) {
		e.ForEach("ul li", func(_ int, el *colly.HTMLElement) {
			size := el.ChildText(".GridSelect11__optionBox")
			unavailable := el.ChildAttr(".GridSelect11__optionBox", "aria-label")
			sizes = append(sizes, size)
			if !strings.Contains(unavailable, "sold out") {
				inventories = append(inventories, domain.InventoryDAO{
					Size:     size,
					Quantity: 10,
				})
			}
		})
	})

	c.OnHTML(".AccordionSection3", func(e *colly.HTMLElement) {
		title := e.ChildText(".EditorialAccordion84__accordionTitle")
		descs := e.ChildText(".EditorialAccordion84__accordionContent")
		description[title] = descs
	})

	err := c.Visit(urlPrefix + productUrl)
	if err != nil {
		log.Println("ERR on detail crawling", err)
	}

	return
}