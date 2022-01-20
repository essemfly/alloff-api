package order

import (
	"errors"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/rs/xid"
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

func (basket *Basket) BuildOrder(user *domain.UserDAO) (*domain.OrderDAO, error) {
	orderItems := []*domain.OrderItemDAO{}
	orderAlloffID := xid.New().String()

	totalProductPrice := 0
	for _, item := range basket.Items {
		orderItemType := domain.NORMAL_ORDER
		if item.ProductGroup != nil {
			// (TODO) 기획전이 생기면 추가되어야함.
			orderItemType = domain.TIMEDEAL_ORDER
		}

		itemCode := utils.CreateShortUUID()

		_, err := ioc.Repo.OrderItems.GetByCode(itemCode)
		for err == nil {
			itemCode = utils.CreateShortUUID()
		}

		orderItems = append(orderItems, &domain.OrderItemDAO{
			OrderItemCode:          itemCode,
			ProductID:              item.Product.ID.Hex(),
			ProductName:            item.Product.AlloffName,
			ProductUrl:             item.Product.ProductInfo.ProductUrl,
			ProductImg:             item.Product.ProductInfo.Images[0],
			BrandKeyname:           item.Product.ProductInfo.Brand.KeyName,
			BrandKorname:           item.Product.ProductInfo.Brand.KorName,
			Removed:                item.Product.Removed,
			SalesPrice:             item.Product.DiscountedPrice,
			CancelDescription:      item.Product.SalesInstruction.CancelDescription,
			DeliveryDescription:    item.Product.SalesInstruction.DeliveryDescription,
			OrderItemType:          orderItemType,
			OrderItemStatus:        domain.ORDER_ITEM_CREATED,
			DeliveryTrackingNumber: []string{},
			DeliveryTrackingUrl:    []string{},
			Size:                   item.Size,
			Quantity:               item.Quantity,
			CreatedAt:              time.Now(),
			UpdatedAt:              time.Now(),
		})

		totalProductPrice += item.Product.DiscountedPrice * item.Quantity
	}

	// (TODO) Delivery Price는 생성시점에 만들어질 예정이다?
	newOrderDao := &domain.OrderDAO{
		AlloffOrderID: orderAlloffID,
		UserID:        user.ID.Hex(),
		User:          user,
		OrderItems:    orderItems,
		TotalPrice:    totalProductPrice,
		ProductPrice:  totalProductPrice,
		DeliveryPrice: 0,
		RefundPrice:   0,
		UserMemo:      "",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return newOrderDao, nil
}
