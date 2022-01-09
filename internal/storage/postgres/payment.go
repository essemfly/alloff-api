package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type paymentRepo struct {
	db *pg.DB
}

func (repo *paymentRepo) Insert(paymentDao *domain.PaymentDAO) (*domain.PaymentDAO, error) {
	_, err := repo.db.Model(paymentDao).Insert()
	if err != nil {
		return nil, err
	}

	return paymentDao, nil
}

func (repo *paymentRepo) GetByOrderIDAndAmount(orderID string, amount int) (*domain.PaymentDAO, error) {
	paymentDao := domain.PaymentDAO{}
	err := repo.db.Model(paymentDao).Where("merchantuid = ?", orderID).Where("amount = ?", amount).Select()
	if err != nil {
		return nil, err
	}

	return &paymentDao, nil
}

func (repo *paymentRepo) GetByImpUID(impUID string) (*domain.PaymentDAO, error) {
	paymentDao := domain.PaymentDAO{}
	err := repo.db.Model(paymentDao).Where("impuid = ?", impUID).Select()
	if err != nil {
		return nil, err
	}

	return &paymentDao, nil
}

func PostgresPaymentRepo(conn *PostgresDB) repository.PaymentsRepository {
	return &paymentRepo{
		db: conn.db,
	}
}
