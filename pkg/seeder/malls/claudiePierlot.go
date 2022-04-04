package malls

import (
	"fmt"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"log"
	"time"
)

func AddClaudiePierlot() {
	modulename := "claudiePierlot"

	categories := map[string]string{
		"코트&자켓":   "coats-et-jackets",
		"드레스":     "dresses",
		"셔츠&티셔츠":  "tops---shirts-and-t-shirts",
		"니트&스웻셔츠": "knits-and-sweatshirts",
		"스커트&팬츠":  "trousers---skirts-and-shorts",
		"악세서리":    "accessories",
	}

	brands := map[string]domain.BrandDAO{
		"claudiePierlot": {
			KorName:       "끌로디 피에로",
			EngName:       "Claudie Pierlot",
			KeyName:       "CLAUDIEPIERLOT",
			Description:   "",
			LogoImgUrl:    "",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의&하의", ImgUrl: ""},
			},
		},
	}

	for brandId, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		for key, val := range categories {
			category := domain.CategoryDAO{
				Name:          key,
				KeyName:       brand.KeyName + "-" + key,
				CatIdentifier: val,
				BrandKeyname:  upsertedBrand.KeyName,
			}

			updatedCat, err := ioc.Repo.Categories.Upsert(&category)
			if err != nil {
				log.Println(err)
			}

			source := domain.CrawlSourceDAO{
				BrandKeyname:         upsertedBrand.KeyName,
				BrandIdentifier:      brandId,
				MainCategoryKey:      val,
				Category:             *updatedCat,
				CrawlUrl:             getClaudiePierlotCrawlSourceUrl(val),
				CrawlModuleName:      modulename,
				IsSalesProducts:      true,
				IsForeignDelivery:    true,
				PriceMarginPolicy:    "CLAUDIEPIERLOT",
				DeliveryPrice:        0,
				EarliestDeliveryDays: 7,
				LatestDeliveryDays:   14,
				DeliveryDesc:         nil,
				RefundAvailable:      true,
				ChangeAvailable:      true,
				RefundFee:            100000,
				ChangeFee:            100000,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("Claudie Pierlot categories & sources are added")
}

func getClaudiePierlotCrawlSourceUrl(cate string) string {
	return fmt.Sprintf("https://de.claudiepierlot.com/de_DE/categories/%s/?sz=10000&start=0", cate)
}
