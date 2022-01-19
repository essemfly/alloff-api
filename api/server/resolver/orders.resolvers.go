package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"

	"github.com/lessbutter/alloff-api/api/middleware"
	"github.com/lessbutter/alloff-api/api/server/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/pkg/order"
)

func (r *mutationResolver) CheckOrder(ctx context.Context, input *model.OrderInput) (*model.OrderValidityResult, error) {
	/*
		1. Baskets을 만든다.
		2. Basket이 Valid한지 Check를 한다.
		3. Errors들을 모아서 보여준다.
		4. Order를 Create해줄 필요는 없다.
	*/

	basketItems, err := BuildBasketItems(input)
	if err != nil {
		return nil, err
	}

	basket := &order.Basket{
		Items:        basketItems,
		ProductPrice: input.ProductPrice,
	}

	orderDao, err := basket.BuildOrder(nil)
	if err != nil {
		return nil, err
	}

	errs := basket.IsValid()

	if len(errs) > 0 {
		var errString = []string{}
		for _, err := range errs {
			errString = append(errString, err.Error())
		}

		return &model.OrderValidityResult{
			Available: false,
			ErrorMsgs: errString,
			Order:     orderDao.ToDTO(),
		}, nil
	}

	return &model.OrderValidityResult{
		Available: true,
		ErrorMsgs: nil,
		Order:     orderDao.ToDTO(),
	}, nil
}

func (r *mutationResolver) RequestOrder(ctx context.Context, input *model.OrderInput) (*model.OrderWithPayment, error) {
	/*
		1. 기본적으로 위와 동일하다 (Order 생성하고, Valid Check하고, Errors들을 모아서 보여준다.)
		2. 여기서 완료되면 Order가 생성이되고, 주문 결제하는 창으로 넘어간다.
	*/

	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	basketItems, err := BuildBasketItems(input)
	if err != nil {
		return nil, err
	}

	basket := &order.Basket{
		Items:        basketItems,
		ProductPrice: input.ProductPrice,
	}

	errs := basket.IsValid()
	if len(errs) > 0 {
		return nil, errs[0]
	}

	orderDao, err := basket.BuildOrder(user)
	if err != nil {
		return nil, err
	}

	newOrderDao, err := ioc.Repo.Orders.Insert(orderDao)
	if err != nil {
		return nil, err
	}

	basePayment := newOrderDao.GetBasePayment()

	return &model.OrderWithPayment{
		Success:        true,
		ErrorMsg:       "",
		PaymentInfo:    basePayment.ToDTO(),
		PaymentMethods: basePayment.GetPaymentMethods(),
		Order:          newOrderDao.ToDTO(),
		User:           user.ToDTO(),
	}, nil
}

func (r *mutationResolver) RequestPayment(ctx context.Context, input *model.PaymentClientInput) (*model.PaymentStatus, error) {
	/*
		type PaymentClientInput struct {
			Pg            string  `json:"pg"`
			PayMethod     string  `json:"payMethod"`
			MerchantUID   string  `json:"merchantUid"`
			Amount        int     `json:"amount"`
			Name          *string `json:"name"`
			BuyerName     *string `json:"buyerName"`
			BuyerMobile   *string `json:"buyerMobile"`
			BuyerAddress  *string `json:"buyerAddress"`
			BuyerPostCode *string `json:"buyerPostCode"`
			Memo          *string `json:"memo"`
			AppScheme     *string `json:"appScheme"`
		}

		1. 결제창 띄우면서, iamport에 등록하는 작업을 해야합니다.
		2. 동시에 상품에서 주문한 상품의 재고를 없애줍니다.
	*/

	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	paymentDao := BuildPaymentDao(input)
	orderDao, err := ioc.Repo.Orders.GetByAlloffID(paymentDao.MerchantUid)
	if err != nil {
		return nil, err
	}

	result := &model.PaymentStatus{
		Success:     false,
		ErrorMsg:    "",
		PaymentInfo: paymentDao.ToDTO(),
		Order:       orderDao.ToDTO(),
	}

	orderDao.UserMemo = *input.Memo
	err = ioc.Service.OrderWithPaymentService.RequestPayment(orderDao, paymentDao)
	if err != nil {
		return result, err
	}

	result.Success = true
	return result, nil
}

