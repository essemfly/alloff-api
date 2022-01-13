package crawler

import "github.com/lessbutter/alloff-api/internal/core/domain"

func GetProductDescription(pd *domain.ProductDAO) *domain.AlloffInstructionDAO {
	// (TODO) 임시로 넣어둠 + foreign 인지 아닌지 체크후에 더 나눠야함
	earlyDays, lateDays := 1, 5
	if pd.ProductInfo.Source.IsForeignDelivery {
		earlyDays, lateDays = 7, 14
	}

	return &domain.AlloffInstructionDAO{
		Description: &domain.ProductDescriptionDAO{
			Images: pd.ProductInfo.Images,
			Texts:  nil,
		},
		DeliveryDescription: &domain.DeliveryDescriptionDAO{
			DeliveryFee:          pd.ProductInfo.Source.DeliveryPrice,
			EarliestDeliveryDays: earlyDays,
			LatestDeliveryDays:   lateDays,
			Texts: []string{
				"결제완료 후 평균 3-4일 이내에 상품을 받아보실 수 있습니다.",
				"배송비는 각 상품 또는 유통사에 따라 다를 수 있으며, 주문/결제 페이지에서 정확한 금액을 확인하실 수 있습니다.",
			},
		},
		CancelDescription: &domain.CancelDescriptionDAO{
			RefundAvailable: false,
			ChangeAvailable: false,
			Texts: []string{
				"주문 취소는 주문 상태가 배송 전일 때 가능합니다. 마이페이지의 주문 내역 확인하기에서 주문 취소를 요청해주세요.",
				"부분 취소의 경우, 마이페이지의 주문 내역 확인하기에서 1:1 문의를 통해 요청해주세요.",
				"교환 및 반품을 원하시는 경우 상품 수령 후 5일 이내에 주문 취소를 요청해야 합니다. 마이페이지의 주문 내역 확인하기에서 주문 취소를 요청해주세요. 주문 취소를 요청해주시면 1:1 고객센터를 통해 빠르게 도와드리겠습니다.",
				"교환 및 반품 요청이 접수되면 영업일 기준 2-3일 내에 택배기사님이 연락 후 방문하여 회수가 진행됩니다.", "상품입고 후 1일의 검수 시간이 소요되며, 검수 완료 후 교환 및 반품 처리가 진행됩니다.",
				"상품 훼손 또는 외부 착용, 브랜드 택/박스 분실 및 훼손, 수령 후 5일이 지난 경우 교환 및 반품이 불가합니다.", "단순 변심에 의한 배송비(왕복 5,000원)는 고객부담입니다.",
			},
		},
	}
}
