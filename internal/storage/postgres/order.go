package postgres

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type orderRepo struct {
	db *pg.DB
}

func (repo *orderRepo) Get(ID int) (*domain.OrderDAO, error) {
	order := new(domain.OrderDAO)

	if err := repo.db.Model(order).Where("id = ?", ID).Select(); err != nil {
		return nil, err
	}

	orderItems := []*domain.OrderItemDAO{}
	if err := repo.db.Model(&orderItems).Where("order_id = ?", order.ID).Order("id ASC").Select(); err != nil {
		return nil, err
	}
	order.OrderItems = orderItems

	return order, nil
}

func (repo *orderRepo) GetByAlloffID(ID string) (*domain.OrderDAO, error) {
	order := new(domain.OrderDAO)

	if err := repo.db.Model(order).
		Where("alloff_order_id = ?", ID).
		Select(); err != nil {
		return nil, err
	}

	orderItems := []*domain.OrderItemDAO{}
	if err := repo.db.Model(&orderItems).Where("order_id = ?", order.ID).Order("id ASC").Select(); err != nil {
		return nil, err
	}
	order.OrderItems = orderItems

	return order, nil
}

func (repo *orderRepo) List(userID string) ([]*domain.OrderDAO, error) {
	orders := []*domain.OrderDAO{}

	if err := repo.db.Model(&orders).Where("user_id = ?", userID).Order("created_at DESC").Select(); err != nil {
		return nil, err
	}
	for _, order := range orders {
		orderItems := []*domain.OrderItemDAO{}
		if err := repo.db.Model(&orderItems).Where("order_id = ?", order.ID).Order("id ASC").Select(); err != nil {
			return nil, err
		}
		order.OrderItems = orderItems
	}

	return orders, nil
}

func (repo *orderRepo) Insert(orderDao *domain.OrderDAO) (*domain.OrderDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repo.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := repo.db.Model(orderDao).Insert()
		if err != nil {
			return err
		}

		for _, item := range orderDao.OrderItems {
			item.OrderID = orderDao.ID
			_, err := repo.db.Model(item).Insert()
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return orderDao, err
	}

	return orderDao, nil
}

func (repo *orderRepo) Update(orderDao *domain.OrderDAO) (*domain.OrderDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repo.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		_, err := repo.db.Model(orderDao).WherePK().Update()
		if err != nil {
			return err
		}

		for _, item := range orderDao.OrderItems {
			_, err := repo.db.Model(item).WherePK().Update()
			if err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return orderDao, err
	}

	return orderDao, nil
}

func PostgresOrderRepo(conn *PostgresDB) repository.OrdersRepository {
	return &orderRepo{
		db: conn.db,
	}
}
