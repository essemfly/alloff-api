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
	sandroExhibitions := []string{
		"outlet",
		"last-chance",
		"sale",
	}
	crawlUrl := "https://de.sandro-paris.com/de/damen/%s/jede-auswahl/?prefn1=smcp_subFamily&prefv1=%s&sz=1000000&format=ajax"

	categories := map[string]string{
		"CLOHTES@%s":                "Kleider",
		"COATS@%s":                  "Mäntel",
		"JACKETS@%s":                "Jacken",
		"TOPS_AND_SHIRTS@%s":        "Tops%20%26%20Hemden",
		"JEANS@%s":                  "Jeans",
		"PANTS@%s":                  "Hosen",
		"T_SHIRTS@%s":               "T-shirts",
		"SKIRTS_AND_SHORTS@%s":      "Röcke%20%26%20Shorts",
		"PULLOVER_AND_CARDIGANS@%s": "Pullover%20%26%20Cardigans",
		"ACCESSORIES@%s":            "Accessoires",
	}

	brands := map[string]domain.BrandDAO{
		"Sandro": {
			KorName:       "산드로",
			EngName:       "sandro",
			KeyName:       "SANDRO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SANDRO.png",
			Onpopular:     true,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/sandro_TOP:BOTTOM.png"},
				{Label: "데님/진", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/sandro_DENIM:JEAN.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/sandro_SHOES.png"},
			},
		},
	}

	for brandId, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		for _categoryKey, categoryValue := range categories {
			for _, exhibition := range sandroExhibitions {
				categoryKey := fmt.Sprintf(_categoryKey, exhibition)
				keyname := brand.KeyName + "-" + categoryKey
				category := domain.CategoryDAO{
					Name:          categoryKey,
					KeyName:       keyname,
					CatIdentifier: keyname,
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
					CrawlUrl:             fmt.Sprintf(crawlUrl, exhibition, categoryValue),
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
	}
	log.Println("Sandro categories & sources are added")
}
