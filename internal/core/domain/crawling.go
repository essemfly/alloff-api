package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CrawlSourceDAO struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// 크롤링하는 브랜드의  Keyname
	BrandKeyname string
	// 브랜드 Identifier
	BrandIdentifier string
	//  여성의류 / 남성의류 / 키즈 같은 분류
	MainCategoryKey string
	// 크롤링하는 세부 카테고리
	Category CategoryDAO
	// 크롤링 main url
	CrawlUrl string
	// 어떤 몰에서 어떤 모듈로 크롤링하는지 내용
	CrawlModuleName string
	// 아울렛 상품인지 유무, 판매 상품인지 여부
	IsSalesProducts bool
	// 해외배송 마크
	IsForeignDelivery bool

	// 배송비 및 가격 추가에 대한 내용 (TODO) 가격정책이 확정되어야함
	DeliveryPrice     int
	PriceMarginPolicy string

	// 상품 정보에 들어갈 내용
	EarliestDeliveryDays int
	LatestDeliveryDays   int
	DeliveryDesc         []string
	RefundAvailable      bool
	ChangeAvailable      bool
	RefundFee            int
	ChangeFee            int
	RefundRoughFee       int
	ChangeRoughFee       int
}

type CrawlRecordDAO struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Date          time.Time          `json:"date"`
	NewProducts   int                `json:"newproducts"`
	OldProducts   int                `json:"oldproducts"`
	CrawledBrands []string           `json:"crawledbrands"`
}
