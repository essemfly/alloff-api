package malls

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.uber.org/zap"
)

// type ProductInfoJson struct {
// 	Name         string   `json:"Name"`
// 	URL          string   `json:"Url"`
// 	DisplayPrice float64  `json:"DisplayPrice"`
// 	WasPrice     float64  `json:"WasPrice"`
// 	Sizes        []string `json:"Sizes"`
// }

type FlannelsListResponse struct {
	Products []struct {
		ImageSashURL          string      `json:"imageSashUrl"`
		URL                   string      `json:"url"`
		Image                 string      `json:"image"`
		ImageLarge            string      `json:"imageLarge"`
		ProductID             string      `json:"productId"`
		ColourID              string      `json:"colourId"`
		HidePrice             bool        `json:"hidePrice"`
		Brand                 string      `json:"brand"`
		Name                  string      `json:"name"`
		ShowFromPriceLabel    bool        `json:"showFromPriceLabel"`
		Price                 string      `json:"price"`
		TicketPrice           string      `json:"ticketPrice"`
		Sizes                 string      `json:"sizes"`
		DiscountPercentage    float64     `json:"discountPercentage"`
		ColourName            string      `json:"colourName"`
		PriceUnFormatted      float64     `json:"priceUnFormatted"`
		PriceLabel            interface{} `json:"priceLabel"`
		ProductSequenceNumber int         `json:"productSequenceNumber"`
		Rank                  int         `json:"rank"`
		ImageSash             string      `json:"imageSash"`
		ProductSizes          struct {
			UseAlternateSizes  bool        `json:"useAlternateSizes"`
			Sizes              string      `json:"sizes"`
			MinSize            interface{} `json:"minSize"`
			MaxSize            interface{} `json:"maxSize"`
			AlternateSizesText interface{} `json:"alternateSizesText"`
		} `json:"productSizes"`
		Category   interface{} `json:"category"`
		CategoryID interface{} `json:"categoryId"`
	} `json:"products"`
	Filters []struct {
		Key      string `json:"key"`
		Title    string `json:"title"`
		IsActive bool   `json:"isActive"`
		Filters  []struct {
			Group  string `json:"group"`
			Key    string `json:"key"`
			Label  string `json:"label"`
			Active bool   `json:"active"`
			Count  int    `json:"count"`
		} `json:"filters"`
		IsCollapsedByDefault bool   `json:"isCollapsedByDefault"`
		Type                 string `json:"type"`
		URLOrder             int    `json:"urlOrder"`
	} `json:"filters"`
	CurrnentPriceFilterLower interface{} `json:"currnentPriceFilterLower"`
	CurrentPriceFilterUpper  interface{} `json:"currentPriceFilterUpper"`
	PriceFilterMin           float64     `json:"priceFilterMin"`
	PriceFilterMax           float64     `json:"priceFilterMax"`
	AvailableSortOptions     []string    `json:"availableSortOptions"`
	CurrentSortOption        string      `json:"currentSortOption"`
	NumberOfPages            int         `json:"numberOfPages"`
	CurrentPage              int         `json:"currentPage"`
	NumberOfProducts         int         `json:"numberOfProducts"`
	RelatedCategories        []struct {
		ID           string      `json:"id"`
		Name         string      `json:"name"`
		URL          string      `json:"url"`
		ProductCount interface{} `json:"productCount"`
	} `json:"relatedCategories"`
	ChildCategories []interface{} `json:"childCategories"`
}

func CrawlFlannels(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	errorMessage := "Crawl Failed: Source " + source.BrandKeyname + " " + source.ID.Hex()
	resp, err := utils.RequestRetryer(source.CrawlUrl, utils.REQUEST_GET, utils.GetFlannelsHeader(), "", errorMessage)
	if err != nil {
		log.Println("flannels fail on", source)
		config.Logger.Error(err.Error())
		<-worker
		done <- true
		return
	}
	defer resp.Body.Close()

	listQueryResp := &FlannelsListResponse{}
	json.NewDecoder(resp.Body).Decode(listQueryResp)

	typeVariations := []string{}
	categoryVariations := []string{}
	for _, filter := range listQueryResp.Filters {
		if filter.Key == "AFLOR" {
			for _, typeFilter := range filter.Filters {
				if typeFilter.Count > 0 {
					typeVariations = append(typeVariations, typeFilter.Key)
				}
			}
		}
		if filter.Key == "1147_726918" {
			for _, catFilter := range filter.Filters {
				if catFilter.Count > 0 {
					categoryVariations = append(categoryVariations, catFilter.Key)
				}
			}
		}
	}

	baseUrl := strings.Replace(source.CrawlUrl, "productsPerPage=10", "productsPerPage=1000", 1)
	productRequests := CrawlFlannelsProducts(baseUrl, typeVariations, categoryVariations)

	log.Println("length of requests for brand", source.BrandKeyname, len(productRequests))
	<-worker
	done <- true
}

func CrawlFlannelsProducts(baseUrl string, productTypes []string, categories []string) []*productinfo.AddMetaInfoRequest {
	productRequests := []*productinfo.AddMetaInfoRequest{}
	for _, cat := range categories {
		for _, ptype := range productTypes {
			paramQuery := "AFLOR^" + ptype + "|1147_726918^" + cat
			catUrl := baseUrl + "&selectedFilters=" + url.QueryEscape(paramQuery)
			productRequests = append(productRequests, GetFlannelProductList(catUrl, ptype, cat)...)
		}
	}
	return productRequests
}

