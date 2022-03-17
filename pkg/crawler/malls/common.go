package malls

import (
	"math"
	"regexp"
	"strconv"
	"strings"

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

func parseEuro(s string) float32 {
	var p float64
	p = 0.0

	priceText := strings.ReplaceAll(s, "€", "")
	trimRe := regexp.MustCompile(`\s*`)
	priceText = trimRe.ReplaceAllString(priceText, "")

	if strings.Contains(priceText, ",") {
		decimals := strings.Split(priceText, ",")[1]
		d, _ := strconv.ParseFloat(decimals, 32)
		p += d * math.Pow(0.1, float64(len(decimals)))
	}

	integer_part := strings.Split(priceText, ",")[0]
	integer_part = strings.ReplaceAll(integer_part, ".", "")
	i, _ := strconv.ParseFloat(integer_part, 32)
	p += i
	return float32(p)
}
