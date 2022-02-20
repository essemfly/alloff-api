package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddIntrend() {
	modulename := "intrend"
	crawlUrl := "https://it.intrend.it/"

	// 카테고리 코드 안먹음.
	categories := map[string]string{
		"SOPRABITI": "cappotti-e-giacche/impermeabili",
		"BORSE":     "borse-e-accessori/borse",
		"SCARPE":    "scarpe/tutte",
		"PANTALONI": "abbigliamento/pantaloni",
		"T-SHIRT":   "abbigliamento/top-e-t-shirt",
		"ABITI":     "abbigliamento/abiti",
	}

	brands := map[string]domain.BrandDAO{
		"InTrend": {
			KorName:       "막스마라(인트렌드)",
			EngName:       "MaxMara(Intrend)",
			KeyName:       "INTREND",
			Description:   "럭셔리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MAXMARA.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "intrend",
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
				CrawlUrl:             crawlUrl + val,
				CrawlModuleName:      modulename,
				IsSalesProducts:      true,
				IsForeignDelivery:    true,
				PriceMarginPolicy:    "INTREND",
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
	log.Println("Intrend categories & sources are added")
}
