package scripts

import "github.com/lessbutter/alloff-api/pkg/brand"

func UpdateBrands() {
	brand.UpdateBrandCategory()
	brand.UpdateBrandDiscountRate()
}
