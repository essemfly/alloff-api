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

	crawlUrl := "https://de.sandro-paris.com/de/damen/%s/%s/?sz=10000&format=ajax"

	categories := map[string]string{
		"원피스":    "kleider",
		"코트/자켓":  "blousons-mantel",
		"니트":     "pullovers-cardigans",
		"상의":     "tops-hemden",
		"바지":     "hosen-jeans",
		"치마/반바지": "rocke-shorts",
		"가방":     "lederwaren",
		"신발":     "schuhe",
		"악세서리":   "andere-accessoires",
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
			IsOpen:        true,
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

		for categoryKey, categoryValue := range categories {
			for _, exhibition := range sandroExhibitions {
				keyname := brand.KeyName + "-" + categoryKey + "-" + exhibition
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
					IsSalesProducts:      false,
					IsForeignDelivery:    true,
					PriceMarginPolicy:    "SANDRO",
					DeliveryPrice:        0,
					EarliestDeliveryDays: 10,
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
	}
	log.Println("Sandro categories & sources are added")
}
