package postgres

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/service"
	"github.com/lessbutter/alloff-api/internal/pkg/alimtalk"
	productinfo "github.com/lessbutter/alloff-api/pkg/productInfo"
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

			orderMsg :=
				"----------결제 취소 요청 ---------- \n" +
					"결제 Order ID: " + orderDao.AlloffOrderID + ": " + orderItemDao.OrderItemCode + "\n" +
					"상품명: " + paymentDao.Name + "\n" +
					"주문자 번호: " + paymentDao.BuyerMobile + "\n" +
					"가격: " + strconv.Itoa(paymentDao.Amount) + "\n" +
					"주소: " + paymentDao.BuyerPostCode + " " + paymentDao.BuyerAddress + "\n" +
					"받는 사람 번호: " + paymentDao.BuyerMobile

			config.WriteCancelMessage(orderMsg)

			refundPrice := orderItemDao.SalesPrice * orderItemDao.Quantity
			_, err := config.PaymentService.CancelPaymentImpUID(paymentDao.ImpUID, orderDao.AlloffOrderID, float64(refundPrice), 0, "cancel before products ready", "", "", "")
			if err != nil {
				log.Println("cancel payment error on iamport")
				return err
			}

			orderItemDao.OrderItemStatus = domain.ORDER_ITEM_CANCEL_FINISHED
			orderItemDao.CancelRequestedAt = time.Now()
			orderItemDao.CancelFinishedAt = time.Now()
			orderItemDao.UpdatedAt = time.Now()
			paymentDao.PaymentStatus = domain.PAYMENT_REFUND_FINISHED
			paymentDao.UpdatedAt = time.Now()
			_, err = ioc.Repo.OrderItems.Update(orderItemDao)
			if err != nil {
				return err
			}

			pd, err := ioc.Repo.Products.Get(orderItemDao.ProductID)
			if err != nil {
				return err
			}
			err = pd.ProductInfo.Revert(orderItemDao.Size, orderItemDao.Quantity)
			if err != nil {
				return err
			}

			pd.ProductInfo.CheckSoldout()
			_, err = productinfo.Update(pd.ProductInfo)
			if err != nil {
				config.Logger.Error("Productinfo update failed")
				return fmt.Errorf("ERR106:product update failed" + err.Error())
			}

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
				log.Println("error on adding refund", err)
				return err
			}
			_, err = ioc.Repo.Payments.Update(paymentDao)
			if err != nil {
				return err
			}

			orderMsg =
				"----------결제 취소 완료 ---------- \n" +
					"결제 Order ID: " + orderDao.AlloffOrderID + ": " + orderItemDao.OrderItemCode + "\n" +
					"상품명: " + paymentDao.Name + "\n" +
					"주문자 번호: " + paymentDao.BuyerMobile + "\n" +
					"가격: " + strconv.Itoa(paymentDao.Amount) + "\n" +
					"주소: " + paymentDao.BuyerPostCode + " " + paymentDao.BuyerAddress + "\n" +
					"받는 사람 번호: " + paymentDao.BuyerMobile

			config.WriteCancelMessage(orderMsg)

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
		orderItemDao.OrderItemStatus = domain.ORDER_ITEM_RETURN_REQUESTED

		_, err := ioc.Repo.OrderItems.Update(orderItemDao)
		if err != nil {
			return err
		}

		orderMsg :=
			"----------결제 취소 요청---------- \n" +
				"결제 Order ID: " + orderDao.AlloffOrderID + ": " + orderItemDao.OrderItemCode + "\n" +
				"상품명: " + paymentDao.Name + "\n" +
				"주문자 번호: " + paymentDao.BuyerMobile + "\n" +
				"가격: " + strconv.Itoa(paymentDao.Amount) + "\n" +
				"주소: " + paymentDao.BuyerPostCode + " " + paymentDao.BuyerAddress + "\n" +
				"받는 사람 번호: " + paymentDao.BuyerMobile

		config.WriteCancelMessage(orderMsg)

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
		return fmt.Errorf("ERR302:failed to find buyer mobile")
	}
	if paymentDao.BuyerAddress == "" || paymentDao.BuyerPostCode == "" {
		return fmt.Errorf("ERR303:failed to find address")
	}
	if orderDao.TotalPrice != paymentDao.Amount {
		return fmt.Errorf("ERR101:invalid total products price order amount")
	}
	if len(orderDao.OrderItems) == 0 {
		return fmt.Errorf("ERR304:empty orders")
	}

	if orderDao.OrderStatus != domain.ORDER_CREATED && orderDao.OrderStatus != domain.ORDER_RECREATED {
		if orderDao.OrderStatus != domain.ORDER_PAYMENT_PENDING {
			return fmt.Errorf("ERR400:already ongoing order finished")
		}
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
				return fmt.Errorf("ERR100:alloffproduct not found")
			}

			if !pd.OnSale {
				return fmt.Errorf("ERR102:alloffproduct is not for sale")
			}

			if pd.ProductInfo.IsRemoved {
				return fmt.Errorf("ERR102:alloffproduct is removed")
			}

			if pd.ProductInfo.IsSoldout {
				return fmt.Errorf("ERR105:product soldout")
			}

			err = pd.ProductInfo.Release(item.Size, item.Quantity)
			if err != nil {
				return fmt.Errorf("ERR106:product update failed" + err.Error())
			}
			totalProductPrices += item.Quantity * item.SalesPrice
			pd.ProductInfo.CheckSoldout()

			_, err = productinfo.Update(pd.ProductInfo)
			if err != nil {
				config.Logger.Error("Productinfo update failed")
				return fmt.Errorf("ERR106:product update failed" + err.Error())
			}

			_, err = tx.Model(item).WherePK().Update()
			if err != nil {
				log.Println("orderitemDao Update")
				return fmt.Errorf("ERR305:order update failed" + err.Error())
			}
		}

		if orderDao.TotalPrice != paymentDao.Amount {
			return fmt.Errorf("ERR101:invalid total products price order amount")
		}

		_, err := tx.Model(orderDao).WherePK().Update()
		if err != nil {
			log.Println("orderDao Update")
			return fmt.Errorf("ERR305:order update failed" + err.Error())
		}

		_, prevFailedErr := ioc.Repo.Payments.GetByOrderIDAndAmount(orderDao.AlloffOrderID, orderDao.TotalPrice)

		if prevFailedErr != nil {
			_, err = config.PaymentService.PreparePayment(orderDao.AlloffOrderID, float64(orderDao.TotalPrice))
			if err != nil {
				log.Println("iamport error", err)
				return fmt.Errorf("ERR401:iamport prepare payment error")
			}

			paymentDao.CreatedAt = time.Now()
			paymentDao.UpdatedAt = time.Now()
			paymentDao.PaymentStatus = domain.PAYMENT_CREATED
			_, err = tx.Model(paymentDao).Insert()
			if err != nil {
				log.Println("paymentDao Insert", err)
				return fmt.Errorf("ERR403:payment create failed")
			}
			return nil
		}

		paymentDao.PaymentStatus = domain.PAYMENT_CREATED
		paymentDao.UpdatedAt = time.Now()
		_, err = tx.Model(paymentDao).WherePK().Update()
		if err != nil {
			log.Println("paymentDao update", err)
			return fmt.Errorf("ERR402:payment update failed")
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
		_, err := tx.Model(orderDao).WherePK().Update()
		if err != nil {
			log.Println("err on orderDAO", err)
			return err
		}

		for _, item := range orderDao.OrderItems {
			item.UpdatedAt = time.Now()
			item.OrderItemStatus = domain.ORDER_ITEM_RECREATED
			_, err := tx.Model(item).WherePK().Update()
			if err != nil {
				log.Println("err on orderItemDAO", err)
				return err
			}

			pd, err := ioc.Repo.Products.Get(item.ProductID)
			if err != nil {
				return err
			}
			err = pd.ProductInfo.Revert(item.Size, item.Quantity)
			if err != nil {
				return err
			}
			pd.ProductInfo.CheckSoldout()
			_, err = ioc.Repo.Products.Upsert(pd)
			if err != nil {
				log.Println("productDao Update", err)
				return err
			}
		}

		paymentDao.UpdatedAt = time.Now()
		paymentDao.PaymentStatus = domain.PAYMENT_CANCELED
		_, err = tx.Model(paymentDao).WherePK().Update()
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

	orderDao.OrderStatus = domain.ORDER_PAYMENT_FINISHED
	orderDao.UpdatedAt = time.Now()
	orderDao.OrderedAt = time.Now()
	_, err = repo.db.Model(orderDao).WherePK().Update()
	if err != nil {
		log.Println("err on orderDAO", err)
		return err
	}
	for _, orderItemDAO := range orderDao.OrderItems {
		orderItemDAO.OrderItemStatus = domain.ORDER_ITEM_PAYMENT_FINISHED
		orderItemDAO.UpdatedAt = time.Now()
		orderItemDAO.OrderedAt = time.Now()
		_, err = repo.db.Model(orderItemDAO).WherePK().Update()
		if err != nil {
			log.Println("err on orderItemDAO", err)
			return err
		}
	}

	paymentDao, err := ioc.Repo.Payments.GetByOrderIDAndAmount(orderDao.AlloffOrderID, int(payment.Amount))
	if err != nil {
		return err
	}
	paymentDao.PaymentStatus = domain.PAYMENT_CONFIRMED
	paymentDao.ImpUID = impUID
	paymentDao.UpdatedAt = time.Now()
	_, err = repo.db.Model(paymentDao).WherePK().Update()
	if err != nil {
		log.Println("err on paymentdao", err)
		return err
	}

	alimtalk.NotifyPaymentSuccessAlarm(paymentDao)
	WritePaymentSuccessSlack(paymentDao)

	return nil
}

