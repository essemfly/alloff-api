package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrencyType string

const (
	CurrencyKRW CurrencyType = "KRW"
	CurrencyUSD CurrencyType = "USD"
	CurrencyEUR CurrencyType = "EUR"
)

type DeliveryType string

const (
	Domestic DeliveryType = "DOMESTIC"
	Foreign  DeliveryType = "FOREIGN"
)

type ProductDescriptionDAO struct {
	Images []string
	Texts  []string
	Infos  map[string]string
}

type DeliveryDescriptionDAO struct {
	DeliveryType         DeliveryType
	DeliveryFee          int
	EarliestDeliveryDays int
	LatestDeliveryDays   int
	Texts                []string
}

type CancelDescriptionDAO struct {
	RefundAvailable bool
	ChangeAvailable bool
	ChangeFee       int
	RefundFee       int
}

type AlloffInstructionDAO struct {
	Description         *ProductDescriptionDAO
	DeliveryDescription *DeliveryDescriptionDAO
	CancelDescription   *CancelDescriptionDAO
	Information         map[string]string
}

type ProductAlloffCategoryDAO struct {
	First   *AlloffCategoryDAO
	Second  *AlloffCategoryDAO
	Done    bool
	Touched bool
}

type ProductMetaInfoDAO struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	Brand                *BrandDAO
	Source               *CrawlSourceDAO
	Category             *CategoryDAO
	AlloffCategory       *ProductAlloffCategoryDAO
	OriginalName         string
	ProductID            string
	ProductUrl           string
	Price                *PriceDAO
	Images               []string
	CachedImages         []string
	Colors               []string
	Sizes                []string
	Inventory            []*InventoryDAO
	SalesInstruction     *AlloffInstructionDAO
	IsImageCached        bool
	IsTranslateRequired  bool
	IsCategoryClassified bool
	IsRemoved            bool
	Created              time.Time
	Updated              time.Time
}

func (pdInfo *ProductMetaInfoDAO) SetBrandAndCategory(brand *BrandDAO, source *CrawlSourceDAO) {
	pdInfo.Brand = brand
	pdInfo.Category = &source.Category
	pdInfo.Source = source
}

func (pdInfo *ProductMetaInfoDAO) SetPrices(origPrice, curPrice float32, currencyType CurrencyType) {
	newHistory := []*PriceHistoryDAO{
		{
			Date:  time.Now(),
			Price: curPrice,
		},
	}

	if pdInfo.Price != nil {
		if pdInfo.Price.CurrentPrice != float32(curPrice) {
			newHistory = append(pdInfo.Price.History, newHistory...)
		} else {
			newHistory = pdInfo.Price.History
		}
	}

	if origPrice == 0 {
		origPrice = curPrice
	}

	pdInfo.Price = &PriceDAO{
		OriginalPrice: float32(origPrice),
		CurrencyType:  currencyType,
		CurrentPrice:  float32(curPrice),
		History:       newHistory,
	}
}

func (pdInfo *ProductMetaInfoDAO) SetGeneralInfo(productName, productID, productUrl string, images, sizes, colors []string, information map[string]string) {
	pdInfo.OriginalName = productName
	pdInfo.ProductID = productID
	pdInfo.ProductUrl = productUrl
	pdInfo.Images = images
	pdInfo.Sizes = sizes
	pdInfo.Colors = colors
}

type PriceHistoryDAO struct {
	Date  time.Time
	Price float32
}

type InventoryDAO struct {
	Size     string
	Quantity int
}

type ProductScoreInfoDAO struct {
	// 신상품 위로 올려줄때 쓰는 필드
	IsNewlyCrawled bool
	ManualScore    int
	AutoScore      int
	TotalScore     int
}
