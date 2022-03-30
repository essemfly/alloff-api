package scripts

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/pkg/amplitude"
)

func MakeOrdersIntoAmplitude() {
	orders, err := ioc.Repo.Orders.ListAllPaid()
	if err != nil {
		log.Println("order listing err", err)
	}
	for _, order := range orders {
		orderDao, paymentDao, err := ioc.Service.OrderWithPaymentService.Find(order.AlloffOrderID)
		if err != nil {
			log.Println("order find error", err, order.ID)
			continue
		}

		amplitude.LogOrderRecord(order.User, orderDao, paymentDao)
	}

	// gracefully shutdown, waiting pending events to be sent
	err = config.AmplitudeClient.Shutdown()
	if err != nil {
		log.Println("Order client start error", err)
	}
}

func MakeCancelOrderItemsIntoAmplitude() {
	orderItems, err := ioc.Repo.OrderItems.ListAllCanceled()
	if err != nil {
		log.Println("order item listing err", err)
	}
	for _, item := range orderItems {
		order, err := ioc.Repo.Orders.Get(item.OrderID)
		if err != nil {
			log.Println("order load failed", err)
		}
		_, paymentDao, err := ioc.Service.OrderWithPaymentService.Find(order.AlloffOrderID)
		if err != nil {
			log.Println("order payment load failed")
		}
		amplitude.LogCancelOrderItemRecord(order.User, item, paymentDao)
	}
}
