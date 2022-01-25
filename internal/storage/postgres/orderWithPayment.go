package postgres

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/service"
	"github.com/lessbutter/alloff-api/internal/pkg/alimtalk"
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
	// 주문취소 바로 가능한 경우
	if orderItemDao.CanCancelPayment() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := repo.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
			orderItemDao.OrderItemStatus = domain.ORDER_ITEM_CANCEL_FINISHED
			orderItemDao.CancelRequestedAt = time.Now()
			orderItemDao.CancelFinishedAt = time.Now()
			orderItemDao.UpdatedAt = time.Now()
			paymentDao.PaymentStatus = domain.PAYMENT_REFUND_FINISHED
			paymentDao.UpdatedAt = time.Now()
			_, err := ioc.Repo.OrderItems.Update(orderItemDao)
			if err != nil {
				return err
			}

			refundPrice := orderItemDao.SalesPrice * orderItemDao.Quantity
			newRefundInfo := &domain.RefundItemDAO{
				OrderID:      orderDao.ID,
				OrderItemID:  orderItemDao.ID,
				RefundFee:    0,
				RefundAmount: refundPrice,
				CreatedAt:    time.Now(),
				UpdatedAt:    time.Now(),
			}
			_, err = ioc.Repo.Refunds.Insert(newRefundInfo)
			if err != nil {
				log.Println("error on adding refund")
				return err
			}

			_, err = config.PaymentService.CancelPaymentImpUID(paymentDao.ImpUID, orderDao.AlloffOrderID, float64(refundPrice), 0, float64(orderDao.TotalPrice), "cancel before products ready", "", "", "")
			if err != nil {
				log.Println("cancel payment error on iamport")
				return err
			}

			_, err = ioc.Repo.Payments.Update(paymentDao)
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			return err
		}
		return nil
	}

	// 주문취소가 바로 안되서 Cancel Requested로 바꿔준다.
	if orderItemDao.CanCancelOrder() {
		orderItemDao.CancelRequestedAt = time.Now()
		orderItemDao.UpdatedAt = time.Now()

		_, err := ioc.Repo.OrderItems.Update(orderItemDao)
		if err != nil {
			return err
		}
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
			item.OrderItemStatus = domain.ORDER_ITEM_PAYMENT_PENDING
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

			_, err = ioc.Repo.Products.Upsert(pd)
			if err != nil {
				log.Println("productDao Update")
				return err
			}

			_, err = repo.db.Model(item).WherePK().Update()
			if err != nil {
				log.Println("orderitemDao Update")
				return err
			}
		}

		if orderDao.TotalPrice != paymentDao.Amount {
			return errors.New("total price not the same")
		}

		_, err := repo.db.Model(orderDao).WherePK().Update()
		if err != nil {
			log.Println("orderDao Update")
			return err
		}

		_, err = config.PaymentService.PreparePayment(orderDao.AlloffOrderID, float64(orderDao.TotalPrice))
		if err != nil {
			log.Println("iamport error")
			return err
		}

		paymentDao.CreatedAt = time.Now()
		paymentDao.UpdatedAt = time.Now()
		paymentDao.PaymentStatus = domain.PAYMENT_CREATED
		_, err = repo.db.Model(paymentDao).Insert()
		if err != nil {
			log.Println("paymentDao Insert")
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *orderPaymentService) CancelPayment(orderDao *domain.OrderDAO, paymentDao *domain.PaymentDAO) error {
	/*
		주문창까지 갔다가 취소 되는 함수
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repo.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		orderDao.UpdatedAt = time.Now()
		orderDao.OrderStatus = domain.ORDER_RECREATED

		for _, item := range orderDao.OrderItems {
			item.UpdatedAt = time.Now()
			item.OrderItemStatus = domain.ORDER_ITEM_RECREATED
			pd, err := ioc.Repo.Products.Get(item.ProductID)
			if err != nil {
				return err
			}

			err = pd.Revert(item.Size, item.Quantity)
			if err != nil {
				return err
			}

			_, err = ioc.Repo.Products.Upsert(pd)
			if err != nil {
				log.Println("productDao Update", err)
				return err
			}
		}

		_, err := repo.db.Model(orderDao).WherePK().Update()
		if err != nil {
			return err
		}

		paymentDao.UpdatedAt = time.Now()
		paymentDao.PaymentStatus = domain.PAYMENT_CANCELED
		_, err = repo.db.Model(paymentDao).WherePK().Update()
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (repo *orderPaymentService) VerifyPayment(orderDao *domain.OrderDAO, impUID string) error {
	/*
		1. Validating: 가격이 맞는지 확인 및 재고 확인
		2. Order와 Payment의 Status 변경, 주문 완료
		3. 알림톡 전송
	*/

	payment, err := config.PaymentService.GetPaymentImpUID(impUID)
	if err != nil {
		return err
	}

	if payment.Amount != int32(orderDao.TotalPrice) {
		return errors.New("payment amount not equal")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := repo.db.RunInTransaction(ctx, func(tx *pg.Tx) error {
		orderDao.OrderStatus = domain.ORDER_PAYMENT_FINISHED
		orderDao.UpdatedAt = time.Now()
		orderDao.OrderedAt = time.Now()
		_, err := repo.db.Model(orderDao).WherePK().Update()
		if err != nil {
			log.Println("err on orderDAO", err)
			return err
		}

		paymentDao, err := ioc.Repo.Payments.GetByOrderIDAndAmount(orderDao.AlloffOrderID, int(payment.Amount))
		if err != nil {
			return err
		}
		paymentDao.PaymentStatus = domain.PAYMENT_CONFIRMED
		paymentDao.UpdatedAt = time.Now()
		_, err = repo.db.Model(paymentDao).WherePK().Update()
		if err != nil {
			log.Println("err on paymentdao", err)
			return err
		}

		alimtalk.NotifyPaymentSuccessAlarm(paymentDao)
		// (TODO) Slack Payment Success Notification

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func PostgresOrderPaymentService(conn *PostgresDB) service.OrderWithPaymentService {
	return &orderPaymentService{
		db: conn.db,
	}
}