func (r *mutationResolver) CancelPayment(ctx context.Context, input *model.PaymentClientInput) (*model.PaymentStatus, error) {
	/*
		0. 주문창까지 넘어갔다가 취소된 경우 발생하는 API
		1. Payment가 취소 되면서, 재고가 다시 회복됩니다.
		2. Order의 Status도 다시 돌아옵니다.
	*/

	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(input.MerchantUID)
	if err != nil {
		return nil, err
	}

	paymentDao, err := ioc.Repo.Payments.GetByOrderIDAndAmount(input.MerchantUID, input.Amount)
	if err != nil {
		return nil, err
	}

	result := &model.PaymentStatus{
		Success:     false,
		ErrorMsg:    "",
		PaymentInfo: paymentDao.ToDTO(),
		Order:       orderDao.ToDTO(),
	}

	err = ioc.Service.OrderWithPaymentService.CancelPayment(orderDao, paymentDao)
	if err != nil {
		return result, err
	}

	result.Success = true
	return result, nil
}

func (r *mutationResolver) HandlePaymentResponse(ctx context.Context, input *model.OrderResponse) (*model.PaymentResult, error) {
	/*
		type OrderResponse struct {
			Success     bool   `json:"success"`
			ImpUID      string `json:"imp_uid"`
			MerchantUID string `json:"merchant_uid"`
			ErrorMsg    string `json:"error_msg"`
		}

		1. Payment의 결과를 앱에서 iamport로 받는다.
	*/

	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(input.MerchantUID)
	if err != nil {
		return nil, err
	}

	err = ioc.Service.OrderWithPaymentService.VerifyPayment(orderDao, input.ImpUID)
	if err != nil {
		return nil, err
	}

	return &model.PaymentResult{
		Success:     true,
		ErrorMsg:    "",
		Order:       nil,
		PaymentInfo: nil,
	}, nil
}

func (r *mutationResolver) CancelOrderItem(ctx context.Context, orderID string, orderItemID string) (*model.PaymentStatus, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	orderDao, paymentDao, err := ioc.Service.OrderWithPaymentService.Find(orderID)
	if err != nil {
		return nil, err
	}

	intOrderItemID, err := strconv.Atoi(orderItemID)
	if err != nil {
		return nil, err
	}

	orderItemDao := orderDao.GetOrderItemByID(intOrderItemID)
	if orderItemDao == nil {
		return nil, err
	}

	result := &model.PaymentStatus{
		Success:     false,
		ErrorMsg:    "",
		PaymentInfo: paymentDao.ToDTO(),
		Order:       orderDao.ToDTO(),
	}

	err = ioc.Service.OrderWithPaymentService.CancelOrderRequest(orderDao, orderItemDao, paymentDao)

	if err != nil {
		return result, err
	}

	result.Success = true
	return result, nil
}

func (r *mutationResolver) ConfirmOrderItem(ctx context.Context, orderID string, orderItemID string) (*model.PaymentStatus, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(orderID)
	if err != nil {
		return nil, err
	}

	intOrderItemID, err := strconv.Atoi(orderItemID)
	if err != nil {
		return nil, err
	}

	orderItemDao := orderDao.GetOrderItemByID(intOrderItemID)
	if orderItemDao == nil {
		return nil, err
	}

	err = orderItemDao.ConfirmOrder()
	if err != nil {
		return nil, err
	}

	_, err = ioc.Repo.OrderItems.Update(orderItemDao)
	if err != nil {
		return nil, err
	}

	newOrderDao, err := ioc.Repo.Orders.GetByAlloffID(orderID)
	if err != nil {
		return nil, err
	}

	return &model.PaymentStatus{
		Success:     true,
		ErrorMsg:    "",
		PaymentInfo: nil,
		Order:       newOrderDao.ToDTO(),
	}, nil
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.OrderInfo, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(id)
	if err != nil {
		return nil, err
	}

	return orderDao.ToDTO(), nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.OrderInfo, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	orderDaos, err := ioc.Repo.Orders.List(user.ID.Hex())
	if err != nil {
		return nil, err
	}

	orders := []*model.OrderInfo{}
	for _, orderDao := range orderDaos {
		orders = append(orders, orderDao.ToDTO())
	}

	return orders, nil
}

