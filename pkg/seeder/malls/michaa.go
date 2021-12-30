package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddMichaa() {
	modulename := "michaa"
	crawlUrl := "https://sisun.com/MICHAA/shop/list?"

	categories := map[string]string{
		"코트":   "110501",
		"자켓":   "110502",
		"점퍼":   "110503",
		"티셔츠":  "110504",
		"블라우스": "110505",
		"니트":   "110506",
		"원피스":  "110507",
		"스커트":  "110508",
		"팬츠":   "110509",
		"액세서리": "110510",
		"베스트":  "110511",
		"프리오픈": "110515",
	}

	brand := domain.BrandDAO{
		KorName:       "미샤",
		EngName:       "MICHAA",
		KeyName:       "MICHAA",
		LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MICHAA.png",
		Onpopular:     false,
		Description:   "컨템포러리",
		Created:       time.Now(),
		IsOpen:        true,
		Modulename:    "michaa",
		InMaintenance: false,
		SizeGuide: []domain.SizeGuideDAO{
			{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MICHAA_TOP:BOTTOM.png"},
		},
	}
	upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
	if err != nil {
		log.Println(err)
	}

	for key, val := range categories {
		// CategoryDAO 먼저 다 업데이트하고 나면, 그 리스트를 받아야 한다.
		// 각각 받은 CategoryDAO들을 바탕으로
		// BrandDAO 와 CrawlSourceDAO를 업데이트하면된다.
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
			CrawlUrl:        crawlUrl + "cate_code=" + val + "&soldout_include=N&upper_cate_code=" + "1105",
			CrawlModuleName: modulename,
		}

		_, err = ioc.Repo.CrawlSources.Upsert(&source)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("MICHAA categories & sources are added")
}
