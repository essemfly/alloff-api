package product

import "github.com/lessbutter/alloff-api/internal/core/domain"

func GetProductPrice(pd *domain.ProductDAO) float32 {
	return pd.ProductInfo.Price.CurrentPrice
}
