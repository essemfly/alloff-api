package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	server "github.com/lessbutter/alloff-api/api/server/model"
)

func (r *mutationResolver) CheckOrder(ctx context.Context, input *server.OrderInput) (*server.OrderValidityResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RequestOrder(ctx context.Context, input *server.OrderInput) (*server.OrderWithPayment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CancelOrder(ctx context.Context, orderID string) (*server.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ConfirmOrder(ctx context.Context, orderID string) (*server.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RequestPayment(ctx context.Context, input *server.PaymentClientInput) (*server.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CancelPayment(ctx context.Context, input *server.PaymentClientInput) (*server.PaymentStatus, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) HandlePaymentResponse(ctx context.Context, input *server.OrderResponse) (*server.PaymentResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Order(ctx context.Context, id string) (*server.OrderInfo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context) ([]*server.OrderInfo, error) {
	panic(fmt.Errorf("not implemented"))
}
