package config

import (
	"log"

	"github.com/renatoaf/amplitude-go/amplitude"
	"github.com/spf13/viper"
)

var AmplitudeClient *amplitude.Client

func InitAmplitude() {
	AmplitudeClient = amplitude.NewDefaultClient(viper.GetString("AMPLITUDE_API_KEY"))

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
