package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapBestProducts(bestProduct *domain.BestProductDAO, brief bool, offset, limit int) []*model.Product {
	pds := []*model.Product{}
	if brief {
		for idx, product := range bestProduct.Products {
			if brief && idx > 9 {
				break
			}
			pds = append(pds, MapProductDaoToProduct(product))
		}
		return pds
	}

	for idx, product := range bestProduct.Products {
		if idx < offset {
			continue
		}
		if idx >= (offset + limit) {
			break
		}
		pds = append(pds, MapProductDaoToProduct(product))
	}

	return pds
}
