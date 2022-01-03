package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
)

type PostgresDB struct {
	db *pg.DB
}

func NewPostgresDB(conf config.Configuration) *PostgresDB {
	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	return &PostgresDB{
		db: db,
	}
}

func (conn *PostgresDB) RegisterRepos() {
	ioc.Repo.Orders = PostgresOrderRepo(conn)
	ioc.Repo.Payments = PostgresPaymentRepo(conn)
}
