package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
)

type PaymentStatusEnum string

const (
	PAYMENT_CREATED          = PaymentStatusEnum("PAYMENT_CREATED")
	PAYMENT_CONFIRMED        = PaymentStatusEnum("PAYMENT_CONFIRMED")
	PAYMENT_TIME_OUT         = PaymentStatusEnum("PAYMENT_TIME_OUT")
	PAYMENT_CANCELED         = PaymentStatusEnum("PAYMENT_CANCELED")
	PAYMENT_REFUND_REQUESTED = PaymentStatusEnum("PAYMENT_REFUND_REQUESTED")
	PAYMENT_REFUND_FINISHED  = PaymentStatusEnum("PAYMENT_REFUND_FINISHED")
)

type PaymentMethod struct {
	Label string
	Code  string
}

type PaymentDAO struct {
	tableName             struct{} `pg:"payments"`
	ID                    int
	ImpUID                string            `pg:"imp_uid"`
	PaymentStatus         PaymentStatusEnum `pg:"payment_status"`
	Pg                    string            `pg:"pg"`
	PayMethod             string            `pg:"pay_method"`
	Name                  string            `pg:"name"`
	MerchantUid           string            `pg:"merchant_uid"`
	Amount                int               `pg:"amount"`
	BuyerName             string            `pg:"buyer_name"`
	BuyerMobile           string            `pg:"buyer_mobile"`
	BuyerAddress          string            `pg:"buyer_address"`
	BuyerPostCode         string            `pg:"buyer_post_code"`
	PersonalCustomsNumber string            `pg:"personal_customs_number"`
	Company               string            `pg:"company"`
	AppScheme             string            `pg:"app_scheme"`
	CreatedAt             time.Time         `pg:"created_at"`
	UpdatedAt             time.Time         `pg:"updated_at"`
}

func (paymentDao *PaymentDAO) GetPaymentMethods() []*model.PaymentMethod {
	// (TODO) To be specified in collection
	paymentMethods := []*model.PaymentMethod{
		{
			Label: "다날",
			Code:  "danal_tpay",
		},
		{
			Label: "카카오페이",
			Code:  "kakaopay",
		},
	}
	return paymentMethods
}
