package malls

import (
	"fmt"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddMaje() {
	modulename := "maje"
	majeExhibitions := []string{
		"outlet",
		"last-chance",
		"sale",
		"pre-sale",
	}

	crawlUrl := "https://de.maje.com/de/%s/kategorien/%s/?&sz=10000&format=ajax"

	categories := map[string]string{
		"원피스":    "kleider",
		"치마/반바지": "rocke-und-shorts",
		"상의":     "tops-und-t-shirts",
		"니트":     "pullover-und-strickjacken",
		"코트":     "mantel",
		"자켓":     "jacken-und-blazers",
		"바지":     "hosen-und-jeans",
		"점프수트":   "jumpshort-und-jumpsuit",
		"가방":     "taschen",
		"악세서리":   "accessoires",
	}

	brands := map[string]domain.BrandDAO{
		"Maje": {
			KorName:       "마쥬",
			EngName:       "Maje",
			KeyName:       "MAJE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MAJE.png",
			Onpopular:     true,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maje_TOP:BOTTOM.png"},
				{Label: "데님/진", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maje_DENIM:JEAN.png"},
				{Label: "슈즈", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maje_SHOES.png"},
			},
		},
	}

	for brandId, brand := range brands {
		upsertedBrand, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}

		for categoryKey, categoryValue := range categories {
			for _, exhibition := range majeExhibitions {
				keyname := brand.KeyName + "-" + categoryKey + "-" + exhibition
				category := domain.CategoryDAO{
					Name:          categoryKey,
					KeyName:       keyname,
					CatIdentifier: keyname,
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
					CrawlUrl:             fmt.Sprintf(crawlUrl, exhibition, categoryValue),
					CrawlModuleName:      modulename,
					IsSalesProducts:      false,
					IsForeignDelivery:    true,
					PriceMarginPolicy:    "MAJE",
					DeliveryPrice:        0,
					EarliestDeliveryDays: 10,
					LatestDeliveryDays:   14,
					DeliveryDesc:         nil,
					RefundAvailable:      true,
					ChangeAvailable:      true,
					RefundFee:            100000,
					ChangeFee:            100000,
				}

				_, err = ioc.Repo.CrawlSources.Upsert(&source)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
	log.Println("Maje categories & sources are added")
}
