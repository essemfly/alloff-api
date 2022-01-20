package config

var NotificationUrl string
var NavigateUrl string

func InitNotification(conf Configuration) {
	NotificationUrl = conf.PUSH_SERVER_URL
	NavigateUrl = conf.PUSH_NAVIGATE_URL
}
