package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddLoungeB() {
	modulename := "loungeb"
	crawlUrl := "https://lounge-b.co.kr/product/list.html?cate_no="

	brands := map[string]domain.BrandDAO{
		"B000000A, B000000C": {
			KorName:       "온앤온",
			EngName:       "ON & ON",
			KeyName:       "ONANDON",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ONANDON.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    modulename,
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ONANDON_TOP:BOTTOM.png"},
			},
		},
		"B000000D": {
			KorName:       "올리브데올리브",
			EngName:       "OLIVE DES OLIVE",
			KeyName:       "OLIVEDESOLIVE",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/OLIVEDESOLIVE.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    modulename,
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/OLIVEDESOLIVE_TOP:BOTTOM.png"},
			},
		},
	}

	// ?keyword=&search_form[option_data][]=brand=B000000C&search_form[option_data][]=brand=B000000D&cate_no=76&sShowOpen=F

	categories := map[string]string{
		"탑웨어":  "76",
		"아우터":  "75",
		"니트웨어": "77",
		"원피스":  "78",
		"스커트":  "79",
		"팬츠":   "80",
		"패션잡화": "281",
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

			brandCrawlUrl := crawlUrl + val
			if brand.KeyName == "ONANDON" {
				brandCrawlUrl += "&search_form[option_data][]=brand=B000000A&search_form[option_data][]=brand=B000000C&sShowOpen"
			} else {
				brandCrawlUrl += "&search_form[option_data][]=brand=B000000D&sShowOpen"
			}

			source := domain.CrawlSourceDAO{
				BrandKeyname:         upsertedBrand.KeyName,
				BrandIdentifier:      brandId,
				MainCategoryKey:      updatedCat.CatIdentifier,
				Category:             *updatedCat,
				CrawlUrl:             brandCrawlUrl,
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
	log.Println("LoungeB categories & sources are added")
}
