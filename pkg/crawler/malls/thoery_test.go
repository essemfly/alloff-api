package malls

import (
	"fmt"
	"testing"
)

func TestGetTheoryDetail(t *testing.T) {
	productUrl := "https://outlet.theory.com/treeca-pull-on/M019206R_E0S.html"
	_, _, _, _, description := getTheoryDetail(productUrl, "M019206R")
	fmt.Println(description["핏"])
}
