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
