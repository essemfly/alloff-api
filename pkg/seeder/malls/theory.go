package malls

import (
	"fmt"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddTheory() {
	modulename := "theory"

	categories := map[string]string{
		"악세서리":     "womens-new",
		"드레스&점프수트": "womens-dresses-and-jumpsuits",
		"자켓":       "womens-jackets",
		"라운지웨어":    "womens-loungewear",
		"아우터":      "womens-outerwear",
		"팬츠":       "womens-pants",
		"스커트":      "womens-skirts",
		"수트":       "womens-suits",
		"스웨터":      "womens-sweaters",
		"티셔츠&스웻셔츠": "womens-sweaters",
		"탑":        "womens-tops",
	}

	brands := map[string]domain.BrandDAO{
		"theory": {
			KorName:       "띠어리",
			EngName:       "Theory",
			KeyName:       "THEORY",
			Description:   "",
			LogoImgUrl:    "",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "드레스", ImgUrl: ""},
				{Label: "자켓", ImgUrl: ""},
				{Label: "하의", ImgUrl: ""},
				{Label: "신발", ImgUrl: ""},
				{Label: "상의", ImgUrl: ""},
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
				CrawlUrl:             getTheoryCrawlSourceUrl(val),
				CrawlModuleName:      modulename,
				IsSalesProducts:      true,
				IsForeignDelivery:    true,
				PriceMarginPolicy:    "THEORY",
				DeliveryPrice:        0,
				EarliestDeliveryDays: 7,
				LatestDeliveryDays:   21,
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
	log.Println("Theory categories & sources are added")
}

func getTheoryCrawlSourceUrl(cate string) string {
	return fmt.Sprintf("https://outlet.theory.com/on/demandware.store/Sites-theory_outlet-Site/default/Search-UpdateGrid?cgid=%s&start=0&sz=10000", cate)
}
