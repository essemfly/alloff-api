package productinfo

import (
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/alloffcategory"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type AddMetaInfoRequest struct {
	AlloffName           string
	ProductID            string
	ProductUrl           string
	ProductType          []domain.AlloffProductType
	OriginalPrice        float32
	DiscountedPrice      float32
	CurrencyType         domain.CurrencyType
	Brand                *domain.BrandDAO
	Source               *domain.CrawlSourceDAO
	AlloffCategory       *domain.AlloffCategoryDAO
	Images               []string
	ThumbnailImage       string
	Colors               []string
	Sizes                []string
	Inventory            []*domain.InventoryDAO
	Description          []string
	DescriptionImages    []string
	DescriptionInfos     map[string]string
	DescriptionRawInfos  map[string]string
	Information          map[string]string
	RawInformation       map[string]string
	IsForeignDelivery    bool
	EarliestDeliveryDays int
	LatestDeliveryDays   int
	IsRefundPossible     bool
	RefundFee            int
	ModuleName           string
	IsTranslateRequired  bool
	IsInventoryMapped    bool
	IsRemoved            bool
	IsSoldout            bool
}

func AddProductInfo(request *AddMetaInfoRequest) (*domain.ProductMetaInfoDAO, error) {
	_, err := ioc.Repo.ProductMetaInfos.GetByProductID(request.Brand.KeyName, request.ProductID)
	if err != mongo.ErrNoDocuments {
		config.Logger.Error("already registered products", zap.Error(err))
		return nil, err
	}

	pdInfo := makeBaseProductInfo(request)
	// 상품 크롤링시 번역은 하지않는다.
	//if pdInfo.IsTranslateRequired {
	//	translated, err := TranslateProductInfo(pdInfo)
	//	if err != nil {
	//		config.Logger.Error("err occurred on translate product info : ", zap.Error(err))
	//	}
	//	if translated != nil {
	//		pdInfo.IsTranslateRequired = false
	//		pdInfo = translated
	//	}
	//}

	pdInfo, err = ioc.Repo.ProductMetaInfos.Insert(pdInfo)
	if err != nil {
		config.Logger.Error("error on adding product infos", zap.Error(err))
	}
	return pdInfo, nil
}

func makeBaseProductInfo(request *AddMetaInfoRequest) *domain.ProductMetaInfoDAO {
	pdInfo := &domain.ProductMetaInfoDAO{
		ID:                   primitive.NewObjectID(),
		Brand:                &domain.BrandDAO{},
		Source:               &domain.CrawlSourceDAO{},
		Category:             &domain.CategoryDAO{},
		AlloffCategory:       &domain.ProductAlloffCategoryDAO{},
		ProductType:          []domain.AlloffProductType{},
		OriginalName:         "",
		AlloffName:           "",
		ProductID:            "",
		ProductUrl:           "",
		Price:                &domain.PriceDAO{},
		Images:               []string{},
		CachedImages:         []string{},
		ThumbnailImage:       "",
		Colors:               []string{},
		Sizes:                []string{},
		Inventory:            []*domain.InventoryDAO{},
		SalesInstruction:     &domain.AlloffInstructionDAO{},
		IsImageCached:        false,
		IsInventoryMapped:    false,
		IsTranslateRequired:  request.IsTranslateRequired,
		IsCategoryClassified: false,
		IsSoldout:            false,
		IsRemoved:            false,
		Created:              time.Now(),
		Updated:              time.Now(),
	}

	pdInfo.SetBrandAndCategory(request.Brand, request.Source)
	pdInfo.SetGeneralInfo(request.ProductType, request.AlloffName, request.ProductID, request.ProductUrl, request.Images, request.Sizes, request.Colors, request.Information)
	alloffOrigPrice, alloffDiscPrice := GetProductPrice(float32(request.OriginalPrice), float32(request.DiscountedPrice), request.CurrencyType, request.Source.PriceMarginPolicy)
	pdInfo.SetPrices(alloffOrigPrice, alloffDiscPrice, domain.CurrencyKRW)
	pdInfo.SetInformation(request.Information)
	//descImages := append(request.DescriptionImages, request.Images...) TODO 이렇게 하면 Images가 수정될때마다 수정되는 친구들이 계속 descImages에 쌓이게 된다.
	descImages := request.DescriptionImages
	if request.ModuleName == "intrend" {
		descImages = append([]string{
			"https://alloff.s3.ap-northeast-2.amazonaws.com/description/Intrend_info.png",
		}, descImages...)
	}

	pdInfo.SetDesc(descImages, request.Description, request.DescriptionInfos)
	pdInfo.SetDeliveryDesc(request.IsForeignDelivery, 0, request.EarliestDeliveryDays, request.LatestDeliveryDays)
	pdInfo.SetCancelDesc(request.IsRefundPossible, request.RefundFee)
	pdInfo.SetThumbnail(request.ThumbnailImage)
	pdInfo.SetCachedImages(request.Images)

	if request.AlloffCategory != nil && request.AlloffCategory.ID != primitive.NilObjectID {
		alloffcat, _ := ioc.Repo.AlloffCategories.Get(request.AlloffCategory.ID.Hex())
		productAlloffCat, err := alloffcategory.BuildProductAlloffCategory(alloffcat, true)
		if err != nil {
			config.Logger.Error("err occured on build product alloff category : alloffcat ID"+request.AlloffCategory.ID.Hex(), zap.Error(err))
		}
		pdInfo.SetAlloffCategory(productAlloffCat)
	}

	if request.AlloffCategory == nil || request.AlloffCategory.ID == primitive.NilObjectID || !pdInfo.AlloffCategory.Done {
		productAlloffCat, err := alloffcategory.InferAlloffCategory(pdInfo)
		if err != nil {
			config.Logger.Error("err occured on infer alloffcategory: pdinfo "+pdInfo.ID.Hex(), zap.Error(err))
		} else {
			pdInfo.SetAlloffCategory(productAlloffCat)
		}
	}

	inventories := AssignAlloffSizesToInventories(request.Inventory, pdInfo.ProductType, pdInfo.AlloffCategory)
	pdInfo.SetInventory(inventories)

	pdInfo.IsRemoved = request.IsRemoved
	pdInfo.IsSoldout = request.IsSoldout
	if !pdInfo.IsSoldout {
		pdInfo.CheckSoldout()
	}

	// if pd.IsTranslateRequired {
	// 	_, err := TranslateProductInfo(pd)
	// 	if err != nil {
	// 		log.Println("Err on translate product info", err)
	// 	}
	// }

	// if !pd.IsImageCached {
	// 	err := CacheProductsImage(pd)
	// 	if err != nil {
	// 		log.Println("Err on cache products image", err)
	// 	}
	// }

	return pdInfo
}
