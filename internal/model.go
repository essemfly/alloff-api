package internal

import "time"

type CurrencyType string

type SizeGuide struct {
	Label  string
	ImgUrl string
}

type Category struct {
	KeyName       string
	CatIdentifier string
}

const (
	CurrencyKRW CurrencyType = "KRW"
	CurrencyUSD CurrencyType = "USD"
)

type Price struct {
	OriginalPrice int
	CurrencyType  CurrencyType
	SellersPrice  int
}

type Brand struct {
	KorName    string
	KeyName    string
	Category   []*Category
	Modulename string
	Created    time.Time
}

type Product struct {
	Name   string
	Images []string
	Brand  []Brand
	Price  Price
}

type ProductSource struct {
	Name    string
	MainUrl string
}
