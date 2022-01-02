package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddDaehyun() {
	modulename := "daehyun"
	crawlUrl := "https://www.daehyuninside.com/main/"

	brands := map[string]domain.BrandDAO{
		"1214": {
			KorName:       "쥬크",
			EngName:       "ZOOC",
			KeyName:       "ZOOC",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ZOOC.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "daehyun",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ZOOC_TOP.png"},
				{Label: "드레스", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ZOOC_DRESS.png"},
				{Label: "팬츠/스커트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ZOOC_PANTS:SKIRTS.png"},
				{Label: "자켓/코트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ZOOC_JACKET:COAT.png"},
			},
		},
		"1213": {
			KorName:       "모조에스핀",
			EngName:       "MOJO.S.PHINE",
			KeyName:       "MOJOSPHINE",
			Description:   "컨템포러리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MOJOSPHINE.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "daehyun",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MOJO.S.PHINE_TOP.png"},
				{Label: "드레스", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MOJO.S.PHINE_DRESS.png"},
				{Label: "팬츠/스커트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MOJO.S.PHINE_PANTS:SKIRTS.png"},
				{Label: "자켓/코트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MOJO.S.PHINE_JACKET:COAT.png"},
			},
		},
		"1215": {
			KorName:       "듀엘",
			EngName:       "DEWL",
			KeyName:       "DEWL",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DEWL.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "daehyun",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DEWL_TOP.png"},
				{Label: "드레스", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DEWL_DRESS.png"},
				{Label: "팬츠/스커트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DEWL_PANTS:SKIRTS.png"},
				{Label: "자켓/코트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DEWL_JACKET:COAT.png"},
			},
		},
		"1216": {
			KorName:       "씨씨콜렉트",
			EngName:       "CC collect",
			KeyName:       "CCCOLLECT",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/CCCOLLECT.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "daehyun",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CC+collect_TOP.png"},
				{Label: "드레스", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CC+collect_DRESS.png"},
				{Label: "팬츠/스커트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CC+collect_PANTS:SKIRTS.png"},
				{Label: "자켓/코트", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CC+collect_JACKET:COAT.png"},
			},
		},
	}

	ZOOC := map[string]string{
		"아우터": "1226",
		"드레스": "1227",
		"탑":   "1228",
		"팬츠":  "1229",
		"스커트": "1230",
	}

	DEWL := map[string]string{
		"아우터": "1232",
		"탑":   "1233",
		"드레스": "1234",
		"스커트": "1235",
		"팬츠":  "1236",
	}

	MOJOSPHINE := map[string]string{
		"아우터": "1218",
		"탑":   "1220",
		"니트":  "1221",
		"팬츠":  "1222",
		"스커트": "1223",
		"드레스": "1224",
	}

	CCCOLLECT := map[string]string{
		"아우터": "1238",
		"탑":   "1240",
		"드레스": "1239",
		"팬츠":  "1241",
		"스커트": "1242",
	}

	for brandId, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		var categories map[string]string

		if brand.KeyName == "ZOOC" {
			categories = ZOOC
		} else if brand.KeyName == "DEWL" {
			categories = DEWL
		} else if brand.KeyName == "MOJOSPHINE" {
			categories = MOJOSPHINE
		} else if brand.KeyName == "CCCOLLECT" {
			categories = CCCOLLECT
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
				BrandKeyname:    upsertedBrand.KeyName,
				BrandIdentifier: brandId,
				MainCategoryKey: updatedCat.CatIdentifier,
				Category:        *updatedCat,
				CrawlUrl:        crawlUrl,
				CrawlModuleName: modulename,
				IsSalesProducts: true,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}
	}

	log.Println("Daehyun categories & sources are added")
}
