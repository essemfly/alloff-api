package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type orderItemRepo struct {
	db *pg.DB
}

func (repo *orderItemRepo) GetByCode(code string) (*domain.OrderItemDAO, error) {
	orderItem := new(domain.OrderItemDAO)

	if err := repo.db.Model(orderItem).Where("order_item_code = ?", code).Select(); err != nil {
		return nil, err
	}

	if orderItem.OrderItemStatus == domain.ORDER_ITEM_RETURN_FINISHED || orderItem.OrderItemStatus == domain.ORDER_ITEM_CANCEL_FINISHED {
		refundInfo := new(domain.RefundItemDAO)
		if err := repo.db.Model(refundInfo).Where("order_item_id = ?", orderItem.ID).Order("id ASC").Select(); err != nil {
			return nil, err
		}
		orderItem.RefundInfo = *refundInfo
	}

	return orderItem, nil
}

func (repo *orderItemRepo) Update(orderItemDao *domain.OrderItemDAO) (*domain.OrderItemDAO, error) {
	_, err := repo.db.Model(orderItemDao).WherePK().Update()
	if err != nil {
		return orderItemDao, err
	}

	return orderItemDao, nil
}

func PostgresOrderItemRepo(conn *PostgresDB) repository.OrderItemsRepository {
	return &orderItemRepo{
		db: conn.db,
	}
}
