package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/server/model"
)

func (r *mutationResolver) CheckOrder(ctx context.Context, input *model.OrderInput) (*model.OrderValidityResult, error) {
	panic(fmt.Errorf("not implemented"))
	/*
		type OrderInput struct {
			Orders       []*ProductOptionInput `json:"orders"`
			ProductPrice int                   `json:"productPrice"`
		}

		type ProductOptionInput struct {
			ProductID  *string `json:"productId"`
			Selectsize string  `json:"selectsize"`
			Quantity   int     `json:"quantity"`
		}

		1. Order를 만든다.
		2. Order가 Valid한지 Check를 한다.
		3. Errors들을 모아서 보여준다.
	*/

	// orderDao, errs := order.CheckValidOrderInput(input)
	// if len(errs) > 0 {
	// 	var errString = []string{}
	// 	for _, err := range errs {
	// 		errString = append(errString, err.Error())
	// 	}

	// 	return &model.OrderValidityResult{
	// 		Available: false,
	// 		ErrorMsgs: errString,
	// 		Order:     mapper.MapOrderToOrderInfo(orderDao),
	// 	}, nil
	// }
	// return &model.OrderValidityResult{
	// 	Available: true,
	// 	ErrorMsgs: nil,
	// 	Order:     mapper.MapOrderToOrderInfo(orderDao),
	// }, nil
}

func (r *mutationResolver) RequestOrder(ctx context.Context, input *model.OrderInput) (*model.OrderWithPayment, error) {
	panic(fmt.Errorf("not implemented"))
	/*
		type OrderInput struct {
				Orders       []*ProductOptionInput `json:"orders"`
				ProductPrice int                   `json:"productPrice"`
			}

		type ProductOptionInput struct {
			ProductID  *string `json:"productId"`
			Selectsize string  `json:"selectsize"`
			Quantity   int     `json:"quantity"`
		}

		1. 기본적으로 위와 동일하다 (Order 생성하고, Valid Check하고, Errors들을 모아서 보여준다.)
		2.
	*/
}

func (r *mutationResolver) CancelOrder(ctx context.Context, orderID string) (*model.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ConfirmOrder(ctx context.Context, orderID string) (*model.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RequestPayment(ctx context.Context, input *model.PaymentClientInput) (*model.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CancelPayment(ctx context.Context, input *model.PaymentClientInput) (*model.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) HandlePaymentResponse(ctx context.Context, input *model.OrderResponse) (*model.PaymentResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context, id string) (*model.OrderInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.OrderInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
