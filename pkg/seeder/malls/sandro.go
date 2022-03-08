package malls

import (
	"fmt"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddSandro() {
	modulename := "sandro"
	crawlUrl := "https://de.sandro-paris.com/de/damen/last-chance/jede-auswahl/?prefn1=smcp_subFamily&prefv1=%s&sz=1000000&format=ajax"

	categories := map[string]string{
		"CLOHTES":                "Kleider",
		"COATS":                  "Mäntel",
		"JACKETS":                "Jacken",
		"TOPS_AND_SHIRTS":        "Tops%20%26%20Hemden",
		"JEANS":                  "Jeans",
		"PANTS":                  "Hosen",
		"T_SHIRTS":               "T-shirts",
		"SKIRTS_AND_SHORTS":      "Röcke%20%26%20Shorts",
		"PULLOVER_AND_CARDIGANS": "Pullover%20%26%20Cardigans",
		"ACCESSORIES":            "Accessoires",
	}

	brands := map[string]domain.BrandDAO{
		"InTrend": {
			KorName:       "산드로",
			EngName:       "Sandro",
			KeyName:       "SANDRO",
			Description:   "산드로입니다.",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MAXMARA.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
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
				CatIdentifier: key,
				BrandKeyname:  upsertedBrand.KeyName,
			}

			updatedCat, err := ioc.Repo.Categories.Upsert(&category)
			if err != nil {
				log.Println(err)
			}

			source := domain.CrawlSourceDAO{
				BrandKeyname:         upsertedBrand.KeyName,
				BrandIdentifier:      brandId,
				MainCategoryKey:      updatedCat.CatIdentifier,
				Category:             *updatedCat,
				CrawlUrl:             fmt.Sprintf(crawlUrl, val),
				CrawlModuleName:      modulename,
				IsSalesProducts:      true,
				IsForeignDelivery:    true,
				PriceMarginPolicy:    "SANDRO",
				DeliveryPrice:        0,
				EarliestDeliveryDays: 14,
				LatestDeliveryDays:   21,
				DeliveryDesc:         nil,
				RefundAvailable:      false,
				ChangeAvailable:      false,
				RefundFee:            5000,
				ChangeFee:            5000,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("Sandro categories & sources are added")
}
