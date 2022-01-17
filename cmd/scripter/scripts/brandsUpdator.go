package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/pkg/brand"
)

func UpdateBrands() {
	log.Println("Update Brands Category")
	brand.UpdateBrandCategory()
	log.Println("Update Brands Discount Rate")
	brand.UpdateBrandDiscountRate()
}
