package resolver

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/order"
)

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

func BuildPaymentDao(input *model.PaymentClientInput) *domain.PaymentDAO {
	personalCustomsNumber := ""
	if input.PersonalCustomsNumber != nil {
		personalCustomsNumber = *input.PersonalCustomsNumber
	}
	return &domain.PaymentDAO{
		Pg:                    input.Pg,
		PayMethod:             input.PayMethod,
		MerchantUid:           input.MerchantUID,
		Amount:                input.Amount,
		Name:                  *input.Name,
		BuyerName:             *input.BuyerName,
		BuyerMobile:           *input.BuyerMobile,
		BuyerAddress:          *input.BuyerAddress,
		BuyerPostCode:         *input.BuyerPostCode,
		PersonalCustomsNumber: personalCustomsNumber,
		Company:               "alloff",
		AppScheme:             "appscheme",
	}
}
