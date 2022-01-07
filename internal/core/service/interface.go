package service

import "github.com/lessbutter/alloff-api/internal/core/domain"

type OrderWithPaymentService interface {
	Find(orderID string) (*domain.OrderDAO, *domain.PaymentDAO, error)
	RequestPayment(*domain.OrderDAO, *domain.PaymentDAO) error
	VerifyPayment(*domain.OrderDAO, *domain.PaymentDAO) error
	CancelOrderRequest(*domain.OrderDAO, *domain.PaymentDAO) error
	CancelPayment(*domain.OrderDAO, *domain.PaymentDAO) error
}
