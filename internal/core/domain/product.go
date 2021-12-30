package domain

import (
	"time"

	"github.com/99designs/gqlgen/example/federation/reviews/graph/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CurrencyType string

const (
	CurrencyKRW CurrencyType = "KRW"
	CurrencyUSD CurrencyType = "USD"
)

type ProductMetaInfoDAO struct {
	ProductID      string
	Category       *CategoryDAO
	Brand          *BrandDAO
	OriginalName   string
	Price          *PriceDAO
	Images         []string
	ProductUrl     string
	Description    map[string]string
	SizeAvailable  []string
	ColorAvailable []string
	PriceHistory   []PriceHistoryDAO
}

type PriceHistoryDAO struct {
	Date  time.Time
	Price float32
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
	ID               string `bson:"_id,omitempty"`
	ProductInfo      *ProductMetaInfoDAO
	DiscountedPrice  int
	DiscountRate     int
	AlloffProductID  string
	AlloffCategories struct {
		First   *AlloffCategoryDAO
		Second  *AlloffCategoryDAO
		Done    bool
		Touched bool
	}
	Soldout   bool
	Removed   bool
	Inventory []struct {
		Size     string
		Quantity int
	}
	Score            *ProductScoreInfoDAO
	SalesInstruction *AlloffInstructionDAO
	Created          time.Time
	Updated          time.Time
}

type PriceDAO struct {
	OriginalPrice float32
	CurrencyType  CurrencyType
	SellersPrice  float32
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
