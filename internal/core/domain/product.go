package domain

import (
	"errors"
	"time"

	"github.com/lessbutter/alloff-api/internal/utils"
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

type ProductMetaInfoDAO struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	ProductID      string
	Category       *CategoryDAO
	Brand          *BrandDAO
	OriginalName   string
	Price          *PriceDAO
	Images         []string
	ProductUrl     string
	Sizes          []string
	Colors         []string
	Created        time.Time
	Updated        time.Time
	Source         *CrawlSourceDAO
	Information    map[string]string
	OriginalImages []string
}

func (pdInfo *ProductMetaInfoDAO) SetBrandAndCategory(brand *BrandDAO, source *CrawlSourceDAO) {
	pdInfo.Brand = brand
	pdInfo.Category = &source.Category
	pdInfo.Source = source
}

func (pdInfo *ProductMetaInfoDAO) SetPrices(origPrice, curPrice int, currencyType CurrencyType) {
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
	pdInfo.Information = information
}

type PriceHistoryDAO struct {
	Date  time.Time
	Price int
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

type ProductAlloffCategoryDAO struct {
	First   *AlloffCategoryDAO
	Second  *AlloffCategoryDAO
	Done    bool
	Touched bool
}

type AlloffInstructionDAO struct {
	Description         *ProductDescriptionDAO
	DeliveryDescription *DeliveryDescriptionDAO
	CancelDescription   *CancelDescriptionDAO
}

type ProductDescriptionDAO struct {
	Images []string
	Texts  []string
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

type ProductDAO struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	ProductInfo         *ProductMetaInfoDAO
	Images              []string
	ProductGroupId      string
	AlloffName          string
	OriginalPrice       int
	DiscountedPrice     int
	DiscountRate        int
	SpecialPrice        int
	AlloffCategories    *ProductAlloffCategoryDAO
	Soldout             bool
	Removed             bool
	Inventory           []InventoryDAO
	Score               *ProductScoreInfoDAO
	SalesInstruction    *AlloffInstructionDAO
	PriceHistory        []PriceHistoryDAO
	IsUpdated           bool
	Created             time.Time
	Updated             time.Time
	IsImageCached       bool
	IsTranslateRequired bool
}

func (pd *ProductDAO) UpdatePrice(origPrice, discountedPrice int) bool {
	lastPrice := pd.DiscountedPrice
	pd.OriginalPrice = origPrice
	pd.DiscountedPrice = discountedPrice
	if origPrice == 0 {
		pd.OriginalPrice = discountedPrice
	}
	pd.DiscountRate = utils.CalculateDiscountRate(origPrice, discountedPrice)

	newHistory := []PriceHistoryDAO{
		{
			Date:  time.Now(),
			Price: discountedPrice,
		},
	}

	if pd.PriceHistory != nil {
		if lastPrice != pd.DiscountedPrice {
			pd.PriceHistory = append(pd.PriceHistory, newHistory...)
		}
	} else {
		pd.PriceHistory = newHistory
	}

	if int(discountedPrice) < lastPrice {
		pd.IsUpdated = true
		return true
	}
	return false
}

func (pd *ProductDAO) UpdateScore(newScore *ProductScoreInfoDAO) {
	pd.Score = newScore
}

func (pd *ProductDAO) UpdateInventory(newInven []InventoryDAO) {
	pd.Inventory = newInven

	isSoldout := true
	for _, inv := range newInven {
		if inv.Quantity > 0 {
			isSoldout = false
			break
		}
	}

	pd.Soldout = isSoldout
}

func (pd *ProductDAO) UpdateAlloffCategory(cat *ProductAlloffCategoryDAO) {
	pd.AlloffCategories = cat
}

func (pd *ProductDAO) UpdateInstruction(instruction *AlloffInstructionDAO) {
	pd.SalesInstruction = instruction
}

func (pd *ProductDAO) Release(size string, quantity int) error {
	for idx, option := range pd.Inventory {
		if option.Size == size {
			if option.Quantity < quantity {
				return errors.New("insufficient product quantity")
			}
			pd.Inventory[idx].Quantity -= quantity

			return nil
		}
	}
	return errors.New("no matched product size option")
}

func (pd *ProductDAO) Revert(size string, quantity int) error {
	for idx, option := range pd.Inventory {
		if option.Size == size {
			if option.Quantity == 0 {
				pd.Soldout = false
			}
			pd.Inventory[idx].Quantity += quantity

			return nil
		}
	}
	return errors.New("no matched product size option")
}

type PriceDAO struct {
	OriginalPrice float32
	CurrencyType  CurrencyType
	CurrentPrice  float32
	History       []*PriceHistoryDAO
}

type LikeProductDAO struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Productid  string
	OldProduct *ProductDAO `bson:"product"`
	Userid     string
	IsPushed   bool
	LastPrice  int
	Removed    bool
	Created    time.Time
	Updated    time.Time
}

type ProductDiffDAO struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	OldPrice   int                `json:"oldprice"`
	NewProduct *ProductDAO        `json:"newproduct"`
	Type       string             `json:"type"`
	IsPushed   bool
}
