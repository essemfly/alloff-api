package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type paymentRepo struct {
	db *pg.DB
}

func PostgresPaymentRepo(conn *PostgresDB) repository.PaymentsRepository {
	return &paymentRepo{
		db: conn.db,
	}
}
