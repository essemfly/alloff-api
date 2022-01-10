package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddIDFMall() {
	modulename := "idfmall"
	crawlUrl := "https://www.idfmall.co.kr/main/"

	brands := map[string]domain.BrandDAO{
		"SHESMISS": {
			KorName:       "쉬즈미스",
			EngName:       "SHESMISS",
			KeyName:       "SHESMISS",
			Description:   "커리어",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SHESMISS.png",
			Onpopular:     true,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SHESMISS_TOP:BOTTOM.png"},
			},
		},
		"LIST": {
			KorName:       "리스트",
			EngName:       "LIST",
			KeyName:       "LIST",
			Description:   "라이프 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/LIST.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "idfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LIST_TOP:BOTTOM.png"},
			},
		},
	}

	shesmissCategory := map[string]string{
		"코트":      "1177",
		"재킷&베스트":  "1178",
		"패딩&점퍼":   "1179",
		"블라우스&셔츠": "1180",
		"티셔츠":     "1181",
		"니트&가디건":  "1182",
		"원피스":     "1183",
		"스커트":     "1184",
		"팬츠":      "1185",
	}

	listCategory := map[string]string{
		"코트":      "1187",
		"재킷&베스트":  "1188",
		"패딩&점퍼":   "1189",
		"블라우스&셔츠": "1190",
		"티셔츠":     "1191",
		"니트&가디건":  "1192",
		"원피스":     "1193",
		"스커트":     "1194",
		"팬츠":      "1195",
	}

	for brandId, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		var categories map[string]string

		if brand.KeyName == "SHESMISS" {
			categories = shesmissCategory
		} else if brand.KeyName == "LIST" {
			categories = listCategory
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
				BrandKeyname:      upsertedBrand.KeyName,
				BrandIdentifier:   brandId,
				MainCategoryKey:   updatedCat.CatIdentifier,
				Category:          *updatedCat,
				CrawlUrl:          crawlUrl,
				CrawlModuleName:   modulename,
				IsSalesProducts:   true,
				IsForeignDelivery: false,
				PriceMarginPolicy: "NORMAL",
				DeliveryPrice:     0,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("IDFmall categories & sources are added")
}
