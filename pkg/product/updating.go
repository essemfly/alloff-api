package product

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.uber.org/zap"
)

func UpdateProductInfo(pdInfo *domain.ProductMetaInfoDAO, request *AddMetaInfoRequest) (*domain.ProductMetaInfoDAO, error) {

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

	pdInfo.SetAlloffInventory(request.Inventory)
	// UpdateAlloffCategory 필요하다

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

	updatedPdInfo, err := ioc.Repo.ProductMetaInfos.Upsert(pdInfo)
	if err != nil {
		config.Logger.Error("error on adding product infos", zap.Error(err))
		return nil, err
	}

	return updatedPdInfo, nil

}
