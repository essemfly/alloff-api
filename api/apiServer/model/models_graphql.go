// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AlloffCategory struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	KeyName  string `json:"keyName"`
	Level    int    `json:"level"`
	ParentID string `json:"parentId"`
	ImgURL   string `json:"imgUrl"`
}

type AlloffInventory struct {
	AlloffSize *AlloffSize `json:"alloffSize"`
	Quantity   int         `json:"quantity"`
}

type AlloffSize struct {
	ID         string `json:"id"`
	SizeName   string `json:"sizeName"`
	GuideImage string `json:"guideImage"`
}

type AppVersion struct {
	LatestVersion     string  `json:"latestVersion"`
	MinVersion        string  `json:"minVersion"`
	SubmissionVersion string  `json:"submissionVersion"`
	Message           *string `json:"message"`
	IsMaintenance     bool    `json:"isMaintenance"`
}

type Brand struct {
	ID              string       `json:"id"`
	KorName         string       `json:"korName"`
	EngName         string       `json:"engName"`
	KeyName         string       `json:"keyName"`
	LogoImgURL      string       `json:"logoImgUrl"`
	BackImgURL      string       `json:"backImgUrl"`
	Categories      []*Category  `json:"categories"`
	OnPopular       bool         `json:"onPopular"`
	Description     string       `json:"description"`
	MaxDiscountRate int          `json:"maxDiscountRate"`
	IsOpen          bool         `json:"isOpen"`
	InMaintenance   bool         `json:"inMaintenance"`
	NumNewProducts  int          `json:"numNewProducts"`
	SizeGuide       []*SizeGuide `json:"sizeGuide"`
}

type BrandInput struct {
	BrandID string `json:"brandId"`
}

type BrandsInput struct {
	OnlyLikes *bool `json:"onlyLikes"`
}

type CancelDescription struct {
	RefundAvailable bool `json:"refundAvailable"`
	ChangeAvailable bool `json:"changeAvailable"`
	ChangeFee       int  `json:"changeFee"`
	RefundFee       int  `json:"refundFee"`
}

type Category struct {
	ID      string `json:"id"`
	KeyName string `json:"keyName"`
	Name    string `json:"name"`
}

type DeliveryDescription struct {
	DeliveryType         DeliveryType `json:"deliveryType"`
	DeliveryFee          int          `json:"deliveryFee"`
	EarliestDeliveryDays int          `json:"earliestDeliveryDays"`
	LatestDeliveryDays   int          `json:"latestDeliveryDays"`
	Texts                []string     `json:"texts"`
}

type Device struct {
	ID                string  `json:"id"`
	DeviceID          string  `json:"deviceId"`
	AllowNotification bool    `json:"allowNotification"`
	UserID            *string `json:"userId"`
}

type Exhibition struct {
	ID             string            `json:"id"`
	ExhibitionType ExhibitionType    `json:"exhibitionType"`
	Title          string            `json:"title"`
	SubTitle       string            `json:"subTitle"`
	Description    string            `json:"description"`
	Tags           []string          `json:"tags"`
	BannerImage    string            `json:"bannerImage"`
	ThumbnailImage string            `json:"thumbnailImage"`
	ProductGroups  []*ProductGroup   `json:"productGroups"`
	StartTime      string            `json:"startTime"`
	FinishTime     string            `json:"finishTime"`
	NumAlarms      int               `json:"numAlarms"`
	MetaInfos      []*ExhibitionInfo `json:"metaInfos"`
}

type ExhibitionInfo struct {
	ProductTypes     []AlloffProductType `json:"productTypes"`
	Brands           []*Brand            `json:"brands"`
	AlloffCategories []*AlloffCategory   `json:"alloffCategories"`
	AlloffSizes      []*AlloffSize       `json:"alloffSizes"`
	MaxDisctounts    int                 `json:"maxDisctounts"`
}

type ExhibitionInput struct {
	Status ExhibitionStatus `json:"status"`
}

