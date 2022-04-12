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
	paymentDao := new(domain.PaymentDAO)
	err := repo.db.Model(paymentDao).Where("merchant_uid = ?", orderID).Where("amount = ?", amount).Select()
	if err != nil {
		return nil, err
	}

	return paymentDao, nil
}

func (repo *paymentRepo) GetByImpUID(impUID string) (*domain.PaymentDAO, error) {
	paymentDao := new(domain.PaymentDAO)
	err := repo.db.Model(paymentDao).Where("imp_uid = ?", impUID).Select()
	if err != nil {
		return nil, err
	}

	return paymentDao, nil
}

func (repo *paymentRepo) Update(paymentDao *domain.PaymentDAO) (*domain.PaymentDAO, error) {
	_, err := repo.db.Model(paymentDao).WherePK().Update()
	if err != nil {
		return paymentDao, err
	}

	return paymentDao, nil
}

func (repo *paymentRepo) ListHolding() ([]*domain.PaymentDAO, error) {
	payments := []*domain.PaymentDAO{}
	query := repo.db.Model(&payments).Where("payment_status = ?", domain.PAYMENT_CREATED)
	if err := query.Order("created_at DESC").Select(); err != nil {
		return nil, err
	}

	return payments, nil
}

func PostgresPaymentRepo(conn *PostgresDB) repository.PaymentsRepository {
	return &paymentRepo{
		db: conn.db,
	}
}
