package config

var OmniousKey string

func InitOmnious(conf Configuration) {
	OmniousKey = conf.OMNIOUS_KEY
}
