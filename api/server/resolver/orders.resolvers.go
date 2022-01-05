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
		4. Order를 Create해줄 필요는 없다.
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
		2. 여기서 완료되면 Order가 생성이되고, 주문 결제하는 창으로 넘어간다.
	*/
}

func (r *mutationResolver) CancelOrder(ctx context.Context, orderID string) (*model.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ConfirmOrder(ctx context.Context, orderID string) (*model.PaymentStatus, error) {
	// Order Confirm하는 API 인데, 이걸 앱에서 쓰는 게 없을 것 같습니다.
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
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.OrderInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
