package postgres

import (
	"context"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
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

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	err := createSchema(db)
	if err != nil {
		log.Println(err)
	}

	return &PostgresDB{
		db: db,
	}
}

func (conn *PostgresDB) RegisterRepos() {
	ioc.Repo.Orders = PostgresOrderRepo(conn)
	ioc.Repo.OrderItems = PostgresOrderItemRepo(conn)
	ioc.Repo.Payments = PostgresPaymentRepo(conn)
	ioc.Repo.Refunds = PostgresRefundItemRepo(conn)
	ioc.Service.OrderWithPaymentService = PostgresOrderPaymentService(conn)
}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*domain.OrderDAO)(nil),
		(*domain.OrderItemDAO)(nil),
		(*domain.PaymentDAO)(nil),
		(*domain.RefundItemDAO)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
