package productinfo

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/exhibition"
	"go.uber.org/zap"
)

func UpdateProductInfo(pdInfo *domain.ProductMetaInfoDAO, request *AddMetaInfoRequest) (*domain.ProductMetaInfoDAO, error) {
	//newPdInfo := makeBaseProductInfo(request)
	//newPdInfo.ID = pdInfo.ID

	inventories := AssignAlloffSizesToInventories(request.Inventory, pdInfo.ProductType, pdInfo.AlloffCategory)
	pdInfo.SetInventory(inventories)
	pdInfo.SetInformation(request.Information)

	if pdInfo.IsTranslateRequired {
		translated, err := TranslateProductInfo(pdInfo)
		if err != nil {
			config.Logger.Error("err occurred on translate product info : ", zap.Error(err))
		}
		if translated != nil {
			pdInfo.IsTranslateRequired = false
			pdInfo = translated
		}
	}

	updatedPdInfo, err := Update(pdInfo)
	if err != nil {
		config.Logger.Error("error on adding product infos", zap.Error(err))
		return nil, err
	}

	return updatedPdInfo, nil
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
		Information:          pdInfoDao.SalesInstruction.Information,
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
		exDao, _ := ioc.Repo.Exhibitions.Get(pd.ExhibitionID)
		go exhibition.ExhibitionSyncer(exDao)
	}

	return updatedPdInfo, nil
}