func (r *queryResolver) OrderItemStatus(ctx context.Context) ([]*model.OrderItemStatusDescription, error) {
	allStatus := []*model.OrderItemStatusDescription{
		{
			DeliveryType: model.DeliveryTypeForeignDelivery,
			StatusEnum:   model.OrderItemStatusEnumPaymentFinished,
			Description:  "고객님의 주문을 확인한 후 현지에서 상품 구매를 진행하고 있습니다. 현지 재고 상황에 따라 3~5일 소요될 수 있습니다.",
		}, {
			DeliveryType: model.DeliveryTypeForeignDelivery,
			StatusEnum:   model.OrderItemStatusEnumForeignProductInspecting,
			Description:  "올오프 현지 물류센터에 상품이 입고되어 순차적인 검수를 진행하고 배송을 준비하고 있습니다.",
		}, {
			DeliveryType: model.DeliveryTypeForeignDelivery,
			StatusEnum:   model.OrderItemStatusEnumForeignDeliveryStatrted,
			Description:  "현지 물류센터에서 한국으로 상품이 출고되었습니다. 인천공항에 도착하여 정식 세관 통관 과정을 거치게 됩니다.",
		}, {
			DeliveryType: model.DeliveryTypeForeignDelivery,
			StatusEnum:   model.OrderItemStatusEnumDeliveryStarted,
			Description:  "통관이 완료된 상품이 국내 물류센터에 입고 후 고객님께 배송 중입니다.",
		}, {
			DeliveryType: model.DeliveryTypeForeignDelivery,
			StatusEnum:   model.OrderItemStatusEnumDeliveryFinished,
			Description:  "주문하신 상품이 고객님께 도착하였습니다.",
		}, {
			DeliveryType: model.DeliveryTypeDomesticDelivery,
			StatusEnum:   model.OrderItemStatusEnumPaymentFinished,
			Description:  "고객님의 주문을 확인한 후 현지에서 상품 구매를 진행하고 있습니다. 현지 재고 상황에 따라 3~5일 소요될 수 있습니다.",
		}, {
			DeliveryType: model.DeliveryTypeDomesticDelivery,
			StatusEnum:   model.OrderItemStatusEnumProductPreparing,
			Description:  "고객님의 주문을 확인한 후 올오프 또는 브랜드에서 상품을 준비하고 있습니다. 재고 상황에 따라 3~5일 소요될 수 있습니다.",
		}, {
			DeliveryType: model.DeliveryTypeDomesticDelivery,
			StatusEnum:   model.OrderItemStatusEnumDeliveryPreparing,
			Description:  "올오프 또는 브랜드 물류센터에 상품이 입고되어 순차적인 검수를 진행하고 배송을 준비하고 있습니다.",
		}, {
			DeliveryType: model.DeliveryTypeDomesticDelivery,
			StatusEnum:   model.OrderItemStatusEnumDeliveryStarted,
			Description:  "올오프 또는 브랜드 물류센터에서 상품이 출고되어 고객님께 배송 중입니다.",
		}, {
			DeliveryType: model.DeliveryTypeDomesticDelivery,
			StatusEnum:   model.OrderItemStatusEnumDeliveryFinished,
			Description:  "주문하신 상품이 고객님께 도착하였습니다.",
		},
	}

	return allStatus, nil
}
