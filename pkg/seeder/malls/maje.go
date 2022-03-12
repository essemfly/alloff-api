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
	}

	crawlUrl := "https://de.maje.com/de/%s/kategorien/alles-einsehen/?prefn1=smcp_subFamily&prefv1=%s&format=ajax&sz=1000000"
	categories := map[string]string{
		"SWEATERS_AND_CARDIGANS@%s": "Pullover%20%26%20Strickjacken",
		"BLAZERS_AND_JACKETS@%s":    "Blazers%20%26%20Jacken",
		"CLOTHES@%s":                "Kleider",
		"TOPS_AND_SHIRTS@%s":        "Tops%20%26%20Hemden",
		"BAGS@%s":                   "Taschen",
		"SKIRTS_AND_SHORTS@%s":      "Röcke%20%26%20Shorts",
		"JEWELRY@%s":                "Schmuck",
		"TROUSERS_AND_JEANS@%s":     "Hosen%20%26%20Jeans",
		"SHOES@%s":                  "Schuhe",
		"BELT@%s":                   "Gürtel",
		"JUMPSUIT@%s":               "Jumpsuit",
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
			IsOpen:        false,
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

		for _categoryKey, categoryValue := range categories {
			for _, exhibition := range majeExhibitions {
				categoryKey := fmt.Sprintf(_categoryKey, exhibition)
				keyname := brand.KeyName + "-" + categoryKey
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
					IsSalesProducts:      true,
					IsForeignDelivery:    true,
					PriceMarginPolicy:    "MAJE",
					DeliveryPrice:        0,
					EarliestDeliveryDays: 14,
					LatestDeliveryDays:   21,
					DeliveryDesc:         nil,
					RefundAvailable:      false,
					ChangeAvailable:      false,
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
	log.Println("Maje categories & sources are added")
}
