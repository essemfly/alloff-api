package domain

import (
	"errors"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/api/server/model"
)

type OrderStatusEnum string

const (
	ORDER_CREATED            = OrderStatusEnum("CREATED")
	ORDER_RECREATED          = OrderStatusEnum("RECREATED")
	ORDER_PAYMENT_PENDING    = OrderStatusEnum("PAYMENT_PENDING")
	ORDER_PAYMENT_FINISHED   = OrderStatusEnum("PAYMENT_FINISHED")
	ORDER_PRODUCT_PREPARING  = OrderStatusEnum("PRODUCT_PREPARING")
	ORDER_DELIVERY_PREPARING = OrderStatusEnum("DELIVERY_PREPARING")
	ORDER_CANCEL_REQUESTED   = OrderStatusEnum("CANCEL_REQUESTED")
	ORDER_CANCEL_PENDING     = OrderStatusEnum("CANCEL_PENDING")
	ORDER_CANCEL_FINISHED    = OrderStatusEnum("CANCEL_FINISHED")
	ORDER_DELIVERY_STARTED   = OrderStatusEnum("DELIVERY_STARTED")
	ORDER_DELIVERY_FINISHED  = OrderStatusEnum("DELIVERY_FINISHED")
	ORDER_CONFIRM_PAYMENT    = OrderStatusEnum("CONFIRM_PAYMENT")
)

type OrderTypeEnum string

const (
	NORMAL_ORDER   = OrderTypeEnum("NORMAL_ORDER")
	TIMEDEAL_ORDER = OrderTypeEnum("TIMEDEAL_ORDER")
	UNKNOWN_ORDER  = OrderTypeEnum("UNKNOWN_ORDER")
)

type OrderDAO struct {
	tableName          struct{} `pg:"orders"`
	ID                 int
	AlloffOrderID      string
	User               *UserDAO
	OrderStatus        OrderStatusEnum
	OrderItems         []*OrderItemDAO `pg:"rel:has-many"`
	TotalPrice         int
	ProductPrice       int
	DeliveryPrice      int
	UserMemo           string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	OrderedAt          time.Time
	DeliveryStartedAt  time.Time
	DeliveryFinishedAt time.Time
	CancelRequestedAt  time.Time
	CancelFinishedAt   time.Time
	ConfirmedAt        time.Time
}

func (orderDao *OrderDAO) ToDTO() *model.OrderInfo {

	orderItems := []*model.OrderItem{}
	for _, orderDao := range orderDao.OrderItems {
		orderItems = append(orderItems, orderDao.ToDTO())
	}

	orderInfo := &model.OrderInfo{
		ID:            orderDao.AlloffOrderID,
		Orders:        orderItems,
		ProductPrice:  orderDao.ProductPrice,
		DeliveryPrice: orderDao.DeliveryPrice,
		TotalPrice:    orderDao.TotalPrice,
		UserMemo:      orderDao.UserMemo,
		CreatedAt:     orderDao.CreatedAt.String(),
		UpdatedAt:     orderDao.UpdatedAt.String(),
		OrderedAt:     orderDao.OrderedAt.String(),
		RefundInfo:    nil, // Refund 정책 구체화 후 모델링
	}

	return orderInfo
}

func (orderDao *OrderDAO) GetBasePayment() *PaymentDAO {
	return &PaymentDAO{
		Pg:            "danal_tpay",
		PayMethod:     "card",
		MerchantUid:   orderDao.AlloffOrderID,
		Amount:        orderDao.TotalPrice,
		Name:          orderDao.GetOrderName(),
		BuyerName:     orderDao.User.Name,
		BuyerMobile:   orderDao.User.Mobile,
		BuyerAddress:  orderDao.User.GetUserAddress(),
		BuyerPostCode: orderDao.User.Postcode,
		Company:       "alloff",
		AppScheme:     "",
	}
}

func (orderDao *OrderDAO) GetOrderName() string {
	if len(orderDao.OrderItems) == 1 {
		return orderDao.OrderItems[0].ProductName
	}

	return orderDao.OrderItems[0].ProductName + "외 " + strconv.Itoa(len(orderDao.OrderItems)-1) + " 건"
}

func (orderDao *OrderDAO) CancelOrder() error {
	orderDao.UpdatedAt = time.Now()
	orderDao.CancelRequestedAt = time.Now()
	if orderDao.OrderStatus == ORDER_PAYMENT_FINISHED || orderDao.OrderStatus == ORDER_PRODUCT_PREPARING {
		orderDao.OrderStatus = ORDER_CANCEL_FINISHED
	} else if orderDao.OrderStatus == ORDER_DELIVERY_PREPARING ||
		orderDao.OrderStatus == ORDER_DELIVERY_STARTED ||
		orderDao.OrderStatus == ORDER_DELIVERY_FINISHED {
		orderDao.OrderStatus = ORDER_CANCEL_REQUESTED
	} else {
		return errors.New("not availabe on order status for cancel")
	}

	return nil
}

