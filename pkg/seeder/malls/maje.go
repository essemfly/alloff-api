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
		"니트":   "Pullover%20%26%20Strickjacken",
		"자켓":   "Blazers%20%26%20Jacken",
		"옷":    "Kleider",
		"상의":   "Tops%20%26%20Hemden",
		"가방":   "Taschen",
		"스커트":  "Röcke%20%26%20Shorts",
		"악세서리": "Schmuck",
		"바지":   "Hosen%20%26%20Jeans",
		"구두":   "Schuhe",
		"벨트":   "Gürtel",
		"점프슈트": "Jumpsuit",
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
