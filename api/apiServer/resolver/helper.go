package resolver

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func BuildBasketItems(input *model.OrderInput) ([]*domain.BasketItem, error) {
	basketItems := []*domain.BasketItem{}
	for _, item := range input.Orders {
		pd, err := ioc.Repo.Products.Get(item.ProductID)
		if err != nil {
			return nil, err
		}

		basketItem := &domain.BasketItem{
			Product:      pd,
			ExhibitionID: pd.ExhibitionID,
			Size:         item.Selectsize,
			Quantity:     item.Quantity,
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
