package product

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
)

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
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateIntrendPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	} else if pd.ProductInfo.Source.PriceMarginPolicy == "THEOUTNET" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateTheoutnetPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	} else if pd.ProductInfo.Source.PriceMarginPolicy == "SANDRO" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateSandroPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	} else if pd.ProductInfo.Source.PriceMarginPolicy == "MAJU" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateMajuPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	}

	return int(origPrice), int(discPrice)
}

func CalculateIntrendPrice(priceKRW int) int {
	priceKRW = priceKRW * 125 / 100 // 수수료
	priceKRW += 40000               // 해외배송비
	priceKRW = priceKRW * 11 / 10   // 마진
	priceKRW = priceKRW + 3000

	priceKRW = priceKRW / 1000
	priceKRW = priceKRW * 1000

	return priceKRW
}

func CalculateTheoutnetPrice(priceKRW int) int {
	luxuryProduct := false
	if (priceKRW / DOLLAR_EXCHANGE_RATE) >= 150 {
		luxuryProduct = true
		priceKRW = priceKRW * 113 / 100
	}
	priceKRW += 25000

	if luxuryProduct {
		priceKRW = priceKRW * 11 / 10 // 통관 부과세
	}
	priceKRW = priceKRW * 11 / 10 // 마진
	priceKRW = priceKRW + 3000

	priceKRW = priceKRW / 1000
	priceKRW = priceKRW * 1000

	return priceKRW
}

func CalculateSandroPrice(priceKRW int) int {
	priceKRW = priceKRW * 100 / 119 // 뉴 공급가
	luxuryProduct := false
	if (priceKRW / DOLLAR_EXCHANGE_RATE) >= 150 {
		luxuryProduct = true
	}
	priceKRW = priceKRW * 109 / 100 // 수수료
	priceKRW = priceKRW + 16000
	if luxuryProduct {
		priceKRW = priceKRW * 11 / 10
	}

	priceKRW = priceKRW * 11 / 10 // 마진
	priceKRW = priceKRW + 3000

	priceKRW = priceKRW / 1000
	priceKRW = priceKRW * 1000

	return priceKRW
}

func CalculateMajuPrice(priceKRW int) int {
	priceKRW = priceKRW * 100 / 119 // 뉴 공급가
	luxuryProduct := false
	if (priceKRW / DOLLAR_EXCHANGE_RATE) >= 150 {
		luxuryProduct = true
	}
	priceKRW = priceKRW * 109 / 100 // 수수료
	priceKRW = priceKRW + 16000
	if luxuryProduct {
		priceKRW = priceKRW * 11 / 10
	}

	priceKRW = priceKRW * 11 / 10 // 마진
	priceKRW = priceKRW + 3000

	priceKRW = priceKRW / 1000
	priceKRW = priceKRW * 1000

	return priceKRW
}
