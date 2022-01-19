package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/front/model"
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
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (paymentDao *PaymentDAO) ToDTO() *model.PaymentInfo {
	return &model.PaymentInfo{
		Pg:            paymentDao.Pg,
		PayMethod:     paymentDao.PayMethod,
		Name:          paymentDao.Name,
		MerchantUID:   paymentDao.MerchantUid,
		Amount:        paymentDao.Amount,
		BuyerName:     paymentDao.BuyerName,
		BuyerMobile:   paymentDao.BuyerMobile,
		BuyerAddress:  paymentDao.BuyerAddress,
		BuyerPostCode: paymentDao.BuyerPostCode,
		Company:       paymentDao.Company,
		AppScheme:     paymentDao.AppScheme,
	}
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
