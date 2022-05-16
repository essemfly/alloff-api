package config

import "github.com/spf13/viper"

var NotificationUrl string
var NavigateUrl string

func InitNotification() {
	NotificationUrl = viper.GetString("PUSH_SERVER_URL")
	NavigateUrl = viper.GetString("PUSH_NAVIGATE_URL")
}