type ExhibitionOutput struct {
	Exhibitions   []*Exhibition    `json:"exhibitions"`
	Status        ExhibitionStatus `json:"status"`
	LiveCounts    int              `json:"liveCounts"`
	NotOpenCounts int              `json:"notOpenCounts"`
}

type InventoryInput struct {
	Size     string `json:"size"`
	Quantity int    `json:"quantity"`
}

type KeyValueInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type LikeBrandInput struct {
	BrandID string `json:"brandId"`
}

type Login struct {
	UUID   string `json:"uuid"`
	Mobile string `json:"mobile"`
}

type NewUser struct {
	UUID                  string  `json:"uuid"`
	Mobile                string  `json:"mobile"`
	Name                  *string `json:"name"`
	Email                 *string `json:"email"`
	BaseAddress           *string `json:"baseAddress"`
	DetailAddress         *string `json:"detailAddress"`
	Postcode              *string `json:"postcode"`
	PersonalCustomsNumber *string `json:"personalCustomsNumber"`
}

type OrderInfo struct {
	ID            string       `json:"id"`
	Orders        []*OrderItem `json:"orders"`
	ProductPrice  int          `json:"productPrice"`
	DeliveryPrice int          `json:"deliveryPrice"`
	TotalPrice    int          `json:"totalPrice"`
	RefundPrice   *int         `json:"refundPrice"`
	UserMemo      string       `json:"userMemo"`
	CreatedAt     string       `json:"createdAt"`
	UpdatedAt     string       `json:"updatedAt"`
	OrderedAt     string       `json:"orderedAt"`
}

type OrderInput struct {
	Orders       []*OrderItemInput `json:"orders"`
	ProductPrice int               `json:"productPrice"`
}

type OrderItem struct {
	ID                     string               `json:"id"`
	ProductID              string               `json:"productId"`
	ProductName            string               `json:"productName"`
	ProductImg             string               `json:"productImg"`
	BrandKeyname           string               `json:"brandKeyname"`
	BrandKorname           string               `json:"brandKorname"`
	Removed                bool                 `json:"removed"`
	SalesPrice             int                  `json:"salesPrice"`
	Selectsize             string               `json:"selectsize"`
	Quantity               int                  `json:"quantity"`
	OrderItemType          OrderItemTypeEnum    `json:"orderItemType"`
	OrderItemStatus        OrderItemStatusEnum  `json:"orderItemStatus"`
	CancelDescription      *CancelDescription   `json:"cancelDescription"`
	DeliveryDescription    *DeliveryDescription `json:"deliveryDescription"`
	RefundInfo             *RefundInfo          `json:"refundInfo"`
	DeliveryTrackingNumber string               `json:"deliveryTrackingNumber"`
	DeliveryTrackingURL    string               `json:"deliveryTrackingUrl"`
	CreatedAt              string               `json:"createdAt"`
	UpdatedAt              string               `json:"updatedAt"`
	OrderedAt              string               `json:"orderedAt"`
	DeliveryStartedAt      string               `json:"deliveryStartedAt"`
	DeliveryFinishedAt     string               `json:"deliveryFinishedAt"`
	CancelRequestedAt      string               `json:"cancelRequestedAt"`
	CancelFinishedAt       string               `json:"cancelFinishedAt"`
	ConfirmedAt            string               `json:"confirmedAt"`
	UserID                 string               `json:"userId"`
	User                   *User                `json:"user"`
}

type OrderItemInput struct {
	ProductID      string `json:"productId"`
	ProductGroupID string `json:"productGroupId"`
	Selectsize     string `json:"selectsize"`
	Quantity       int    `json:"quantity"`
}

type OrderItemStatusDescription struct {
	DeliveryType DeliveryType        `json:"deliveryType"`
	StatusEnum   OrderItemStatusEnum `json:"statusEnum"`
	Description  string              `json:"description"`
}

type OrderResponse struct {
	Success     bool   `json:"success"`
	ImpUID      string `json:"imp_uid"`
	MerchantUID string `json:"merchant_uid"`
	ErrorMsg    string `json:"error_msg"`
}

type OrderValidityResult struct {
	Available bool       `json:"available"`
	ErrorMsgs []string   `json:"errorMsgs"`
	Order     *OrderInfo `json:"order"`
}

