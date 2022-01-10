package brand

import (
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func UpdateBrandDiscountRate() {
	offset, limit := 0, 1000
	brandDaos, _, err := ioc.Repo.Brands.List(offset, limit, nil, nil)
	if err != nil {
		log.Println(err)
	}

	warnings := ""
	newProductsMessage := ""
	for _, brand := range brandDaos {
		offset, limit := 0, 1000
		numProducts := 0

		newProductCriterion := time.Now().Add(-2 * time.Hour)
		newlyCrawledProductsIn2Days := 0
		newlyAdded := 0

		maxDiscountRate := 0

		for {
			products, totalCount, err := product.ProductsListing(0, 1, brand.ID.Hex(), "", "", nil)
			if err != nil {
				log.Println(err)
			}

			numProducts += len(products)

			for _, product := range products {
				if maxDiscountRate < product.DiscountRate {
					maxDiscountRate = product.DiscountRate
				}
				if product.Score.IsNewlyCrawled && !product.Soldout {
					newlyCrawledProductsIn2Days += 1
				}
				if product.Created.After(newProductCriterion) {
					newlyAdded += 1
				}
			}

			if totalCount < limit {
				break
			} else {
				offset += limit
			}
		}

		brand.InMaintenance = false
		if numProducts == 0 && brand.IsOpen {
			brand.InMaintenance = true
		}

		brand.MaxDiscountRate = maxDiscountRate
		brand.NumNewProductsIn3days = newlyCrawledProductsIn2Days
		_, err := ioc.Repo.Brands.Upsert(brand)
		if err != nil {
			log.Println(err)
		}

		if brand.IsOpen && brand.MaxDiscountRate == 0 {
			warnings += "<크롤링 확인 필요> " + brand.KorName + " discount rate = 0 \n"
		}

		if newlyAdded > 0 {
			newProductsMessage += brand.KorName + ": " + strconv.Itoa(newlyAdded) + " new products(in this crawling) crawled \n"
		}
	}

	log.Println(warnings, newProductsMessage)
	// config.WriteSlackMessage(warnings)
	// config.WriteSlackMessage(newProductsMessage)
}
