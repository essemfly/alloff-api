package main

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/config/ioc"
)

func main() {
	cmd.SetBaseConfig()

	payments, err := ioc.Repo.Payments.ListHolding()
	if err != nil {
		log.Println("err on listing payments", err)
	}

	log.Println("length of pending", len(payments))
	for _, paymentDao := range payments {
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
