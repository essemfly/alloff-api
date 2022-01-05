package order

import (
	"errors"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type OrderService struct {
	OrderItems []*domain.OrderItemDAO
	Errors     error
}

func (os *OrderService) IsValidItem(orderItem *domain.OrderItemDAO) error {
	pd, err := ioc.Repo.Products.Get(orderItem.ProductID)
	if err != nil {
		return errors.New("product not found on this product id")
	}

	if pd.Removed || pd.Soldout {
		return errors.New("product is sold out or removed")
	}

	isValidSize, isValidQuantity := false, false
	for _, inv := range pd.Inventory {
		if inv.Size == orderItem.Size {
			isValidSize = true
			if inv.Quantity > orderItem.Quantity {
				isValidQuantity = true
			}
		}
	}

	if !isValidSize {
		return errors.New("invalid product option")
	}

	if !isValidQuantity {
		return errors.New("invalid product option quantity")
	}

	return nil
}

func (os *OrderService) AddProduct(productID, size string, quantity int) error {
	if productID == "" {
		return errors.New("no product id provided")
	}

	pd, err := ioc.Repo.Products.Get(productID)
	if err != nil {
		return errors.New("product not found on this product id")
	}

	orderItemDao := &domain.OrderItemDAO{
		ProductID:           pd.ID.Hex(),
		ProductName:         pd.AlloffName,
		BrandKeyname:        pd.ProductInfo.Brand.KeyName,
		SalesPrice:          pd.DiscountedPrice,
		Instruction:         nil, // Instruction 어떤것?
		CancelDescription:   pd.SalesInstruction.CancelDescription,
		DeliveryDescription: pd.SalesInstruction.DeliveryDescription,
		OrderType:           domain.OrderTypeEnum(domain.NORMAL_ORDER),
		OrderStatus:         domain.OrderStatusEnum(domain.ORDER_CREATED),
		Size:                size,
		Quantity:            quantity,
	}

	err = os.IsValidItem(orderItemDao)
	if err != nil {
		return err
	}

	os.OrderItems = append(os.OrderItems, orderItemDao)

	return nil
}

func CheckValidOrderInput()
