package utils

import (
	"strconv"
	"strings"
	"unicode"
)

func ParsePriceString(letter string) int {
	ret := ""
	for _, val := range letter {
		if unicode.IsDigit(val) {
			ret += string(val)
		}
	}
	retInt, _ := strconv.Atoi(ret)
	return retInt
}

func CalculateDiscountRate(originalPrice float32, salesPrice float32) int {
	return int(((originalPrice - salesPrice) / originalPrice) * 100)
}

func StandardizeSpaces(s string) []string {
	return strings.Fields(s)
}

func RemoveDuplicates(sizes []string) []string {
	var newSizes []string
	for _, size := range sizes {
		isThere := false
		for _, ns := range newSizes {
			if ns == string(size) {
				isThere = true
				break
			}
		}
		if !isThere {
			newSizes = append(newSizes, size)
		}
	}
	return newSizes
}
