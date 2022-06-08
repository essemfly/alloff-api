package domain

import (
	"errors"
	"time"

	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrencyType string

const (
	CurrencyKRW   CurrencyType = "KRW"
	CurrencyUSD   CurrencyType = "USD"
	CurrencyEUR   CurrencyType = "EUR"
	CurrencyPOUND CurrencyType = "POUND"
)

type DeliveryType string

const (
	Domestic DeliveryType = "DOMESTIC"
	Foreign  DeliveryType = "FOREIGN"
)

type ProductDescriptionDAO struct {
	Images   []string
	Texts    []string
	Infos    map[string]string
	RawInfos map[string]string
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
	RawInformation      map[string]string
}

type TaggingResultDAO struct {
	Item         EstimateModelType   `json:"item"`
	Colors       []EstimateModelType `json:"colors"`
	ColorDetails []EstimateModelType `json:"colorDetails"`
	Prints       []EstimateModelType `json:"prints"`
	Looks        []EstimateModelType `json:"looks"`
	Textures     []EstimateModelType `json:"textures"`
	Details      []EstimateModelType `json:"details"`
	Length       EstimateModelType   `json:"length"`
	SleeveLength EstimateModelType   `json:"sleeveLength"`
	NeckLine     EstimateModelType   `json:"neckLine"`
	Fit          EstimateModelType   `json:"fit"`
	Shape        EstimateModelType   `json:"shape"`
}

type EstimateModelType struct {
	Id         string
	Name       string
	Confidence float64
}

type InventoryDAO struct {
	Size        string
	Quantity    int
	AlloffSizes []*AlloffSizeDAO
}

type ProductScoreInfoDAO struct {
	// 신상품 위로 올려줄때 쓰는 필드
	IsNewlyCrawled bool
	ManualScore    int
	AutoScore      int
	TotalScore     int
}

type ProductAlloffCategoryDAO struct {
	TaggingResults *TaggingResultDAO
	First          *AlloffCategoryDAO
	Second         *AlloffCategoryDAO
	Done           bool
	Touched        bool
}

type PriceDAO struct {
	CurrencyType  CurrencyType
	OriginalPrice int
	CurrentPrice  int
	DiscountRate  int
	History       []*PriceHistoryDAO
}

type PriceHistoryDAO struct {
	Date  time.Time
	Price int
}

type AlloffProductType string

const (
	Male   = AlloffProductType("MALE")
	Female = AlloffProductType("FEMALE")
	Boy    = AlloffProductType("BOY")
	Girl   = AlloffProductType("GIRL")
	Kids   = AlloffProductType("KIDS")
)

func ListProductTypes() []AlloffProductType {
	return []AlloffProductType{
		Male,
		Female,
		Boy,
		Girl,
		Kids,
	}
}

type ProductMetaInfoDAO struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	Brand                *BrandDAO
	Source               *CrawlSourceDAO
	Category             *CategoryDAO
	AlloffCategory       *ProductAlloffCategoryDAO
	ProductType          []AlloffProductType
	OriginalName         string
	AlloffName           string
	ProductID            string
	ProductUrl           string
	Price                *PriceDAO
	Images               []string
	CachedImages         []string
	ThumbnailImage       string
	Colors               []string
	Sizes                []string
	Inventory            []*InventoryDAO
	SalesInstruction     *AlloffInstructionDAO
	IsImageCached        bool
	IsInventoryMapped    bool
	IsTranslateRequired  bool
	IsCategoryClassified bool
	IsSoldout            bool
	IsRemoved            bool
	Created              time.Time
	Updated              time.Time
	LastCrawled          time.Time
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
		if pdInfo.Price.CurrentPrice != curPrice {
			newHistory = append(pdInfo.Price.History, newHistory...)
		} else {
			newHistory = pdInfo.Price.History
		}
	}

	if origPrice == 0 {
		origPrice = curPrice
	}

	pdInfo.Price = &PriceDAO{
		OriginalPrice: origPrice,
		CurrencyType:  currencyType,
		CurrentPrice:  curPrice,
		DiscountRate:  utils.CalculateDiscountRate(int(origPrice), int(curPrice)),
		History:       newHistory,
	}
}

