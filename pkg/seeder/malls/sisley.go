package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddSisley() {
	modulename := "sisley"
	crawlUrl := "https://www.sisleymall.com/product/list"
	categories := map[string]string{
		"TOPS":    "001",
		"BOTTOMS": "002",
		"OUTER":   "005",
		"DRESS":   "006",
		"BAG":     "003",
		"ACC":     "004",
	}

	brands := map[string]domain.BrandDAO{
		"sisley": {
			KorName:       "시슬리",
			EngName:       "sisley",
			KeyName:       "SISLEY",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SISLEY.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SISLEY_TOP.png"},
				{Label: "드레스", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SISLEY_DRESS.png"},
				{Label: "아우터", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SISLEY_OUTER.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SISLEY_BOTTOM.png"},
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
				CrawlUrl:             crawlUrl,
				CrawlModuleName:      modulename,
				IsSalesProducts:      true,
				IsForeignDelivery:    false,
				PriceMarginPolicy:    "NORMAL",
				DeliveryPrice:        0,
				EarliestDeliveryDays: 2,
				LatestDeliveryDays:   7,
				DeliveryDesc:         nil,
				RefundAvailable:      true,
				ChangeAvailable:      true,
				RefundFee:            5000,
				ChangeFee:            5000,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("Sisley categories & sources are added")
}
