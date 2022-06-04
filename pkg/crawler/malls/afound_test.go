package malls

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAfoundCrawler(t *testing.T) {
	t.Run("test getAfoundDetail", func(t *testing.T) {

		productUrl := "https://www.afound.com/de-de/produkte/cropped-teddy-bomber-jacket-dark-khaki-green_9fae92hb"
		images, colors, infos, desc, notSale := getAfoundDetail(productUrl)
		if !notSale {
			require.Equal(t, []string{"Grün"}, colors)
			require.NotEqual(t, "", desc["사이즈 및 핏"])
			require.NotEqual(t, "", desc["제품설명"])
			require.NotEqual(t, "", infos["색상"])
			require.NotEqual(t, "", infos["소재"])
			require.NotEqual(t, 0, len(images))
		}
	})
}
