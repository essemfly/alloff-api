package crawler

import "github.com/lessbutter/alloff-api/internal/core/domain"

type ProductsAddRequest struct {
	Brand         *domain.BrandDAO
	Source        *domain.CrawlSourceDAO
	ProductID     string
	ProductName   string
	ProductUrl    string
	Images        []string
	Sizes         []string
	Inventories   []domain.InventoryDAO
	Colors        []string
	Description   map[string]string
	OriginalPrice float32
	SalesPrice    float32
	CurrencyType  domain.CurrencyType
}
