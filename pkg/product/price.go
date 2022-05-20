package product

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func GetCurrentPrice(pdDao *domain.ProductDAO) int {
	alloffPrice := pdDao.DiscountedPrice
	if alloffPrice == 0 {
		alloffPrice = int(pdDao.OriginalPrice)
	} else if pdDao.OriginalPrice == 0 {
		pdDao.OriginalPrice = alloffPrice
	}

	return alloffPrice
}
