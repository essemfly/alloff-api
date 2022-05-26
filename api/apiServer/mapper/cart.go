package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapCart(basket *domain.Basket) *model.Cart {
	items := []*model.CartItem{}
	for _, item := range basket.Items {
		items = append(items, MapCartItem(item))
	}
	return &model.Cart{
		ID:    basket.ID.Hex(),
		Items: items,
	}
}

func MapCartItem(basketItem *domain.BasketItem) *model.CartItem {
	return &model.CartItem{
		Product:   MapProduct(basketItem.Product),
		Quantity:  basketItem.Quantity,
		Size:      basketItem.Size,
		ErrorMsgs: basketItem.Errors,
	}
}
