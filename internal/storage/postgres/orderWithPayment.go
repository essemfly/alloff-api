package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/service"
)

type orderPaymentService struct {
	db *pg.DB
}

func (repo *orderPaymentService) Insert(*domain.PaymentDAO) (*domain.PaymentDAO, error) {
	panic("work in progress")
}

func (repo *orderPaymentService) Find(orderID string) (*domain.OrderDAO, *domain.PaymentDAO, error) {
	panic("work in progress")
}

func (repo *orderPaymentService) CancelOrderRequest(*domain.OrderDAO, *domain.PaymentDAO) error {
	panic("work in progress")
}

func (repo *orderPaymentService) CancelPayment(*domain.OrderDAO, *domain.PaymentDAO) error {
	panic("work in progress")
}

func (repo *orderPaymentService) RequestPayment(*domain.OrderDAO, *domain.PaymentDAO) error {
	/*
		1. Validating: 가격이 맞는지 확인 및 재고 확인
		2. Start Payment: 재고 수량 조절, Order와 Payment의 Status 변경
	*/
	panic("work in progress")
}

func (repo *orderPaymentService) VerifyPayment(*domain.OrderDAO, *domain.PaymentDAO) error {
	/*
		1. Validating: 가격이 맞는지 확인 및 재고 확인
		2. Start Payment: 재고 수량 조절
		3. Order와 Payment의 Status 변경, 주문 완료
	*/
	panic("work in progress")
}

func PostgresOrderPaymentRepo(conn *PostgresDB) service.OrderWithPaymentService {
	return &orderPaymentService{
		db: conn.db,
	}
}
