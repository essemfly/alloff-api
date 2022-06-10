package productinfo

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPriceCalculate(t *testing.T) {
	// expected data from https://docs.google.com/spreadsheets/d/1RREZMtodluk_tv17V3ISzRtgzUtDnZ0IMrOPdoTvj30/edit#gid=0
	t.Run("calculate theoutnet fashion", func(t *testing.T) {
		currencyType := domain.CurrencyUSD
		originalPrice := float32(520.00)
		discountPriceWithTax := float32(520.00)
		marginPolicy := "THEOUTNET"
		_, disTax := GetProductPrice(originalPrice, discountPriceWithTax, currencyType, marginPolicy)
		require.Equal(t, 923000, disTax)

		discountPriceWithoutTax := float32(158.00)
		_, disNoTax := GetProductPrice(originalPrice, discountPriceWithoutTax, currencyType, marginPolicy)
		require.Equal(t, 240000, disNoTax)
	})

	t.Run("calculate theoutnet nonfashion", func(t *testing.T) {
		currencyType := domain.CurrencyUSD
		originalPrice := float32(500.00)
		discountPriceWithTax := float32(82.00)
		marginPolicy := "THEOUTNET_NON_FASHION"
		_, disNoTax := GetProductPrice(originalPrice, discountPriceWithTax, currencyType, marginPolicy)
		require.Equal(t, 134000, disNoTax)

		discountPriceWithoutTax := float32(540.00)
		_, disTax := GetProductPrice(originalPrice, discountPriceWithoutTax, currencyType, marginPolicy)
		require.Equal(t, 917000, disTax)
	})

	t.Run("calculate flannels", func(t *testing.T) {
		currencyType := domain.CurrencyPOUND
		originalPrice := float32(105.00)
		discountPrice := float32(105.00)
		marginPolicy := "FLANNELS"
		_, disPrice := GetProductPrice(originalPrice, discountPrice, currencyType, marginPolicy)
		require.Equal(t, 214000, disPrice)
	})

	t.Run("calculate afound", func(t *testing.T) {
		currencyType := domain.CurrencyEUR
		originalPrice := float32(35.00)
		discountPrice := float32(35.00)
		marginPolicy := "AFOUND"
		_, disPrice := GetProductPrice(originalPrice, discountPrice, currencyType, marginPolicy)
		require.Equal(t, 71000, disPrice)
	})
}
