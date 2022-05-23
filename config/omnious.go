package config

import "github.com/spf13/viper"

var OmniousKey string

func InitOmnious() {
	OmniousKey = viper.GetString("OMNIOUS_KEY")
}