func (pdInfo *ProductMetaInfoDAO) SetGeneralInfo(productTypes []AlloffProductType, productName, productID, productUrl string, images, sizes, colors []string, information map[string]string) {
	pdInfo.OriginalName = productName
	pdInfo.AlloffName = productName
	pdInfo.ProductID = productID
	pdInfo.ProductUrl = productUrl
	pdInfo.Images = images
	pdInfo.Sizes = sizes
	pdInfo.Colors = colors
	pdInfo.ProductType = productTypes
}

func (pdInfo *ProductMetaInfoDAO) SetDesc(descImages, texts []string, infos map[string]string) {
	pdInfo.SalesInstruction.Description = &ProductDescriptionDAO{
		Images: descImages,
		Texts:  texts,
		Infos:  infos,
	}
}

func (pdInfo *ProductMetaInfoDAO) SetInformation(information map[string]string) {
	pdInfo.SalesInstruction.Information = information
}

func (pdInfo *ProductMetaInfoDAO) SetDeliveryDesc(isForeignDelivery bool, deliveryPrice, earliestDeliveryDays, latestDeliveryDays int) {
	deliveryType := Domestic
	deliveryTexts := []string{
		"도착 예정일은 택배사의 사정이나 주문량에 따라 변동될 수 있습니다.",
		"브랜드 및 제품에 따라 입점 업체(브랜드) 배송과 올오프 자체 배송으로 나뉩니다.",
	}

	if isForeignDelivery {
		deliveryType = Foreign
		deliveryTexts = []string{
			"도착 예정일은 현지 택배사의 사정이나 통관 과정에서 변동될 수 있습니다.",
			"배송기간에 현지 및 한국의 공휴일, 연말이 포함된 경우 배송이 지연될 수 있습니다.",
		}
	}
	pdInfo.SalesInstruction.DeliveryDescription = &DeliveryDescriptionDAO{
		DeliveryType:         deliveryType,
		DeliveryFee:          deliveryPrice,
		Texts:                deliveryTexts,
		EarliestDeliveryDays: earliestDeliveryDays,
		LatestDeliveryDays:   latestDeliveryDays,
	}
}

func (pdInfo *ProductMetaInfoDAO) SetCancelDesc(isRefundPossible bool, refundFee int) {
	pdInfo.SalesInstruction.CancelDescription = &CancelDescriptionDAO{
		RefundAvailable: isRefundPossible,
		ChangeAvailable: isRefundPossible,
		ChangeFee:       refundFee,
		RefundFee:       refundFee,
	}
}

func (pdInfo *ProductMetaInfoDAO) SetAlloffCategory(cat *ProductAlloffCategoryDAO) {
	pdInfo.AlloffCategory = cat
	pdInfo.IsCategoryClassified = true
}

func (pdInfo *ProductMetaInfoDAO) SetInventory(inventories []*InventoryDAO) {
	pdInfo.Inventory = inventories
	pdInfo.IsInventoryMapped = true
}

func (pdInfo *ProductMetaInfoDAO) Release(size string, quantity int) error {
	for idx, option := range pdInfo.Inventory {
		if option.Size == size {
			if option.Quantity < quantity {
				return errors.New("insufficient product quantity")
			}
			pdInfo.Inventory[idx].Quantity -= quantity
			if pdInfo.Inventory[idx].Quantity == 0 {
				pdInfo.IsSoldout = true
			}

			return nil
		}
	}
	return errors.New("no matched product size option")
}

func (pdInfo *ProductMetaInfoDAO) Revert(size string, quantity int) error {
	for idx, option := range pdInfo.Inventory {
		if option.Size == size {
			if option.Quantity == 0 {
				pdInfo.IsSoldout = false
			}
			pdInfo.Inventory[idx].Quantity += quantity

			return nil
		}
	}
	return errors.New("no matched product size option")
}

func (pdInfo *ProductMetaInfoDAO) CheckSoldout() {
	for _, inv := range pdInfo.Inventory {
		if inv.Quantity > 0 {
			pdInfo.IsSoldout = false
		}
	}
}
