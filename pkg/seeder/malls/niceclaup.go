package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddNiceClaup() {
	modulename := "niceclaup"
	crawlUrl := "https://www.niceclaup.co.kr/category?id="
	categories := map[string]string{
		"아우터":  "150000231",
		"원피스":  "150000232",
		"상의":   "150000233",
		"니트웨어": "150000234",
		"하의":   "150000235",
		"신발":   "150000236",
		"가방":   "150000237",
		"기타소품": "150000238",
	}

	brands := map[string]domain.BrandDAO{
		"niceclaup": {
			KorName:       "나이스클랍",
			EngName:       "NICE CLAUP",
			KeyName:       "NICECLAUP",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/NICECLAUP.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "niceclaup",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/NICECLAUP_TOP:BOTTOM.png"},
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
				CrawlUrl:             crawlUrl + val,
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
				RefundRoughFee:       5500,
				ChangeRoughFee:       5500,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("Niceclaup categories & sources are added")
}
