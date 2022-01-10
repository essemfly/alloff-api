package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddHfashion() {
	modulename := "hfashion"
	categories := map[string]string{
		"티셔츠":     "HFMA02A01",
		"셔츠/블라우스": "HFMA02A02",
		"니트/스웨터":  "HFMA02A03",
		"아우터":     "HFMA02A04",
		"팬츠":      "HFMA02A05",
		"데님":      "HFMA02A06",
		"수트정장":    "HFMA02A07",
		"언더웨어":    "HFMA02A08",
		"원피스":     "HFMA02A09",
		"스커트":     "HFMA02A10",
		"스포츠웨어":   "HFMA02A16",
	}

	brands := map[string]domain.BrandDAO{
		"T2HGR": {
			KorName:       "타미 힐피거",
			EngName:       "TOMMY HILFIGER",
			KeyName:       "TOMMYHILFIGERW",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/TOMMYHILFIGERW.png",
			Onpopular:     false,
			Description:   "트래디셔널 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/TOMMY+HILFIGER_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/TOMMY+HILFIGER_BOTTOM.png"},
			},
		},
		"VWHGR": {
			KorName:       "캘빈 클라인",
			EngName:       "Calvin Klein",
			KeyName:       "CALVINKLEINW",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/CALVINKLEINW.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Calvin+Klein_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Calvin+Klein_BOTTOM.png"},
			},
		},
		"PWHGR": {
			KorName:       "에스제이와이피",
			EngName:       "SJYP",
			KeyName:       "SJYP",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SJYP.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SJYP_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SJYP_BOTTOM.png"},
				{Label: "데님", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SJYP_DENIM.png"},
				{Label: "유니섹스 상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SJYP_UNISEX+TOP.png"},
			},
		},
		"DWHGR": {
			KorName:       "디케이앤와이",
			EngName:       "DKNY",
			KeyName:       "DKNY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DKNY.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        false,
			Modulename:    "handsome",
			InMaintenance: true,
			SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DKNY_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DKNY_BOTTOM.png"},
			},
		},
	}

	for shopId, brand := range brands {
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
				BrandKeyname:      upsertedBrand.KeyName,
				BrandIdentifier:   shopId,
				MainCategoryKey:   updatedCat.CatIdentifier,
				Category:          *updatedCat,
				CrawlUrl:          buildHfashionCrawlUrl(shopId, updatedCat.CatIdentifier),
				CrawlModuleName:   modulename,
				IsSalesProducts:   false,
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

	log.Println("HFashion malls, categories & sources are added")
}

func buildHfashionCrawlUrl(shopId, category string) string {
	return "https://www.hfashionmall.com/display/category/list?otltYn=Y&brndId=" + shopId + "&dspCtgryNo=" + category
}