func GetFlannelProductList(categoryUrl, productType, categoryName string) []*productinfo.AddMetaInfoRequest {
	productTypes := map[string][]domain.AlloffProductType{
		"Mens":          {domain.Male},
		"Womens":        {domain.Female},
		"Boys":          {domain.Boy},
		"Girls":         {domain.Girl},
		"Unisex Kids":   {domain.Boy, domain.Girl},
		"Unisex Adults": {domain.Male, domain.Female},
	}

	errorMessage := "Crawl Failed: ProductURL: " + categoryUrl
	resp, err := utils.RequestRetryer(categoryUrl, utils.REQUEST_GET, utils.GetFlannelsHeader(), "", errorMessage)
	if err != nil {
		log.Println("flannels fail on: ", categoryUrl)
		config.Logger.Error("product list crawling error "+categoryUrl, zap.Error(err))
		return nil
	}
	defer resp.Body.Close()

	listQueryResp := &FlannelsListResponse{}
	json.NewDecoder(resp.Body).Decode(listQueryResp)

	log.Println("length of products: ", categoryName+" "+productType, len(listQueryResp.Products))

	baseURL := "https://www.flannels.com"
	requests := []*productinfo.AddMetaInfoRequest{}
	for _, pd := range listQueryResp.Products {
		newRequest := GetFlannelsDetail(baseURL + pd.URL)
		newRequest.ProductType = productTypes[productType]
		requests = append(requests, newRequest)
	}

	return requests
}

func GetFlannelsDetail(productURL string) *productinfo.AddMetaInfoRequest {
	c := colly.NewCollector(
		colly.AllowedDomains("www.flannels.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.SetRequestTimeout(10 * time.Second)

	productRequest := &productinfo.AddMetaInfoRequest{
		// AlloffName:          pd.ProductName,
		// ProductID:           pd.ProductCode,
		ProductUrl: productURL,
		// OriginalPrice:       float32(originalPriceInt),
		// DiscountedPrice:     float32(discountedPriceInt),
		// CurrencyType:        domain.CurrencyKRW,
		// Brand:               brand,
		// Source:              source,
		// AlloffCategory:      &domain.AlloffCategoryDAO{},
		// Images:              images,
		// Colors:              colors,
		// Sizes:               sizes,
		// Inventory:           inventories,
		// Information:         description,
		// IsForeignDelivery:   false,
		// IsTranslateRequired: false,
		// ModuleName:          source.CrawlModuleName,
		// IsRemoved:           false,
		// IsSoldout:           false,
	}

	c.Visit(productURL)

	return productRequest
}

// func getFlannelsDetail(productUrl string) (images, colors []string, infos, description map[string]string, notSale bool) {
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("www.afound.com"),
// 		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
// 	)
// 	c.SetRequestTimeout(10 * time.Second)

// 	//isDigit := regexp.MustCompile(`^\d*\.?\d+$`)

// 	infos = map[string]string{
// 		"소재": "",
// 		"색상": "",
// 	}
// 	description = map[string]string{
// 		"사이즈 및 핏": "",
// 		"제품설명":    "",
// 	}
// 	images = []string{}
// 	colors = []string{}
// 	sizeAndFit := ""

// 	// 상품 url에 접속했지만 상품 정보가 없는경우 리스트로 리다이렉트된다. 이떄 리스트로 리다이렉트 됐는지를 판단하여 상품 유효성을 검증한다.
// 	c.OnHTML(".af-product-listing__category-filters__section__header--wrapper", func(e *colly.HTMLElement) {
// 		if e.Text != "" {
// 			notSale = true
// 			return
// 		}
// 	})

// 	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
// 		"section:nth-child(1) > div > p", func(e *colly.HTMLElement) {
// 		desc := e.Text
// 		desc = strings.TrimSpace(desc)
// 		description["제품설명"] = desc
// 	})

// 	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
// 		"section:nth-child(1) > div > div > dl:nth-child(6) > dd", func(e *colly.HTMLElement) {
// 		sizeAndFit += e.Text + ", "
// 	})
// 	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
// 		"section:nth-child(1) > div > div > dl:nth-child(5) > dd", func(e *colly.HTMLElement) {
// 		sizeAndFit += e.Text
// 	})

// 	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
// 		"section:nth-child(1) > div > div > dl:nth-child(1) > dd", func(e *colly.HTMLElement) {
// 		infos["소재"] = e.Text
// 	})

// 	c.OnHTML("div.product-page__grid > section.product-page__grid-col-right > div.af-accordion.pdp-accordion.spaced--xmedium > "+
// 		"section:nth-child(1) > div > div > dl:nth-child(2) > dd", func(e *colly.HTMLElement) {
// 		infos["색상"] = e.Text
// 		colors = append(colors, e.Text)
// 	})

// 	c.OnHTML("div.product-page__grid > section.product-page__grid-col-left", func(e *colly.HTMLElement) {
// 		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
// 			imgUrl := el.ChildAttr("img", "src")
// 			imgUrl = strings.Split(imgUrl, "?")[0]
// 			imgUrl += "?preset=product-details-desktop"
// 			images = append(images, imgUrl)
// 		})
// 	})

// 	c.Visit(productUrl)
// 	description["사이즈 및 핏"] = sizeAndFit

// 	return
// }
