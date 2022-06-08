package malls

import (
	"log"
	"testing"
)

func TestTheoutnet(t *testing.T) {
	t.Run("get theoutnet detail", func(t *testing.T) {
		productUrl := "/oscar-de-la-renta/dresses/knee-length-dress/one-shoulder-bow-embellished-silk-taffeta-dress/1647597283535844"
		composition, sizes, invs, desc, imgs := CrawlTheoutnetDetail(productUrl)
		log.Println(composition)
		log.Println(sizes)
		log.Println(invs)
		log.Println(desc)
		log.Println(imgs)

	})
}
