package malls

import (
	"fmt"
	"testing"
)

func TestGetTheoryDetail(t *testing.T) {
	productUrl := "https://outlet.theory.com/power-jkt-r/L081103R_G0F.html"
	imageUrls, sizes, colors, inventories, description := getTheoryDetail(productUrl, "L081103R")
	fmt.Println(description)
	fmt.Println(imageUrls)
	fmt.Println(sizes)
	fmt.Println(colors)
	fmt.Println(inventories)
}
