package order

import (
	"fmt"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

func IsValid(basket *domain.Basket) []error {
	errs := []error{}
	totalPrices := 0

	for _, item := range basket.Items {
		currentPrice := item.Product.ProductInfo.Price.CurrentPrice
		totalPrices += currentPrice * item.Quantity

		if item.Product.IsNotSale {
			errs = append(errs, fmt.Errorf("ERR200:productgroup timeout"+item.Product.ID.Hex()))
		}

		isValidSize, isValidQuantity := false, false
		for _, inv := range item.Product.ProductInfo.AlloffInventory {
			if inv.AlloffSize.AlloffSizeName == item.Size {
				isValidSize = true
				if inv.Quantity >= item.Quantity {
					isValidQuantity = true
				} else {
					item.Quantity = inv.Quantity
				}
			}
		}

		if item.Product.ProductInfo.IsSoldout {
			item.Quantity = 0
			errs = append(errs, fmt.Errorf("ERR105:product soldout"))
		}

		if item.Product.IsNotSale {
			item.Quantity = 0
			errs = append(errs, fmt.Errorf("ERR102:alloffproduct is not for sale"))
		}

		if !isValidSize {
			item.Quantity = 0
			errs = append(errs, fmt.Errorf("ERR104:invalid product option size"+item.Size))
		}

		if !isValidQuantity {
			errs = append(errs, fmt.Errorf("ERR103:invalid product option quantity"+item.Product.ID.Hex()))
		}
	}

	if basket.ProductPrice != totalPrices {
		errs = append(errs, fmt.Errorf("ERR101:invalid total products price order amount"))
	}

	return errs
}

func BuildOrder(user *domain.UserDAO, basket *domain.Basket) (*domain.OrderDAO, error) {
	orderItems := []*domain.OrderItemDAO{}
	orderAlloffID := xid.New().String()

	totalProductPrice := 0
	for _, item := range basket.Items {
		orderItemType := domain.NORMAL_ORDER
		productPrice := item.Product.ProductInfo.Price.CurrentPrice
		exDao, err := ioc.Repo.Exhibitions.Get(item.ExhibitionID)
		if err != nil {
			config.Logger.Error("pg not found in build order", zap.Error(err))
			return nil, err
		}
		if exDao.ExhibitionType == domain.EXHIBITION_NORMAL {
			orderItemType = domain.EXHIBITION_ORDER
		} else if exDao.ExhibitionType == domain.EXHIBITION_TIMEDEAL {
			orderItemType = domain.TIMEDEAL_ORDER
		} else if exDao.ExhibitionType == domain.EXHIBITION_GROUPDEAL {
			orderItemType = domain.GROUPDEAL_ORDER
		}

		itemCode := utils.CreateShortUUID()

		_, err = ioc.Repo.OrderItems.GetByCode(itemCode)
		for err == nil {
			itemCode = utils.CreateShortUUID()
		}

		exhibitionID := item.ExhibitionID

		orderItems = append(orderItems, &domain.OrderItemDAO{
			OrderItemCode:          itemCode,
			ProductID:              item.Product.ID.Hex(),
			ProductName:            item.Product.ProductInfo.AlloffName,
			ProductUrl:             item.Product.ProductInfo.ProductUrl,
			ProductImg:             item.Product.ProductInfo.Images[0],
			BrandKeyname:           item.Product.ProductInfo.Brand.KeyName,
			BrandKorname:           item.Product.ProductInfo.Brand.KorName,
			Removed:                item.Product.IsNotSale,
			SalesPrice:             productPrice,
			CancelDescription:      item.Product.ProductInfo.SalesInstruction.CancelDescription,
			DeliveryDescription:    item.Product.ProductInfo.SalesInstruction.DeliveryDescription,
			OrderItemType:          orderItemType,
			OrderItemStatus:        domain.ORDER_ITEM_CREATED,
			DeliveryTrackingNumber: []string{},
			DeliveryTrackingUrl:    []string{},
			Size:                   item.Size,
			Quantity:               item.Quantity,
			CreatedAt:              time.Now(),
			UpdatedAt:              time.Now(),
			ExhibitionID:           exhibitionID,
			CompanyKeyname:         item.Product.ProductInfo.Source.CrawlModuleName,
		})

		totalProductPrice += productPrice * item.Quantity
	}

	userID := ""
	if user != nil {
		userID = user.ID.Hex()
	}
	// (TODO) Delivery Price는 생성시점에 만들어질 예정이다?
	newOrderDao := &domain.OrderDAO{
		AlloffOrderID: orderAlloffID,
		OrderStatus:   domain.ORDER_CREATED,
		UserID:        userID,
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
