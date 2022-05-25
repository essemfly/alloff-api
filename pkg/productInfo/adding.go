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
	AlloffInventory      []*domain.AlloffInventoryDAO
	Description          []string
	DescriptionImages    []string
	DescriptionInfos     map[string]string
	Information          map[string]string
	IsForeignDelivery    bool
	EarliestDeliveryDays int
	LatestDeliveryDays   int
	IsRefundPossible     bool
	RefundFee            int
	ModuleName           string
	IsTranslateRequired  bool
	IsInventoryMapped    bool
}

func AddProductInfo(request *AddMetaInfoRequest) (*domain.ProductMetaInfoDAO, error) {
	_, err := ioc.Repo.ProductMetaInfos.GetByProductID(request.Brand.KeyName, request.ProductID)
	if err != mongo.ErrNoDocuments {
		config.Logger.Error("already registered products", zap.Error(err))
		return nil, err
	}

	pdInfo := makeBaseProductInfo(request)
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
		AlloffInventory:      []*domain.AlloffInventoryDAO{},
		SalesInstruction:     &domain.AlloffInstructionDAO{},
		IsImageCached:        false,
		IsInventoryMapped:    false,
		IsTranslateRequired:  false,
		IsCategoryClassified: false,
		IsSoldout:            false,
		IsRemoved:            false,
		Created:              time.Now(),
		Updated:              time.Now(),
	}

	pdInfo.SetBrandAndCategory(request.Brand, request.Source)
	pdInfo.SetGeneralInfo(request.AlloffName, request.ProductID, request.ProductUrl, request.Images, request.Sizes, request.Colors, request.Information)
	alloffOrigPrice, alloffDiscPrice := GetProductPrice(float32(request.OriginalPrice), float32(request.DiscountedPrice), request.CurrencyType, request.Source.PriceMarginPolicy)
	pdInfo.SetPrices(alloffOrigPrice, alloffDiscPrice, domain.CurrencyKRW)
	pdInfo.ProductType = request.ProductType

	descImages := append(request.DescriptionImages, request.Images...)
	if request.ModuleName == "intrend" {
		descImages = append([]string{
			"https://alloff.s3.ap-northeast-2.amazonaws.com/description/Intrend_info.png",
		}, descImages...)
	}
	if request.ModuleName == "theoutnet" || request.ModuleName == "sandro" || request.ModuleName == "maje" || request.ModuleName == "intrend" {
		descImages = append(descImages, "https://alloff.s3.ap-northeast-2.amazonaws.com/description/size_guide.png")
	}

	pdInfo.SetDesc(descImages, request.Description, request.DescriptionInfos)
	pdInfo.SetDeliveryDesc(request.IsForeignDelivery, 0, request.EarliestDeliveryDays, request.LatestDeliveryDays)
	pdInfo.SetCancelDesc(request.IsRefundPossible, request.RefundFee)

	if request.AlloffCategory != nil {
		productAlloffCat, err := alloffcategory.BuildProductAlloffCategory(request.AlloffCategory.ID.Hex(), true)
		if err != nil {
			config.Logger.Error("err occured on build product alloff category : alloffcat ID"+request.AlloffCategory.ID.Hex(), zap.Error(err))
		}
		pdInfo.SetAlloffCategory(productAlloffCat)
	}

	if !pdInfo.AlloffCategory.Done {
		productAlloffCat, err := alloffcategory.InferAlloffCategory(pdInfo)
		if err != nil {
			config.Logger.Error("err occured on infer alloffcategory: pdinfo "+pdInfo.ID.Hex(), zap.Error(err))
		}
		pdInfo.SetAlloffCategory(productAlloffCat)
	}

	if request.IsInventoryMapped {
		pdInfo.AlloffInventory = request.AlloffInventory
	} else {
		pdInfo.SetAlloffInventory(request.Inventory)
	}

	pdInfo.IsInventoryMapped = true

	return pdInfo
}
