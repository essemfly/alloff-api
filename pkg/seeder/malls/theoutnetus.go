package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddTheoutnet() {
	modulename := "theoutnet"
	// dspCtgryNo = 아우터, 베스트, 이런 것들
	categories := map[string]string{
		"티셔츠":     "SFMA44A02A01",
		"셔츠/블라우스": "SFMA44A02A02",
		"니트":      "SFMA44A02A03",
		"팬츠":      "SFMA44A02A04",
		"스커트":     "SFMA44A02A05",
		"원피스":     "SFMA44A02A06",
		"아우터":     "SFMA44A02A07",
		"재킷/베스트":  "SFMA44A02A17",
	}

	brands := map[string]domain.BrandDAO{
		"BDMA07A34/ECBPC": {
			KorName:       "띠어리",
			EngName:       "Theory",
			KeyName:       "THEORY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/THEORY.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Theory_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Theory_BOTTOM.png"},
			},
		},
	}

	for shopId, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		if brand.IsOpen {
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
					BrandIdentifier:   shopId,
					MainCategoryKey:   updatedCat.CatIdentifier,
					Category:          *updatedCat,
					CrawlUrl:          buildCrawlUrl(shopId, updatedCat.CatIdentifier),
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
	}
	log.Println("SSF mall brands, categories & sources are added")
}
