package domain

import "time"

type PaymentStatusEnum string

const (
	PAYMENT_CREATED          = PaymentStatusEnum("PAYMENT_CREATED")
	PAYMENT_CONFIRMED        = PaymentStatusEnum("PAYMENT_CONFIRMED")
	PAYMENT_TIME_OUT         = PaymentStatusEnum("PAYMENT_TIME_OUT")
	PAYMENT_REFUND_REQUESTED = PaymentStatusEnum("PAYMENT_REFUND_REQUESTED")
	PAYMENT_REFUND_FINISHED  = PaymentStatusEnum("PAYMENT_REFUND_FINISHED")
)

type PaymentDAO struct {
	tableName     struct{} `pg:"payments"`
	ID            int      `bson:"_id, omitempty"`
	ImpUID        string
	PaymentStatus PaymentStatusEnum
	Pg            string
	PayMethod     string
	Name          string
	MerchantUid   string
	Amount        int
	BuyerName     string
	BuyerMobile   string
	BuyerAddress  string
	BuyerPostCode string
	Company       string
	AppScheme     string
	Created       time.Time
	Updated       time.Time
}
