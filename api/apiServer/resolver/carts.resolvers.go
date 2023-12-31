package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/basket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (r *mutationResolver) AddCartItem(ctx context.Context, cartID string, items []*model.CartItemInput) (*model.Cart, error) {
	cart, err := ioc.Repo.Carts.Get(cartID)
	if err != nil {
		return nil, err
	}

	productPrices := 0
	for _, cartItem := range items {
		product, err := ioc.Repo.Products.Get(cartItem.ProductID)
		if err != nil {
			return nil, err
		}

		isNewProduct := true
		for _, item := range cart.Items {
			if item.Product.ID == product.ID && item.Size == cartItem.Selectsize {
				item.Quantity += cartItem.Quantity
				isNewProduct = false
			}
			productPrices += item.Quantity * item.Product.ProductInfo.Price.CurrentPrice
		}

		if isNewProduct {
			cart.Items = append(cart.Items, &domain.BasketItem{
				Product:      product,
				ExhibitionID: product.ExhibitionID,
				Size:         cartItem.Selectsize,
				Quantity:     cartItem.Quantity,
			})
			productPrices += cartItem.Quantity * product.ProductInfo.Price.CurrentPrice
		}
	}

	cart.ProductPrice = productPrices
	newCart, _, _ := basket.Refresh(cart)
	if err != nil {
		return nil, err
	}

	return mapper.MapCart(newCart), nil
}

func (r *mutationResolver) RemoveCartItem(ctx context.Context, cartID string, items []*model.CartItemInput) (*model.Cart, error) {
	cart, err := ioc.Repo.Carts.Get(cartID)
	if err != nil {
		return nil, err
	}

	productPrices := 0
	newItems := []*domain.BasketItem{}
	for _, item := range cart.Items {
		isRemovedProduct := false
		for _, removedItem := range items {
			if removedItem.ProductID == item.Product.ID.Hex() && removedItem.Selectsize == item.Size {
				isRemovedProduct = true
				item.Quantity -= removedItem.Quantity
				if item.Quantity > 0 {
					newItems = append(newItems, item)
					productPrices += item.Quantity * item.Product.ProductInfo.Price.CurrentPrice
				}
				break
			}
		}
		if !isRemovedProduct {
			newItems = append(newItems, item)
			productPrices += item.Quantity * item.Product.ProductInfo.Price.CurrentPrice
		}
	}

	cart.Items = newItems
	cart.ProductPrice = productPrices
	newCart, _, _ := basket.Refresh(cart)
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
		_, err := ioc.Repo.Carts.Upsert(newCartDAO)
		if err != nil {
			config.Logger.Error("cart create err", zap.Error(err))
			return nil, err
		}
		return mapper.MapCart(newCartDAO), nil
	}

	cart, err := ioc.Repo.Carts.Get(id)
	if err != nil {
		return nil, err
	}

	newCart, _, _ := basket.Refresh(cart)

	return mapper.MapCart(newCart), nil
}
