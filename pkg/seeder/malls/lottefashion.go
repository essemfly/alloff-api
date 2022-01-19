package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddLotteFashion() {
	modulename := "lottefashion"
	// 여성의류 Category Number
	mainCategoryNumber := "2080451"
	crawlUrl := "https://mapi.lfmall.co.kr/api/search/v2/outlet/categories"

	categories := map[string]string{
		"점퍼":       "2080521",
		"니트/스웨터":   "2080457",
		"자켓":       "2080515",
		"베스트":      "2080465",
		"가디건":      "2080452",
		"원피스":      "2080500",
		"팬츠":       "2080553",
		"티셔츠":      "2080536",
		"셔츠":       "2080474",
		"블라우스":     "2080468",
		"스커트":      "2080484",
		"언더/라운지웨어": "2080495",
		"여성비치웨어":   "2092139",
		"코트":       "2080527",
	}

	brands := map[string]domain.BrandDAO{
		"1002": {
			KorName:       "헤지스 레이디스",
			EngName:       "HAZZYS LADIES",
			KeyName:       "HAZZYSL",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/HAZZYSL.png",
			Onpopular:     true,
			Description:   "트레디셔널 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/HAZZYS+LADIES_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/HAZZYS+LADIES_BOTTOM.png"},
			},
		},
		"1017": {
			KorName:       "질스튜어트 뉴욕",
			EngName:       "JILLSTUART NEW YORK",
			KeyName:       "JILLSTUARTNY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JILLSTUARTNY.png",
			Onpopular:     true,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/JILLSTUART+NEW+YORK_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/JILLSTUART+NEW+YORK_BOTTOM.png"},
			},
		},
		"1664": {
			KorName:       "조셉",
			EngName:       "JOSEPH",
			KeyName:       "JOSEPH",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JOSEPH.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/JOSEPH_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/JOSEPH_BOTTOM.png"},
			},
		},
		"3132": {
			KorName:       "이자벨 마랑",
			EngName:       "ISABEL MARANT",
			KeyName:       "ISABELMARANT",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ISABELMARANT.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ISABEL+MARANT_TOP:BOTTOM.png"},
			},
		},
		"2637": {
			KorName:       "오피신 제네랄",
			EngName:       "Officine Generale",
			KeyName:       "OFFICINEGENERALE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/OFFICINEGENERALE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Officine+Generale_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Officine+Generale_BOTTOM.png"},
			},
		},
		"1011": {
			KorName:       "빈스",
			EngName:       "VINCE",
			KeyName:       "VINCE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/VINCE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VINCE_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VINCE_BOTTOM.png"},
			},
		},
		"1009": {
			KorName:       "바네사브루노",
			EngName:       "VANESSABRUNO",
			KeyName:       "VANESSABRUNO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/VANESSABRUNO.png",
			Onpopular:     true,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VANESSABRUNO_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VANESSABRUNO_BOTTOM.png"},
			},
		},
		"1001": {
			KorName:       "닥스 레이디스",
			EngName:       "DAKS LADIES",
			KeyName:       "DAKSL",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DAKSL.png",
			Onpopular:     false,
			Description:   "커리어",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DAKS+LADIES_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DAKS+LADIES_BOTTOM.png"},
			},
		},
		"2110": {
			KorName:       "넘버21",
			EngName:       "N21",
			KeyName:       "N21",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/N21.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/N21_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/N21_BOTTOM.png"},
			},
		},
		"13977": {
			KorName:       "제이에스엔와이",
			EngName:       "JSNY",
			KeyName:       "JSNY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JSNY.png",
			Onpopular:     false,
			Description:   "트레디셔널 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/JSNY_TOP:BOTTOM.png"},
			},
		},
		"11324": {
			KorName:       "쥬시쥬디",
			EngName:       "JUCY JUDY",
			KeyName:       "JUCYJUDY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JUCYJUDY.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "lottefashion",
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
	}
	for shopId, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		if brand.IsOpen {
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
					BrandKeyname:         upsertedBrand.KeyName,
					BrandIdentifier:      shopId,
					MainCategoryKey:      mainCategoryNumber,
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
				}

				_, err = ioc.Repo.CrawlSources.Upsert(&source)
				if err != nil {
					log.Println(err)
				}
			}
		}

	}
	log.Println("Lotte fashion brands, categories & sources are added")
}
