package malls

import (
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
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

	log.Println("length of products", len(listQueryResp.Products))

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
	CrawlFlannelsProducts(baseUrl, typeVariations, categoryVariations)

	// crawlURL, err := url.ParseQuery(source.CrawlUrl)
	// log.Println("crwlUrl", crawlURL)
	// if err != nil {
	// 	config.Logger.Error("Parse error in source crawling "+source.ID.Hex(), zap.Error(err))
	// }

	// _, err = ioc.Repo.Brands.GetByKeyname(source.BrandKeyname)
	// if err != nil {
	// 	config.Logger.Error("error occurred on get brand by key name : ", zap.Error(err))
	// }

	<-worker
	done <- true
}

func CrawlFlannelsProducts(baseUrl string, productTypes []string, categories []string) {
	urls := []string{}
	if len(productTypes) == 0 {
		for _, cat := range categories {
			urlparams = append(urlparams, "&selectedFilters=1147_726918^"+cat)
		}
		return
	}

	for _, cat := range categories {
		for _, ptype := range productTypes {
			urlparams = append(urlparams, "&selectedFilters=AFLOR^"+ptype+"|1147_726918^"+cat)
		}
	}

	return
}

func getFlannelsDetail(productUrl string) (images, colors []string, infos, description map[string]string, notSale bool) {
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
