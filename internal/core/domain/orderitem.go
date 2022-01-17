package domain

import (
	"errors"
	"time"

	"github.com/lessbutter/alloff-api/api/server/model"
)

type OrderItemStatusEnum string

const (
	ORDER_ITEM_CREATED                    = OrderItemStatusEnum("ORDER_ITEM_CREATED")
	ORDER_ITEM_RECREATED                  = OrderItemStatusEnum("ORDER_ITEM_RECREATED")
	ORDER_ITEM_PAYMENT_PENDING            = OrderItemStatusEnum("ORDER_ITEM_PAYMENT_PENDING")
	ORDER_ITEM_PAYMENT_FINISHED           = OrderItemStatusEnum("ORDER_ITEM_PAYMENT_FINISHED")
	ORDER_ITEM_PRODUCT_PREPARING          = OrderItemStatusEnum("ORDER_ITEM_PRODUCT_PREPARING")
	ORDER_ITEM_FOREIGN_PRODUCT_INSPECTING = OrderItemStatusEnum("ORDER_ITEM_FOREIGN_PRODUCT_INSPECTING")
	ORDER_ITEM_DELIVERY_PREPARING         = OrderItemStatusEnum("ORDER_ITEM_DELIVERY_PREPARING")
	ORDER_ITEM_FOREIGN_DELIVERY_STARTED   = OrderItemStatusEnum("ORDER_ITEM_FOREIGN_DELIVERY_STARTED")
	ORDER_ITEM_DELIVERY_STARTED           = OrderItemStatusEnum("ORDER_ITEM_DELIVERY_STARTED")
	ORDER_ITEM_DELIVERY_FINISHED          = OrderItemStatusEnum("ORDER_ITEM_DELIVERY_FINISHED")
	ORDER_ITEM_CONFIRM_PAYMENT            = OrderItemStatusEnum("ORDER_ITEM_CONFIRM_PAYMENT")
	ORDER_ITEM_CANCEL_FINISHED            = OrderItemStatusEnum("ORDER_ITEM_CANCEL_FINISHED")
	ORDER_ITEM_EXCHANGE_REQUESTED         = OrderItemStatusEnum("ORDER_ITEM_EXCHANGE_REQUESTED")
	ORDER_ITEM_EXCHANGE_PENDING           = OrderItemStatusEnum("ORDER_ITEM_EXCHANGE_PENDING")
	ORDER_ITEM_EXCHANGE_FINISHED          = OrderItemStatusEnum("ORDER_ITEM_EXCHANGE_FINISHED")
	ORDER_ITEM_RETURN_REQUESTED           = OrderItemStatusEnum("ORDER_ITEM_RETURN_REQUESTED")
	ORDER_ITEM_RETURN_PENDING             = OrderItemStatusEnum("ORDER_ITEM_RETURN_PENDING")
	ORDER_ITEM_RETURN_FINISHED            = OrderItemStatusEnum("ORDER_ITEM_RETURN_FINISHED")
)

type OrderItemTypeEnum string

const (
	NORMAL_ORDER     = OrderItemTypeEnum("NORMAL_ORDER")
	TIMEDEAL_ORDER   = OrderItemTypeEnum("TIMEDEAL_ORDER")
	EXHIBITION_ORDER = OrderItemTypeEnum("EXHIBITION_ORDER")
	UNKNOWN_ORDER    = OrderItemTypeEnum("UNKNOWN_ORDER")
)

