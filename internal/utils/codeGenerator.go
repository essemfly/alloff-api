package utils

import (
	"math/rand"
	"time"
)

func CreateShortUUID() string {
	rand.Seed(time.Now().UnixNano())
	CODE_CHARSET := []rune("346789ABCDEFGHJKLMNPQRTUVWXY")

	b := make([]rune, 6)
	for i := range b {
		b[i] = CODE_CHARSET[rand.Intn(len(CODE_CHARSET))]
	}
	return string(b)
}

func CreateMockMobile() string {
	rand.Seed(time.Now().UnixNano())
	MOBILE_CHARSET := []rune("0123456789")

	b := make([]rune, 7)
	for i := range b {
		b[i] = MOBILE_CHARSET[rand.Intn(len(MOBILE_CHARSET))]
	}
	return "0101" + string(b)
}