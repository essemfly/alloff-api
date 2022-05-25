package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) AddCartItem(ctx context.Context, input *model.AddCartItemInput) (*model.Cart, error) {
	cart, err := ioc.Repo.Carts.Get(input.CartID)
	if err != nil {
		return nil, err
	}

	product, err := ioc.Repo.Products.Get(input.ProductID)
	if err != nil {
		return nil, err
	}

	cart.Items = append(cart.Items, &domain.BasketItem{
		Product:        product,
		ProductGroupID: product.ProductGroupID,
		ExhibitionID:   product.ExhibitionID,
		Size:           input.Selectsize,
		Quantity:       input.Quantity,
	})

	newCart, err := ioc.Repo.Carts.Upsert(cart)
	if err != nil {
		return nil, err
	}

	return mapper.MapCart(newCart), nil
}

func (r *mutationResolver) RemoveCartItem(ctx context.Context, cartID string, productID string) (*model.Cart, error) {
	cart, err := ioc.Repo.Carts.Get(cartID)
	if err != nil {
		return nil, err
	}

	newItems := []*domain.BasketItem{}
	for _, item := range cart.Items {
		if item.Product.ID.Hex() != productID {
			newItems = append(newItems, item)
		}
	}

	cart.Items = newItems
	newCart, err := ioc.Repo.Carts.Upsert(cart)
	if err != nil {
		return nil, err
	}

	return mapper.MapCart(newCart), nil
}

func (r *queryResolver) Cart(ctx context.Context, id string) (*model.Cart, error) {
	if id == "" {
		newCartDAO := &domain.Basket{
			ID:           primitive.NewObjectID(),
			Items:        []*domain.BasketItem{},
			ProductPrice: 0,
		}
		return mapper.MapCart(newCartDAO), nil
	}

	cart, err := ioc.Repo.Carts.Get(id)
	if err != nil {
		return nil, err
	}
	return mapper.MapCart(cart), nil
}
