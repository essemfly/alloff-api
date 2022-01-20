package utils

import "github.com/lithammer/shortuuid/v3"

const (
	CODE_CHARSET = "346789ABCDEFGHJKLMNPQRTUVWXY"
	CODE_LENGTH  = 6
)

func CreateShortUUID() string {
	return shortuuid.NewWithAlphabet(CODE_CHARSET)[:CODE_LENGTH]
}
