package config

import (
	"log"

	"github.com/renatoaf/amplitude-go/amplitude"
)

var AmplitudeClient *amplitude.Client

func InitAmplitude(conf Configuration) {
	AmplitudeClient = amplitude.NewDefaultClient(conf.AMPLITUDE_API_KEY)

	err := AmplitudeClient.Start()
	if err != nil {
		log.Println("Order client start error", err, AmplitudeClient.State())
	}
}

func Destroy() {
	err := AmplitudeClient.Shutdown()
	if err != nil {
		log.Println("Order client start error", err)
	}
}
