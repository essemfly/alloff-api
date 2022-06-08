package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlimtalkStatus string

const (
	DELIVERY_START       = "deliveryStarted"
	DELIVERY_START_TRACK = "deliveryStartedTrack"
	PAYMENT_CANCEL       = "paymentCanceled2"
	PAYMENT_OK           = "paymentOk2"
	PAYMENT_OK_OVERSEAS  = "paymentOkOverseas"
	EXHIBITION_ALARM     = "exhibitionAlarmOk"
	DEAL_OPEN            = "timedealOpenNoti2"
)

const (
	ALIMTALK_STATUS_READY     = AlimtalkStatus("ALIMTALK_STATUS_READY")
	ALIMTALK_STATUS_CANCLED   = AlimtalkStatus("ALIMTALK_STATUS_CANCLED")
	ALIMTALK_STATUS_FAILED    = AlimtalkStatus("ALIMTALK_STATUS_FAILED")
	ALIMTALK_STATUS_COMPLETED = AlimtalkStatus("ALIMTALK_STATUS_COMPLETED")
)

// ReferenceID could be productGroupId or paymentId
type AlimtalkDAO struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Source         string
	UserID         string
	Mobile         string
	TemplateCode   string
	ReferenceID    string // 알림톡이 참조하는 무언가의 id (Payment id 같은것)
	ToastRequestID string
	TemplateParams map[string]string
	Status         AlimtalkStatus
	SendDate       *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
