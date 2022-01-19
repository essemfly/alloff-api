package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddTheAMall() {
	modulename := "theamall"
	crawlUrl := "http://www.theamall.com/product/list?catno="

	brands := map[string]domain.BrandDAO{
		"2102": {
			KorName:       "플라스틱 아일랜드",
			EngName:       "PLASTIC ISLAND",
			KeyName:       "PLASTICISLAND",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/PLASTICISLAND.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    modulename,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/PLASTICISLAND_TOP:BOTTOM.png"},
			},
		},
		"2101": {
			KorName:       "에고이스트",
			EngName:       "EGOIST",
			KeyName:       "EGOIST",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/EGOIST.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    modulename,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/EGOIST_TOP:BOTTOM.png"},
			},
		},
		"2103": {
			KorName:       "랩",
			EngName:       "LAP",
			KeyName:       "LAP",
			Description:   "라이프 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/LAP.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    modulename,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LAP_TOP:BOTTOM.png"},
			},
		},
	}

	// 카테고리 코드 안먹음.
	categories := map[string]string{
		"OUTER":  "01",
		"TOP":    "02",
		"DRESS":  "03",
		"BOTTOM": "04",
		"KIDS":   "07",
		"ACC":    "05",
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
				CrawlUrl:             crawlUrl + brandId + val,
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
	log.Println("Theamall categories & sources are added")
}
