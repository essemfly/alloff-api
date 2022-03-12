package malls

import (
	"github.com/gocolly/colly"
)

const (
	defaultStock   = 100
	collyUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11"
)

func getCollyCollector(allowedDomain string) *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains(allowedDomain),
		colly.UserAgent(collyUserAgent),
	)
	return c
}
