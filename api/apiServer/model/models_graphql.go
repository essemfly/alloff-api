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

type AlloffCategoryID struct {
	ID string `json:"id"`
}

type AlloffCategoryInput struct {
	ParentID *string `json:"parentId"`
}

type AlloffCategoryProducts struct {
	Alloffcategory *AlloffCategory `json:"alloffcategory"`
	Products       []*Product      `json:"products"`
	AllBrands      []*Brand        `json:"allBrands"`
	SelectedBrands []string        `json:"selectedBrands"`
	TotalCount     int             `json:"totalCount"`
	Offset         int             `json:"offset"`
	Limit          int             `json:"limit"`
}

type AlloffCategoryProductsInput struct {
	Offset           int           `json:"offset"`
	Limit            int           `json:"limit"`
	AlloffcategoryID string        `json:"alloffcategoryId"`
	BrandIds         []string      `json:"brandIds"`
	Sorting          []SortingType `json:"sorting"`
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

type BrandItem struct {
	ImgURL string `json:"imgUrl"`
	Brand  *Brand `json:"brand"`
}

type BrandsInput struct {
	OnlyLikes *bool `json:"onlyLikes"`
}

type BrandsResult struct {
	Brands      []*Brand `json:"brands"`
	LastUpdated string   `json:"lastUpdated"`
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

type CommunityItem struct {
	Name       string            `json:"name"`
	Target     string            `json:"target"`
	TargetType CommunityItemType `json:"targetType"`
	ImgURL     string            `json:"imgUrl"`
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
	ID                 string              `json:"id"`
	BannerImage        string              `json:"bannerImage"`
	ThumbnailImage     string              `json:"thumbnailImage"`
	Title              string              `json:"title"`
	SubTitle           string              `json:"subTitle"`
	Description        string              `json:"description"`
	ProductGroups      []*ProductGroup     `json:"productGroups"`
	StartTime          string              `json:"startTime"`
	FinishTime         string              `json:"finishTime"`
	TargetSales        int                 `json:"targetSales"`
	CurrentSales       int                 `json:"currentSales"`
	ExhibitionType     ExhibitionType      `json:"exhibitionType"`
	Banners            []*ExhibitionBanner `json:"banners"`
	TotalProducts      int                 `json:"totalProducts"`
	TotalProductGroups int                 `json:"totalProductGroups"`
	TotalParticipants  int                 `json:"totalParticipants"`
	NumUsersRequired   int                 `json:"numUsersRequired"`
	TotalUserGroups    int                 `json:"totalUserGroups"`
	CheapestPrice      int                 `json:"cheapestPrice"`
	UserGroup          *UserGroup          `json:"userGroup"`
	LatestPurchase     []*OrderItem        `json:"latestPurchase"`
}

type ExhibitionBanner struct {
	ImgURL         string `json:"imgUrl"`
	ProductGroupID string `json:"productGroupId"`
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle"`
}

type FeaturedItem struct {
	ID       string    `json:"id"`
	Order    int       `json:"order"`
	Brand    *Brand    `json:"brand"`
	Img      string    `json:"img"`
	Category *Category `json:"category"`
}

type Group struct {
	ID               string  `json:"id"`
	ExhibitionID     string  `json:"exhibitionId"`
	NumUsersRequired int     `json:"numUsersRequired"`
	Users            []*User `json:"users"`
}

type HomeItem struct {
	ID             string           `json:"id"`
	Priority       int              `json:"priority"`
	Title          string           `json:"title"`
	ItemType       HomeItemType     `json:"itemType"`
	TargetID       string           `json:"targetId"`
	Sorting        []SortingType    `json:"sorting"`
	Images         []string         `json:"images"`
	CommunityItems []*CommunityItem `json:"communityItems"`
	Brands         []*BrandItem     `json:"brands"`
	Products       []*Product       `json:"products"`
	ProductGroups  []*ProductGroup  `json:"productGroups"`
}

type HomeTabItem struct {
	ID           string              `json:"id"`
	Title        string              `json:"title"`
	Description  string              `json:"description"`
	Tags         []string            `json:"tags"`
	BackImageURL string              `json:"backImageUrl"`
	ItemType     HomeTabItemTypeEnum `json:"itemType"`
	Products     []*Product          `json:"products"`
	Brands       []*Brand            `json:"brands"`
	Exhibitions  []*Exhibition       `json:"exhibitions"`
	Reference    *ItemReference      `json:"reference"`
}

type Inventory struct {
	Size     string `json:"size"`
	Quantity int    `json:"quantity"`
}

type InventoryInput struct {
	Size     string `json:"size"`
	Quantity int    `json:"quantity"`
}

type ItemReference struct {
	Path    string        `json:"path"`
	Params  string        `json:"params"`
	Options []SortingType `json:"options"`
}

type KeyValueInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type LikeBrandInput struct {
	BrandID string `json:"brandId"`
}

type LikeProductInput struct {
	ProductID string `json:"productId"`
}

type LikeProductOutput struct {
	OldProduct *Product `json:"oldProduct"`
	NewProduct *Product `json:"newProduct"`
}

type Login struct {
	UUID   string `json:"uuid"`
	Mobile string `json:"mobile"`
}

type MyGroupDeal struct {
	User              *User `json:"user"`
	NumParticipates   int   `json:"numParticipates"`
	NumLiveGroupdeals int   `json:"numLiveGroupdeals"`
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
	Category            *Category            `json:"category"`
	Brand               *Brand               `json:"brand"`
	Name                string               `json:"name"`
	ProductGroupID      string               `json:"productGroupId"`
	OriginalPrice       int                  `json:"originalPrice"`
	Soldout             bool                 `json:"soldout"`
	Images              []string             `json:"images"`
	DiscountedPrice     int                  `json:"discountedPrice"`
	DiscountRate        int                  `json:"discountRate"`
	SpecialPrice        *int                 `json:"specialPrice"`
	SpecialDiscountRate *int                 `json:"specialDiscountRate"`
	ProductURL          string               `json:"productUrl"`
	Inventory           []*Inventory         `json:"inventory"`
	IsUpdated           bool                 `json:"isUpdated"`
	IsNewProduct        bool                 `json:"isNewProduct"`
	Removed             bool                 `json:"removed"`
	Information         []*KeyValueInfo      `json:"information"`
	Description         *ProductDescription  `json:"description"`
	CancelDescription   *CancelDescription   `json:"cancelDescription"`
	DeliveryDescription *DeliveryDescription `json:"deliveryDescription"`
	ThumbnailImage      string               `json:"thumbnailImage"`
}

type ProductDescription struct {
	Images []string        `json:"images"`
	Texts  []string        `json:"texts"`
	Infos  []*KeyValueInfo `json:"infos"`
}

type ProductGroup struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	ShortTitle    string     `json:"shortTitle"`
	Instruction   []string   `json:"instruction"`
	ImgURL        string     `json:"imgUrl"`
	Products      []*Product `json:"products"`
	StartTime     string     `json:"startTime"`
	FinishTime    string     `json:"finishTime"`
	NumAlarms     int        `json:"numAlarms"`
	SetAlarm      bool       `json:"setAlarm"`
	Brand         *Brand     `json:"brand"`
	TotalProducts int        `json:"totalProducts"`
}

type ProductQueryInput struct {
	Offset  int           `json:"offset"`
	Limit   int           `json:"limit"`
	Keyword string        `json:"keyword"`
	Sorting []SortingType `json:"sorting"`
}

type ProductsInput struct {
	Offset         int           `json:"offset"`
	Limit          int           `json:"limit"`
	Brand          *string       `json:"brand"`
	Category       *string       `json:"category"`
	Sorting        []SortingType `json:"sorting"`
	ProductGroupID *string       `json:"productGroupId"`
	ExhibitionID   *string       `json:"exhibitionId"`
}

type ProductsOutput struct {
	TotalCount int        `json:"totalCount"`
	Offset     int        `json:"offset"`
	Limit      int        `json:"limit"`
	Products   []*Product `json:"products"`
}

type ProductsResult struct {
	Products    []*Product `json:"products"`
	LastUpdated string     `json:"lastUpdated"`
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

type TopBanner struct {
	ID             string         `json:"id"`
	ImageURL       string         `json:"imageUrl"`
	ExhibitionID   string         `json:"exhibitionId"`
	ExhibitionType ExhibitionType `json:"exhibitionType"`
	Title          string         `json:"title"`
	SubTitle       string         `json:"subTitle"`
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

type UserGroup struct {
	MyInfo  *User   `json:"myInfo"`
	GroupID string  `json:"groupId"`
	Users   []*User `json:"users"`
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

type CommunityItemType string

const (
	CommunityItemTypeOutlink  CommunityItemType = "OUTLINK"
	CommunityItemTypeInternal CommunityItemType = "INTERNAL"
)

var AllCommunityItemType = []CommunityItemType{
	CommunityItemTypeOutlink,
	CommunityItemTypeInternal,
}

func (e CommunityItemType) IsValid() bool {
	switch e {
	case CommunityItemTypeOutlink, CommunityItemTypeInternal:
		return true
	}
	return false
}

func (e CommunityItemType) String() string {
	return string(e)
}

func (e *CommunityItemType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommunityItemType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommunityItemType", str)
	}
	return nil
}

func (e CommunityItemType) MarshalGQL(w io.Writer) {
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

type ExhibitionType string

const (
	ExhibitionTypeGroupdeal     ExhibitionType = "GROUPDEAL"
	ExhibitionTypeTimedeal      ExhibitionType = "TIMEDEAL"
	ExhibitionTypeNormal        ExhibitionType = "NORMAL"
	ExhibitionTypeBrandTimedeal ExhibitionType = "BRAND_TIMEDEAL"
)

var AllExhibitionType = []ExhibitionType{
	ExhibitionTypeGroupdeal,
	ExhibitionTypeTimedeal,
	ExhibitionTypeNormal,
	ExhibitionTypeBrandTimedeal,
}

func (e ExhibitionType) IsValid() bool {
	switch e {
	case ExhibitionTypeGroupdeal, ExhibitionTypeTimedeal, ExhibitionTypeNormal, ExhibitionTypeBrandTimedeal:
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

type GroupdealStatus string

const (
	GroupdealStatusPending GroupdealStatus = "PENDING"
	GroupdealStatusOpen    GroupdealStatus = "OPEN"
	GroupdealStatusClosed  GroupdealStatus = "CLOSED"
)

var AllGroupdealStatus = []GroupdealStatus{
	GroupdealStatusPending,
	GroupdealStatusOpen,
	GroupdealStatusClosed,
}

func (e GroupdealStatus) IsValid() bool {
	switch e {
	case GroupdealStatusPending, GroupdealStatusOpen, GroupdealStatusClosed:
		return true
	}
	return false
}

func (e GroupdealStatus) String() string {
	return string(e)
}

func (e *GroupdealStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = GroupdealStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GroupdealStatus", str)
	}
	return nil
}

func (e GroupdealStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HomeItemType string

const (
	HomeItemTypeCommunity HomeItemType = "COMMUNITY"
	HomeItemTypeImage     HomeItemType = "IMAGE"
	HomeItemTypeProduct   HomeItemType = "PRODUCT"
	HomeItemTypeBrand     HomeItemType = "BRAND"
	HomeItemTypeTimedeal  HomeItemType = "TIMEDEAL"
)

var AllHomeItemType = []HomeItemType{
	HomeItemTypeCommunity,
	HomeItemTypeImage,
	HomeItemTypeProduct,
	HomeItemTypeBrand,
	HomeItemTypeTimedeal,
}

func (e HomeItemType) IsValid() bool {
	switch e {
	case HomeItemTypeCommunity, HomeItemTypeImage, HomeItemTypeProduct, HomeItemTypeBrand, HomeItemTypeTimedeal:
		return true
	}
	return false
}

func (e HomeItemType) String() string {
	return string(e)
}

func (e *HomeItemType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HomeItemType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HomeItemType", str)
	}
	return nil
}

func (e HomeItemType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HomeTabItemTypeEnum string

const (
	HomeTabItemTypeEnumHometabItemBrands             HomeTabItemTypeEnum = "HOMETAB_ITEM_BRANDS"
	HomeTabItemTypeEnumHometabItemBrandExhibition    HomeTabItemTypeEnum = "HOMETAB_ITEM_BRAND_EXHIBITION"
	HomeTabItemTypeEnumHometabItemExhibitions        HomeTabItemTypeEnum = "HOMETAB_ITEM_EXHIBITIONS"
	HomeTabItemTypeEnumHometabItemExhibition         HomeTabItemTypeEnum = "HOMETAB_ITEM_EXHIBITION"
	HomeTabItemTypeEnumHometabItemProductsBrands     HomeTabItemTypeEnum = "HOMETAB_ITEM_PRODUCTS_BRANDS"
	HomeTabItemTypeEnumHometabItemProductsCategories HomeTabItemTypeEnum = "HOMETAB_ITEM_PRODUCTS_CATEGORIES"
)

var AllHomeTabItemTypeEnum = []HomeTabItemTypeEnum{
	HomeTabItemTypeEnumHometabItemBrands,
	HomeTabItemTypeEnumHometabItemBrandExhibition,
	HomeTabItemTypeEnumHometabItemExhibitions,
	HomeTabItemTypeEnumHometabItemExhibition,
	HomeTabItemTypeEnumHometabItemProductsBrands,
	HomeTabItemTypeEnumHometabItemProductsCategories,
}

func (e HomeTabItemTypeEnum) IsValid() bool {
	switch e {
	case HomeTabItemTypeEnumHometabItemBrands, HomeTabItemTypeEnumHometabItemBrandExhibition, HomeTabItemTypeEnumHometabItemExhibitions, HomeTabItemTypeEnumHometabItemExhibition, HomeTabItemTypeEnumHometabItemProductsBrands, HomeTabItemTypeEnumHometabItemProductsCategories:
		return true
	}
	return false
}

func (e HomeTabItemTypeEnum) String() string {
	return string(e)
}

func (e *HomeTabItemTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HomeTabItemTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HomeTabItemTypeEnum", str)
	}
	return nil
}

func (e HomeTabItemTypeEnum) MarshalGQL(w io.Writer) {
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

type SortingType string

const (
	SortingTypePriceAscending         SortingType = "PRICE_ASCENDING"
	SortingTypePriceDescending        SortingType = "PRICE_DESCENDING"
	SortingTypeDiscount0_30           SortingType = "DISCOUNT_0_30"
	SortingTypeDiscount30_50          SortingType = "DISCOUNT_30_50"
	SortingTypeDiscount50_70          SortingType = "DISCOUNT_50_70"
	SortingTypeDiscount70_100         SortingType = "DISCOUNT_70_100"
	SortingTypeDiscountrateAscending  SortingType = "DISCOUNTRATE_ASCENDING"
	SortingTypeDiscountrateDescending SortingType = "DISCOUNTRATE_DESCENDING"
)

var AllSortingType = []SortingType{
	SortingTypePriceAscending,
	SortingTypePriceDescending,
	SortingTypeDiscount0_30,
	SortingTypeDiscount30_50,
	SortingTypeDiscount50_70,
	SortingTypeDiscount70_100,
	SortingTypeDiscountrateAscending,
	SortingTypeDiscountrateDescending,
}

func (e SortingType) IsValid() bool {
	switch e {
	case SortingTypePriceAscending, SortingTypePriceDescending, SortingTypeDiscount0_30, SortingTypeDiscount30_50, SortingTypeDiscount50_70, SortingTypeDiscount70_100, SortingTypeDiscountrateAscending, SortingTypeDiscountrateDescending:
		return true
	}
	return false
}

func (e SortingType) String() string {
	return string(e)
}

func (e *SortingType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortingType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortingType", str)
	}
	return nil
}

func (e SortingType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
