package basket

import (
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

func Refresh(cart *domain.Basket) (*domain.Basket, bool, string) {
	totalPrices := 0
	globalErrs := []string{}

	for _, item := range cart.Items {
		item.Product, _ = ioc.Repo.Products.Get(item.Product.ID.Hex())

		errs := []string{}
		currentPrice := item.Product.ProductInfo.Price.CurrentPrice
		totalPrices += currentPrice * item.Quantity

		if !item.Product.OnSale {
			errs = append(errs, "ERR200:productgroup timeout"+item.Product.ID.Hex())
		}

		isValidSize, isValidQuantity := false, false
		for _, inv := range item.Product.ProductInfo.Inventory {
			if inv.Size == item.Size {
				isValidSize = true
				if inv.Quantity >= item.Quantity {
					isValidQuantity = true
				}
			}
		}

		if item.Product.ProductInfo.IsSoldout {
			errs = append(errs, "ERR105:product soldout")
		}

		if !item.Product.OnSale {
			errs = append(errs, "ERR102:alloffproduct is not for sale")
		}

		if !isValidSize {
			errs = append(errs, "ERR104:invalid product option size "+item.Size)
		}

		if !isValidQuantity {
			errs = append(errs, "ERR103:invalid product option quantity "+item.Size)
		}

		item.Errors = errs
		globalErrs = append(globalErrs, errs...)
	}

	if cart.ProductPrice != totalPrices {
		globalErrs = append(globalErrs, "ERR101:invalid total products price order amount")
		cart.ProductPrice = totalPrices
	}

	newBasket, err := ioc.Repo.Carts.Upsert(cart)
	if err != nil {
		config.Logger.Error("cart update error in valid check", zap.Error(err))
		return newBasket, false, err.Error()
	}

	if len(globalErrs) > 0 {
		return newBasket, false, globalErrs[0]
	}

	return newBasket, true, ""
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
			Removed:                item.Product.IsRemoved,
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
			UserID:                 user.ID.Hex(),
			User:                   user,
		})

		totalProductPrice += productPrice * item.Quantity
	}

	userID := ""
	if user != nil {
		userID = user.ID.Hex()
	}

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
