package config

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func InitSentry() {
	if viper.GetString("ENVIRONMENT") == "prod" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: "https://3bb7aa9c71b44397928e0101ebfecef2:dd53f2cf3eec49ea8e1ce7967f1fe3ce@o1042350.ingest.sentry.io/6011225",
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
	}
}
