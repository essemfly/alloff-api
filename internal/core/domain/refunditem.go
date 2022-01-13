package domain

import "time"

type RefundItemDAO struct {
	tableName           struct{} `pg:"refundItems"`
	ID                  int
	OrderItemID         int
	RefundDeliveryPrice int
	RefundPrice         int
	RefundAmount        int
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
