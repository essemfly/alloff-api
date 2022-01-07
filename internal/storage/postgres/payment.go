package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type paymentRepo struct {
	db *pg.DB
}

func (repo *paymentRepo) Insert(*domain.PaymentDAO) (*domain.PaymentDAO, error) {
	panic("work in progress")
}

func (repo *paymentRepo) GetByOrderIDAndAmount(orderID string, amount int) (*domain.PaymentDAO, error) {
	panic("work in progress")
}

func (repo *paymentRepo) GetByImpUID(impUID string) (*domain.PaymentDAO, error) {
	panic("work in progress")
}

func PostgresPaymentRepo(conn *PostgresDB) repository.PaymentsRepository {
	return &paymentRepo{
		db: conn.db,
	}
}
