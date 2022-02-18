package iamport

import (
	"errors"
	"net/http"
	"time"

	TypePayment "github.com/iamport/interface/gen_src/go/v1/payment"
)

const (
	DefaultURL = "https://api.iamport.kr"

	ErrMustExistImpUID              = "iamport: imp_uid must be exist"
	ErrMustExistMerchantUID         = "iamport: Merchant UID must be exist"
	ErrMustExistImpUIDorMerchantUID = "iamport: imp_uid or Merchant UID must be exist"
	ErrMustExistCustomerUID         = "iamport: customer_uid must be exist"
	ErrInvalidStatusParam           = "iamport: status parmeter is invalid. must be all, ready, paid, failed and cancelled"
	ErrInvalidSortParam             = "iamport: sort parmeter is invalid. must be -started, started, -paid, paid, -updated and updated"
	ErrInvalidPage                  = "iamport: page is more than 1"
	ErrInvalidLimit                 = "iamport: limit is more than 0"
	ErrInvalidFrom                  = "iamport: 'from' cannot be more future than 'to'"
	ErrInvalidTo                    = "iamport: 'to' date cannot be more than 3 months."
	ErrInvalidAmount                = "iamport: amount is more than 0"
)

type Iamport struct {
	Authenticate *Authenticate
}

func NewIamport(apiURL string, restAPIKey string, restAPISecret string) (*Iamport, error) {
	client := &http.Client{}

	auth, err := NewAuthenticate(apiURL, client, restAPIKey, restAPISecret)
	if err != nil {
		return nil, err
	}

	iamport := &Iamport{
		Authenticate: auth,
	}

	return iamport, nil
}

// GetPaymentImpUID imp_uid로 결제 정보 가져오기
//
// GET /payments/{imp_uid}
func (iamport *Iamport) GetPaymentImpUID(iuid string) (*TypePayment.Payment, error) {
	if iuid == "" {
		return nil, errors.New(ErrMustExistImpUID)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	reqPaymentImpUID := &TypePayment.PaymentRequest{
		ImpUid: iuid,
	}

	res, err := GetByImpUID(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, reqPaymentImpUID,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentsImpUIDs 여러개의 imp_uid로 결제 정보 가져오기
//
// GET /payments
func (iamport *Iamport) GetPaymentsImpUIDs(iuids []string) ([]*TypePayment.Payment, error) {
	if len(iuids) < 0 {
		return nil, errors.New(ErrMustExistImpUID)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentsRequest{
		ImpUid: iuids,
	}

	res, err := GetByImpUIDs(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentMerchantUID merchant_uid로 결제 정보 가져오기
//
// GET /payments/find/{merchant_uid}
func (iamport *Iamport) GetPaymentMerchantUID(muid string, status string, sorting string) (*TypePayment.Payment, error) {
	if muid == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	if !ValidateStatusParameter(status) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	if !ValidateSortParameter(sorting) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	merchantUIDPaymentReq := &TypePayment.PaymentMerchantUidRequest{
		MerchantUid: muid,
		Status:      status,
		Sorting:     sorting,
	}

	res, err := GetByMerchantUID(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, merchantUIDPaymentReq,
	)

	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentsMerchantUID merchant_uid로 모든 결제 정보 가져오기
//
// GET /payments/find/{merchant_uid}
func (iamport *Iamport) GetPaymentsMerchantUID(muid string, status string, sorting string, page int) (*TypePayment.PaymentPage, error) {
	if muid == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	if !ValidateStatusParameter(status) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	if !ValidateSortParameter(sorting) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	if page < 0 {
		return nil, errors.New(ErrInvalidPage)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	merchantUIDPaymentReq := &TypePayment.PaymentsMerchantUidRequest{
		MerchantUid: muid,
		Status:      status,
		Sorting:     sorting,
		Page:        int32(page),
	}

	res, err := GetByMerchantUIDs(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, merchantUIDPaymentReq,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentsStatus 결제 상태에 따른 결제 정보들 가져오기
//
// GET /payments/status/{payment_status}
func (iamport *Iamport) GetPaymentsStatus(status string, page int, limit int, from time.Time, to time.Time, sorting string) (*TypePayment.PaymentPage, error) {
	if !ValidateSortParameter(sorting) {
		return nil, errors.New(ErrInvalidSortParam)
	}

	if !ValidateStatusParameter(status) {
		return nil, errors.New(ErrInvalidStatusParam)
	}

	if page < 0 {
		return nil, errors.New(ErrInvalidPage)
	}

	if limit < 0 {
		return nil, errors.New(ErrInvalidLimit)
	}

	if from.After(to) {
		return nil, errors.New(ErrInvalidFrom)
	}

	if from.AddDate(0, 3, 0).Before(to) {
		return nil, errors.New(ErrInvalidTo)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentStatusRequest{

		Status:  status,
		Page:    int32(page),
		From:    int32(from.Unix()),
		Limit:   int32(limit),
		Sorting: sorting,
		To:      int32(to.Unix()),
	}

	res, err := GetByStatus(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)

	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPaymentBalanceImpUID imp_uid로 결제 정보 가져오기
//
// GET /payments/{imp_uid}
func (iamport *Iamport) GetPaymentBalanceImpUID(iuid string) (*TypePayment.PaymentBalance, error) {
	if iuid == "" {
		return nil, errors.New(ErrMustExistImpUID)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	reqPaymentImpUID := &TypePayment.PaymentBalanceRequest{
		ImpUid: iuid,
	}

	res, err := GetBalanceByImpUID(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, reqPaymentImpUID,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// CancelPaymentImpUID imp_uid로 결제 취소하기
//
// POST /payments/cancel
func (iamport *Iamport) CancelPaymentImpUID(iuid string, merchantUID string, amount float64, taxFree float64, reason string, refundHolder string, refundBank string, refundAccount string) (*TypePayment.Payment, error) {
	if iuid == "" && merchantUID == "" {
		return nil, errors.New(ErrMustExistImpUIDorMerchantUID)
	}

	if amount < 0 {
		return nil, errors.New(ErrInvalidAmount)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentCancelRequest{
		ImpUid:        iuid,
		MerchantUid:   merchantUID,
		Amount:        amount,
		TaxFree:       taxFree,
		Reason:        reason,
		RefundHolder:  refundHolder,
		RefundBank:    refundBank,
		RefundAccount: refundAccount,
	}

	res, err := Cancel(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// PreparePayment 결제 정보 사전 등록하기
//
// POST /payments/prepare
func (iamport *Iamport) PreparePayment(merchantUID string, amount float64) (*TypePayment.Prepare, error) {
	if merchantUID == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	if amount < 0 {
		return nil, errors.New(ErrInvalidAmount)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentPrepareRequest{
		MerchantUid: merchantUID,
		Amount:      amount,
	}

	res, err := Prepare(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}

// GetPreparePayment 사전 등록된 결제 정보 보기
//
// GET /payments/prepare/{merchant_uid}
func (iamport *Iamport) GetPreparePayment(merchantUID string) (*TypePayment.Prepare, error) {
	if merchantUID == "" {
		return nil, errors.New(ErrMustExistMerchantUID)
	}

	token, err := iamport.Authenticate.GetToken()
	if err != nil {
		return nil, err
	}

	req := &TypePayment.PaymentGetPrepareRequest{
		MerchantUid: merchantUID,
	}

	res, err := GetPrepareByMerchantUID(
		iamport.Authenticate.Client, iamport.Authenticate.APIUrl,
		token, req,
	)
	if err != nil {
		return nil, err
	}

	if res.Code != CodeOK {
		return nil, errors.New(res.Message)
	}

	return res.Response, nil
}
