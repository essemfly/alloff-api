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
}

func (r *mutationResolver) RequestOrder(ctx context.Context, input *model.OrderInput) (*model.OrderWithPayment, error) {
	panic(fmt.Errorf("not implemented"))
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
