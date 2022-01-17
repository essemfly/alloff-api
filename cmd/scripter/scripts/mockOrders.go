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
	mobileNumber := "010-0000-0000"
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
		domain.ORDER_ITEM_EXCHANGE_RETURN_PENDING,
		domain.ORDER_ITEM_EXCHANGE_PENDING,
		domain.ORDER_ITEM_EXCHANGE_FINISHED,
		domain.ORDER_ITEM_RETURN_REQUESTED,
		domain.ORDER_ITEM_RETURN_PENDING,
		domain.ORDER_ITEM_RETURN_FINISHED}

	brandIdx := 0

	for brandIdx < 5 {
		log.Println("mock order created # " + strconv.Itoa(brandIdx))
		brandDao, _, _ := ioc.Repo.Brands.List(0, 100, nil, nil)
		products, cnt, _ := product.ProductsListing(0, 3, brandDao[brandIdx].ID.Hex(), "", "", nil)
		for cnt < 1 {
			brandIdx += 1
			log.Println("HOIT?", brandDao[brandIdx].ID.Hex())
			products, cnt, _ = product.ProductsListing(0, len(allstatus), brandDao[brandIdx].ID.Hex(), "", "", nil)
		}

		basket := BuildBasket(products)

		orderDao, err := basket.BuildOrder(userDao)
		if err != nil {
			log.Println("err occured in build order", err)
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
		BuyerMobile:   "010-9771-1882",
		BuyerAddress:  "서울광역시 강남구 역삼동 스파크플러스 선릉3호점",
		BuyerPostCode: "08365",
		Company:       "alloff",
		AppScheme:     "appscheme",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
