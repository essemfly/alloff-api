package malls

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/crawler"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
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
	Price struct {
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
	urlPrefix := "https://www.theoutnet.com/en-us/shop/product"

	for _, pd := range pds {
		composition, sizes, inventories, description, images := CrawlTheoutnetDetail(pd.Seo.SeoUrl)
		colors := []string{}
		for _, colorResp := range pd.Colors {
			colors = append(colors, colorResp.Label)
		}

		if len(sizes) == 0 && len(colors) > 0 {
			for _, color := range colors {
				inventories = append(inventories, &domain.InventoryDAO{
					Quantity: 1,
					Size:     color,
				})
			}
		}

		if len(images) == 8 {
			newImages := []string{
				images[3],
				images[0],
				images[2],
				images[1],
			}
			images = newImages
		}

		infos := map[string]string{
			"소재": composition,
			"색상": colors[0],
		}

		addRequest := &productinfo.AddMetaInfoRequest{
			AlloffName:           pd.Name,
			ProductID:            pd.ProductID,
			ProductUrl:           urlPrefix + pd.Seo.SeoUrl,
			ProductType:          []domain.AlloffProductType{domain.Female},
			OriginalPrice:        float32(pd.Price.WasPrice.Amount) / float32(pd.Price.WasPrice.Divisor),
			DiscountedPrice:      float32(pd.Price.SellingPrice.Amount) / float32(pd.Price.SellingPrice.Divisor),
			CurrencyType:         domain.CurrencyEUR,
			Brand:                brand,
			Source:               source,
			AlloffCategory:       &domain.AlloffCategoryDAO{},
			Images:               images,
			Colors:               colors,
			Sizes:                sizes,
			Inventory:            inventories,
			DescriptionInfos:     infos,
			Information:          description,
			IsTranslateRequired:  true,
			ModuleName:           source.CrawlModuleName,
			IsRemoved:            false,
			IsSoldout:            false,
			DescriptionImages:    images,
			IsForeignDelivery:    true,
			EarliestDeliveryDays: 14,
			LatestDeliveryDays:   21,
			IsRefundPossible:     true,
			RefundFee:            100000,
		}

		log.Println(urlPrefix + pd.Seo.SeoUrl)
		numProducts += 1
		productinfo.ProcessCrawlingInfoRequests(addRequest)

	}
	return numProducts
}

func CrawlTheoutnetDetail(productUrl string) (composition string, sizes []string, inventories []*domain.InventoryDAO, description map[string]string, images []string) {
	c := colly.NewCollector(
		colly.AllowedDomains("www.theoutnet.com", "www.theoutnet.com:443"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)

	isDigit := regexp.MustCompile(`^\d*\.?\d+$`)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("x-ibm-client-id", "19c36e19-5bc7-4de4-a4a9-65ffb9dcb727")
		r.Headers.Set("accept", "*/*")
		r.Headers.Set("accept-encoding", "gzip, deflate, br")
		r.Headers.Set("connection", "keep-alive")
		r.Headers.Set("user-agent", "PostmanRuntime/7.29.0")
		r.Headers.Set("content-type", "application/x-www-form-urlencoded")
	})

	urlPrefix := "https://www.theoutnet.com/en-us/shop/product"
	description = map[string]string{}

	c.OnHTML(".multipleSizes", func(e *colly.HTMLElement) {
		e.ForEach("ul li", func(_ int, el *colly.HTMLElement) {
			size := el.ChildText(".GridSelect11__optionBox")
			size = strings.Replace(size, " ", "", -1) // alloffSize 와 맞추기 위해 중간에 공백을 없애준다 (FR 32 -> FR32)
			if isDigit.MatchString(size) {
				size = "EU" + size
			}
			unavailable := el.ChildAttr(".GridSelect11__optionBox", "aria-label")
			sizes = append(sizes, size)
			if !strings.Contains(unavailable, "sold out") {
				inventories = append(inventories, &domain.InventoryDAO{
					Size:     size,
					Quantity: 1,
				})
			}
		})
	})

	if len(sizes) == 0 {
		c.OnHTML(" div.SizeSelect84__size.ProductDetails84__sizeSelect.ProductDetails84__sizeSelect--isOneSize > div > span.SizeSelect84__oneSizeLabel", func(e *colly.HTMLElement) {
			sizes = append(sizes, e.Text)
			inventories = append(inventories, &domain.InventoryDAO{
				Size:     sizes[0],
				Quantity: 1,
			})
		})
	}

	c.OnHTML("#SIZE_AND_FIT", func(e *colly.HTMLElement) {
		description["사이즈 및 핏"] = e.ChildText(".EditorialAccordion84__accordionContent--size_and_fit > p")
	})

	c.OnHTML("#TECHNICAL_DESCRIPTION", func(e *colly.HTMLElement) {
		desc := ""
		e.ForEach(".EditorialAccordion84__accordionContent--technical_description p", func(_ int, el *colly.HTMLElement) {
			desc += el.Text + "\n"
		})
		desc = strings.TrimRight(desc, "\n")
		description["제품설명"] = desc

		composition = e.ChildText(".EditorialAccordion84__composition")
	})

	c.OnHTML("ul.ImageCarousel84__track", func(e *colly.HTMLElement) {
		e.ForEach("li.ImageCarousel84__slide", func(_ int, el *colly.HTMLElement) {
			imageUrlBeforeParsing := el.ChildAttr(".ZoomedImage84", "style")

			newString := strings.TrimPrefix(imageUrlBeforeParsing, "background-image:url(//")
			newString = strings.TrimSuffix(newString, ")")
			alreadyRegistered := false
			for _, alreadyImage := range images {
				if alreadyImage == newString {
					alreadyRegistered = true
					break
				}
			}
			if !alreadyRegistered {
				images = append(images, "https://"+newString)
			}
		})
	})

	err := c.Visit(urlPrefix + productUrl)
	if err != nil {
		log.Println("ERR on detail crawling", err)
	}

	return
}
