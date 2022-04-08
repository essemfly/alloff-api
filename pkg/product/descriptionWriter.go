package product

import "github.com/lessbutter/alloff-api/internal/core/domain"

func GetProductDescription(pd *domain.ProductDAO, source *domain.CrawlSourceDAO) *domain.AlloffInstructionDAO {
	deliveryType := domain.Domestic
	deliveryTexts := []string{
		"도착 예정일은 택배사의 사정이나 주문량에 따라 변동될 수 있습니다.",
		"브랜드 및 제품에 따라 입점 업체(브랜드) 배송과 올오프 자체 배송으로 나뉩니다.",
	}

	if source.IsForeignDelivery {
		deliveryType = domain.Foreign
		deliveryTexts = []string{
			"도착 예정일은 현지 택배사의 사정이나 통관 과정에서 변동될 수 있습니다.",
			"배송기간에 현지 및 한국의 공휴일, 연말이 포함된 경우 배송이 지연될 수 있습니다.",
		}
	}

	descImages := pd.ProductInfo.Images
	if source.CrawlModuleName == "intrend" {
		descImages = append([]string{
			"https://alloff.s3.ap-northeast-2.amazonaws.com/description/Intrend_info.png",
		}, descImages...)
	}
	if source.CrawlModuleName == "theoutnet" || source.CrawlModuleName == "sandro" || source.CrawlModuleName == "maje" || source.CrawlModuleName == "intrend" {
		descImages = append(descImages, "https://alloff.s3.ap-northeast-2.amazonaws.com/description/size_guide.png")
	}

	descriptionText := []string{}
	descriptionInfo := map[string]string{}
	if pd.SalesInstruction != nil {
		descriptionText = pd.SalesInstruction.Description.Texts
		descriptionInfo = pd.SalesInstruction.Description.Infos
	}

	return &domain.AlloffInstructionDAO{
		Description: &domain.ProductDescriptionDAO{
			Images: descImages,
			Texts:  descriptionText,
			Infos:  descriptionInfo,
		},
		DeliveryDescription: &domain.DeliveryDescriptionDAO{
			DeliveryType:         deliveryType,
			DeliveryFee:          source.DeliveryPrice,
			EarliestDeliveryDays: source.EarliestDeliveryDays,
			LatestDeliveryDays:   source.LatestDeliveryDays,
			Texts:                deliveryTexts,
		},
		CancelDescription: &domain.CancelDescriptionDAO{
			RefundAvailable: source.RefundAvailable,
			ChangeAvailable: source.ChangeAvailable,
			ChangeFee:       source.ChangeFee,
			RefundFee:       source.RefundFee,
		},
	}
}

func GetManualProductDescription(pd *domain.ProductDAO, request *ProductManualAddRequest) *domain.AlloffInstructionDAO {
	deliveryType := domain.Domestic
	deliveryTexts := []string{
		"도착 예정일은 택배사의 사정이나 주문량에 따라 변동될 수 있습니다.",
		"브랜드 및 제품에 따라 입점 업체(브랜드) 배송과 올오프 자체 배송으로 나뉩니다.",
	}

	if request.IsForeignDelivery {
		deliveryType = domain.Foreign
		deliveryTexts = []string{
			"도착 예정일은 현지 택배사의 사정이나 통관 과정에서 변동될 수 있습니다.",
			"배송기간에 현지 및 한국의 공휴일, 연말이 포함된 경우 배송이 지연될 수 있습니다.",
		}
	}

	descImages := append(pd.ProductInfo.Images, request.DescriptionImages...)
	return &domain.AlloffInstructionDAO{
		Description: &domain.ProductDescriptionDAO{
			Images: descImages,
			Texts:  request.Description,
			Infos:  request.DescriptionInfos,
		},
		DeliveryDescription: &domain.DeliveryDescriptionDAO{
			DeliveryType:         deliveryType,
			DeliveryFee:          0,
			EarliestDeliveryDays: request.EarliestDeliveryDays,
			LatestDeliveryDays:   request.LatestDeliveryDays,
			Texts:                deliveryTexts,
		},
		CancelDescription: &domain.CancelDescriptionDAO{
			RefundAvailable: request.IsRefundPossible,
			ChangeAvailable: request.IsRefundPossible,
			ChangeFee:       request.RefundFee,
			RefundFee:       request.RefundFee,
		},
	}
}
