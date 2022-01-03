package domain

import (
	"time"

	"github.com/99designs/gqlgen/example/federation/reviews/graph/model"
	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrencyType string

const (
	CurrencyKRW CurrencyType = "KRW"
	CurrencyUSD CurrencyType = "USD"
)

type ProductMetaInfoDAO struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	ProductID    string
	Category     *CategoryDAO
	Brand        *BrandDAO
	OriginalName string
	Price        *PriceDAO
	Images       []string
	ProductUrl   string
	Description  map[string]string
	Sizes        []string
	Colors       []string
	Created      time.Time
	Updated      time.Time
	Source       *CrawlSourceDAO
}

func (pdInfo *ProductMetaInfoDAO) SetBrandAndCategory(brand *BrandDAO, source *CrawlSourceDAO) {
	pdInfo.Brand = brand
	pdInfo.Category = &source.Category
	pdInfo.Source = source
}

func (pdInfo *ProductMetaInfoDAO) SetPrices(origPrice, curPrice int, currencyType CurrencyType) {
	newHistory := []PriceHistoryDAO{
		{
			Date:  time.Now(),
			Price: float32(curPrice),
		},
	}

	if pdInfo.Price != nil {
		newHistory = append(pdInfo.Price.History, newHistory...)
	}

	pdInfo.Price = &PriceDAO{
		OriginalPrice: float32(origPrice),
		CurrencyType:  currencyType,
		CurrentPrice:  float32(curPrice),
		History:       newHistory,
	}
}

func (pdInfo *ProductMetaInfoDAO) SetGeneralInfo(productName, productID, productUrl string, images, sizes, colors []string, description map[string]string) {
	pdInfo.OriginalName = productName
	pdInfo.ProductID = productID
	pdInfo.ProductUrl = productUrl
	pdInfo.Images = images
	pdInfo.Sizes = sizes
	pdInfo.Colors = colors
	pdInfo.Description = description
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
	// 상품의 가격/재고가 업데이트 될때 쓰는 필드
	IsUpdated   bool
	ManualScore int
	AutoScore   int
	TotalScore  int
}

type ProductAlloffCategoryDAO struct {
	First   *AlloffCategoryDAO
	Second  *AlloffCategoryDAO
	Done    bool
	Touched bool
}

type AlloffInstructionDAO struct {
	ProductType string
	Instruction struct {
		Title       string
		Thumbnail   string
		Description []string
		Images      []string
	}
	Faults []struct {
		Image       string
		Description string
	}
	SizeDescription     []string
	DeliveryDescription []string
	CancelDescription   []string
}

type ProductDAO struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ProductInfo      *ProductMetaInfoDAO
	DiscountedPrice  int
	DiscountRate     int
	AlloffCategories *ProductAlloffCategoryDAO
	Soldout          bool
	Removed          bool
	Inventory        []InventoryDAO
	Score            *ProductScoreInfoDAO
	SalesInstruction *AlloffInstructionDAO
	PriceHistory     []PriceHistoryDAO
	Created          time.Time
	Updated          time.Time
}

func (pd *ProductDAO) UpdatePrice(origPrice, alloffPrice float32) {
	pd.DiscountedPrice = int(alloffPrice)
	pd.DiscountRate = utils.CalculateDiscountRate(origPrice, alloffPrice)

	newHistory := []PriceHistoryDAO{
		{
			Date:  time.Now(),
			Price: alloffPrice,
		},
	}

	if pd.PriceHistory != nil {
		newHistory = append(pd.PriceHistory, newHistory...)
	}

	pd.PriceHistory = append(pd.PriceHistory, newHistory...)
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

type PriceDAO struct {
	OriginalPrice float32
	CurrencyType  CurrencyType
	CurrentPrice  float32
	History       []PriceHistoryDAO
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

type ProductGroupDAO struct {
	ID          primitive.ObjectID `bson:"_id, omitempty"`
	Hidden      bool               `json:"hidden"`
	Instruction []string           `json:"instruction"`
	ShortTitle  string             `json:"shorttitle"`
	Title       string             `json:"title"`
	ImgUrl      string             `json:"imgurl"`
	NumAlarms   int
	Products    []*ProductPriorityDAO
	StartTime   time.Time
	FinishTime  time.Time
	Created     time.Time
}

type ProductPriorityDAO struct {
	Priority int
	Product  *ProductDAO
}

type ProductDiffDAO struct {
	ID         string      `bson:"_id,omitempty"`
	OldProduct *ProductDAO `json:"oldproduct"`
	NewProduct *ProductDAO `json:"newproduct"`
	Type       string      `json:"type"`
	IsPushed   bool
}

func (pdDao *ProductDAO) ToEntity() *model.Product {
	var pd *model.Product
	return pd
}
