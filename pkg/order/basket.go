package order

import (
	"fmt"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/lessbutter/alloff-api/pkg/product"
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
		currentPrice := product.GetCurrentPrice(item.Product)
		totalPrices += currentPrice * item.Quantity

		if item.ProductGroup != nil {
			if !item.ProductGroup.IsLive() {
				errs = append(errs, fmt.Errorf("ERR200:productgroup timeout"+item.Product.ID.Hex()))
			}
		}

		isValidSize, isValidQuantity := false, false
		for _, inv := range item.Product.Inventory {
			if inv.Size == item.Size {
				isValidSize = true
				if inv.Quantity >= item.Quantity {
					isValidQuantity = true
				} else {
					item.Quantity = inv.Quantity
				}
			}
		}

		if item.Product.Soldout {
			item.Quantity = 0
			errs = append(errs, fmt.Errorf("ERR105:product soldout"))
		}

		if item.Product.Removed {
			item.Quantity = 0
			errs = append(errs, fmt.Errorf("ERR102:alloffproduct is removed"))
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

func (basket *Basket) BuildOrder(user *domain.UserDAO) (*domain.OrderDAO, error) {
	orderItems := []*domain.OrderItemDAO{}
	orderAlloffID := xid.New().String()

	totalProductPrice := 0
	for _, item := range basket.Items {
		orderItemType := domain.NORMAL_ORDER
		productPrice := product.GetCurrentPrice(item.Product)
		if item.ProductGroup != nil {
			if item.ProductGroup.GroupType == domain.PRODUCT_GROUP_EXHIBITION {
				orderItemType = domain.EXHIBITION_ORDER
			} else if item.ProductGroup.GroupType == domain.PRODUCT_GROUP_TIMEDEAL {
				orderItemType = domain.TIMEDEAL_ORDER
			} else if item.ProductGroup.GroupType == domain.PRODUCT_GROUP_GROUPDEAL {
				orderItemType = domain.GROUPDEAL_ORDER
			}
		}

		itemCode := utils.CreateShortUUID()

		_, err := ioc.Repo.OrderItems.GetByCode(itemCode)
		for err == nil {
			itemCode = utils.CreateShortUUID()
		}

		exhibitionID := ""
		if item.ProductGroup != nil {
			exhibitionID = item.ProductGroup.ExhibitionID
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
			SalesPrice:             productPrice,
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
