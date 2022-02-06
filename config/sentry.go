package config

import (
	"log"

	"github.com/getsentry/sentry-go"
)

func InitSentry(conf Configuration) {
	if conf.ENVIRONMENT == "prod" {
		err := sentry.Init(sentry.ClientOptions{
			Dsn: "https://3bb7aa9c71b44397928e0101ebfecef2@o306501.ingest.sentry.io/6011225",
		})
		if err != nil {
			log.Fatalf("sentry.Init: %s", err)
		}
	}
}
