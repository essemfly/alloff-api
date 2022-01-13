package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type refundItemRepo struct {
	db *pg.DB
}

func (refundRepo *refundItemRepo) Insert(*domain.RefundItemDAO) (*domain.RefundItemDAO, error) {
	return nil, nil
}

func PostgresRefundItemRepo(conn *PostgresDB) repository.RefundItemsRepository {
	return &refundItemRepo{
		db: conn.db,
	}
}
