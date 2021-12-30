package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddLacoste() {
	modulename := "lacoste"

	crawlUrl := "https://www.lacoste.com/kr/lacoste/sale/%EC%97%AC%EC%84%B1-%EC%84%B8%EC%9D%BC/"

	categories := map[string]string{
		"폴로":       "폴로",
		"셔츠":       "셔츠",
		"티셔츠":      "티셔츠",
		"스웨터&스웻셔츠": "스웨터-스웻셔츠",
		"드레스&스커트":  "드레스-스커트",
		"바지&반바지":   "바지-반바지",
		"자켓&코트":    "자켓-코트",
		"신발":       "신발",
		"레더굿":      "레더굿",
	}

	brand := domain.BrandDAO{
		KorName:       "라코스테",
		EngName:       "LACOSTE",
		KeyName:       "LACOSTE",
		Description:   "트래디셔널 캐주얼",
		LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/LACOSTE.png",
		Onpopular:     false,
		Created:       time.Now(),
		IsOpen:        true,
		Modulename:    "lacoste",
		InMaintenance: false,
		SizeGuide: []domain.SizeGuideDAO{
			{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LACOSTE_TOP.png"},
			{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LACOSTE_BOTTOM.png"},
			{Label: "데님/진", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LACOSTE_DENIM.png"},
			{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LACOSTE_SHOES.png"},
			{Label: "유니섹스 상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LACOSTE_TOP_UNISEX.png"},
		},
	}

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
			BrandKeyname:    upsertedBrand.KeyName,
			BrandIdentifier: brand.KeyName,
			MainCategoryKey: val,
			Category:        *updatedCat,
			CrawlUrl:        crawlUrl + val,
			CrawlModuleName: modulename,
		}

		_, err = ioc.Repo.CrawlSources.Upsert(&source)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("Lacoste categories & sources are added")
}
