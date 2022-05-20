package product

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type AddRequest struct {
	AlloffName           string
	ProductID            string
	ProductUrl           string
	OriginalPrice        int
	DiscountedPrice      int
	SpecialPrice         int
	CurrencyType         domain.CurrencyType
	Brand                *domain.BrandDAO
	Source               *domain.CrawlSourceDAO
	AlloffCategory       *domain.AlloffCategoryDAO
	Images               []string
	ThumbnailImage       string
	Information          map[string]string
	Description          []string
	DescriptionImages    []string
	DescriptionInfos     map[string]string
	IsForeignDelivery    bool
	EarliestDeliveryDays int
	LatestDeliveryDays   int
	IsRefundPossible     bool
	RefundFee            int
	Inventory            []domain.InventoryDAO
	ModuleName           string

	IsSpecial bool
}

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
	ThumbnailImage       string
	IsSpecial            bool
}
