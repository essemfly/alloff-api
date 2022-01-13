package domain

import (
	"errors"
	"time"

	"github.com/lessbutter/alloff-api/api/server/model"
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
	Sizes        []string
	Colors       []string
	Created      time.Time
	Updated      time.Time
	Source       *CrawlSourceDAO
	Information  map[string]string
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

func (productDesc *ProductDescriptionDAO) ToDTO() *model.ProductDescription {
	return &model.ProductDescription{
		Images: productDesc.Images,
		Texts:  productDesc.Texts,
	}
}

type DeliveryDescriptionDAO struct {
	DeliveryFee          int
	EarliestDeliveryDays int
	LatestDeliveryDays   int
	Texts                []string
}

func (deliveryDesc *DeliveryDescriptionDAO) ToDTO() *model.DeliveryDescription {
	return &model.DeliveryDescription{
		DeliveryFee:         deliveryDesc.DeliveryFee,
		EarlistDeliveryDays: deliveryDesc.EarliestDeliveryDays,
		LatestDeliveryDays:  deliveryDesc.LatestDeliveryDays,
		Texts:               deliveryDesc.Texts,
	}
}

type CancelDescriptionDAO struct {
	RefundAvailable bool
	ChangeAvailable bool
	Texts           []string
}

func (cancelDesc *CancelDescriptionDAO) ToDTO() *model.CancelDescription {
	return &model.CancelDescription{
		RefundAvailable: cancelDesc.RefundAvailable,
		ChangeAvailable: cancelDesc.ChangeAvailable,
		Texts:           cancelDesc.Texts,
	}
}

type ProductDAO struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	ProductInfo      *ProductMetaInfoDAO
	ProductGroupId   *primitive.ObjectID
	AlloffName       string
	DiscountedPrice  int
	DiscountRate     int
	AlloffCategories *ProductAlloffCategoryDAO
	Soldout          bool
	Removed          bool
	Inventory        []InventoryDAO
	Score            *ProductScoreInfoDAO
	SalesInstruction *AlloffInstructionDAO
	PriceHistory     []PriceHistoryDAO
	IsUpdated        bool
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

	pd.PriceHistory = newHistory

	if alloffPrice < origPrice {
		pd.IsUpdated = true
	}
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
	for _, option := range pd.Inventory {
		if option.Size == size {
			if option.Quantity < quantity {
				return errors.New("insufficient product quantity")
			}
			option.Quantity -= quantity

			if option.Quantity == 0 {
				pd.Soldout = true
			}
			return nil
		}
	}
	return errors.New("no matched product size option")
}

func (pd *ProductDAO) Revert(size string, quantity int) error {
	for _, option := range pd.Inventory {
		if option.Size == size {
			if option.Quantity == 0 {
				pd.Soldout = false
			}
			option.Quantity += quantity
			return nil
		}
	}
	return errors.New("no matched product size option")
}

func (pdDao *ProductDAO) ToDTO() *model.Product {
	inventories := []*model.Inventory{}

	for _, inv := range pdDao.Inventory {
		inventories = append(inventories, &model.Inventory{
			Quantity: inv.Quantity,
			Size:     inv.Size,
		})
	}

	var information []*model.KeyValueInfo
	for k, v := range pdDao.ProductInfo.Information {
		var newInfo model.KeyValueInfo
		newInfo.Key = k
		newInfo.Value = v
		information = append(information, &newInfo)
	}

	return &model.Product{
		ID:                  pdDao.ID.Hex(),
		Category:            pdDao.ProductInfo.Category.ToDTO(),
		Brand:               pdDao.ProductInfo.Brand.ToDTO(false),
		Name:                pdDao.AlloffName,
		OriginalPrice:       int(pdDao.ProductInfo.Price.OriginalPrice),
		Soldout:             pdDao.Soldout,
		Images:              pdDao.ProductInfo.Images,
		DiscountedPrice:     &pdDao.DiscountedPrice,
		DiscountRate:        &pdDao.DiscountRate,
		ProductURL:          pdDao.ProductInfo.ProductUrl,
		Inventory:           inventories,
		IsUpdated:           pdDao.IsUpdated,
		IsNewProduct:        pdDao.Score.IsNewlyCrawled,
		Removed:             pdDao.Removed,
		Information:         information,
		Description:         pdDao.SalesInstruction.Description.ToDTO(),
		DeliveryDescription: pdDao.SalesInstruction.DeliveryDescription.ToDTO(),
		CancelDescription:   pdDao.SalesInstruction.CancelDescription.ToDTO(),
	}
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
