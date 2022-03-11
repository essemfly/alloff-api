package mapper

import (
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapOrder(orderDao *domain.OrderDAO) *model.OrderInfo {
	orderItems := []*model.OrderItem{}
	for _, itemDao := range orderDao.OrderItems {
		orderItems = append(orderItems, MapOrderItem(itemDao))
	}

	orderInfo := &model.OrderInfo{
		ID:            orderDao.AlloffOrderID,
		Orders:        orderItems,
		ProductPrice:  orderDao.ProductPrice,
		DeliveryPrice: orderDao.DeliveryPrice,
		TotalPrice:    orderDao.TotalPrice,
		RefundPrice:   &orderDao.RefundPrice,
		UserMemo:      orderDao.UserMemo,
		CreatedAt:     orderDao.CreatedAt.Add(9 * time.Hour).String(),
		UpdatedAt:     orderDao.UpdatedAt.Add(9 * time.Hour).String(),
		OrderedAt:     orderDao.OrderedAt.Add(9 * time.Hour).String(),
	}

	return orderInfo
}

func MapOrderItem(orderItemDao *domain.OrderItemDAO) *model.OrderItem {
	trackingNumber := ""
	trackingUrl := ""

	if len(orderItemDao.DeliveryTrackingNumber) > 0 {
		trackingNumber = orderItemDao.DeliveryTrackingNumber[len(orderItemDao.DeliveryTrackingNumber)-1]
		trackingUrl = orderItemDao.DeliveryTrackingUrl[len(orderItemDao.DeliveryTrackingUrl)-1]
	}
	return &model.OrderItem{
		ID:                     strconv.Itoa(orderItemDao.ID),
		ProductID:              orderItemDao.ProductID,
		ProductName:            orderItemDao.ProductName,
		ProductImg:             orderItemDao.ProductImg,
		BrandKeyname:           orderItemDao.BrandKeyname,
		BrandKorname:           orderItemDao.BrandKorname,
		Removed:                orderItemDao.Removed,
		SalesPrice:             orderItemDao.SalesPrice,
		Selectsize:             orderItemDao.Size,
		Quantity:               orderItemDao.Quantity,
		OrderItemType:          MapOrderItemType(orderItemDao.OrderItemType),
		OrderItemStatus:        MapOrderItemStatus(orderItemDao.OrderItemStatus),
		CancelDescription:      MapCancelDescription(orderItemDao.CancelDescription),
		DeliveryDescription:    MapDeliveryDescription(orderItemDao.DeliveryDescription),
		DeliveryTrackingNumber: trackingNumber,
		DeliveryTrackingURL:    trackingUrl,
		RefundInfo:             MapRefund(orderItemDao.RefundInfo),
		CreatedAt:              orderItemDao.CreatedAt.Add(9 * time.Hour).String(),
		UpdatedAt:              orderItemDao.UpdatedAt.Add(9 * time.Hour).String(),
		OrderedAt:              orderItemDao.OrderedAt.Add(9 * time.Hour).String(),
		DeliveryStartedAt:      orderItemDao.DeliveryStartedAt.Add(9 * time.Hour).String(),
		DeliveryFinishedAt:     orderItemDao.DeliveryFinishedAt.Add(9 * time.Hour).String(),
		CancelRequestedAt:      orderItemDao.CancelRequestedAt.Add(9 * time.Hour).String(),
		CancelFinishedAt:       orderItemDao.CancelFinishedAt.Add(9 * time.Hour).String(),
		ConfirmedAt:            orderItemDao.ConfirmedAt.Add(9 * time.Hour).String(),
	}
}

func MapOrderItemStatus(enum domain.OrderItemStatusEnum) model.OrderItemStatusEnum {
	switch enum {
	case domain.ORDER_ITEM_PAYMENT_FINISHED:
		return model.OrderItemStatusEnumPaymentFinished
	case domain.ORDER_ITEM_PRODUCT_PREPARING:
		return model.OrderItemStatusEnumProductPreparing
	case domain.ORDER_ITEM_FOREIGN_PRODUCT_INSPECTING:
		return model.OrderItemStatusEnumForeignProductInspecting
	case domain.ORDER_ITEM_DELIVERY_PREPARING:
		return model.OrderItemStatusEnumDeliveryPreparing
	case domain.ORDER_ITEM_FOREIGN_DELIVERY_STARTED:
		return model.OrderItemStatusEnumForeignDeliveryStatrted
	case domain.ORDER_ITEM_DELIVERY_STARTED:
		return model.OrderItemStatusEnumDeliveryStarted
	case domain.ORDER_ITEM_DELIVERY_FINISHED:
		return model.OrderItemStatusEnumDeliveryFinished
	case domain.ORDER_ITEM_CONFIRM_PAYMENT:
		return model.OrderItemStatusEnumConfirmPayment
	case domain.ORDER_ITEM_CANCEL_FINISHED:
		return model.OrderItemStatusEnumCancelFinished
	case domain.ORDER_ITEM_EXCHANGE_REQUESTED:
		return model.OrderItemStatusEnumExchangeRequested
	case domain.ORDER_ITEM_EXCHANGE_PENDING:
		return model.OrderItemStatusEnumExchangePending
	case domain.ORDER_ITEM_EXCHANGE_FINISHED:
		return model.OrderItemStatusEnumExchangeFinished
	case domain.ORDER_ITEM_RETURN_REQUESTED:
		return model.OrderItemStatusEnumReturnRequested
	case domain.ORDER_ITEM_RETURN_PENDING:
		return model.OrderItemStatusEnumReturnPending
	case domain.ORDER_ITEM_RETURN_FINISHED:
		return model.OrderItemStatusEnumReturnFinished
	default:
		return model.OrderItemStatusEnumUnknown
	}
}

func MapOrderItemType(enum domain.OrderItemTypeEnum) model.OrderItemTypeEnum {
	switch enum {
	case domain.NORMAL_ORDER:
		return model.OrderItemTypeEnumNormal
	case domain.TIMEDEAL_ORDER:
		return model.OrderItemTypeEnumTimedeal
	case domain.EXHIBITION_ORDER:
		return model.OrderItemTypeEnumExhibition
	default:
		return model.OrderItemTypeEnumUnknown
	}
}

func MapPayment(paymentDao *domain.PaymentDAO) *model.PaymentInfo {
	return &model.PaymentInfo{
		Pg:            paymentDao.Pg,
		PayMethod:     paymentDao.PayMethod,
		Name:          paymentDao.Name,
		MerchantUID:   paymentDao.MerchantUid,
		Amount:        paymentDao.Amount,
		BuyerName:     paymentDao.BuyerName,
		BuyerMobile:   paymentDao.BuyerMobile,
		BuyerAddress:  paymentDao.BuyerAddress,
		BuyerPostCode: paymentDao.BuyerPostCode,
		Company:       paymentDao.Company,
		AppScheme:     paymentDao.AppScheme,
	}
}

func MapRefund(itemDao domain.RefundItemDAO) *model.RefundInfo {
	return &model.RefundInfo{
		RefundFee:    itemDao.RefundFee,
		RefundAmount: itemDao.RefundAmount,
		CreatedAt:    itemDao.CreatedAt.Add(9 * time.Hour).String(),
		UpdatedAt:    itemDao.UpdatedAt.Add(9 * time.Hour).String(),
	}
}
