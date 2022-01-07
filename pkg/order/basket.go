package order

import (
	"errors"

	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type Basket struct {
	Items        []*BasketItem
	ProductPrice int
}

type BasketItem struct {
	Product      *domain.ProductDAO
	ProductGroup *domain.ProductGroupDAO
	Size         string
	Quantity     int
}

func (basket *Basket) IsValid() []error {
	errs := []error{}
	totalPrices := 0

	for _, item := range basket.Items {
		totalPrices += item.Product.DiscountedPrice
		isValidSize, isValidQuantity := false, false
		for _, inv := range item.Product.Inventory {
			if inv.Size == item.Size {
				isValidSize = true
				if inv.Quantity > item.Quantity {
					isValidQuantity = true
				}
			}
		}
		if !isValidSize {
			errs = append(errs, errors.New("invalid product option "+item.Product.ProductInfo.ProductID))
		}

		if !isValidQuantity {
			errs = append(errs, errors.New("invalid product option quantity "+item.Product.ProductInfo.ProductID))
		}
	}

	if basket.ProductPrice != totalPrices {
		errs = append(errs, errors.New("invalid total products price"))
	}

	return errs
}

func (basket *Basket) BuildOrder() (*domain.OrderDAO, error) {
	newOrderDao := &domain.OrderDAO{}

	return newOrderDao, nil
}
