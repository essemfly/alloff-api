package domain

import (
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
)

type OrderStatusEnum string

const (
	ORDER_CREATED          = OrderStatusEnum("ORDER_CREATED")
	ORDER_RECREATED        = OrderStatusEnum("ORDER_RECREATED")
	ORDER_PAYMENT_PENDING  = OrderStatusEnum("ORDER_PAYMENT_PENDING")
	ORDER_PAYMENT_FINISHED = OrderStatusEnum("ORDER_PAYMENT_FINISHED")
)

type OrderDAO struct {
	tableName     struct{} `pg:"orders"`
	ID            int
	AlloffOrderID string          `pg:"alloff_order_id"`
	UserID        string          `pg:"user_id"`
	User          *UserDAO        `pg:"user"`
	OrderStatus   OrderStatusEnum `pg:"order_status"`
	OrderItems    []*OrderItemDAO `pg:"rel:has-many"`
	TotalPrice    int             `pg:"total_price"`
	ProductPrice  int             `pg:"product_price"`
	DeliveryPrice int             `pg:"delivery_price"`
	RefundPrice   int             `pg:"refund_price"`
	UserMemo      string          `pg:"user_memo"`
	CreatedAt     time.Time       `pg:"created_at"`
	UpdatedAt     time.Time       `pg:"updated_at"`
	OrderedAt     time.Time       `pg:"ordered_at"`
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

func (orderDao *OrderDAO) GetOrderItem(productID string) *OrderItemDAO {
	for _, item := range orderDao.OrderItems {
		if item.ProductID == productID {
			return item
		}
	}
	return nil
}

func (orderDao *OrderDAO) GetOrderItemByID(orderItemID int) *OrderItemDAO {
	for _, item := range orderDao.OrderItems {
		if item.ID == orderItemID {
			return item
		}
	}
	return nil
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
	default:
		return model.OrderStatusEnumUnknown
	}
}
