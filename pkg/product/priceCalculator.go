package product

import (
	"math"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
)

const (
	EURO_EXCHANGE_RATE   = 1350
	DOLLAR_EXCHANGE_RATE = 1200
)

func GetProductPrice(origPrice, discPrice float32, currencyType domain.CurrencyType, marginPolicy string) (int, int) {
	if currencyType == domain.CurrencyEUR {
		origPrice *= EURO_EXCHANGE_RATE
		discPrice *= EURO_EXCHANGE_RATE
	} else if currencyType == domain.CurrencyUSD {
		origPrice *= DOLLAR_EXCHANGE_RATE
		discPrice *= DOLLAR_EXCHANGE_RATE
	}

	if marginPolicy == "INTREND" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateIntrendPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	} else if marginPolicy == "THEOUTNET" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateTheoutnetPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	} else if marginPolicy == "SANDRO" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateSandroPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	} else if marginPolicy == "MAJE" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateMajuPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	} else if marginPolicy == "THEORY" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateTheoryPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	} else if marginPolicy == "CLAUDIEPIERLOT" {
		discountRate := utils.CalculateDiscountRate(int(origPrice), int(discPrice))
		discPriceKRW := CalculateClaudiePierlotPrice(int(discPrice))
		origPriceKRW := 100 * discPriceKRW / (100 - discountRate)
		origPriceKRW = origPriceKRW / 1000
		origPriceKRW = origPriceKRW * 1000
		return origPriceKRW, discPriceKRW
	}

	return int(origPrice), int(discPrice)
}

func CalculateIntrendPrice(priceKRW int) int {
	priceKRW = priceKRW + 35000
	priceKRW = priceKRW * 130 / 100
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
	customTax := 0
	luxuryProduct := false
	priceKRW = priceKRW * 100 / 119 // 뉴 공급가 ( 19% 환급금 적용)

	if (priceKRW / DOLLAR_EXCHANGE_RATE) >= 150 {
		luxuryProduct = true // 관부가세 부과 상품 대상 여부
	}

	// 관세금액은 수수료 9%가 부과되기 전 && 환급금 19%가 적용된 금액에 부과된다. -> 관세청이 딴지걸면 바꿔야할수 있음
	if luxuryProduct {
		customTax = priceKRW * 13 / 100 // 관세액
	}

	priceKRW = priceKRW * 109 / 100 // 수수료 9%

	if luxuryProduct {
		priceKRW = priceKRW + customTax // 관세 부과된 물품가액
		priceKRW += 10000               // 해외배송비
		priceKRW = priceKRW * 11 / 10   // 부가세 10%는 관세 및 현지 배송업체 납부 비용에 추가 10%가 붙는다.
	} else {
		priceKRW += 10000 // 해외배송비							// 관부가세가 없는 경우, 해외배송비 10000원만 붙는다.
	}

	priceKRW = priceKRW * 105 / 100 // 마진
	priceKRW = priceKRW + 3000      // 국내 배송비는 함격에 포함

	priceKRW = int(math.Round(float64(priceKRW)/1000)) * 1000 // 1000원단위 가격 반올림
	return priceKRW
}

func CalculateMajuPrice(priceKRW int) int {
	customTax := 0
	luxuryProduct := false
	priceKRW = priceKRW * 100 / 119 // 뉴 공급가 ( 19% 환급금 적용)

	if (priceKRW / DOLLAR_EXCHANGE_RATE) >= 150 {
		luxuryProduct = true // 관부가세 부과 상품 대상 여부
	}

	// 관세금액은 수수료 9%가 부과되기 전 && 환급금 19%가 적용된 금액에 부과된다. -> 관세청이 딴지걸면 바꿔야할수 있음
	if luxuryProduct {
		customTax = priceKRW * 13 / 100 // 관세액
	}

	priceKRW = priceKRW * 109 / 100 // 수수료 9%

	if luxuryProduct {
		priceKRW = priceKRW + customTax // 관세 부과된 물품가액
		priceKRW += 10000               // 해외배송비
		priceKRW = priceKRW * 11 / 10   // 부가세 10%는 관세 및 현지 배송업체 납부 비용에 추가 10%가 붙는다.
	} else {
		priceKRW += 10000 // 해외배송비							// 관부가세가 없는 경우, 해외배송비 10000원만 붙는다.
	}

	priceKRW = priceKRW * 105 / 100 // 마진
	priceKRW = priceKRW + 3000      // 국내 배송비는 함격에 포함

	priceKRW = int(math.Round(float64(priceKRW)/1000)) * 1000 // 1000원단위 가격 반올림
	return priceKRW
}

func CalculateTheoryPrice(priceKRW int) int {
	// 원가 + (원가가 200불 넘을 때 관세 13%) + 15000 해외 배송비 + (원가가 200불 넘을 때 부가세 10%) + 총 가격의 10% + 3000원

	luxuryProduct := false
	if (priceKRW / DOLLAR_EXCHANGE_RATE) >= 200 {
		luxuryProduct = true
		priceKRW = priceKRW * 113 / 100
	}
	priceKRW = priceKRW + 15000

	if luxuryProduct {
		priceKRW = priceKRW * 11 / 10
	}
	priceKRW = priceKRW * 11 / 10 // 마진
	priceKRW = priceKRW + 3000

	priceKRW = priceKRW / 1000
	priceKRW = priceKRW * 1000

	return priceKRW
}

func CalculateClaudiePierlotPrice(priceKRW int) int {
	customTax := 0
	luxuryProduct := false
	priceKRW = priceKRW * 100 / 119 // 뉴 공급가 ( 19% 환급금 적용)

	if (priceKRW / DOLLAR_EXCHANGE_RATE) >= 150 {
		luxuryProduct = true // 관부가세 부과 상품 대상 여부
	}

	// 관세금액은 수수료 9%가 부과되기 전 && 환급금 19%가 적용된 금액에 부과된다. -> 관세청이 딴지걸면 바꿔야할수 있음
	if luxuryProduct {
		customTax = priceKRW * 13 / 100 // 관세액
	}

	priceKRW = priceKRW * 109 / 100 // 수수료 9%

	if luxuryProduct {
		priceKRW = priceKRW + customTax // 관세 부과된 물품가액
		priceKRW += 10000               // 해외배송비
		priceKRW = priceKRW * 11 / 10   // 부가세 10%는 관세 및 현지 배송업체 납부 비용에 추가 10%가 붙는다.
	} else {
		priceKRW += 10000 // 해외배송비							// 관부가세가 없는 경우, 해외배송비 10000원만 붙는다.
	}

	priceKRW = priceKRW * 105 / 100 // 마진
	priceKRW = priceKRW + 3000      // 국내 배송비는 함격에 포함

	priceKRW = int(math.Round(float64(priceKRW)/1000)) * 1000 // 1000원단위 가격 반올림
	return priceKRW
}
