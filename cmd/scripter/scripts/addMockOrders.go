package scripts

import (
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/order"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func AddMockOrders() {
	log.Println("Add Mock Orders")
	mobileNumber := "01000000000"
	userDao, _ := ioc.Repo.Users.GetByMobile(mobileNumber)

	allstatus := []domain.OrderItemStatusEnum{
		domain.ORDER_ITEM_PAYMENT_FINISHED,
		domain.ORDER_ITEM_PRODUCT_PREPARING,
		domain.ORDER_ITEM_FOREIGN_PRODUCT_INSPECTING,
		domain.ORDER_ITEM_DELIVERY_PREPARING,
		domain.ORDER_ITEM_FOREIGN_DELIVERY_STARTED,
		domain.ORDER_ITEM_DELIVERY_STARTED,
		domain.ORDER_ITEM_DELIVERY_FINISHED,
		domain.ORDER_ITEM_CONFIRM_PAYMENT,
		domain.ORDER_ITEM_CANCEL_FINISHED,
		domain.ORDER_ITEM_EXCHANGE_REQUESTED,
		domain.ORDER_ITEM_EXCHANGE_PENDING,
		domain.ORDER_ITEM_EXCHANGE_FINISHED,
		domain.ORDER_ITEM_RETURN_REQUESTED,
		domain.ORDER_ITEM_RETURN_PENDING,
		domain.ORDER_ITEM_RETURN_FINISHED}

	brandIdx := 0

	for brandIdx < 1 {
		log.Println("mock order created # " + strconv.Itoa(brandIdx))
		products, _, _ := product.Listing(product.ProductListInput{
			Offset:                    0,
			Limit:                     len(allstatus),
			BrandID:                   "61d699ec74b2b71fe80ff58a",
			IncludeSpecialProductType: product.NOT_SPECIAL_PRODUCTS,
			IncludeClassifiedType:     product.NO_MATTER_CLASSIFIED,
		})

		basket := BuildBasket(products)

		orderDao, err := basket.BuildOrder(userDao)
		if err != nil {
			log.Println("err occured in build order", err)
		}

		orderDao.OrderStatus = domain.ORDER_PAYMENT_FINISHED
		_, err = ioc.Repo.Orders.Insert(orderDao)
		if err != nil {
			log.Println("err occured in adding order", err)
		}

		paymentDao := BuildPaymentDao(orderDao)
		_, err = ioc.Repo.Payments.Insert(paymentDao)
		if err != nil {
			log.Println("err occured in build payments", err)
		}
		for idx, item := range orderDao.OrderItems {
			item.OrderItemStatus = allstatus[idx]
			_, err := ioc.Repo.OrderItems.Update(item)
			if err != nil {
				log.Println("err occured in changing orderitems")
			}
		}
		brandIdx += 1
	}

}

func BuildBasket(products []*domain.ProductDAO) *order.Basket {
	basketItems := []*order.BasketItem{}
	totalPrices := 0
	for idx, item := range products {
		quantity := 1
		if idx == 1 {
			quantity = 2
		}

		basketItem := &order.BasketItem{
			Product:      item,
			ProductGroup: nil,
			Size:         item.Inventory[0].Size,
			Quantity:     quantity,
		}
		totalPrices += item.DiscountedPrice * quantity
		basketItems = append(basketItems, basketItem)
	}

	return &order.Basket{
		Items:        basketItems,
		ProductPrice: totalPrices,
	}

}

func BuildPaymentDao(orderDao *domain.OrderDAO) *domain.PaymentDAO {
	return &domain.PaymentDAO{
		PaymentStatus: domain.PAYMENT_CONFIRMED,
		ImpUID:        "test_" + orderDao.AlloffOrderID,
		Pg:            "mocktest",
		PayMethod:     "card",
		MerchantUid:   orderDao.AlloffOrderID,
		Amount:        orderDao.TotalPrice,
		Name:          "테스트 상품 2개 이상?",
		BuyerName:     "테스트 이석민",
		BuyerMobile:   "01097711882",
		BuyerAddress:  "서울광역시 강남구 역삼동 스파크플러스 선릉3호점",
		BuyerPostCode: "08365",
		Company:       "alloff",
		AppScheme:     "appscheme",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
