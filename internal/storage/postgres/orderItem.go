package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type orderItemRepo struct {
	db *pg.DB
}

func (repo *orderItemRepo) Get(ID int) (*domain.OrderItemDAO, error) {
	orderItem := new(domain.OrderItemDAO)

	if err := repo.db.Model(orderItem).Where("id = ?", ID).Select(); err != nil {
		return nil, err
	}

	return orderItem, nil
}

func (repo *orderItemRepo) GetByCode(code string) (*domain.OrderItemDAO, error) {
	orderItem := new(domain.OrderItemDAO)

	if err := repo.db.Model(orderItem).Where("order_item_code = ?", code).Select(); err != nil {
		return nil, err
	}

	return orderItem, nil
}

func (repo *orderItemRepo) ListByOrderID(orderID int) ([]*domain.OrderItemDAO, error) {
	orderItems := []*domain.OrderItemDAO{}
	if err := repo.db.Model(&orderItems).Where("order_id = ?", orderID).Select(); err != nil {
		return nil, err
	}

	return orderItems, nil
}

func (repo *orderItemRepo) ListByExhibitionID(exhibitionID string) ([]*domain.OrderItemDAO, error) {
	orderItems := []*domain.OrderItemDAO{}
	if err := repo.db.Model(&orderItems).Where("exhibition_id = ?", exhibitionID).Order("exhibition_id DESC").Limit(5).Select(); err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (repo *orderItemRepo) Update(orderItemDao *domain.OrderItemDAO) (*domain.OrderItemDAO, error) {
	_, err := repo.db.Model(orderItemDao).WherePK().Update()
	if err != nil {
		return orderItemDao, err
	}

	return orderItemDao, nil
}

func (repo *orderItemRepo) ListAllCanceled() ([]*domain.OrderItemDAO, error) {
	orderItems := []*domain.OrderItemDAO{}
	if err := repo.db.Model(&orderItems).Where("order_item_status = ?", domain.ORDER_ITEM_CANCEL_FINISHED).Select(); err != nil {
		return nil, err
	}

	return orderItems, nil
}

func (repo *orderItemRepo) ListByProductGroupID(pgID string) ([]*domain.OrderItemDAO, int, error) {
	orderItems := []*domain.OrderItemDAO{}
	if err := repo.db.Model(&orderItems).Where("order_item_status = ?", domain.ORDER_ITEM_PAYMENT_FINISHED).Select(); err != nil {
		return nil, 0, err
	}
	return orderItems, len(orderItems), nil
}

func PostgresOrderItemRepo(conn *PostgresDB) repository.OrderItemsRepository {
	return &orderItemRepo{
		db: conn.db,
	}
}
