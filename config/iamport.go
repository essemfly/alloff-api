package config

import (
	"log"

	"github.com/lessbutter/alloff-api/internal/pkg/iamport"
	"github.com/spf13/viper"
)

var PaymentService *iamport.Iamport

func InitIamPort() {
	iamport, err := iamport.NewIamport("https://api.iamport.kr", viper.GetString("IAMPORT_API_KEY"), viper.GetString("IAMPORT_SECRET_KEY"))
	if err != nil {
		log.Println(err)
	}

	PaymentService = iamport
}
