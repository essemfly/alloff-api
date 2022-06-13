package alimtalk

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type NormalRequest struct {
	SenderKey     string      `json:"senderKey"`
	TemplateCode  string      `json:"templateCode"`
	RequestDate   string      `json:"requestDate"`
	RecipientList []Recipient `json:"recipientList"`
}

type RecipientTemplate struct {
}

type Recipient struct {
	RecipientNo       string            `json:"recipientNo"`
	TemplateParameter map[string]string `json:"templateParameter"`
	// ResendParam       ResendParameter   `json:"resendParameter"`
}

type ResendParameter struct {
	IsResend      bool   `json:"isResend"`
	ResendType    string `json:"resendType"`
	ResendTitle   string `json:"resendTitle"`
	ResendContent string `json:"resendContent"`
	ResendSendNo  string `json:"resendSendNo"`
}

type RequestResponse struct {
	Header struct {
		ResultCode    int    `json:"resultCode"`
		ResultMessage string `json:"resultMessage"`
		IsSuccessful  bool   `json:"isSuccessful"`
	} `json:"header"`
	Message struct {
		RequestId         string       `json:"requestId"`
		SenderGroupingKey string       `json:"senderGroupingKey"`
		SendResults       []SendResult `json:"sendResults"`
	} `json:"message"`
}

type DeleteResponse struct {
	Header struct {
		ResultCode    int    `json:"resultCode"`
		ResultMessage string `json:"resultMessage"`
		IsSuccessful  bool   `json:"isSuccessful"`
	} `json:"header"`
}

type SendResult struct {
	RecipientSeq         int    `json:"recipientSeq"`
	RecipientNo          string `json:"recipientNo"`
	ResultCode           int    `json:"resultCode"`
	ResultMessage        string `json:"resultMessage"`
	RecipientGroupingKey string `json:"recipientGroupingKey"`
}

func SendMessage(alimtalk *domain.AlimtalkDAO) (string, error) {
	baseUrl := "https://api-alimtalk.cloud.toast.com"
	appKey := "sj4UJFouCcvOHajL"
	alloffSenderKey := "63949416400523d6fbe6fa9112644ab359710b74"
	realUrl := baseUrl + "/alimtalk/v2.0/appkeys/" + appKey + "/messages"

	recipient := Recipient{
		RecipientNo:       alimtalk.Mobile,
		TemplateParameter: alimtalk.TemplateParams,
	}

	// 요청 일시 (yyyy-MM-dd HH:mm)
	requestDate := ""
	if alimtalk.SendDate != nil {
		loc, _ := time.LoadLocation("Asia/Seoul")
		korDate := alimtalk.SendDate.In(loc)
		requestDate = korDate.Format("2006-01-02 15:04")
	}
	reqBody := NormalRequest{
		SenderKey:     alloffSenderKey,
		TemplateCode:  alimtalk.TemplateCode,
		RequestDate:   requestDate,
		RecipientList: []Recipient{recipient},
	}

	reqBodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		log.Println(err)
		return "", err
	}

	resp, err := MakePostRequest(realUrl, reqBodyBytes)
	if err != nil {
		log.Println(err)
		return "", err
	}

	var respBody RequestResponse
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return respBody.Message.RequestId, nil
}

func DeleteMessage(alimtalk *domain.AlimtalkDAO) error {
	baseUrl := "https://api-alimtalk.cloud.toast.com"
	appKey := "sj4UJFouCcvOHajL"
	realUrl := baseUrl + "/alimtalk/v2.0/appkeys/" + appKey + "/messages/" + alimtalk.ToastRequestID
	resp, err := MakeDeleteRequest(realUrl)
	if err != nil {
		log.Println(err)
		return err
	}

	var deleteBody DeleteResponse
	err = json.NewDecoder(resp.Body).Decode(&deleteBody)
	if err != nil {
		return err
	}

	if !deleteBody.Header.IsSuccessful {
		return errors.New("delete body is not successful")
	}
	return nil
}
