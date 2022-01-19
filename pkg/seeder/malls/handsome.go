package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddHandsome() {
	modulename := "handsome"
	categories := map[string]string{
		"탑":            "we01",
		"팬츠":           "we02",
		"스커트":          "we03",
		"드레스":          "we04",
		"아우터":          "we05",
		"Special Shop": "we09",
	}

	brands := map[string]domain.BrandDAO{
		"br01": {
			KorName:       "타임",
			EngName:       "TIME",
			KeyName:       "TIME",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/TIME.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/TIME_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/TIME_BOTTOM.png"},
			},
		},
		"br02": {
			KorName:       "마인",
			EngName:       "MINE",
			KeyName:       "MINE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MINE.png",
			Onpopular:     true,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MINE_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MINE_BOTTOM.png"},
			},
		},
		"br03": {
			KorName:       "시스템",
			EngName:       "SYSTEM",
			KeyName:       "SYSTEM",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SYSTEM.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SYSTEM_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SYSTEM_BOTTOM.png"},
			},
		},
		"br04": {
			KorName:       "에스제이에스제이",
			EngName:       "SJSJ",
			KeyName:       "SJSJ",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SJSJ.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SJSJ_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SJSJ_BOTTOM.png"},
			},
		},
		"br08": {
			KorName:       "더 캐시미어",
			EngName:       "the CASHMERE",
			KeyName:       "THECASHMERE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/THECASHMERE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/the+CASHMERE_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/the+CASHMERE_BOTTOM.png"},
			},
		},
		"br19": {
			KorName:       "랑방컬렉션",
			EngName:       "LANVIN COLLECTION",
			KeyName:       "LANVINCOLLECTION",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/LANVINCOLLECTION.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LANVIN+COLLECTION_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LANVIN+COLLECTION_BOTTOM.png"},
			},
		},
		"br31": {
			KorName:       "래트",
			EngName:       "LÄTT",
			KeyName:       "LATT",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/LATT.png",
			Onpopular:     false,
			Description:   "커리어",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LATT_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/LATT_BOTTOM.png"},
			},
		},
		"br43": {
			KorName:       "오브제",
			EngName:       "OBZEE",
			KeyName:       "OBZEE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/OBZEE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/OBZEE_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/OBZEE_BOTTOM.png"},
			},
		},
		"br44": {
			KorName:       "클럽 모나코",
			EngName:       "CLUB MONACO",
			KeyName:       "CLUBMONACO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/CLUBMONACO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CLUB+MONACO_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/CLUB+MONACO_BOTTOM.png"},
			},
		},
		"br45": {
			KorName:       "오즈세컨",
			EngName:       "O'2nd",
			KeyName:       "O2ND",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/O2ND.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/O'2nd_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/O'2nd_BOTTOM.png"},
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
				CrawlUrl:             "http://www.thehandsome.com/ko/c/categoryList?brandCode=ou&categoryCode=ou_" + brandId + "_" + val,
				CrawlModuleName:      modulename,
				IsSalesProducts:      false,
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

	log.Println("Handsome mall brands, categories & sources are added")
}
