package malls

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
	"go.uber.org/zap"
)

type BrandPage struct {
	ProductType  []*domain.AlloffProductType
	Brand        *domain.BrandDAO
	BrandKeyname string
	URL          string
}

const IMAGE_BASE_URL = "http://colognese.atelier98.info"
const BASE_URL = "http://colognese.atelier98.info/en"
const BRAND_URL = "/brand.html"
const ALLOWED_DOMAIN = "colognese.atelier98.info"
const MODULE_NAME = "colognese"
const COOKIE = "sic=; ASPSESSIONIDSCRCBTRS=AAJDBLBCMIAAJKMEOKCOLEPA; can=; ASPSESSIONIDSASBDSQS=PKHNKECCIAODADMLKBBHCEAP; ASPSESSIONIDQCRDATQR=CCJDGANCLCKINHALAFHCBBAA; ASPSESSIONIDQCQCAQQS=AHNHBHOCEMOFIPOIEMMDAFFG; ASPSESSIONIDSARDARRS=NELBJAPCNMJNMBCMDDKCMGCO; ASPSESSIONIDSASCCRQQ=FIAHLIPCGOKLMCNMKCGMDLCG; impostazioni=trasporto=0&settore=WOMAN&lingua=en&nazione=ITALY&n=30&list=ITA&idvaluta=1&valuta=%E2%82%AC&idnazione=1&carrellonew=3102618%2D%2D%2D2%2D%2D%2D10A%2D%2D%2D10A%2D%2D%2D0%7C%7C%7C3102618%2D%2D%2D2%2D%2D%2D8A%2D%2D%2D8A%2D%2D%2D0%7C%7C%7C4091092%2D%2D%2D1%2D%2D%2D12A%2D%2D%2D12A%2D%2D%2D0%7C%7C%7C"

func CrawlColognese(worker chan bool, done chan bool, source *domain.CrawlSourceDAO) {
	c := colly.NewCollector(
		colly.AllowedDomains(ALLOWED_DOMAIN),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", COOKIE)
	})
	c.SetRequestTimeout(60 * time.Second)

	pageNum := 1
	url := source.CrawlUrl + "?page=" + strconv.Itoa(pageNum)
	productRequests := []*productinfo.AddMetaInfoRequest{}

	pageRequests := ListProductsFromPage(url)
	productRequests = append(productRequests, pageRequests...)
	for len(pageRequests) >= 30 {
		pageNum += 1
		url = source.CrawlUrl + "?page=" + strconv.Itoa(pageNum)
		pageRequests = ListProductsFromPage(url)
		productRequests = append(productRequests, pageRequests...)
	}

	for _, productRequest := range productRequests {
		GetProductInfo(productRequest)
	}

	brand, err := ioc.Repo.Brands.GetByKeyname(source.BrandKeyname)
	if err != nil {
		config.Logger.Error("cannot find brane keyname in cologness crawler", zap.Error(err))
	}

	for _, req := range productRequests {
		req.Brand = brand
		req.Source = source
		req.ProductType = source.ProductType

		if req.AlloffCategory != nil {
			if req.AlloffCategory.KeyName == "1_bags" || req.AlloffCategory.KeyName == "1_shoes" || req.AlloffCategory.KeyName == "1_accessory" || req.AlloffCategory.KeyName == "1_jewelry" {
				req.Source.PriceMarginPolicy = "COLOGNESE_NON_FASHION"
			}
		}

		productinfo.ProcessCrawlingInfoRequests(req)
	}

	<-worker
	done <- true
}

func ListProductsFromPage(pageURL string) (requests []*productinfo.AddMetaInfoRequest) {
	c := colly.NewCollector(
		colly.AllowedDomains(ALLOWED_DOMAIN),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", COOKIE)
	})
	c.SetRequestTimeout(60 * time.Second)

	c.OnHTML(".contfoto", func(e *colly.HTMLElement) {
		originPriceStr := e.ChildText(".testofoto .prezzo .saldi")
		originPriceStr = strings.ReplaceAll(originPriceStr, "€", "")
		originPriceStr = strings.ReplaceAll(originPriceStr, " ", "")
		originPriceStr = strings.ReplaceAll(originPriceStr, ",", ".")
		originalPrice, _ := strconv.ParseFloat(originPriceStr, 32)

		currentPriceStr := e.ChildText(".testofoto .prezzo .saldi2")
		currentPriceStr = strings.ReplaceAll(currentPriceStr, "€", "")
		currentPriceStr = strings.ReplaceAll(currentPriceStr, " ", "")
		currentPriceStr = strings.ReplaceAll(currentPriceStr, ",", ".")
		currentPrice, _ := strconv.ParseFloat(currentPriceStr, 32)

		information := map[string]string{}
		information["SEASON"] = e.ChildText(".percstagione")
		imageURL := IMAGE_BASE_URL + e.ChildAttr(".cotienifoto a img", "src")
		imageURL = strings.ReplaceAll(imageURL, "Thumbs_", "")
		imageURL = strings.ReplaceAll(imageURL, "thumbs_", "")
		requests = append(requests, &productinfo.AddMetaInfoRequest{
			ProductUrl:      BASE_URL + "/" + e.ChildAttr(".cotienifoto a", "href"),
			Images:          []string{imageURL},
			AlloffName:      e.ChildText(".testofoto a .notranslate"),
			ProductID:       e.ChildText(".testofoto div strong"),
			OriginalPrice:   float32(originalPrice),
			DiscountedPrice: float32(currentPrice),
			CurrencyType:    domain.CurrencyEUR,
			Information:     information,
			Inventory:       []*domain.InventoryDAO{},
		})
	})

	c.Visit(pageURL)

	return
}

