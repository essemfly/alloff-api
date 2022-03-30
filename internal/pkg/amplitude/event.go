package amplitude

import (
	"log"
	"strings"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/renatoaf/amplitude-go/amplitude/data"
)

func LogOrderRecord(user *domain.UserDAO, order *domain.OrderDAO, payment *domain.PaymentDAO) {

	addressSlice := strings.Split(payment.BuyerAddress, " ")
	region := payment.BuyerAddress
	city := payment.BuyerAddress
	if len(addressSlice) >= 2 {
		region = addressSlice[0]
		city = addressSlice[0] + addressSlice[1]
	}

	orderEvent := logOrder(user, order, payment)
	orderEvent.Region = region
	orderEvent.City = city
	err := config.AmplitudeClient.LogEvent(
		orderEvent,
	)
	if err != nil {
		log.Println("err occured on log order event", err, config.AmplitudeClient.State())
	}

	for _, item := range order.OrderItems {
		orderItemEvent := logOrderItem(user, item)
		orderItemEvent.Region = region
		orderItemEvent.City = city
		err := config.AmplitudeClient.LogEvent(
			orderItemEvent,
		)
		if err != nil {
			log.Println("err occured on log order item event", err)
		}
	}

}

func logOrder(user *domain.UserDAO, order *domain.OrderDAO, payment *domain.PaymentDAO) *data.Event {
	return &data.Event{
		DeviceID:  user.Uuid,
		EventType: "[Server]CreateOrder220330",
		EventProperties: map[string]interface{}{
			"order":   order,
			"payment": payment,
		},
		Time: order.OrderedAt.Unix(),
		UserProperties: map[string]interface{}{
			"mobile": user.Mobile,
		},
	}
}

func logOrderItem(user *domain.UserDAO, item *domain.OrderItemDAO) *data.Event {
	return &data.Event{
		DeviceID:  user.Uuid,
		EventType: "[Server]CreateOrderItem220330",
		EventProperties: map[string]interface{}{
			"orderitem": item,
		},
		Time:        item.OrderedAt.Unix(),
		Price:       float64(item.SalesPrice),
		Quantity:    int32(item.Quantity),
		Revenue:     float64(item.SalesPrice) * float64(item.Quantity),
		RevenueType: string(item.OrderItemType),
		ProductID:   item.ProductID,
		Carrier:     item.BrandKorname,
		UserProperties: map[string]interface{}{
			"mobile": user.Mobile,
		},
	}
}

func LogCancelOrderItemRecord(user *domain.UserDAO, orderItem *domain.OrderItemDAO, payment *domain.PaymentDAO) {
	config.AmplitudeClient.Start()

	addressSlice := strings.Split(payment.BuyerAddress, " ")
	region := payment.BuyerAddress
	city := payment.BuyerAddress
	if len(addressSlice) >= 2 {
		region = addressSlice[0]
		city = addressSlice[0] + addressSlice[1]
	}

	cancelOrderItemEvent := logCancelOrderItem(user, orderItem)
	cancelOrderItemEvent.Region = region
	cancelOrderItemEvent.City = city
	config.AmplitudeClient.LogEvent(
		cancelOrderItemEvent,
	)

	// gracefully shutdown, waiting pending events to be sent
	config.AmplitudeClient.Shutdown()
}

func logCancelOrderItem(user *domain.UserDAO, item *domain.OrderItemDAO) *data.Event {
	return &data.Event{
		DeviceID:  user.Uuid,
		EventType: "[Server]CancelOrderItem",
		EventProperties: map[string]interface{}{
			"orderitem": item,
		},
		Time:        item.CancelFinishedAt.Unix(),
		Price:       float64(item.SalesPrice),
		Quantity:    int32(item.Quantity),
		Revenue:     -1 * float64(item.SalesPrice) * float64(item.Quantity),
		RevenueType: string(item.OrderItemType),
		ProductID:   item.ProductID,
		Carrier:     item.BrandKorname,
		UserProperties: map[string]interface{}{
			"mobile": user.Mobile,
		},
	}
}
