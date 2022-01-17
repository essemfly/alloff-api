package slack

import (
	"strconv"
	"strings"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func WritePaymentSuccessSlack(payment *domain.PaymentDAO) {
	order, _ := ioc.Repo.Orders.GetByAlloffID(payment.MerchantUid)
	orderProducts := []string{}
	for _, orderItem := range order.OrderItems {
		orderProducts = append(orderProducts, string(orderItem.OrderItemType)+" "+orderItem.ProductName+": "+orderItem.Size+" "+strconv.Itoa(orderItem.Quantity)+"개"+" : "+orderItem.ProductUrl)
	}

	orderMsg :=
		"**결제 완료** \n" +
			"결제 ID: " + payment.ImpUID + "\n" +
			"주문 ID: " + payment.MerchantUid + "\n" +
			"주문명: " + payment.Name + "\n" +
			"주문 정보: \n" + strings.Join(orderProducts[:], ", ") + "\n" +
			"주문자 번호: " + payment.BuyerMobile + "\n" +
			"가격: " + strconv.Itoa(payment.Amount) + "\n" +
			"주소: " + payment.BuyerPostCode + " " + payment.BuyerAddress + "\n" +
			"받는 사람 번호: " + payment.BuyerMobile

	config.WriteOrderMessage(orderMsg)
}
