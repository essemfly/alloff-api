package malls

import (
	"fmt"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddTheoryOutlet() {
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
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Theory_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Theory_BOTTOM.png"},
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
				CrawlUrl:             getTheoryOutletCrawlSourceUrl(val),
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

func AddTheory() {
	modulename := "theory"
	categories := map[string]string{
		"여성세일": "womens-sale",
		"남성세일": "mens-sale",
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
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Theory_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Theory_BOTTOM.png"},
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
}

func getTheoryOutletCrawlSourceUrl(cate string) string {
	return fmt.Sprintf("https://outlet.theory.com/on/demandware.store/Sites-theory_outlet-Site/default/Search-UpdateGrid?cgid=%s&start=0&sz=10000", cate)
}

func getTheoryCrawlSourceUrl(cate string) string {
	return fmt.Sprintf("https://theory.com/on/demandware.store/Sites-theory2_US-Site/default/Search-UpdateGrid?cgid=%s&start=0&sz=10000", cate)
}