type OrderWithPayment struct {
	Success        bool             `json:"success"`
	ErrorMsg       string           `json:"errorMsg"`
	PaymentMethods []*PaymentMethod `json:"paymentMethods"`
	User           *User            `json:"user"`
	PaymentInfo    *PaymentInfo     `json:"paymentInfo"`
	Order          *OrderInfo       `json:"order"`
}

type PaymentClientInput struct {
	Pg                    string  `json:"pg"`
	PayMethod             string  `json:"payMethod"`
	MerchantUID           string  `json:"merchantUid"`
	Amount                int     `json:"amount"`
	Name                  *string `json:"name"`
	BuyerName             *string `json:"buyerName"`
	BuyerMobile           *string `json:"buyerMobile"`
	BuyerAddress          *string `json:"buyerAddress"`
	BuyerPostCode         *string `json:"buyerPostCode"`
	Memo                  *string `json:"memo"`
	AppScheme             *string `json:"appScheme"`
	PersonalCustomsNumber *string `json:"personalCustomsNumber"`
}

type PaymentInfo struct {
	Pg                    string  `json:"pg"`
	PayMethod             string  `json:"payMethod"`
	MerchantUID           string  `json:"merchantUid"`
	Amount                int     `json:"amount"`
	Name                  string  `json:"name"`
	BuyerName             string  `json:"buyerName"`
	BuyerMobile           string  `json:"buyerMobile"`
	BuyerAddress          string  `json:"buyerAddress"`
	BuyerPostCode         string  `json:"buyerPostCode"`
	Company               string  `json:"company"`
	AppScheme             string  `json:"appScheme"`
	PersonalCustomsNumber *string `json:"personalCustomsNumber"`
}

type PaymentMethod struct {
	Label string `json:"label"`
	Code  string `json:"code"`
}

type PaymentResult struct {
	Success     bool         `json:"success"`
	ErrorMsg    string       `json:"errorMsg"`
	Order       *OrderInfo   `json:"order"`
	PaymentInfo *PaymentInfo `json:"paymentInfo"`
}

type PaymentStatus struct {
	Success     bool         `json:"success"`
	ErrorMsg    string       `json:"errorMsg"`
	Order       *OrderInfo   `json:"order"`
	PaymentInfo *PaymentInfo `json:"paymentInfo"`
}

type Product struct {
	ID                  string               `json:"id"`
	Brand               *Brand               `json:"brand"`
	AlloffCategory      *AlloffCategory      `json:"alloffCategory"`
	Name                string               `json:"name"`
	OriginalPrice       int                  `json:"originalPrice"`
	DiscountedPrice     int                  `json:"discountedPrice"`
	DiscountRate        int                  `json:"discountRate"`
	Images              []string             `json:"images"`
	ThumbnailImage      string               `json:"thumbnailImage"`
	Inventory           []*AlloffInventory   `json:"inventory"`
	IsSoldout           bool                 `json:"isSoldout"`
	Description         *ProductDescription  `json:"description"`
	DeliveryDescription *DeliveryDescription `json:"deliveryDescription"`
	CancelDescription   *CancelDescription   `json:"cancelDescription"`
	Information         []*KeyValueInfo      `json:"information"`
}

type ProductDescription struct {
	Images []string        `json:"images"`
	Texts  []string        `json:"texts"`
	Infos  []*KeyValueInfo `json:"infos"`
}

type ProductGroup struct {
	ID            string     `json:"id"`
	Brand         *Brand     `json:"brand"`
	Title         string     `json:"title"`
	ShortTitle    string     `json:"shortTitle"`
	ImgURL        string     `json:"imgUrl"`
	Products      []*Product `json:"products"`
	TotalProducts int        `json:"totalProducts"`
}

type ProductsInput struct {
	Offset           int                   `json:"offset"`
	Limit            int                   `json:"limit"`
	ProductType      string                `json:"productType"`
	ExhibitionID     string                `json:"exhibitionId"`
	AlloffCategoryID *string               `json:"alloffCategoryId"`
	BrandIds         []string              `json:"brandIds"`
	AlloffSizeIds    []string              `json:"alloffSizeIds"`
	Sorting          []ProductsSortingType `json:"sorting"`
}

