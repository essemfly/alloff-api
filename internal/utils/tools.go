package utils

import (
	"math"
	"reflect"
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

func CalculateDiscountRate(originalPrice int, salesPrice int) int {
	if originalPrice == 0 {
		return 0
	}

	return int(math.Round((float64(originalPrice-salesPrice) / float64(originalPrice)) * 100))
}

func StandardizeSpaces(s string) []string {
	return strings.Fields(s)
}

func ItemExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func RemoveDuplicateString(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
