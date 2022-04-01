package product

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type ProductCrawlingAddRequest struct {
	Brand               *domain.BrandDAO
	Source              *domain.CrawlSourceDAO
	ProductID           string
	ProductName         string
	ProductUrl          string
	Images              []string
	Sizes               []string
	Inventories         []domain.InventoryDAO
	Colors              []string
	Description         map[string]string
	OriginalPrice       float32
	SalesPrice          float32
	CurrencyType        domain.CurrencyType
	IsTranslateRequired bool
}

type ProductManualAddRequest struct {
	AlloffName           string
	IsForeignDelivery    bool
	ProductID            string
	ProductUrl           string
	OriginalPrice        int
	DiscountedPrice      int
	SpecialPrice         int
	BrandKeyName         string
	Inventory            []domain.InventoryDAO
	Description          []string
	EarliestDeliveryDays int
	LatestDeliveryDays   int
	IsRefundPossible     bool
	RefundFee            int
	Images               []string
	DescriptionImages    []string
	ModuleName           string
	AlloffCategoryID     string
	DescriptionInfos     map[string]string
	ProductInfos         map[string]string
}
