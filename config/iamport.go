package config

import (
	"log"

	"github.com/lessbutter/alloff-api/internal/pkg/iamport"
)

var PaymentService *iamport.Iamport

func InitIamPort(conf Configuration) {
	iamport, err := iamport.NewIamport("https://api.iamport.kr", conf.IAMPORT_API_KEY, conf.IAMPORT_SECRET_KEY)
	if err != nil {
		log.Println(err)
	}

	PaymentService = iamport
}
