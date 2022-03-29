package malls

import (
	"fmt"
	"testing"
)

func TestGetTheoryDetail(t *testing.T) {
	productUrl := "https://outlet.theory.com/line-dtl-dress/L096715R_QWG.html"
	imageUrls, sizes, colors, inventories, description := getTheoryDetail(productUrl, "L096715R")
	fmt.Println(description["설명"])
	fmt.Println(description["핏"])
	fmt.Println(description["제품 주 소재"])
	fmt.Println(description["취급시 주의사항"])
	fmt.Println(imageUrls)
	fmt.Println(sizes)
	fmt.Println(colors)
	fmt.Println(inventories)
}
