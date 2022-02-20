package product

import "github.com/lessbutter/alloff-api/internal/core/domain"

const (
	EURO_EXCHANGE_RATE   = 1380
	DOLLAR_EXCHANGE_RATE = 1200
)

func GetProductPrice(pd *domain.ProductDAO) (int, int) {
	origPrice, discPrice := pd.ProductInfo.Price.OriginalPrice, pd.ProductInfo.Price.CurrentPrice

	if pd.ProductInfo.Price.CurrencyType == domain.CurrencyEUR {
		origPrice *= EURO_EXCHANGE_RATE
		discPrice *= EURO_EXCHANGE_RATE
	} else if pd.ProductInfo.Price.CurrencyType == domain.CurrencyUSD {
		origPrice *= DOLLAR_EXCHANGE_RATE
		discPrice *= DOLLAR_EXCHANGE_RATE
	}

	if pd.ProductInfo.Source.PriceMarginPolicy == "INTREND" {
		origPriceKRW := CalculateIntrendPrice(int(origPrice))
		discPriceKRW := CalculateIntrendPrice(int(discPrice))
		return origPriceKRW, discPriceKRW
	}

	return int(origPrice), int(discPrice)
}

func CalculateIntrendPrice(priceKRW int) int {
	priceKRW *= 89 / 100 //  뉴 공급가
	if (priceKRW / DOLLAR_EXCHANGE_RATE) > 150 {
		priceKRW *= 11 / 10
	} // 관세 포함 공급가
	priceKRW += 16000   // 해외 배송비 추가
	priceKRW *= 11 / 10 // 마진
	priceKRW += 3000    // 국내 배송비
	priceKRW += 13000   // 사업자 통관

	return priceKRW
}