type ProductsOutput struct {
	TotalCount   int        `json:"totalCount"`
	Offset       int        `json:"offset"`
	Limit        int        `json:"limit"`
	ExhibitionID string     `json:"exhibitionId"`
	Products     []*Product `json:"products"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type RefundInfo struct {
	RefundFee    int    `json:"refundFee"`
	RefundAmount int    `json:"refundAmount"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type SizeGuide struct {
	Label  string `json:"label"`
	ImgURL string `json:"imgUrl"`
}

type User struct {
	ID                    string  `json:"id"`
	UUID                  string  `json:"uuid"`
	Mobile                string  `json:"mobile"`
	Name                  *string `json:"name"`
	Email                 *string `json:"email"`
	BaseAddress           *string `json:"baseAddress"`
	DetailAddress         *string `json:"detailAddress"`
	Postcode              *string `json:"postcode"`
	PersonalCustomsNumber *string `json:"personalCustomsNumber"`
}

type UserInfoInput struct {
	UUID                  *string `json:"uuid"`
	Name                  *string `json:"name"`
	Mobile                *string `json:"mobile"`
	Email                 *string `json:"email"`
	BaseAddress           *string `json:"baseAddress"`
	DetailAddress         *string `json:"detailAddress"`
	Postcode              *string `json:"postcode"`
	PersonalCustomsNumber *string `json:"personalCustomsNumber"`
}

type AlloffProductType string

const (
	AlloffProductTypeMale   AlloffProductType = "MALE"
	AlloffProductTypeFemale AlloffProductType = "FEMALE"
	AlloffProductTypeKids   AlloffProductType = "KIDS"
	AlloffProductTypeSports AlloffProductType = "SPORTS"
)

var AllAlloffProductType = []AlloffProductType{
	AlloffProductTypeMale,
	AlloffProductTypeFemale,
	AlloffProductTypeKids,
	AlloffProductTypeSports,
}

func (e AlloffProductType) IsValid() bool {
	switch e {
	case AlloffProductTypeMale, AlloffProductTypeFemale, AlloffProductTypeKids, AlloffProductTypeSports:
		return true
	}
	return false
}

func (e AlloffProductType) String() string {
	return string(e)
}

func (e *AlloffProductType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AlloffProductType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AlloffProductType", str)
	}
	return nil
}

func (e AlloffProductType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DeliveryType string

const (
	DeliveryTypeDomesticDelivery DeliveryType = "DOMESTIC_DELIVERY"
	DeliveryTypeForeignDelivery  DeliveryType = "FOREIGN_DELIVERY"
)

var AllDeliveryType = []DeliveryType{
	DeliveryTypeDomesticDelivery,
	DeliveryTypeForeignDelivery,
}

func (e DeliveryType) IsValid() bool {
	switch e {
	case DeliveryTypeDomesticDelivery, DeliveryTypeForeignDelivery:
		return true
	}
	return false
}

func (e DeliveryType) String() string {
	return string(e)
}

func (e *DeliveryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeliveryType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeliveryType", str)
	}
	return nil
}

func (e DeliveryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ExhibitionStatus string

const (
	ExhibitionStatusLive    ExhibitionStatus = "LIVE"
	ExhibitionStatusNotOpen ExhibitionStatus = "NOT_OPEN"
	ExhibitionStatusClosed  ExhibitionStatus = "CLOSED"
)

var AllExhibitionStatus = []ExhibitionStatus{
	ExhibitionStatusLive,
	ExhibitionStatusNotOpen,
	ExhibitionStatusClosed,
}

func (e ExhibitionStatus) IsValid() bool {
	switch e {
	case ExhibitionStatusLive, ExhibitionStatusNotOpen, ExhibitionStatusClosed:
		return true
	}
	return false
}

func (e ExhibitionStatus) String() string {
	return string(e)
}

func (e *ExhibitionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ExhibitionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ExhibitionStatus", str)
	}
	return nil
}

