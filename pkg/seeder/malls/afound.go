package malls

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"log"
	"strings"
	"time"
)

func AddAfound() {
	modulename := "afound"
	baseUrl := "https://www.afound.com/de-de/marken/"

	brands := map[string]domain.BrandDAO{
		"cos": {
			KorName:               "코스",
			EngName:               "COS",
			KeyName:               "COS",
			Description:           "컨템포러리",
			Onpopular:             false,
			MaxDiscountRate:       0,
			IsOpen:                false,
			IsHide:                false,
			InMaintenance:         false,
			NumNewProductsIn3days: 0,
			UseAlloffCategory:     false,
		},
		"arket": {
			KorName:               "아르켓",
			EngName:               "ARKET",
			KeyName:               "ARKET",
			Description:           "컨템포러리",
			Onpopular:             false,
			MaxDiscountRate:       0,
			IsOpen:                false,
			IsHide:                false,
			InMaintenance:         false,
			NumNewProductsIn3days: 0,
			UseAlloffCategory:     false,
		},
		"otherStories": {
			KorName:               "앤아더스토리즈",
			EngName:               "&other stories",
			KeyName:               "OTHERSTORIES",
			Description:           "컨템포러리",
			LogoImgUrl:            "",
			BackImgUrl:            "",
			SizeGuide:             []domain.SizeGuideDAO{},
			Created:               time.Now(),
			Onpopular:             false,
			MaxDiscountRate:       0,
			IsOpen:                false,
			IsHide:                false,
			InMaintenance:         false,
			NumNewProductsIn3days: 0,
			UseAlloffCategory:     false,
		},
	}

	for brandId, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		categories := map[string]string{}
		switch brandId {
		case "cos":
			categories = map[string]string{
				"cos-m-knitwear":           "cos?department=1073743391_on_M_U&p_categoryid=1073743409",
				"cos-m-bottoms":            "cos?department=1073743391_on_M_U&p_categoryid=1073743441",
				"cos-m-functionalclothing": "cos?department=1073743391_on_M_U&p_categoryid=1073743418", // 영어번역시 functional clothing으로 나오는 부분, 아우터만 모여있다.
				"cos-m-shorts":             "cos?department=1073743391_on_M_U&p_categoryid=1073743456",
				"cos-m-shirts":             "cos?department=1073743391_on_M_U&p_categoryid=1073743405",
				"cos-m-sweatshirts":        "cos?department=1073743391_on_M_U&p_categoryid=1073743471",
				"cos-m-shoes":              "cos?department=1073743391_on_M_U&p_categoryid=1073743457",
				"cos-m-tops":               "cos?department=1073743391_on_M_U&p_categoryid=1073743402",
				"cos-m-accessories":        "cos?department=1073743391_on_M_U&p_categoryid=1073743474",
				"cos-m-jeans":              "cos?department=1073743391_on_M_U&p_categoryid=1073743412",
				"cos-m-underwear":          "cos?department=1073743391_on_M_U&p_categoryid=1073743433",

				"cos-f-dresses":            "cos?department=1073743391_on_F_U&p_categoryid=1073743392",
				"cos-f-knitwear":           "cos?department=1073743391_on_F_U&p_categoryid=1073743409",
				"cos-f-bottoms":            "cos?department=1073743391_on_F_U&p_categoryid=1073743441",
				"cos-f-functionalclothing": "cos?department=1073743391_on_F_U&p_categoryid=1073743418",
				"cos-f-shorts":             "cos?department=1073743391_on_F_U&p_categoryid=1073743456",
				"cos-f-shirts":             "cos?department=1073743391_on_F_U&p_categoryid=1073743405",
				"cos-f-sweatshirts":        "cos?department=1073743391_on_F_U&p_categoryid=1073743471",
				"cos-f-shoes":              "cos?department=1073743391_on_F_U&p_categoryid=1073743457",
				"cos-f-tops":               "cos?department=1073743391_on_F_U&p_categoryid=1073743402",
				"cos-f-accessories":        "cos?department=1073743391_on_F_U&p_categoryid=1073743474",
				"cos-f-jeans":              "cos?department=1073743391_on_F_U&p_categoryid=1073743412",
				"cos-f-underwear":          "cos?department=1073743391_on_F_U&p_categoryid=1073743433",
				"cos-f-skirts":             "cos?department=1073743391_on_F_U&p_categoryid=1073743455",
			}
		case "arket":
			categories = map[string]string{
				"arket-m-bottoms":            "arket?department=1073743391_on_M_U&p_categoryid=1073743441",
				"arket-m-functionalclothing": "arket?department=1073743391_on_M_U&p_categoryid=1073743418",
				"arket-m-knitwear":           "arket?department=1073743391_on_M_U&p_categoryid=1073743409",
				"arket-m-shirts":             "arket?department=1073743391_on_M_U&p_categoryid=1073743405",
				"arket-m-shorts":             "arket?department=1073743391_on_M_U&p_categoryid=1073743456",
				"arket-m-tops":               "arket?department=1073743391_on_M_U&p_categoryid=1073743402",
				"arket-m-beachwear":          "arket?department=1073743391_on_M_U&p_categoryid=1073743424",
				"arket-m-sweatshirts":        "arket?department=1073743391_on_M_U&p_categoryid=1073743471",
				"arket-m-shoes":              "arket?department=1073743391_on_M_U&p_categoryid=1073743457",
				"arket-m-accessories":        "arket?department=1073743391_on_M_U&p_categoryid=1073743474",
				"arket-m-sports":             "arket?department=1073743391_on_M_U&p_categoryid=1073743451",
				"arket-m-jeans":              "arket?department=1073743391_on_M_U&p_categoryid=1073743412",

				"arket-f-dresses":            "arket?department=1073743391_on_F_U&p_categoryid=1073743392",
				"arket-f-bottoms":            "arket?department=1073743391_on_F_U&p_categoryid=1073743441",
				"arket-f-functionalclothing": "arket?department=1073743391_on_F_U&p_categoryid=1073743418",
				"arket-f-knitwear":           "arket?department=1073743391_on_F_U&p_categoryid=1073743409",
				"arket-f-shirts":             "arket?department=1073743391_on_F_U&p_categoryid=1073743405",
				"arket-f-shorts":             "arket?department=1073743391_on_F_U&p_categoryid=1073743456",
				"arket-f-tops":               "arket?department=1073743391_on_F_U&p_categoryid=1073743402",
				"arket-f-beachwear":          "arket?department=1073743391_on_F_U&p_categoryid=1073743424",
				"arket-f-sweatshirts":        "arket?department=1073743391_on_F_U&p_categoryid=1073743471",
				"arket-f-shoes":              "arket?department=1073743391_on_F_U&p_categoryid=1073743457",
				"arket-f-accessories":        "arket?department=1073743391_on_F_U&p_categoryid=1073743474",
				"arket-f-sports":             "arket?department=1073743391_on_F_U&p_categoryid=1073743451",
				"arket-f-jeans":              "arket?department=1073743391_on_F_U&p_categoryid=1073743412",
				"arket-f-skirts":             "arket?department=1073743391_on_F_U&p_categoryid=1073743455",

				"arket-kids-sweater":            "arket?department=1073743355&p_categoryid=1073743519",
				"arket-kids-bottoms":            "arket?department=1073743355&p_categoryid=1073743390",
				"arket-kids-sweatshirts":        "arket?department=1073743355&p_categoryid=1073743518",
				"arket-kids-accessories":        "arket?department=1073743355&p_categoryid=1073743373",
				"arket-kids-tops":               "arket?department=1073743355&p_categoryid=1073743367",
				"arket-kids-dresses":            "arket?department=1073743355&p_categoryid=1073743381",
				"arket-kids-jumpsuits":          "arket?department=1073743355&p_categoryid=1073743520",
				"arket-kids-bodysuits":          "arket?department=1073743355&p_categoryid=1073743384",
				"arket-kids-shorts":             "arket?department=1073743355&p_categoryid=1073743386",
				"arket-kids-functionalclothing": "arket?department=1073743355&p_categoryid=1073743361",
			}
		case "otherStories":
			categories = map[string]string{
				"otherStories-f-dresses":            "other-stories?department=1073743391&p_categoryid=1073743392",
				"otherStories-f-bottoms":            "other-stories?department=1073743391&p_categoryid=1073743441",
				"otherStories-f-jeans":              "other-stories?department=1073743391&p_categoryid=1073743412",
				"otherStories-f-knitwear":           "other-stories?department=1073743391&p_categoryid=1073743409",
				"otherStories-f-tops":               "other-stories?department=1073743391&p_categoryid=1073743402",
				"otherStories-f-functionalclothing": "other-stories?department=1073743391&p_categoryid=1073743418",
				"otherStories-f-shoes":              "other-stories?department=1073743391&p_categoryid=1073743457",
				"otherStories-f-skirts":             "other-stories?department=1073743391&p_categoryid=1073743455",
				"otherStories-f-accessories":        "other-stories?department=1073743391&p_categoryid=1073743474",
				"otherStories-f-shirts":             "other-stories?department=1073743391&p_categoryid=1073743405",
				"otherStories-f-beachwear":          "other-stories?department=1073743391&p_categoryid=1073743424",
				"otherStories-f-sports":             "other-stories?department=1073743391&p_categoryid=1073743451",
				"otherStories-f-shorts":             "other-stories?department=1073743391&p_categoryid=1073743456",
			}
		}

		for key, val := range categories {
			productType := []domain.AlloffProductType{}
			switch strings.Split(key, "-")[1] {
			case "m":
				productType = []domain.AlloffProductType{domain.Male}
			case "f":
				productType = []domain.AlloffProductType{domain.Female}
			case "kids":
				productType = []domain.AlloffProductType{domain.Kids}
			}

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
				CrawlUrl:             baseUrl + val,
				CrawlModuleName:      modulename,
				IsSalesProducts:      true,
				IsForeignDelivery:    false,
				DeliveryPrice:        0,
				PriceMarginPolicy:    "AFOUND",
				EarliestDeliveryDays: 7,
				LatestDeliveryDays:   21,
				DeliveryDesc:         nil,
				RefundAvailable:      true,
				ChangeAvailable:      true,
				RefundFee:            100000,
				ChangeFee:            100000,
				ProductType:          productType,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}
	}
	log.Println("Afound categories & sources are added")
}
