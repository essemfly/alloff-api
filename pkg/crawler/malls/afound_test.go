package malls

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAfoundCrawler(t *testing.T) {
	t.Run("test getAfoundDetail", func(t *testing.T) {

		productUrl := "https://www.afound.com/de-de/produkte/cropped-teddy-bomber-jacket-dark-khaki-green_9fae92hb"
		images, colors, _, _, notSale := getAfoundDetail(productUrl)
		if !notSale {
			require.Equal(t, []string{"Grün"}, colors)
			require.NotEqual(t, 0, len(images))
		}
	})
}