func (e ExhibitionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ExhibitionType string

const (
	ExhibitionTypeNormal    ExhibitionType = "NORMAL"
	ExhibitionTypeGroupdeal ExhibitionType = "GROUPDEAL"
	ExhibitionTypeTimedeal  ExhibitionType = "TIMEDEAL"
	ExhibitionTypeAll       ExhibitionType = "ALL"
)

var AllExhibitionType = []ExhibitionType{
	ExhibitionTypeNormal,
	ExhibitionTypeGroupdeal,
	ExhibitionTypeTimedeal,
	ExhibitionTypeAll,
}

func (e ExhibitionType) IsValid() bool {
	switch e {
	case ExhibitionTypeNormal, ExhibitionTypeGroupdeal, ExhibitionTypeTimedeal, ExhibitionTypeAll:
		return true
	}
	return false
}

func (e ExhibitionType) String() string {
	return string(e)
}

func (e *ExhibitionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ExhibitionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ExhibitionType", str)
	}
	return nil
}

func (e ExhibitionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderItemStatusEnum string

const (
	OrderItemStatusEnumUnknown                  OrderItemStatusEnum = "UNKNOWN"
	OrderItemStatusEnumCreated                  OrderItemStatusEnum = "CREATED"
	OrderItemStatusEnumRecreated                OrderItemStatusEnum = "RECREATED"
	OrderItemStatusEnumPaymentPending           OrderItemStatusEnum = "PAYMENT_PENDING"
	OrderItemStatusEnumPaymentFinished          OrderItemStatusEnum = "PAYMENT_FINISHED"
	OrderItemStatusEnumProductPreparing         OrderItemStatusEnum = "PRODUCT_PREPARING"
	OrderItemStatusEnumForeignProductInspecting OrderItemStatusEnum = "FOREIGN_PRODUCT_INSPECTING"
	OrderItemStatusEnumDeliveryPreparing        OrderItemStatusEnum = "DELIVERY_PREPARING"
	OrderItemStatusEnumForeignDeliveryStatrted  OrderItemStatusEnum = "FOREIGN_DELIVERY_STATRTED"
	OrderItemStatusEnumDeliveryStarted          OrderItemStatusEnum = "DELIVERY_STARTED"
	OrderItemStatusEnumDeliveryFinished         OrderItemStatusEnum = "DELIVERY_FINISHED"
	OrderItemStatusEnumConfirmPayment           OrderItemStatusEnum = "CONFIRM_PAYMENT"
	OrderItemStatusEnumCancelFinished           OrderItemStatusEnum = "CANCEL_FINISHED"
	OrderItemStatusEnumExchangeRequested        OrderItemStatusEnum = "EXCHANGE_REQUESTED"
	OrderItemStatusEnumExchangePending          OrderItemStatusEnum = "EXCHANGE_PENDING"
	OrderItemStatusEnumExchangeFinished         OrderItemStatusEnum = "EXCHANGE_FINISHED"
	OrderItemStatusEnumReturnRequested          OrderItemStatusEnum = "RETURN_REQUESTED"
	OrderItemStatusEnumReturnPending            OrderItemStatusEnum = "RETURN_PENDING"
	OrderItemStatusEnumReturnFinished           OrderItemStatusEnum = "RETURN_FINISHED"
)

var AllOrderItemStatusEnum = []OrderItemStatusEnum{
	OrderItemStatusEnumUnknown,
	OrderItemStatusEnumCreated,
	OrderItemStatusEnumRecreated,
	OrderItemStatusEnumPaymentPending,
	OrderItemStatusEnumPaymentFinished,
	OrderItemStatusEnumProductPreparing,
	OrderItemStatusEnumForeignProductInspecting,
	OrderItemStatusEnumDeliveryPreparing,
	OrderItemStatusEnumForeignDeliveryStatrted,
	OrderItemStatusEnumDeliveryStarted,
	OrderItemStatusEnumDeliveryFinished,
	OrderItemStatusEnumConfirmPayment,
	OrderItemStatusEnumCancelFinished,
	OrderItemStatusEnumExchangeRequested,
	OrderItemStatusEnumExchangePending,
	OrderItemStatusEnumExchangeFinished,
	OrderItemStatusEnumReturnRequested,
	OrderItemStatusEnumReturnPending,
	OrderItemStatusEnumReturnFinished,
}

func (e OrderItemStatusEnum) IsValid() bool {
	switch e {
	case OrderItemStatusEnumUnknown, OrderItemStatusEnumCreated, OrderItemStatusEnumRecreated, OrderItemStatusEnumPaymentPending, OrderItemStatusEnumPaymentFinished, OrderItemStatusEnumProductPreparing, OrderItemStatusEnumForeignProductInspecting, OrderItemStatusEnumDeliveryPreparing, OrderItemStatusEnumForeignDeliveryStatrted, OrderItemStatusEnumDeliveryStarted, OrderItemStatusEnumDeliveryFinished, OrderItemStatusEnumConfirmPayment, OrderItemStatusEnumCancelFinished, OrderItemStatusEnumExchangeRequested, OrderItemStatusEnumExchangePending, OrderItemStatusEnumExchangeFinished, OrderItemStatusEnumReturnRequested, OrderItemStatusEnumReturnPending, OrderItemStatusEnumReturnFinished:
		return true
	}
	return false
}

func (e OrderItemStatusEnum) String() string {
	return string(e)
}

func (e *OrderItemStatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderItemStatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderItemStatusEnum", str)
	}
	return nil
}

