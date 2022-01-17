package config

import (
	"log"

	"github.com/lessbutter/alloff-api/internal/utils"
)

var OnWriteSlackMessage bool

func InitSlack(conf Configuration) {
	OnWriteSlackMessage = false
	if conf.ENVIRONMENT == "prod" {
		OnWriteSlackMessage = true
	}
}

func WriteSlackMessage(message string) {
	if !OnWriteSlackMessage {
		return
	}

	headers := map[string]string{
		"accept":       "*/*",
		"content-type": "application/json",
		"connection":   "keep-alive",
		"user-agent":   "Crawler",
	}

	url := "https://hooks.slack.com/services/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"

	bodymsg := `{"text": "` + message + `"}`
	_, err := utils.MakeRequest(url, utils.REQUEST_POST, headers, bodymsg)
	if err != nil {
		log.Println("Error occured in writing slack")
	}
}

func WriteOrderMessage(message string) {
	if !OnWriteSlackMessage {
		return
	}
	headers := map[string]string{
		"accept":       "*/*",
		"content-type": "application/json",
		"connection":   "keep-alive",
		"user-agent":   "Crawler",
	}
	url := "https://hooks.slack.com/services/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"

	bodymsg := `{"text": "` + message + `"}`
	_, err := utils.MakeRequest(url, utils.REQUEST_POST, headers, bodymsg)
	if err != nil {
		log.Println("Error occured in writing order")
	}
}

func WriteCancelMessage(message string) {
	if !OnWriteSlackMessage {
		return
	}
	headers := map[string]string{
		"accept":       "*/*",
		"content-type": "application/json",
		"connection":   "keep-alive",
		"user-agent":   "Crawler",
	}
	url := "https://hooks.slack.com/services/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"

	bodymsg := `{"text": "` + message + `"}`
	_, err := utils.MakeRequest(url, utils.REQUEST_POST, headers, bodymsg)
	if err != nil {
		log.Println("Error occured in writing order")
	}
}