func GetProductInfo(request *productinfo.AddMetaInfoRequest) *productinfo.AddMetaInfoRequest {
	c := colly.NewCollector(
		colly.AllowedDomains(ALLOWED_DOMAIN),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", COOKIE)
	})
	c.SetRequestTimeout(60 * time.Second)

	c.OnHTML(".bloccodett", func(e *colly.HTMLElement) {
		newInv := &domain.InventoryDAO{}
		e.ForEach("div #taglia", func(i int, h *colly.HTMLElement) {
			newInv.Quantity = 1
			newInv.Size = h.Attr("value")
		})
		request.Inventory = append(request.Inventory, newInv)
	})

	keys := []string{}
	values := []string{}
	c.OnHTML("h2", func(e *colly.HTMLElement) {
		request.AlloffCategory = cologneseCategoryMapper(e.Text)
	})

	c.OnHTML(".blacktxt.col3", func(e *colly.HTMLElement) {
		keys = append(keys, e.Text)
	})
	c.OnHTML(".col9.last", func(e *colly.HTMLElement) {
		values = append(values, e.Text)
	})

	c.Visit(request.ProductUrl)

	for i, key := range keys {
		request.Information[key] = values[i]
	}

	return nil
}

func cologneseCategoryMapper(categoryKeyname string) *domain.AlloffCategoryDAO {
	if strings.Contains(categoryKeyname, "TRENCH") || strings.Contains(categoryKeyname, "COAT") || strings.Contains(categoryKeyname, "JACKET") || strings.Contains(categoryKeyname, "BOMBER") || strings.Contains(categoryKeyname, "PARKA") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_outer")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "CARDIGAN") || strings.Contains(categoryKeyname, "POLO") || strings.Contains(categoryKeyname, "BLOUSE") || strings.Contains(categoryKeyname, "KNITWEAR") || strings.Contains(categoryKeyname, "SWEATER") || strings.Contains(categoryKeyname, "TOP") || strings.Contains(categoryKeyname, "SHIRT") || strings.Contains(categoryKeyname, "HOODIE") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_top")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "PANTS") || strings.Contains(categoryKeyname, "TROUSERS") || strings.Contains(categoryKeyname, "SHORTS") || strings.Contains(categoryKeyname, "JEANS") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_bottom")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "DRESS") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_onePiece")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "SKIRT") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_skirt")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "BRAS") || strings.Contains(categoryKeyname, "LINGERIE") || strings.Contains(categoryKeyname, "SOCKS") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_underwear")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "HANDBAGS") || strings.Contains(categoryKeyname, "BAG") || strings.Contains(categoryKeyname, "CLUTCH") || strings.Contains(categoryKeyname, "POUCH") || strings.Contains(categoryKeyname, "CROSSBODY") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_bags")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "BALLET") || strings.Contains(categoryKeyname, "SNEAKERS") || strings.Contains(categoryKeyname, "SHOES") || strings.Contains(categoryKeyname, "SLIPPERS") || strings.Contains(categoryKeyname, "TRAINERS") || strings.Contains(categoryKeyname, "BOOTS") || strings.Contains(categoryKeyname, "SANDALS") || strings.Contains(categoryKeyname, "DECOLLETE") || strings.Contains(categoryKeyname, "LOAFER") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_shoes")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "GLOVES") || strings.Contains(categoryKeyname, "SCARF") || strings.Contains(categoryKeyname, "DIFFUSER") || strings.Contains(categoryKeyname, "BEAUTY") || strings.Contains(categoryKeyname, "ACCESSORY") || strings.Contains(categoryKeyname, "ACCESSORIES") || strings.Contains(categoryKeyname, "SUNGLASS") || strings.Contains(categoryKeyname, "BELT") || strings.Contains(categoryKeyname, "HAT") || strings.Contains(categoryKeyname, "CAP") || strings.Contains(categoryKeyname, "WALLET") || strings.Contains(categoryKeyname, "PERFUME") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_accessory")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "RING") || strings.Contains(categoryKeyname, "EARRINGS") || strings.Contains(categoryKeyname, "NECKLACE") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_jewelry")
		return alloffcat
	}
	if strings.Contains(categoryKeyname, "JUMPSUIT") || strings.Contains(categoryKeyname, "SWIMWEAR") || strings.Contains(categoryKeyname, "BEACHWEAR") {
		alloffcat, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_beachwear")
		return alloffcat
	}

	config.Logger.Warn("Not matched category keyname in colognese: " + categoryKeyname)
	return &domain.AlloffCategoryDAO{}
}
