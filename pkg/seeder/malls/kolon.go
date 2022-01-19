package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddKolonMall() {
	modulename := "kolon"
	categories := map[string]string{
		"티셔츠":     "133010080601",
		"블라우스/셔츠": "133010080602",
		"니트/가디건":  "133010080603",
		"스커트/팬츠":  "133010080604",
		"원피스":     "133010080605",
		// 다운, 패딩의 경우 겨울에만 있나봄?
		// "다운/패딩": "133010080606",
		"점퍼": "133010080607",
		"코트": "133010080608",
		"자켓": "133010080609",
	}

	brands := map[string]domain.BrandDAO{
		"IO_B1": {
			KorName:       "이로",
			EngName:       "IRO",
			KeyName:       "IRO",
			Description:   "컨템포러리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/IRO.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "kolon",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/IRO_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/IRO_BOTTOM.png"},
			},
		},

		"EP_B1": {
			KorName:       "에피그램",
			EngName:       "epigram",
			KeyName:       "EPIGRAM",
			Description:   "라이프 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/EPIGRAM.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "kolon",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/epigram_TOP:BOTTOM.png"},
			},
		},
		"JA_B191231114653": {
			KorName:       "세인트제임스",
			EngName:       "Saint James",
			KeyName:       "SAINTJAMES",
			Description:   "라이프 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SAINTJAMES.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "kolon",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Saint+James_TOP.png"},
			},
		},

		"LC_B1": {
			KorName:       "럭키슈에뜨",
			EngName:       "LUCKY CHOUETTE",
			KeyName:       "LUCKYCHOUETTE",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/LUCKYCHOUETTE.png",
			Onpopular:     true,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "kolon",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LUCKY+CHOUETTE_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LUCKY+CHOUETTE_BOTTOM.png"},
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
				MainCategoryKey:      updatedCat.CatIdentifier,
				Category:             *updatedCat,
				CrawlUrl:             "https://www.kolonmall.com/Category/List/" + val + "?supplierBrands=" + brandId,
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
	log.Println("Kolon mall brands, categories & sources are added")
}
