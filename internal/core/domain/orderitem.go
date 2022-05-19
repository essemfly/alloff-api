package domain

import (
	"errors"
	"time"
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
	GROUPDEAL_ORDER  = OrderItemTypeEnum("GROUPDEAL_ORDER")
	UNKNOWN_ORDER    = OrderItemTypeEnum("UNKNOWN_ORDER")
)

type OrderItemDAO struct {
	tableName              struct{} `pg:"order_items"`
	ID                     int
	OrderID                int                     `pg:"order_id"`
	OrderItemCode          string                  `pg:"order_item_code"`
	ProductID              string                  `pg:"product_id"`
	ProductUrl             string                  `pg:"product_url"`
	ProductImg             string                  `pg:"product_img"`
	ProductName            string                  `pg:"product_name"`
	BrandKeyname           string                  `pg:"brand_keyname"`
	BrandKorname           string                  `pg:"brand_korname"`
	Removed                bool                    `pg:"is_removed,use_zero"`
	SalesPrice             int                     `pg:"sales_price"`
	CancelDescription      *CancelDescriptionDAO   `pg:"cancel_description"`
	DeliveryDescription    *DeliveryDescriptionDAO `pg:"delivery_description"`
	OrderItemType          OrderItemTypeEnum       `pg:"order_item_type"`
	OrderItemStatus        OrderItemStatusEnum     `pg:"order_item_status"`
	DeliveryTrackingNumber []string                `pg:"tracking_number,array"`
	DeliveryTrackingUrl    []string                `pg:"tracking_url,array"`
	Size                   string                  `pg:"size"`
	Quantity               int                     `pg:"quantity"`
	RefundInfo             RefundItemDAO           `pg:"rel:has-one"`
	CreatedAt              time.Time               `pg:"created_at"`
	UpdatedAt              time.Time               `pg:"updated_at"`
	OrderedAt              time.Time               `pg:"ordered_at"`
	DeliveryStartedAt      time.Time               `pg:"delivery_started_at"`
	DeliveryFinishedAt     time.Time               `pg:"delivery_finished_at"`
	CancelRequestedAt      time.Time               `pg:"cancel_requested_at"`
	CancelFinishedAt       time.Time               `pg:"cancel_finished_at"`
	ConfirmedAt            time.Time               `pg:"confirmed_at"`
	ExhibitionID           string                  `pg:"exhibition_id"`
	CompanyKeyname         string                  `pg:"company_keyname"`
	UserID                 string                  `pg:"user_id"`
	User                   *UserDAO                `pg:"user"`
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