func PostgresOrderPaymentService(conn *PostgresDB) service.OrderWithPaymentService {
	return &orderPaymentService{
		db: conn.db,
	}
}

func WritePaymentSuccessSlack(payment *domain.PaymentDAO) {
	order, _ := ioc.Repo.Orders.GetByAlloffID(payment.MerchantUid)
	orderProducts := []string{}
	itemBackOfficeUrl := "https://office.alloff.co/items/"
	for _, orderItem := range order.OrderItems {
		orderProducts = append(orderProducts, orderItem.ProductUrl+": "+orderItem.Size+" "+strconv.Itoa(orderItem.Quantity)+"개 - "+itemBackOfficeUrl+orderItem.OrderItemCode)
	}

	orderMsg := "**결제 완료** \n" +
		"결제 ID: " + payment.ImpUID + "\n" +
		"주문 ID: " + payment.MerchantUid + "\n" +
		"주문명: " + payment.Name + "\n" +
		"주문 정보: \n" + strings.Join(orderProducts[:], ", ") + "\n" +
		"주문자 번호: " + payment.BuyerMobile + "\n" +
		"가격: " + strconv.Itoa(payment.Amount) + "\n" +
		"주소: " + payment.BuyerPostCode + " " + payment.BuyerAddress + "\n" +
		"받는 사람 번호: " + payment.BuyerMobile

	config.WriteOrderMessage(orderMsg)
}
