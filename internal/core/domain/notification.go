package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotificationType string

const (
	NOTIFICATION_PRODUCT_DIFF_NOTIFICATION      = NotificationType("PRODUCT_DIFF_NOTIFICATION")
	NOTIFICATION_BRAND_NEW_PRODUCT_NOTIFICATION = NotificationType("BRAND_NEW_PRODUCT_NOTIFICATION")
	NOTIFICATION_TIMEDEAL_OPEN_NITIFICATION     = NotificationType("TIMEDEAL_OPEN_NITIFICATION")
	NOTIFICATION_BRAND_OPEN_NOTIFICATION        = NotificationType("BRAND_OPEN_NOTIFICATION")
	NOTIFICATION_EVENT_NOTIFICATION             = NotificationType("EVENT_NOTIFICATION")
	NOTIFICATION_GENERAL_NOTIFICATION           = NotificationType("GENERAL_NOTIFICATION")
)

type NotificationStatus string

const (
	NOTIFICATION_READY     = NotificationStatus("READY")
	NOTIFICATION_IN_QUEUE  = NotificationStatus("IN_QUEUE")
	NOTIFICATION_CANCELED  = NotificationStatus("CANCELED")
	NOTIFICATION_SUCCEEDED = NotificationStatus("SUCCEEDED")
	NOTIFICATION_FAILED    = NotificationStatus("FAILED")
)

type NotificationDAO struct {
	ID               primitive.ObjectID `bson:"_id, omitempty"`
	Status           NotificationStatus
	NotificationType NotificationType
	Title            string
	Message          string
	Notificationid   string
	DeviceIDs        []string
	NavigateTo       string
	ReferenceID      string
	Created          time.Time
	Updated          time.Time
	Sended           time.Time
}
