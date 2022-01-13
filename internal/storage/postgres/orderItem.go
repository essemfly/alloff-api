package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type orderItemRepo struct {
	db *pg.DB
}

func (orderItemRepo *orderItemRepo) Get(ID int) (*domain.OrderItemDAO, error) {
	return nil, nil
}
func (orderItemRepo *orderItemRepo) ListByOrderID(orderID int) ([]*domain.OrderItemDAO, error) {
	return nil, nil
}
func (orderItemRepo *orderItemRepo) Insert(*domain.OrderItemDAO) (*domain.OrderItemDAO, error) {
	return nil, nil
}
func (orderItemRepo *orderItemRepo) Update(*domain.OrderItemDAO) (*domain.OrderItemDAO, error) {
	return nil, nil
}

func PostgresOrderItemRepo(conn *PostgresDB) repository.OrderItemsRepository {
	return &orderItemRepo{
		db: conn.db,
	}
}