type OrderItemDAO struct {
	tableName              struct{} `pg:"orderItems"`
	ID                     int
	OrderID                int
	OrderItemCode          string // Item 사서함번호
	ProductID              string
	ProductUrl             string
	ProductImg             string
	ProductName            string
	BrandKeyname           string
	BrandKorname           string
	Removed                bool
	SalesPrice             int
	CancelDescription      *CancelDescriptionDAO
	DeliveryDescription    *DeliveryDescriptionDAO
	OrderItemType          OrderItemTypeEnum
	OrderItemStatus        OrderItemStatusEnum
	DeliveryTrackingNumber string
	DeliveryTrackingUrl    string
	Size                   string
	Quantity               int
	RefundInfo             []*RefundItemDAO `pg:"rel:has-many"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
	OrderedAt              time.Time
	DeliveryStartedAt      time.Time
	DeliveryFinishedAt     time.Time
	CancelRequestedAt      time.Time
	CancelFinishedAt       time.Time
	ConfirmedAt            time.Time
}

func (orderItemDao *OrderItemDAO) ToDTO() *model.OrderItem {

	refundInfos := []*model.RefundInfo{}

	for _, item := range orderItemDao.RefundInfo {
		refundInfos = append(refundInfos, item.ToDTO())
	}

	return &model.OrderItem{
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
		CancelDescription:      orderItemDao.CancelDescription.ToDTO(),
		DeliveryDescription:    orderItemDao.DeliveryDescription.ToDTO(),
		DeliveryTrackingNumber: orderItemDao.DeliveryTrackingNumber,
		DeliveryTrackingURL:    orderItemDao.DeliveryTrackingUrl,
		RefundInfo:             refundInfos,
		CreatedAt:              orderItemDao.CreatedAt.String(),
		UpdatedAt:              orderItemDao.UpdatedAt.String(),
		OrderedAt:              orderItemDao.OrderedAt.String(),
		DeliveryStartedAt:      orderItemDao.DeliveryStartedAt.String(),
		DeliveryFinishedAt:     orderItemDao.DeliveryFinishedAt.String(),
		CancelRequestedAt:      orderItemDao.CancelRequestedAt.String(),
		CancelFinishedAt:       orderItemDao.CancelFinishedAt.String(),
		ConfirmedAt:            orderItemDao.ConfirmedAt.String(),
	}
}

func MapOrderItemStatus(enum OrderItemStatusEnum) model.OrderItemStatusEnum {
	switch enum {
	case ORDER_ITEM_PAYMENT_FINISHED:
		return model.OrderItemStatusEnumPaymentFinished
	case ORDER_ITEM_PRODUCT_PREPARING:
		return model.OrderItemStatusEnumProductPreparing
	case ORDER_ITEM_FOREIGN_PRODUCT_INSPECTING:
		return model.OrderItemStatusEnumForeignProductInspecting
	case ORDER_ITEM_DELIVERY_PREPARING:
		return model.OrderItemStatusEnumDeliveryPreparing
	case ORDER_ITEM_FOREIGN_DELIVERY_STARTED:
		return model.OrderItemStatusEnumForeignDeliveryStatrted
	case ORDER_ITEM_DELIVERY_STARTED:
		return model.OrderItemStatusEnumDeliveryStarted
	case ORDER_ITEM_DELIVERY_FINISHED:
		return model.OrderItemStatusEnumDeliveryFinished
	case ORDER_ITEM_CONFIRM_PAYMENT:
		return model.OrderItemStatusEnumConfirmPayment
	case ORDER_ITEM_CANCEL_FINISHED:
		return model.OrderItemStatusEnumCancelFinished
	case ORDER_ITEM_EXCHANGE_REQUESTED:
		return model.OrderItemStatusEnumExchangeRequested
	case ORDER_ITEM_EXCHANGE_PENDING:
		return model.OrderItemStatusEnumExchangePending
	case ORDER_ITEM_EXCHANGE_FINISHED:
		return model.OrderItemStatusEnumExchangeFinished
	case ORDER_ITEM_RETURN_REQUESTED:
		return model.OrderItemStatusEnumReturnRequested
	case ORDER_ITEM_RETURN_PENDING:
		return model.OrderItemStatusEnumReturnPending
	case ORDER_ITEM_RETURN_FINISHED:
		return model.OrderItemStatusEnumReturnFinished
	default:
		return model.OrderItemStatusEnumUnknown
	}
}

func MapOrderItemType(enum OrderItemTypeEnum) model.OrderItemTypeEnum {
	switch enum {
	case NORMAL_ORDER:
		return model.OrderItemTypeEnumNormal
	case TIMEDEAL_ORDER:
		return model.OrderItemTypeEnumTimedeal
	case EXHIBITION_ORDER:
		return model.OrderItemTypeEnumExhibition
	default:
		return model.OrderItemTypeEnumUnknown
	}
}

func (orderItemDao *OrderItemDAO) ConfirmOrder() error {
	if orderItemDao.OrderItemStatus == ORDER_ITEM_CONFIRM_PAYMENT {
		return errors.New("order already confirmed")
	}

	if orderItemDao.OrderItemStatus != ORDER_ITEM_DELIVERY_FINISHED {
		return errors.New("not available on order status for confirm")
	}

	orderItemDao.OrderItemStatus = ORDER_ITEM_CONFIRM_PAYMENT
	orderItemDao.ConfirmedAt = time.Now()
	orderItemDao.UpdatedAt = time.Now()

	return nil
}

func (orderItemDao *OrderItemDAO) CanCancelOrder() bool {
	if orderItemDao.OrderItemStatus == ORDER_ITEM_DELIVERY_PREPARING ||
		orderItemDao.OrderItemStatus == ORDER_ITEM_DELIVERY_STARTED ||
		orderItemDao.OrderItemStatus == ORDER_ITEM_DELIVERY_FINISHED {
		return true
	}
	return false
}

func (orderItemDao *OrderItemDAO) CanCancelPayment() bool {
	if orderItemDao.OrderItemStatus == ORDER_ITEM_PAYMENT_FINISHED || orderItemDao.OrderItemStatus == ORDER_ITEM_PRODUCT_PREPARING {
		return true
	}
	return false
}
