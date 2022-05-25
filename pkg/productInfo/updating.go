package productinfo

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/alloffcategory"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// TODO: Crawling 다시 했을때 부분인데, Reset을 어떻게까지 정의해야할지 아직 좀 감이 안온다.
func Reset(pdInfo *domain.ProductMetaInfoDAO, request *AddMetaInfoRequest) (*domain.ProductMetaInfoDAO, error) {
	pdInfo.SetBrandAndCategory(request.Brand, request.Source)
	pdInfo.SetGeneralInfo(request.AlloffName, request.ProductID, request.ProductUrl, request.Images, request.Sizes, request.Colors, request.Information)
	alloffOrigPrice, alloffDiscPrice := GetProductPrice(float32(request.OriginalPrice), float32(request.DiscountedPrice), request.CurrencyType, request.Source.PriceMarginPolicy)

	// Price 크롤링되어서 오고 있는 것들을 수동으로 만약 바꾼 상태면, 이걸 다시 크롤링 돌면 크롤링 가격으로 바꿔?
	pdInfo.SetPrices(alloffOrigPrice, alloffDiscPrice, domain.CurrencyKRW)

	// Images들을 Intrend, theoutnet 등의 특별한 이미지들은 여기서는 업데이트 하지 않는다.
	descImages := append(request.DescriptionImages, request.Images...)
	pdInfo.SetDesc(descImages, request.Description, request.DescriptionInfos)
	pdInfo.SetDeliveryDesc(request.IsForeignDelivery, 0, request.EarliestDeliveryDays, request.LatestDeliveryDays)
	pdInfo.SetCancelDesc(request.IsRefundPossible, request.RefundFee)

	if request.AlloffCategory.ID != primitive.NilObjectID {
		productAlloffCat, err := alloffcategory.BuildProductAlloffCategory(request.AlloffCategory.ID.Hex(), true)
		if err != nil {
			config.Logger.Error("err occured on build product alloff category : alloffcat ID"+request.AlloffCategory.ID.Hex(), zap.Error(err))
		}
		pdInfo.SetAlloffCategory(productAlloffCat)
	}

	if request.AlloffCategory.ID == primitive.NilObjectID || !pdInfo.AlloffCategory.Done {
		productAlloffCat, err := alloffcategory.InferAlloffCategory(pdInfo)
		if err != nil {
			config.Logger.Error("err occured on infer alloffcategory: pdinfo "+pdInfo.ID.Hex(), zap.Error(err))
		}
		pdInfo.SetAlloffCategory(productAlloffCat)
	}

	pdInfo.SetAlloffInventory(request.Inventory)

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

	updatedPdInfo, err := Update(pdInfo)
	if err != nil {
		config.Logger.Error("error on adding product infos", zap.Error(err))
		return nil, err
	}

	return updatedPdInfo, nil

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
	}

	return updatedPdInfo, nil
}
