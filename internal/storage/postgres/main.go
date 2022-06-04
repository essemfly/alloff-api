package postgres

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/spf13/viper"
)

type PostgresDB struct {
	db *pg.DB
}

func NewPostgresDB() *PostgresDB {
	db := pg.Connect(&pg.Options{
		Addr:     viper.GetString("POSTGRES_URL"),
		User:     viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		Database: viper.GetString("POSTGRES_DB_NAME"),
	})

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
