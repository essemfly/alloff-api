package malls

import (
	"encoding/json"
	"log"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"github.com/lessbutter/alloff-api/pkg/seeder/malls"
	"go.uber.org/zap"
)

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

	brand, err := ioc.Repo.Brands.GetByKeyname(source.BrandKeyname)
	if err != nil {
		log.Println(err)
	}

	baseUrl := strings.Replace(source.CrawlUrl, "productsPerPage=10", "productsPerPage=1000", 1)
	productRequests := CrawlFlannelsProducts(baseUrl, typeVariations, categoryVariations)
	for _, req := range productRequests {
		req.Brand = brand
		req.Source = source
		productinfo.ProcessCrawlingInfoRequests(req)
	}

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
		newRequest.AlloffCategory = &domain.AlloffCategoryDAO{}
		requests = append(requests, newRequest)
	}

	return requests
}

func GetFlannelsDetail(productURL string) *productinfo.AddMetaInfoRequest {
	productRequest := &productinfo.AddMetaInfoRequest{
		ProductUrl:          productURL,
		IsForeignDelivery:   true,
		IsTranslateRequired: true,
		IsRemoved:           false,
		IsSoldout:           false,
		ModuleName:          malls.FLANNELS_MODULE_NAME,
	}

	colorSplits := strings.Split(productURL, "colcode=")
	colorCode := ""
	if len(colorSplits) > 1 {
		colorCode = colorSplits[1]
	}

	c := colly.NewCollector(
		colly.AllowedDomains("www.flannels.com"),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"),
	)
	c.SetRequestTimeout(10 * time.Second)

	infos := map[string]string{}
	c.OnHTML("#lblProductName", func(e *colly.HTMLElement) {
		log.Println("name", e.Text)
		productRequest.AlloffName = e.Text
	})
	c.OnHTML(".product-detail__price", func(e *colly.HTMLElement) {
		productRequest.OriginalPrice = parseOnlyNumbers(e.ChildText("#lblSellingPrice"))
		productRequest.DiscountedPrice = parseOnlyNumbers(e.ChildText(".originalprice #lblTicketPrice"))
		productRequest.CurrencyType = domain.CurrencyPOUND
	})
	c.OnHTML("#lblProductCode", func(e *colly.HTMLElement) {
		productCodeDiv := e.Text
		productCode := parseOnlyNumbers(productCodeDiv)
		productRequest.ProductID = strconv.Itoa(int(productCode))
		infos["제품코드"] = strconv.Itoa(int(productCode))
	})
	c.OnHTML("#divColour #divColourImageDropdownGroup .dropdown-images .image-dropdown #btnImageDropdown #spanDropdownSelectedText", func(e *colly.HTMLElement) {
		color := e.Text
		productRequest.Colors = []string{color}
		infos["색상"] = color
	})
	c.OnHTML(".productImage #productRollOverPanel_"+colorCode, func(e *colly.HTMLElement) {
		images := []string{}
		e.ForEach(".swiper-wrapper .swiper-slide a", func(i int, el *colly.HTMLElement) {
			images = append(images, el.Attr("href"))
		})
		productRequest.Images = images
	})

	sizes := []string{}
	c.OnHTML("#divSize #spanSize", func(e *colly.HTMLElement) {
		s := strings.TrimSpace(e.Text)
		sizes = append(sizes, s)
	})
	c.OnHTML("#sizeDdl", func(e *colly.HTMLElement) {
		e.ForEach("option", func(i int, el *colly.HTMLElement) {
			if el.Attr("value") != "0" {
				sizes = append(sizes, el.Attr("value"))
			}
		})
	})

	c.OnHTML("#DisplayAttributes", func(e *colly.HTMLElement) {
		keys, values := []string{}, []string{}
		e.ForEach("dt", func(i int, el *colly.HTMLElement) {
			keys = append(keys, el.Text)
		})
		e.ForEach("dd", func(_ int, el *colly.HTMLElement) {
			values = append(values, el.Text)
		})

		for i, val := range values {
			if keys[i] == "Fabric" {
				infos["소재"] = val
			} else if keys[i] == "Style" {
				infos["스타일"] = val
			} else {
				infos[keys[i]] = val
			}
		}
	})

	c.Visit(productURL)

	productRequest.Information = infos
	productRequest.DescriptionInfos = infos
	productRequest.Sizes = sizes
	invs := []*domain.InventoryDAO{}
	for _, size := range sizes {
		invs = append(invs, &domain.InventoryDAO{
			Size:     size,
			Quantity: 1,
		})
	}
	productRequest.Inventory = invs
	return productRequest
}

func parseOnlyNumbers(texts string) float32 {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	onlyNumbers := re.FindAllString(texts, -1)
	num, _ := strconv.ParseFloat(onlyNumbers[0], 32)
	return float32(num)
}

func flannelsCategoryMapper(categoryKeyname string) *domain.AlloffCategoryDAO {

	return &domain.AlloffCategoryDAO{}
}
