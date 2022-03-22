package config

import "github.com/renatoaf/amplitude-go/amplitude"

var AmplitudeClient *amplitude.Client

func InitAmplitude(conf Configuration) {
	AmplitudeClient = amplitude.NewDefaultClient(conf.AMPLITUDE_API_KEY)
}
