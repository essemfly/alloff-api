package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlimtalkStatus string

const (
	DELIVERY_START       = "deliveryStarted"
	DELIVERY_START_TRACK = "deliveryStartedTrack"
	PAYMENT_CANCEL       = "paymentCanceled"
	PAYMENT_OK           = "paymentOk"
	TIMEDEAL_OPEN_NOTI   = "tdNotifyOpen"
	TIMEDEAL_NOTI_IN     = "tdOptInOk"
	TIMEDEAL_NOTI_OUT    = "tdOptOutOk"
)

const (
	ALIMTALK_STATUS_READY     = AlimtalkStatus("ALIMTALK_STATUS_READY")
	ALIMTALK_STATUS_CANCLED   = AlimtalkStatus("ALIMTALK_STATUS_CANCLED")
	ALIMTALK_STATUS_FAILED    = AlimtalkStatus("ALIMTALK_STATUS_CANCLED")
	ALIMTALK_STATUS_COMPLETED = AlimtalkStatus("ALIMTALK_STATUS_COMPLETED")
)

// ReferenceID could be productGroupId or paymentId
type AlimtalkDAO struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Source         string
	UserID         string
	Mobile         string
	TemplateCode   string
	ReferenceID    string
	ToastRequestID string
	TemplateParams map[string]string
	Status         AlimtalkStatus
	SendDate       *time.Time
	Created        time.Time
	Updated        time.Time
}

// func (alim *AlimtalkDAO) FillTemplate() *AlimtalkDAO {
// 	switch alim.TemplateCode {
// 	case PAYMENT_OK:
// 		payment, err := ioc.Repo.Payments.GetByImpUID(alim.ReferenceID)
// 		if err != nil {
// 			return nil
// 		}
// 		order, err = ioc.Repo.Orders.GetByAlloffID(payment.Mer)
// 		alim.TemplateParams = map[string]string{
// 			"orderedAt":   utils.TimeFormatterForKorea(order.Updated),
// 			"productName": payment.Name,
// 			"amount":      utils.PriceFormatter(payment.Amount),
// 		}
// 	default:
// 		break
// 	}

// 	return alim
// }
