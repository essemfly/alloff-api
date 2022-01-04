package postgres

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
)

type PostgresDB struct {
	db *pg.DB
}

func NewPostgresDB(conf config.Configuration) *PostgresDB {
	db := pg.Connect(&pg.Options{
		Addr:     conf.POSTGRES_URL,
		User:     conf.POSTGRES_USER,
		Password: conf.POSTGRES_PASSWORD,
		Database: conf.POSTGRES_DB_NAME,
	})
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	return &PostgresDB{
		db: db,
	}
}

func (conn *PostgresDB) RegisterRepos() {
	ioc.Repo.Orders = PostgresOrderRepo(conn)
	ioc.Repo.Payments = PostgresPaymentRepo(conn)
}
