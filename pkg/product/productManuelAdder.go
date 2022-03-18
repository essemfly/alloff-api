package product

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/classifier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProductManually(request *ProductManualAddRequest) (*domain.ProductDAO, error) {
	pdInfo, err := AddManualProductInfo(request)
	if err != nil {
		return nil, err
	}

	pd := &domain.ProductDAO{
		AlloffName:    pdInfo.OriginalName,
		ProductInfo:   pdInfo,
		Removed:       false,
		Created:       time.Now(),
		Updated:       time.Now(),
		IsImageCached: true,
	}

	return ProcessManualProductRequest(pd, request)
}

func AddManualProductInfo(request *ProductManualAddRequest) (*domain.ProductMetaInfoDAO, error) {
	pdInfo := &domain.ProductMetaInfoDAO{
		Created: time.Now(),
		Updated: time.Now(),
	}

	brand, err := ioc.Repo.Brands.GetByKeyname(request.BrandKeyName)
	if err != nil {
		return nil, err
	}

	source := GetManualSource(request)
	newSource, err := ioc.Repo.CrawlSources.Upsert(source)
	if err != nil {
		return nil, err
	}

	sizes := []string{}
	for _, inv := range request.Inventory {
		sizes = append(sizes, inv.Size)
	}

	pdInfo.Images = request.Images
	pdInfo.SetBrandAndCategory(brand, newSource)
	pdInfo.SetGeneralInfo(request.AlloffName, request.ProductID, "", request.Images, sizes, nil, nil)
	pdInfo.SetPrices(float32(request.OriginalPrice), float32(request.DiscountedPrice), domain.CurrencyKRW)

	newPdInfo, err := ioc.Repo.ProductMetaInfos.Insert(pdInfo)
	if err != nil {
		log.Println("productinfo 1", err)
		return nil, err
	}
	return newPdInfo, nil
}

func GetManualSource(req *ProductManualAddRequest) *domain.CrawlSourceDAO {
	return &domain.CrawlSourceDAO{
		BrandKeyname:         req.BrandKeyName,
		BrandIdentifier:      "",
		MainCategoryKey:      "",
		Category:             domain.CategoryDAO{},
		CrawlUrl:             "",
		CrawlModuleName:      req.ModuleName,
		IsSalesProducts:      true,
		IsForeignDelivery:    req.IsForeignDelivery,
		PriceMarginPolicy:    "NORMAL",
		DeliveryPrice:        0,
		EarliestDeliveryDays: req.EarliestDeliveryDays,
		LatestDeliveryDays:   req.LatestDeliveryDays,
		DeliveryDesc:         nil,
		RefundAvailable:      req.IsRefundPossible,
		ChangeAvailable:      req.IsRefundPossible,
		ChangeFee:            req.RefundFee,
		RefundFee:            req.RefundFee,
	}
}

func ProcessManualProductRequest(pd *domain.ProductDAO, request *ProductManualAddRequest) (*domain.ProductDAO, error) {
	if request.AlloffCategoryID != "" {
		productCatDao := classifier.ClassifyProducts(request.AlloffCategoryID)
		pd.UpdateAlloffCategory(productCatDao)
	}

	alloffInstruction := GetManualProductDescription(pd, request)
	pd.UpdateInstruction(alloffInstruction)

	alloffScore := GetProductScore(pd)
	pd.UpdateScore(alloffScore)
	pd.UpdateInventory(request.Inventory)

	pd.UpdatePrice(request.OriginalPrice, request.DiscountedPrice)
	pd.SpecialPrice = request.SpecialPrice

	pd.Images = request.Images
	pd.IsTranslateRequired = false

	if pd.ID == primitive.NilObjectID {
		newPd, err := ioc.Repo.Products.Insert(pd)
		if err != nil {
			log.Println("product 1", err)
			return nil, err
		}
		return newPd, nil
	}

	newPd, err := ioc.Repo.Products.Upsert(pd)
	if err != nil {
		log.Println("product 2", err)
		return nil, err
	}
	return newPd, nil
}

func UpdateManuelProducts() {
	offset, limit := 0, 10000
	filter := bson.M{
		"productinfo.source.crawlmodulename": "manuel",
		"removed":                            false,
	}
	twoDaysAgo := time.Now().Add(-48 * time.Hour)
	pds, _, err := ioc.Repo.Products.List(offset, limit, filter, nil)
	if err != nil {
		log.Println("err on update manuel products", err)
	}
	for _, pd := range pds {
		isNewProduct := true
		if pd.Created.Before(twoDaysAgo) {
			isNewProduct = false
		}
		pd.Score.IsNewlyCrawled = isNewProduct
		_, err := ioc.Repo.Products.Upsert(pd)
		if err != nil {
			log.Println("err on update products", err)
		}
	}
}
