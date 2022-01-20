package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
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

func (itemDao *RefundItemDAO) ToDTO() *model.RefundInfo {
	refundInfo := &model.RefundInfo{
		RefundFee:    itemDao.RefundFee,
		RefundAmount: itemDao.RefundAmount,
		CreatedAt:    itemDao.CreatedAt.String(),
		UpdatedAt:    itemDao.UpdatedAt.String(),
	}

	return refundInfo
}
