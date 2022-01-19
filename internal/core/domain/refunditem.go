package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
)

type RefundItemDAO struct {
	tableName    struct{} `pg:"refundItems"`
	ID           int
	OrderID      string
	OrderItemID  int
	RefundFee    int
	RefundAmount int
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
