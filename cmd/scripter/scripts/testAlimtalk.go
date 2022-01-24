package scripts

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/alimtalk"
	"github.com/lessbutter/alloff-api/internal/utils"
)

func TestAlimtalk() {
	user, _ := ioc.Repo.Users.GetByMobile("01097711882")
	newAlimtalk := &domain.AlimtalkDAO{
		UserID:       user.ID.Hex(),
		Mobile:       user.Mobile,
		TemplateCode: domain.PAYMENT_OK,
		ReferenceID:  "Test alimtalk payment ref id",
		SendDate:     nil,
		TemplateParams: map[string]string{
			"orderedAt":   utils.TimeFormatterForKorea(time.Now()),
			"productName": "test payment id",
			"amount":      utils.PriceFormatter(135000),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	requestId, err := alimtalk.SendMessage(newAlimtalk)
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