func (e OrderItemStatusEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderItemTypeEnum string

const (
	OrderItemTypeEnumUnknown    OrderItemTypeEnum = "UNKNOWN"
	OrderItemTypeEnumTimedeal   OrderItemTypeEnum = "TIMEDEAL"
	OrderItemTypeEnumExhibition OrderItemTypeEnum = "EXHIBITION"
	OrderItemTypeEnumGroupdeal  OrderItemTypeEnum = "GROUPDEAL"
	OrderItemTypeEnumNormal     OrderItemTypeEnum = "NORMAL"
)

var AllOrderItemTypeEnum = []OrderItemTypeEnum{
	OrderItemTypeEnumUnknown,
	OrderItemTypeEnumTimedeal,
	OrderItemTypeEnumExhibition,
	OrderItemTypeEnumGroupdeal,
	OrderItemTypeEnumNormal,
}

func (e OrderItemTypeEnum) IsValid() bool {
	switch e {
	case OrderItemTypeEnumUnknown, OrderItemTypeEnumTimedeal, OrderItemTypeEnumExhibition, OrderItemTypeEnumGroupdeal, OrderItemTypeEnumNormal:
		return true
	}
	return false
}

func (e OrderItemTypeEnum) String() string {
	return string(e)
}

func (e *OrderItemTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderItemTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderItemTypeEnum", str)
	}
	return nil
}

func (e OrderItemTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderStatusEnum string

const (
	OrderStatusEnumUnknown         OrderStatusEnum = "UNKNOWN"
	OrderStatusEnumCreated         OrderStatusEnum = "CREATED"
	OrderStatusEnumRecreated       OrderStatusEnum = "RECREATED"
	OrderStatusEnumPaymentPending  OrderStatusEnum = "PAYMENT_PENDING"
	OrderStatusEnumPaymentFinished OrderStatusEnum = "PAYMENT_FINISHED"
)

var AllOrderStatusEnum = []OrderStatusEnum{
	OrderStatusEnumUnknown,
	OrderStatusEnumCreated,
	OrderStatusEnumRecreated,
	OrderStatusEnumPaymentPending,
	OrderStatusEnumPaymentFinished,
}

func (e OrderStatusEnum) IsValid() bool {
	switch e {
	case OrderStatusEnumUnknown, OrderStatusEnumCreated, OrderStatusEnumRecreated, OrderStatusEnumPaymentPending, OrderStatusEnumPaymentFinished:
		return true
	}
	return false
}

func (e OrderStatusEnum) String() string {
	return string(e)
}

func (e *OrderStatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderStatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderStatusEnum", str)
	}
	return nil
}

