package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/service"
)

type orderPaymentService struct {
	db *pg.DB
}

func (repo *orderPaymentService) Find(orderID string) (*domain.OrderDAO, *domain.PaymentDAO, error) {
	orderDao, err := ioc.Repo.Orders.GetByAlloffID(orderID)
	if err != nil {
		return nil, nil, err
	}

	paymentDao, err := ioc.Repo.Payments.GetByOrderIDAndAmount(orderID, orderDao.TotalPrice)
	if err != nil {
		return orderDao, nil, err
	}

	return orderDao, paymentDao, nil
}

func (repo *orderPaymentService) CancelOrderRequest(orderDao *domain.OrderDAO, orderItemDao *domain.OrderItemDAO, paymentDao *domain.PaymentDAO) error {
	/*
		주문 취소요청시 실행되는 함수
		1. 주문 취소가 가능한 Status면 취소 잘 되게끔 만들어준다. + 환불까지
		2. 주문 취소가 가능한 Status가 아니면, Cancel Requested로 바꿔준다.
	*/
	// (다시) OrderItem 별로 Check해야 될 것 같은 느낌이 드네요?
	if orderDao.CanCancelPayment() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := repo.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
			// TODO: Cancel Order ITEMS one by one
			return nil
		}); err != nil {
			panic(err)
		}
		return nil
	}

	if orderDao.CanCancelOrder() {
		return nil
	}

	return errors.New("order status not available")
}

func (repo *orderPaymentService) RequestPayment(orderDao *domain.OrderDAO, paymentDao *domain.PaymentDAO) error {
	/*
		1. Validating: 가격이 맞는지 확인 및 재고 확인
		2. Start Payment: 재고 수량 조절, Order와 Payment의 Status 변경
	*/

	if paymentDao.BuyerMobile == "" {
		return errors.New("invalid mobile error")
	}
	if paymentDao.BuyerAddress == "" || paymentDao.BuyerPostCode == "" {
		return errors.New("invalid order address error")
	}
	if orderDao.TotalPrice != paymentDao.Amount {
		return errors.New("order amount not matched")
	}
	if len(orderDao.OrderItems) == 0 {
		return errors.New("empty orders")
	}
	if orderDao.OrderStatus != domain.ORDER_CREATED && orderDao.OrderStatus != domain.ORDER_RECREATED {
		return errors.New("already ongoing order exists")
	}

	// 이제 Stock 옵션 줄이면 된다. + Order의 상태 및 timestamp찍으면 된다.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := repo.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		orderDao.UpdatedAt = time.Now()
		orderDao.OrderStatus = domain.ORDER_PAYMENT_PENDING
		totalProductPrices := 0

		for _, item := range orderDao.OrderItems {
			item.UpdatedAt = time.Now()
			item.OrderStatus = domain.ORDER_PAYMENT_PENDING
			pd, err := ioc.Repo.Products.Get(item.ProductID)
			if err != nil {
				return err
			}

			if pd.Removed || pd.Soldout {
				return errors.New("product sold out or removed")
			}

			err = pd.Release(item.Size, item.Quantity)
			if err != nil {
				return err
			}
			totalProductPrices += item.Quantity * item.SalesPrice
			repo.db.Model(pd).Update()
		}

		if orderDao.TotalPrice != paymentDao.Amount {
			return errors.New("total price not the same")
		}

		repo.db.Model(orderDao).Update()
		paymentDao.Updated = time.Now()
		paymentDao.PaymentStatus = domain.PAYMENT_CONFIRMED
		repo.db.Model(paymentDao).Update()

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *orderPaymentService) CancelPayment(*domain.OrderDAO, *domain.PaymentDAO) error {
	/*
		주문창까지 갔다가 취소 되는 함수
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

func PostgresOrderPaymentService(conn *PostgresDB) service.OrderWithPaymentService {
	return &orderPaymentService{
		db: conn.db,
	}
}