func (orderDao *OrderDAO) GetOrderItem(productID string) *OrderItemDAO {
	for _, item := range orderDao.OrderItems {
		if item.ProductID == productID {
			return item
		}
	}
	return nil
}

func (orderDao *OrderDAO) ConfirmOrder() error {
	if orderDao.OrderStatus == ORDER_CONFIRM_PAYMENT {
		return errors.New("order already confirmed")
	}

	if orderDao.OrderStatus != ORDER_DELIVERY_FINISHED {
		return errors.New("not available on order status for confirm")
	}

	orderDao.OrderStatus = ORDER_CONFIRM_PAYMENT
	orderDao.UpdatedAt = time.Now()
	orderDao.ConfirmedAt = time.Now()

	return nil
}

func (orderDao *OrderDAO) CanCancelOrder() bool {
	if orderDao.OrderStatus == ORDER_DELIVERY_PREPARING ||
		orderDao.OrderStatus == ORDER_DELIVERY_STARTED ||
		orderDao.OrderStatus == ORDER_DELIVERY_FINISHED {
		return true
	}
	return false
}

func (orderDao *OrderDAO) CanCancelPayment() bool {
	if orderDao.OrderStatus == ORDER_PAYMENT_FINISHED || orderDao.OrderStatus == ORDER_PRODUCT_PREPARING {
		return true
	}
	return false
}

func (orderDao *OrderDAO) ValidateOrder() error {
	prices := 0
	for _, orderItem := range orderDao.OrderItems {
		// (TODO) Product가 Soldout이거나, removed인 경우는 어떻게 처리할 것인가?
		prices += orderItem.SalesPrice * orderItem.Quantity
	}
	return nil
}

type OrderItemDAO struct {
	tableName              struct{} `pg:"orderItems"`
	ID                     int
	OrderID                int
	OrderItemCode          string
	ProductID              string
	ProductImg             string
	ProductName            string
	BrandKeyname           string
	BrandKorname           string
	Removed                bool
	SalesPrice             int
	SizeDescription        []string
	CancelDescription      []string
	DeliveryDescription    []string
	OrderType              OrderTypeEnum
	OrderStatus            OrderStatusEnum
	DeliveryTrackingNumber string
	DeliveryTrackingUrl    string
	Size                   string
	Quantity               int
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
		OrderType:              MapOrderType(orderItemDao.OrderType),
		OrderStatus:            MapOrderStatus(orderItemDao.OrderStatus),
		SizeDescription:        orderItemDao.SizeDescription,
		CancelDescription:      orderItemDao.CancelDescription,
		DeliveryDescription:    orderItemDao.DeliveryDescription,
		DeliveryTrackingNumber: orderItemDao.DeliveryTrackingNumber,
		DeliveryTrackingURL:    orderItemDao.DeliveryTrackingUrl,
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

func MapOrderStatus(enum OrderStatusEnum) model.OrderStatusEnum {
	switch enum {
	case ORDER_CREATED:
		return model.OrderStatusEnumCreated
	case ORDER_RECREATED:
		return model.OrderStatusEnumRecreated
	case ORDER_PAYMENT_PENDING:
		return model.OrderStatusEnumPaymentPending
	case ORDER_PAYMENT_FINISHED:
		return model.OrderStatusEnumPaymentFinished
	case ORDER_PRODUCT_PREPARING:
		return model.OrderStatusEnumProductPreparing
	case ORDER_DELIVERY_PREPARING:
		return model.OrderStatusEnumDeliveryPreparing
	case ORDER_CANCEL_REQUESTED:
		return model.OrderStatusEnumCancelRequested
	case ORDER_CANCEL_PENDING:
		return model.OrderStatusEnumCancelPending
	case ORDER_CANCEL_FINISHED:
		return model.OrderStatusEnumCancelFinished
	case ORDER_DELIVERY_STARTED:
		return model.OrderStatusEnumDeliveryStarted
	case ORDER_DELIVERY_FINISHED:
		return model.OrderStatusEnumDeliveryFinished
	case ORDER_CONFIRM_PAYMENT:
		return model.OrderStatusEnumConfirmPayment
	default:
		return model.OrderStatusEnumUnknown
	}
}

func MapOrderType(enum OrderTypeEnum) model.OrderTypeEnum {
	switch enum {
	case NORMAL_ORDER:
		return model.OrderTypeEnumNormal
	case TIMEDEAL_ORDER:
		return model.OrderTypeEnumTimedeal
	default:
		return model.OrderTypeEnumUnknown
	}
}
