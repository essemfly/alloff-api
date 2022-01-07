package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

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

	errs := basket.IsValid()

	if len(errs) > 0 {
		var errString = []string{}
		for _, err := range errs {
			errString = append(errString, err.Error())
		}

		return &model.OrderValidityResult{
			Available: false,
			ErrorMsgs: errString,
		}, nil
	}

	return &model.OrderValidityResult{
		Available: true,
		ErrorMsgs: nil,
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

	orderDao, err := basket.BuildOrder()
	if err != nil {
		return nil, err
	}

	return &model.OrderWithPayment{
		Success:        true,
		ErrorMsg:       "",
		PaymentInfo:    mapper.MapPaymentToPaymentInfo(paymentDao),
		PaymentMethods: order.GetPaymentMethods(),
		Order:          mapper.MapOrderToOrderInfo(newOrderDao),
		User:           mapper.MapUserDaoToUser(user),
	}, nil

}

func (r *mutationResolver) CancelOrder(ctx context.Context, orderID string) (*model.PaymentStatus, error) {
	// 유저가 Order를 취소하고 싶을때 하는 취소요청 API
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ConfirmOrder(ctx context.Context, orderID string) (*model.PaymentStatus, error) {
	// Order Confirm하는 API 인데, 유저가 앱에서 구매확정을 누르는 API
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RequestPayment(ctx context.Context, input *model.PaymentClientInput) (*model.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
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
}

func (r *mutationResolver) CancelPayment(ctx context.Context, input *model.PaymentClientInput) (*model.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
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

		0. 주문창까지 넘어갔다가 취소된 경우 발생하는 API
		1. Payment가 취소 되면서, 재고가 다시 회복됩니다.
		2. Order의 Status도 다시 돌아옵니다.
	*/
}

func (r *mutationResolver) HandlePaymentResponse(ctx context.Context, input *model.OrderResponse) (*model.PaymentResult, error) {
	panic(fmt.Errorf("not implemented"))
	/*
		type OrderResponse struct {
			Success     bool   `json:"success"`
			ImpUID      string `json:"imp_uid"`
			MerchantUID string `json:"merchant_uid"`
			ErrorMsg    string `json:"error_msg"`
		}

		1. Payment의 결과를 앱에서 iamport로 받는다.
	*/
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

func BuildBasketItems(input *model.OrderInput) ([]*order.BasketItem, error) {
	basketItems := []*order.BasketItem{}
	for _, item := range input.Orders {
		pd, err := ioc.Repo.Products.Get(item.ProductID)
		if err != nil {
			return nil, err
		}

		basketItem := &order.BasketItem{
			Product:      pd,
			ProductGroup: nil,
			Size:         item.Selectsize,
			Quantity:     item.Quantity,
		}

		if item.ProductGroupID != "" {
			pg, err := ioc.Repo.ProductGroups.Get(item.ProductGroupID)
			basketItem.ProductGroup = pg
			if err != nil {
				return nil, err
			}
		}
		basketItems = append(basketItems, basketItem)
	}
	return basketItems, nil
}
