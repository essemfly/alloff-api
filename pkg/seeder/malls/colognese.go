package malls

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type BrandPage struct {
	ProductType  []domain.AlloffProductType
	Brand        *domain.BrandDAO
	BrandKeyname string
	URL          string
}

const IMAGE_BASE_URL = "http://colognese.atelier98.info"
const BASE_URL = "http://colognese.atelier98.info/en"
const BRAND_URL = "/brand.html"
const ALLOWED_DOMAIN = "colognese.atelier98.info"
const MODULE_NAME = "colognese"
const COOKIE = "sic=; ASPSESSIONIDSCRCBTRS=AAJDBLBCMIAAJKMEOKCOLEPA; can=; ASPSESSIONIDSASBDSQS=PKHNKECCIAODADMLKBBHCEAP; ASPSESSIONIDQCRDATQR=CCJDGANCLCKINHALAFHCBBAA; ASPSESSIONIDQCQCAQQS=AHNHBHOCEMOFIPOIEMMDAFFG; ASPSESSIONIDSARDARRS=NELBJAPCNMJNMBCMDDKCMGCO; ASPSESSIONIDSASCCRQQ=FIAHLIPCGOKLMCNMKCGMDLCG; impostazioni=trasporto=0&settore=woman&lingua=en&nazione=ITALY&n=30&list=ITA&idvaluta=1&valuta=%E2%82%AC&idnazione=1&carrellonew=3102618%2D%2D%2D2%2D%2D%2D10A%2D%2D%2D10A%2D%2D%2D0%7C%7C%7C3102618%2D%2D%2D2%2D%2D%2D8A%2D%2D%2D8A%2D%2D%2D0%7C%7C%7C4091092%2D%2D%2D1%2D%2D%2D12A%2D%2D%2D12A%2D%2D%2D0%7C%7C%7C"

func AddColognese() {
	var brandPages []*BrandPage

	c := colly.NewCollector(
		colly.AllowedDomains(ALLOWED_DOMAIN),
	)
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Cookie", COOKIE)
	})
	c.SetRequestTimeout(60 * time.Second)

	c.OnHTML(".elenco", func(e *colly.HTMLElement) {
		brandUrl := e.ChildAttr("a", "href")
		brandRawName := e.ChildText("a")
		brandKeyname := strings.ReplaceAll(brandRawName, ".", "")
		brandKeyname = strings.ReplaceAll(brandKeyname, " ", "")
		brand, err := ioc.Repo.Brands.GetByKeyname(brandKeyname)
		if err != nil {
			config.Logger.Info("Brand name: " + brandRawName + "newly added")
			brand, err = ioc.Repo.Brands.Upsert(
				&domain.BrandDAO{
					ID:                    primitive.NewObjectID(),
					KeyName:               brandKeyname,
					EngName:               brandRawName,
					IsOpen:                false,
					IsHide:                false,
					InMaintenance:         true,
					NumNewProductsIn3days: 0,
					UseAlloffCategory:     true,
					Created:               time.Now(),
				},
			)
			if err != nil {
				config.Logger.Error("err occured on brands upsert", zap.Error(err))
			}
		}

		productType := domain.Male
		if strings.Contains(brandUrl, "woman") {
			productType = domain.Female
		} else if strings.Contains(brandUrl, "kids") {
			productType = domain.Kids
		}

		brandPages = append(brandPages, &BrandPage{
			BrandKeyname: brandKeyname,
			URL:          BASE_URL + "/" + brandUrl,
			ProductType:  []domain.AlloffProductType{productType},
			Brand:        brand,
		})
	})

	c.Visit(BASE_URL + BRAND_URL)

	for idx, brandPage := range brandPages {
		addCologneseSource(idx, brandPage)
	}
}

func addCologneseSource(index int, brandPage *BrandPage) {
	source := domain.CrawlSourceDAO{
		ProductType:          brandPage.ProductType,
		BrandKeyname:         brandPage.BrandKeyname,
		BrandIdentifier:      brandPage.BrandKeyname + "-" + MODULE_NAME + "-" + strconv.Itoa(index),
		CrawlUrl:             brandPage.URL,
		CrawlModuleName:      MODULE_NAME,
		IsSalesProducts:      true,
		IsForeignDelivery:    true,
		PriceMarginPolicy:    "COLOGNESE",
		DeliveryPrice:        0,
		EarliestDeliveryDays: 14,
		LatestDeliveryDays:   21,
		DeliveryDesc:         nil,
		RefundAvailable:      true,
		ChangeAvailable:      true,
		RefundFee:            100000,
		ChangeFee:            100000,
	}

	_, err := ioc.Repo.CrawlSources.Upsert(&source)
	if err != nil {
		log.Println(err)
	}
}
