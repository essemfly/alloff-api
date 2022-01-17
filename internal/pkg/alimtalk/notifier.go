package alimtalk

import (
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
)

func NotifyPaymentSuccessAlarm(payment *domain.PaymentDAO) {
	order, _ := ioc.Repo.Orders.GetByAlloffID(payment.MerchantUid)
	user, _ := ioc.Repo.Users.Get(order.User.ID.Hex())

	newAlimtalk := &domain.AlimtalkDAO{
		UserID:       user.ID.Hex(),
		Mobile:       user.Mobile,
		TemplateCode: domain.PAYMENT_OK,
		ReferenceID:  strconv.Itoa(payment.ID),
		SendDate:     nil,
		TemplateParams: map[string]string{
			"orderedAt":   utils.TimeFormatterForKorea(order.UpdatedAt),
			"productName": payment.Name,
			"amount":      utils.PriceFormatter(payment.Amount),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	requestId, err := SendMessage(newAlimtalk)
	if err != nil {
		newAlimtalk.Status = domain.ALIMTALK_STATUS_FAILED
	} else {
		newAlimtalk.ToastRequestID = requestId
		newAlimtalk.Status = domain.ALIMTALK_STATUS_COMPLETED
	}

	_, err = ioc.Repo.Alimtalks.Insert(newAlimtalk)
	if err != nil {
		log.Println(err)
	}
}
