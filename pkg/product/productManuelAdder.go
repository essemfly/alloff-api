package product

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddProductInManuel(request *ProductManuelAddRequest) (*domain.ProductDAO, error) {
	pdInfo, err := AddManuelProductInfo(request)
	if err != nil {
		return nil, err
	}

	pd := &domain.ProductDAO{
		AlloffName:  pdInfo.OriginalName,
		ProductInfo: pdInfo,
		Removed:     false,
		Created:     time.Now(),
		Updated:     time.Now(),
	}

	return ProcessManuelProductRequest(pd, request)
}

func AddManuelProductInfo(request *ProductManuelAddRequest) (*domain.ProductMetaInfoDAO, error) {
	pdInfo := &domain.ProductMetaInfoDAO{
		Created: time.Now(),
		Updated: time.Now(),
	}

	brand, err := ioc.Repo.Brands.GetByKeyname(request.BrandKeyName)
	if err != nil {
		return nil, err
	}

	source := GetManuelSource(request)
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
	pdInfo.SetPrices(int(request.OriginalPrice), int(request.DiscountedPrice), domain.CurrencyKRW)

	newPdInfo, err := ioc.Repo.ProductMetaInfos.Insert(pdInfo)
	if err != nil {
		log.Println("productinfo 1", err)
		return nil, err
	}
	return newPdInfo, nil
}

func GetManuelSource(req *ProductManuelAddRequest) *domain.CrawlSourceDAO {
	return &domain.CrawlSourceDAO{
		BrandKeyname:         req.BrandKeyName,
		BrandIdentifier:      "",
		MainCategoryKey:      "",
		Category:             domain.CategoryDAO{},
		CrawlUrl:             "",
		CrawlModuleName:      "manuel",
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

func ProcessManuelProductRequest(pd *domain.ProductDAO, request *ProductManuelAddRequest) (*domain.ProductDAO, error) {
	alloffInstruction := GetManuelProductDescription(pd, request)
	pd.UpdateInstruction(alloffInstruction)

	alloffScore := GetProductScore(pd)
	pd.UpdateScore(alloffScore)
	pd.UpdateInventory(request.Inventory)

	pd.UpdatePrice(float32(request.DiscountedPrice))
	pd.SpecialPrice = request.SpecialPrice

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
