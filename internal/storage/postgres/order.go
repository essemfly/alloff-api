package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type orderRepo struct {
	db *pg.DB
}

func PostgresOrderRepo(conn *PostgresDB) repository.OrdersRepository {
	return &orderRepo{
		db: conn.db,
	}
}
