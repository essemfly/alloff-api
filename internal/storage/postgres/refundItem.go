package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type refundItemRepo struct {
	db *pg.DB
}

func (repo *refundItemRepo) Insert(refundDao *domain.RefundItemDAO) (*domain.RefundItemDAO, error) {
	_, err := repo.db.Model(refundDao).Insert()
	if err != nil {
		return nil, err
	}

	return refundDao, nil
}

func PostgresRefundItemRepo(conn *PostgresDB) repository.RefundItemsRepository {
	return &refundItemRepo{
		db: conn.db,
	}
}
