package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/alimtalk"
	"github.com/lessbutter/alloff-api/pkg/basket"
	"go.uber.org/zap"
)

func (r *mutationResolver) RequestOrder(ctx context.Context, input *model.OrderInput) (*model.OrderWithPayment, error) {
	/*
		1. 기본적으로 위와 동일하다 (Order 생성하고, Valid Check하고, Errors들을 모아서 보여준다.)
		2. 여기서 완료되면 Order가 생성이되고, 주문 결제하는 창으로 넘어간다.
	*/

	user, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, err
	}

	basketItems, err := BuildBasketItems(input)
	if err != nil {
		return nil, err
	}

	basketDao := &domain.Basket{
		Items:        basketItems,
		ProductPrice: input.ProductPrice,
	}

	newCart, isValid, reason := basket.Refresh(basketDao)
	if !isValid {
		return nil, errors.New(reason)
	}

	orderDao, _ := basket.BuildOrder(user, newCart)
	newOrderDao, err := ioc.Repo.Orders.Insert(orderDao)
	if err != nil {
		config.Logger.Error("order insert fail", zap.Error(err))
		return nil, fmt.Errorf("ERR300:failed to create order not found" + err.Error())
	}

	basePayment := newOrderDao.GetBasePayment()

	return &model.OrderWithPayment{
		Success:        true,
		ErrorMsg:       "",
		PaymentInfo:    mapper.MapPayment(basePayment),
		PaymentMethods: basePayment.GetPaymentMethods(),
		Order:          mapper.MapOrder(newOrderDao),
		User:           mapper.MapUserDaoToUser(user),
	}, nil
}

func (r *mutationResolver) RequestPayment(ctx context.Context, input *model.PaymentClientInput) (*model.PaymentStatus, error) {
	/*
		1. 결제창 띄우면서, iamport에 등록하는 작업을 해야합니다.
		2. 동시에 상품에서 주문한 상품의 재고를 없애줍니다.
	*/

	_, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, err
	}

	paymentDao := BuildPaymentDao(input)
	orderDao, err := ioc.Repo.Orders.GetByAlloffID(paymentDao.MerchantUid)
	if err != nil {
		return nil, fmt.Errorf("ERR301:failed to find order order not found")
	}

	result := &model.PaymentStatus{
		Success:     false,
		ErrorMsg:    "",
		PaymentInfo: mapper.MapPayment(paymentDao),
		Order:       mapper.MapOrder(orderDao),
	}

	orderDao.UserMemo = *input.Memo
	err = ioc.Service.OrderWithPaymentService.RequestPayment(orderDao, paymentDao)
	if err != nil {
		return result, err
	}

	result.Success = true
	return result, nil
}

