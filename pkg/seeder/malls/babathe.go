package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddBabathe() {
	modulename := "babathe"
	crawlUrl := "https://pc.babathe.com/search/getGoodsListAjax"

	categories := map[string]string{
		"상의":  "100010083",
		"아우터": "100010081",
		"원피스": "100010082",
		"잡화":  " 100010089",
		"하의":  " 100010084",
		"주얼리": " 100010103",
	}

	brands := map[string]domain.BrandDAO{
		"10387": {
			KorName:       "지고트",
			EngName:       "JIGOTT",
			KeyName:       "JIGOTT",
			Description:   "컨템포러리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JIGOTT.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "babathe",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/JIGOTT_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/JIGOTT_BOTTOM.png"},
			},
		},
		"10183": {
			KorName:       "아이잗 바바",
			EngName:       "IZZAT BABA",
			KeyName:       "IZZATBABA",
			Description:   "커리어",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/IZZATBABA.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "babathe",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/IZZAT+BABA_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/IZZAT+BABA_BOTTOM.png"},
			},
		},
		"10388": {
			KorName:       "더 아이잗 컬렉션",
			EngName:       "THE IZZAT COLLECTION",
			KeyName:       "THEIZZATCOLLECTION",
			Description:   "컨템포러리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/THEIZZATCOLLECTION.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "babathe",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/THE+IZZAT+COLLECTION_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/THE+IZZAT+COLLECTION_BOTTOM.png"},
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
				RefundRoughFee:       5500,
				ChangeRoughFee:       5500,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}

	}
	log.Println("Babathe categories & sources are added")
}
