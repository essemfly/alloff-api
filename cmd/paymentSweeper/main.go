package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/config/ioc"
)

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "prod"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	cmd.SetBaseConfig(Env)

	payments, err := ioc.Repo.Payments.ListHolding()
	if err != nil {
		log.Println("err on listing payments", err)
	}

	log.Println("length of pending", len(payments))
	for idx, paymentDao := range payments {
		log.Println("idx", idx, paymentDao.ID)
		if paymentDao.UpdatedAt.Before(time.Now().Add(time.Minute * -10)) {
			orderDao, err := ioc.Repo.Orders.GetByAlloffID(paymentDao.MerchantUid)
			if err != nil {
				log.Println("err on getting order", err)
			}
			err = ioc.Service.OrderWithPaymentService.CancelPayment(orderDao, paymentDao)
			if err != nil {
				log.Println("err on cancel payment", err)
			}
		}
	}
}
