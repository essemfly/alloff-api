package domain

import "time"

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
	tableName              struct{} `pg:"orders"`
	ID                     int
	User                   *UserDAO
	OrderStatus            OrderStatusEnum
	OrderItems             []*OrderItemDAO
	CancelOrderItems       []*OrderItemDAO
	TotalPrice             int
	ProductPrice           int
	DeliveryPrice          int
	DeliveryTrackingNumber []string
	DeliveryTrackingUrl    []string
	UserMemo               string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	OrderedAt              time.Time
	DeliveryStartedAt      time.Time
	DeliveryFinishedAt     time.Time
	CancelRequestedAt      time.Time
	CancelFinishedAt       time.Time
	ConfirmedAt            time.Time
}

type OrderItemDAO struct {
	tableName           struct{} `pg:"orderItems"`
	ID                  int
	OrderID             *OrderDAO `pg:"rel:has-one"`
	OrderItemCode       string
	ProductID           string
	ProductName         string
	BrandKeyname        string
	SalesPrice          int
	Instruction         []string
	CancelDescription   []string
	DeliveryDescription []string
	OrderType           OrderTypeEnum
	OrderStatus         OrderStatusEnum
	Size                string
	Quantity            int
	Created             time.Time
	Updated             time.Time
	OrderedAt           time.Time
	DeliveryStartedAt   time.Time
	DeliveryFinishedAt  time.Time
	CancelRequestedAt   time.Time
	CancelFinishedAt    time.Time
	ConfirmedAt         time.Time
}
