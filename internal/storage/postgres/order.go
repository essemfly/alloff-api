package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type orderRepo struct {
	db *pg.DB
}

func (repo *orderRepo) Get(ID int) (*domain.OrderDAO, error) {
	var order *domain.OrderDAO
	if err := repo.db.Model(order).Where("id = ?", ID).Select(); err != nil {
		return nil, err
	}

	return order, nil
}

func (repo *orderRepo) GetByAlloffID(ID string) (*domain.OrderDAO, error) {
	var order *domain.OrderDAO
	if err := repo.db.Model(order).Where("alloff_order_id = ?", ID).Select(); err != nil {
		return nil, err
	}

	return order, nil
}

func (repo *orderRepo) List(userID string) ([]*domain.OrderDAO, error) {
	orders := []*domain.OrderDAO{}
	if err := repo.db.Model(&orders).Order("created_at DESC").Select(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (repo *orderRepo) Insert(orderDao *domain.OrderDAO) (*domain.OrderDAO, error) {
	_, err := repo.db.Model(orderDao).Insert()
	if err != nil {
		return nil, err
	}

	return orderDao, nil
}

func (repo *orderRepo) Update(orderDao *domain.OrderDAO) (*domain.OrderDAO, error) {
	_, err := repo.db.Model(orderDao).WherePK().Update()
	if err != nil {
		return nil, err
	}

	return orderDao, nil
}

func PostgresOrderRepo(conn *PostgresDB) repository.OrdersRepository {
	return &orderRepo{
		db: conn.db,
	}
}
