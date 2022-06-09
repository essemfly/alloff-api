package productinfo

import (
	"fmt"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.uber.org/zap"
	"time"
)

func UpdateProductInfo(pdInfo *domain.ProductMetaInfoDAO, request *AddMetaInfoRequest, requestFrom string) (*domain.ProductMetaInfoDAO, error) {
	switch requestFrom {
	// only update inventories for requests from crawler
	case "CRAWLER":
		inventories := AssignAlloffSizesToInventories(request.Inventory, pdInfo.ProductType, pdInfo.AlloffCategory)
		pdInfo.SetInventory(inventories)
		pdInfo.IsRemoved = false

		updatedPdInfo, err := Update(pdInfo)
		if err != nil {
			config.Logger.Error("error on adding product infos", zap.Error(err))
			return nil, err
		}

		return updatedPdInfo, nil
	// update data on requests from grpc
	case "GRPC":
		newPdInfo := makeBaseProductInfo(request)
		newPdInfo.ID = pdInfo.ID

		updatedPdInfo, err := Update(newPdInfo)
		if err != nil {
			config.Logger.Error("error on adding product infos", zap.Error(err))
			return nil, err
		}

		return updatedPdInfo, nil
	default:
		return nil, fmt.Errorf("requests not supported")
	}
}

func LoadMetaInfoRequest(pdInfoDao *domain.ProductMetaInfoDAO) *AddMetaInfoRequest {
	return &AddMetaInfoRequest{
		AlloffName:           pdInfoDao.AlloffName,
		ProductID:            pdInfoDao.ProductID,
		ProductUrl:           pdInfoDao.ProductUrl,
		ProductType:          pdInfoDao.ProductType,
		OriginalPrice:        float32(pdInfoDao.Price.OriginalPrice),
		DiscountedPrice:      float32(pdInfoDao.Price.CurrentPrice),
		CurrencyType:         pdInfoDao.Price.CurrencyType,
		Brand:                pdInfoDao.Brand,
		Source:               pdInfoDao.Source,
		AlloffCategory:       pdInfoDao.AlloffCategory.First,
		Images:               pdInfoDao.Images,
		ThumbnailImage:       pdInfoDao.ThumbnailImage,
		Colors:               pdInfoDao.Colors,
		Sizes:                pdInfoDao.Sizes,
		Inventory:            pdInfoDao.Inventory,
		Description:          pdInfoDao.SalesInstruction.Description.Texts,
		DescriptionImages:    pdInfoDao.SalesInstruction.Description.Images,
		DescriptionInfos:     pdInfoDao.SalesInstruction.Description.Infos,
		DescriptionRawInfos:  pdInfoDao.SalesInstruction.Description.RawInfos,
		Information:          pdInfoDao.SalesInstruction.Information,
		RawInformation:       pdInfoDao.SalesInstruction.RawInformation,
		IsForeignDelivery:    pdInfoDao.SalesInstruction.DeliveryDescription.DeliveryType == domain.Foreign,
		EarliestDeliveryDays: pdInfoDao.SalesInstruction.DeliveryDescription.EarliestDeliveryDays,
		LatestDeliveryDays:   pdInfoDao.SalesInstruction.DeliveryDescription.LatestDeliveryDays,
		IsRefundPossible:     pdInfoDao.SalesInstruction.CancelDescription.RefundAvailable,
		RefundFee:            pdInfoDao.SalesInstruction.CancelDescription.RefundFee,
		ModuleName:           pdInfoDao.Source.CrawlModuleName,
		IsTranslateRequired:  pdInfoDao.IsTranslateRequired,
		IsInventoryMapped:    pdInfoDao.IsInventoryMapped,
		IsSoldout:            pdInfoDao.IsSoldout,
		IsRemoved:            pdInfoDao.IsRemoved,
	}

}

func Update(pdInfo *domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error) {
	pdInfo.Updated = time.Now()
	updatedPdInfo, err := ioc.Repo.ProductMetaInfos.Upsert(pdInfo)
	if err != nil {
		config.Logger.Error("error on adding product infos", zap.Error(err))
		return nil, err
	}

	pds, err := ioc.Repo.Products.ListByMetaID(pdInfo.ID.Hex())
	if err != nil {
		config.Logger.Error("error on listing products by product infos", zap.Error(err))
		return nil, err
	}

	for _, pd := range pds {
		pd.ProductInfo = updatedPdInfo
		_, err = ioc.Repo.Products.Upsert(pd)
		if err != nil {
			config.Logger.Error("error on updating products"+pd.ID.Hex(), zap.Error(err))
		}
		// from alloff 1.0 exhibition will not persist products data in product groups
		//exDao, _ := ioc.Repo.Exhibitions.Get(pd.ExhibitionID)
		//go exhibition.ExhibitionSyncer(exDao)
	}

	return updatedPdInfo, nil
}
