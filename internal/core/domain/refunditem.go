package domain

import (
	"time"
)

type RefundItemDAO struct {
	tableName    struct{} `pg:"refund_items"`
	ID           int
	OrderID      int       `pg:"order_id"`
	OrderItemID  int       `pg:"order_item_id"`
	RefundFee    int       `pg:"refund_fee"`
	RefundAmount int       `pg:"refund_amount"`
	CreatedAt    time.Time `pg:"created_at"`
	UpdatedAt    time.Time `pg:"updated_at`
}