func (r *mutationResolver) CancelPayment(ctx context.Context, orderID string) (*model.PaymentStatus, error) {
	/*
		0. 주문창까지 넘어갔다가 취소된 경우 발생하는 API
		1. Payment가 취소 되면서, 재고가 다시 회복됩니다.
		2. Order의 Status도 다시 돌아옵니다.
	*/

	_, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, err
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(orderID)
	if err != nil {
		return nil, fmt.Errorf("ERR301:failed to find order order not found")
	}

	paymentDao, err := ioc.Repo.Payments.GetByOrderIDAndAmount(orderID, orderDao.TotalPrice)
	if err != nil {
		return nil, fmt.Errorf("ERR404:failed to find payment order not found")
	}

	result := &model.PaymentStatus{
		Success:     false,
		ErrorMsg:    "",
		PaymentInfo: mapper.MapPayment(paymentDao),
		Order:       mapper.MapOrder(orderDao),
	}

	err = ioc.Service.OrderWithPaymentService.CancelPayment(orderDao, paymentDao)
	if err != nil {
		return result, fmt.Errorf("ERR306:failed to cancel payment")
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

	_, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, err
	}

	if !input.Success {
		return &model.PaymentResult{
			Success:     false,
			ErrorMsg:    input.ErrorMsg,
			Order:       nil,
			PaymentInfo: nil,
		}, nil
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(input.MerchantUID)
	if err != nil {
		return nil, fmt.Errorf("ERR301:failed to find order order not found")
	}

	if orderDao.OrderStatus == domain.ORDER_PAYMENT_FINISHED {
		return &model.PaymentResult{
			Success:     true,
			ErrorMsg:    "",
			Order:       nil,
			PaymentInfo: nil,
		}, nil
	}

	paymentDao, err := ioc.Repo.Payments.GetByOrderIDAndAmount(input.MerchantUID, orderDao.TotalPrice)
	if err != nil {
		return nil, fmt.Errorf("ERR404:failed to find payment order not found")
	}

	return &model.PaymentResult{
		Success:     true,
		ErrorMsg:    "",
		Order:       mapper.MapOrder(orderDao),
		PaymentInfo: mapper.MapPayment(paymentDao),
	}, nil
}

func (r *mutationResolver) CancelOrderItem(ctx context.Context, orderID string, orderItemID string) (*model.PaymentStatus, error) {
	user, _ := middleware.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("ERR000:invalid token")
	}

	orderDao, paymentDao, err := ioc.Service.OrderWithPaymentService.Find(orderID)
	if err != nil {
		return nil, fmt.Errorf("ERR301:failed to find order order not found")
	}

	intOrderItemID, err := strconv.Atoi(orderItemID)
	if err != nil {
		return nil, fmt.Errorf("invalid orderitemId")
	}

	orderItemDao := orderDao.GetOrderItemByID(intOrderItemID)
	if orderItemDao == nil {
		return nil, fmt.Errorf("ERR307:failed to find order item order not found")
	}

	result := &model.PaymentStatus{
		Success:     false,
		ErrorMsg:    "",
		PaymentInfo: mapper.MapPayment(paymentDao),
		Order:       mapper.MapOrder(orderDao),
	}

	err = ioc.Service.OrderWithPaymentService.CancelOrderRequest(orderDao, orderItemDao, paymentDao)

	if err != nil {
		return result, fmt.Errorf("ERR308:failed to cancel order " + err.Error())
	}

	alimtalk.NotifyOrderCancelAlarm(orderItemDao)

	result.Success = true
	return result, nil
}

func (r *mutationResolver) ConfirmOrderItem(ctx context.Context, orderID string, orderItemID string) (*model.PaymentStatus, error) {
	_, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, err
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(orderID)
	if err != nil {
		return nil, fmt.Errorf("ERR301:failed to find order order not found")
	}

	intOrderItemID, err := strconv.Atoi(orderItemID)
	if err != nil {
		return nil, fmt.Errorf("invalid orderitemId")
	}

	orderItemDao := orderDao.GetOrderItemByID(intOrderItemID)
	if orderItemDao == nil {
		return nil, fmt.Errorf("ERR307:failed to find order item order not found")
	}

	err = orderItemDao.ConfirmOrder()
	if err != nil {
		return nil, fmt.Errorf("ERR309:failed to confirm order " + err.Error())
	}

	_, err = ioc.Repo.OrderItems.Update(orderItemDao)
	if err != nil {
		return nil, fmt.Errorf("ERR305:order update failed")
	}

	newOrderDao, err := ioc.Repo.Orders.GetByAlloffID(orderID)
	if err != nil {
		return nil, fmt.Errorf("ERR301:failed to find order order not found")
	}

	return &model.PaymentStatus{
		Success:     true,
		ErrorMsg:    "",
		PaymentInfo: nil,
		Order:       mapper.MapOrder(newOrderDao),
	}, nil
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.OrderInfo, error) {
	_, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, err
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(id)
	if err != nil {
		return nil, fmt.Errorf("ERR301:failed to find order order not found")
	}

	return mapper.MapOrder(orderDao), nil
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.OrderInfo, error) {
	user, err := middleware.ForContext(ctx)
	if err != nil {
		return nil, err
	}

	onlyPaid := true
	orderDaos, err := ioc.Repo.Orders.List(user.ID.Hex(), onlyPaid)
	if err != nil {
		return nil, fmt.Errorf("ERR301:failed to find order order not found")
	}

	orders := []*model.OrderInfo{}
	for _, orderDao := range orderDaos {
		orders = append(orders, mapper.MapOrder(orderDao))
	}

	return orders, nil
}

func (r *queryResolver) OrderItemStatus(ctx context.Context) ([]*model.OrderItemStatusDescription, error) {
	allStatus := []*model.OrderItemStatusDescription{
		{
			DeliveryType: model.DeliveryTypeForeignDelivery,
			StatusEnum:   model.OrderItemStatusEnumProductPreparing,
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