func (e OrderStatusEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PaymentStatusEnum string

const (
	PaymentStatusEnumCreated         PaymentStatusEnum = "CREATED"
	PaymentStatusEnumConfirmed       PaymentStatusEnum = "CONFIRMED"
	PaymentStatusEnumTimeOut         PaymentStatusEnum = "TIME_OUT"
	PaymentStatusEnumCancled         PaymentStatusEnum = "CANCLED"
	PaymentStatusEnumRefundRequested PaymentStatusEnum = "REFUND_REQUESTED"
	PaymentStatusEnumRefundFinished  PaymentStatusEnum = "REFUND_FINISHED"
)

var AllPaymentStatusEnum = []PaymentStatusEnum{
	PaymentStatusEnumCreated,
	PaymentStatusEnumConfirmed,
	PaymentStatusEnumTimeOut,
	PaymentStatusEnumCancled,
	PaymentStatusEnumRefundRequested,
	PaymentStatusEnumRefundFinished,
}

func (e PaymentStatusEnum) IsValid() bool {
	switch e {
	case PaymentStatusEnumCreated, PaymentStatusEnumConfirmed, PaymentStatusEnumTimeOut, PaymentStatusEnumCancled, PaymentStatusEnumRefundRequested, PaymentStatusEnumRefundFinished:
		return true
	}
	return false
}

func (e PaymentStatusEnum) String() string {
	return string(e)
}

func (e *PaymentStatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PaymentStatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PaymentStatusEnum", str)
	}
	return nil
}

func (e PaymentStatusEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProductsSortingType string

const (
	ProductsSortingTypePriceAscending         ProductsSortingType = "PRICE_ASCENDING"
	ProductsSortingTypePriceDescending        ProductsSortingType = "PRICE_DESCENDING"
	ProductsSortingTypeDiscount0_30           ProductsSortingType = "DISCOUNT_0_30"
	ProductsSortingTypeDiscount30_50          ProductsSortingType = "DISCOUNT_30_50"
	ProductsSortingTypeDiscount50_70          ProductsSortingType = "DISCOUNT_50_70"
	ProductsSortingTypeDiscount70_100         ProductsSortingType = "DISCOUNT_70_100"
	ProductsSortingTypeDiscountrateAscending  ProductsSortingType = "DISCOUNTRATE_ASCENDING"
	ProductsSortingTypeDiscountrateDescending ProductsSortingType = "DISCOUNTRATE_DESCENDING"
	ProductsSortingTypeInventoryAscending     ProductsSortingType = "INVENTORY_ASCENDING"
	ProductsSortingTypeInventoryDescending    ProductsSortingType = "INVENTORY_DESCENDING"
)

var AllProductsSortingType = []ProductsSortingType{
	ProductsSortingTypePriceAscending,
	ProductsSortingTypePriceDescending,
	ProductsSortingTypeDiscount0_30,
	ProductsSortingTypeDiscount30_50,
	ProductsSortingTypeDiscount50_70,
	ProductsSortingTypeDiscount70_100,
	ProductsSortingTypeDiscountrateAscending,
	ProductsSortingTypeDiscountrateDescending,
	ProductsSortingTypeInventoryAscending,
	ProductsSortingTypeInventoryDescending,
}

func (e ProductsSortingType) IsValid() bool {
	switch e {
	case ProductsSortingTypePriceAscending, ProductsSortingTypePriceDescending, ProductsSortingTypeDiscount0_30, ProductsSortingTypeDiscount30_50, ProductsSortingTypeDiscount50_70, ProductsSortingTypeDiscount70_100, ProductsSortingTypeDiscountrateAscending, ProductsSortingTypeDiscountrateDescending, ProductsSortingTypeInventoryAscending, ProductsSortingTypeInventoryDescending:
		return true
	}
	return false
}

func (e ProductsSortingType) String() string {
	return string(e)
}

func (e *ProductsSortingType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProductsSortingType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProductsSortingType", str)
	}
	return nil
}

func (e ProductsSortingType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
