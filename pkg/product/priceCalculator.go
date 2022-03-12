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
		origPriceKRW = origPriceKRW * 100
		return origPriceKRW, discPriceKRW
	} else if pd.ProductInfo.Source.PriceMarginPolicy == "THEOUTNET" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateTheoutnetPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW * 100
		return origPriceKRW, discPriceKRW
	}

	return int(origPrice), int(discPrice)
}

func CalculateIntrendPrice(priceKRW int) int {
	if (priceKRW / DOLLAR_EXCHANGE_RATE) >= 150 {
		priceKRW += 16000
		priceKRW = priceKRW * 11 / 10
	} else {
		priceKRW += 16000
	}
	priceKRW = priceKRW * 11 / 10 // 마진
	priceKRW += 3000              // 국내 배송비
	priceKRW += 13000             // 사업자 통관

	remainder := priceKRW % 1000
	priceKRW = priceKRW / 1000
	if remainder >= 500 {
		priceKRW += 1
	}
	priceKRW = priceKRW * 1000

	return priceKRW
}

func CalculateTheoutnetPrice(priceKRW int) int {
	if (priceKRW / EURO_EXCHANGE_RATE) >= 150 {
		priceKRW += 25000
		priceKRW = priceKRW * 11 / 10
	} else {
		priceKRW += 25000
	}
	priceKRW = priceKRW * 11 / 10 // 마진
	priceKRW += 3000              // 국내 배송비

	remainder := priceKRW % 1000
	priceKRW = priceKRW / 1000
	if remainder >= 500 {
		priceKRW += 1
	}
	priceKRW = priceKRW * 1000

	return priceKRW
}
