package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapBestProducts(bestProduct *domain.BestProductDAO, brief bool) []*model.Product {
	pds := []*model.Product{}
	for idx, product := range bestProduct.Products {
		if brief && idx > 9 {
			break
		}
		pds = append(pds, MapProductDaoToProduct(product))
	}
	return pds
}
