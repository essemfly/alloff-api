package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddBenetton() {
	modulename := "benetton"
	crawlUrl := "https://www.benettonmall.com/product/list"

	categories := map[string]string{
		"TOPS":    "001",
		"BOTTOMS": "002",
		"ACC":     "003",
	}

	brands := map[string]domain.BrandDAO{
		"benetton": {
			KorName:       "베네통",
			EngName:       "BENETTON",
			KeyName:       "BENETTON",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BENETTON.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/BENETTON_TOP.png"},
				{Label: "드레스/코트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/BENETTON_DRESS:COAT.png"},
				{Label: "아우터", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/BENETTON_OUTER.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/BENETTON_BOTTOM.png"},
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
	log.Println("Benetton categories & sources are added")
}
