package amplitude

import (
	"strings"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/renatoaf/amplitude-go/amplitude/data"
)

func LogOrderRecord(user *domain.UserDAO, order *domain.OrderDAO, payment *domain.PaymentDAO) {
	config.AmplitudeClient.Start()

	addressSlice := strings.Split(payment.BuyerAddress, " ")
	region := payment.BuyerAddress
	city := payment.BuyerAddress
	if len(addressSlice) >= 2 {
		region = addressSlice[0]
		city = addressSlice[0] + addressSlice[1]
	}

	orderEvent := LogOrder(user, order, payment)
	orderEvent.Region = region
	orderEvent.City = city
	config.AmplitudeClient.LogEvent(
		orderEvent,
	)

	for _, item := range order.OrderItems {
		orderItemEvent := LogOrderItem(user, item)
		orderItemEvent.Region = region
		orderItemEvent.City = city
		config.AmplitudeClient.LogEvent(
			orderItemEvent,
		)
	}

	// gracefully shutdown, waiting pending events to be sent
	config.AmplitudeClient.Shutdown()
}

func LogOrder(user *domain.UserDAO, order *domain.OrderDAO, payment *domain.PaymentDAO) *data.Event {
	return &data.Event{
		DeviceID:  user.Uuid,
		EventType: "[Server]OrderCreation",
		EventProperties: map[string]interface{}{
			"order":   order,
			"payment": payment,
		},
		UserProperties: map[string]interface{}{
			"mobile": user.Mobile,
		},
	}
}

func LogOrderItem(user *domain.UserDAO, item *domain.OrderItemDAO) *data.Event {
	return &data.Event{
		DeviceID:  user.Uuid,
		EventType: "[Server]OrderItemCreation",
		EventProperties: map[string]interface{}{
			"orderitem": item,
		},
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

	cancelOrderItemEvent := LogCancelOrderItem(user, orderItem)
	cancelOrderItemEvent.Region = region
	cancelOrderItemEvent.City = city
	config.AmplitudeClient.LogEvent(
		cancelOrderItemEvent,
	)

	// gracefully shutdown, waiting pending events to be sent
	config.AmplitudeClient.Shutdown()
}

func LogCancelOrderItem(user *domain.UserDAO, item *domain.OrderItemDAO) *data.Event {
	return &data.Event{
		DeviceID:  user.Uuid,
		EventType: "[Server]CancelOrderItem",
		EventProperties: map[string]interface{}{
			"orderitem": item,
		},
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
