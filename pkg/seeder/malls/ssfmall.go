package malls

import (
	"log"
	"strings"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddSSFMall() {
	modulename := "ssfmall"
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
			KorName:       "플랜 씨",
			EngName:       "Plan C",
			KeyName:       "PLANC",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/PLANC.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/PLAN+C_TOP.png"},
			},
		},
		"BDMA07A25/OA": {
			KorName:       "오이아우어",
			EngName:       "OIAUER",
			KeyName:       "OIAUER",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/OIAUER.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/OIAUER_TOP:BOTTOM.png"},
			},
		},
		"BDMA09D06/BQAAW": {
			KorName:       "아스페시",
			EngName:       "Aspesi",
			KeyName:       "ASPESI",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ASPESI.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Aspesi_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Aspesi_BOTTOM.png"},
			},
		},
		"BDMA07A32/ECBSV": {
			KorName:       "슬로웨어",
			EngName:       "SLOWEAR",
			KeyName:       "SLOWEAR",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SLOWEAR.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SLOWEAR_TOP:BOTTOM.png"},
			},
		},
		"BDMA01A02/BPBBF": {
			KorName:       "빈폴 레이디스",
			EngName:       "Beanpole Ladies",
			KeyName:       "BEANPOLEL",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BEANPOLEL.png",
			Onpopular:     true,
			Description:   "트레디셔널 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Beanpole+Ladies_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Beanpole+Ladies_BOTTOM.png"},
			},
		},
		"BDMA09D08/BQWPB": {
			KorName:       "비이커",
			EngName:       "BEAKER",
			KeyName:       "BEAKER",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BEAKER.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/BEAKER_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/BEAKER_BOTTOM.png"},
			},
		},
		"BDMA09A13/BQWBV": {
			KorName:       "바레나 베네치아",
			EngName:       "Barena Venezia",
			KeyName:       "BARENAVENEZIA",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BARENAVENEZIA.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Barena+Venezia_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Barena+Venezia_BOTTOM.png"},
			},
		},
		"BDMA07A21/ECBGF": {
			KorName:       "멜리사",
			EngName:       "Melissa",
			KeyName:       "MELISSA",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MELISSA.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Melissa_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Melissa_BOTTOM.png"},
			},
		},
		"BDMA07A22/BQMKT": {
			KorName:       "메종키츠네",
			EngName:       "Maison Kitsune",
			KeyName:       "MAISONKITSUNE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MAISONKITSUNE.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maison+Kitsune_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maison+Kitsune_BOTTOM.png"},
			},
		},
		"BDMA19Y39/RTMRC": {
			KorName:       "마리끌레르",
			EngName:       "marie claire",
			KeyName:       "MARIECLAIRE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MARIECLAIRE.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		"11": {
			KorName:       "르메르",
			EngName:       "Lemaire",
			KeyName:       "LEMAIRE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/LEMAIRE.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		"BDMA07A40/RAB": {
			KorName:       "랙앤본",
			EngName:       "rag & bone",
			KeyName:       "RAGANDBONE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/RAGANDBONE.png",
			Onpopular:     true,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/rag+%26+bone_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/rag+%26+bone_BOTTOM.png"},
			},
		},
		"BDMA07A30/ECBTM": {
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
		"BDMA09F02/BQMDT": {
			KorName:       "단톤",
			EngName:       "Danton",
			KeyName:       "DANTON",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DANTON.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Danton_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Danton_BOTTOM.png"},
			},
		},
		"BDMA07A02/WMBKF": {
			KorName:       "구호",
			EngName:       "KUHO",
			KeyName:       "KUHO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/KUHO.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			Modulename:    "ssfmall",
			InMaintenance: false,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/KUHO_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/KUHO_BOTTOM.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/KUHO_SHOES.png"},
			},
		},
		"KUHOKUHO": {
			KorName:       "준지",
			EngName:       "Juun.J",
			KeyName:       "JUUNJ",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JUUNJ.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "ssfmall",
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
					MainCategoryKey:      updatedCat.CatIdentifier,
					Category:             *updatedCat,
					CrawlUrl:             buildCrawlUrl(shopId, updatedCat.CatIdentifier),
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
	log.Println("SSF mall brands, categories & sources are added")
}

func buildCrawlUrl(shopid, category string) string {
	shopIdentifiers := strings.Split(shopid, "/")
	return "https://www.ssfshop.com/selectProductList?dspCtgryNo=" + category + "&brandShopNo=" + shopIdentifiers[0] + "&brndShopId=" + shopIdentifiers[1]
}
