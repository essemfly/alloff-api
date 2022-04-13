package notification

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

const (
	ErrStatusUnauthorized = "push_notification: unauthorized"
	ErrStatusNotFound     = "push_notification: status_not_found"
	ErrUnknown            = "push_notification: unknown error"

	HeaderContentType     = "Content-Type"
	HeaderContentTypeForm = "application/x-www-form-urlencoded"
	HeaderContentTypeJson = "application/json"
	HeaderAuthorization   = "Authorization"

	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type Method string

type NotificationPayload struct {
	Notifications []NotificationStruct `json:"notifications"`
}

type NotificationStruct struct {
	Tokens   []string         `json:"tokens"`
	Platform int              `json:"platform"`
	Message  string           `json:"message"`
	Title    string           `json:"title"`
	Data     NavigationStruct `json:"data"`
}

type NavigationStruct struct {
	Type       string `json:"type"`
	NavigateTo string `json:"navigateTo"`
}

type NotificationResponse struct {
	Counts  int
	Logs    []NotificationResult
	Success string
}

type NotificationResult struct {
	Type     string
	Platform string
	Token    string
	Message  string
	Error    string
}

// {
//     "counts": 2,
//     "logs": [
//         {
//             "type": "failed-push",
//             "platform": "android",
//             "token": "**********q2rPnePTmxtz:APA91bGFl4h8CxwD6rEB1bkFknetUt3lQ5i9FEmF7bQLdNrkVUGaOIj6zM8fnpoNchcj3149K7qZ9wEgbocti_jGwhqNtigUvys2AHCil6qRf6oFEbZm7Tg4qTkUDK4u17**********",
//             "message": "아무거나 다 담을 수 있져요 이번건 주문내역으로 갈게요",
//             "error": "unregistered device"
//         }
//     ],
//     "success": "ok"
// }

// 현재 동작 가능한 딥링크 alloff://app.alloff.co
// Dev alloffdev://dev.app.alloff.co/cart
// 홈 /home
// 주문내역 /orders
// 장바구니 /cart
// 타임딜 목록 /timedeals
// 타임딜 상세 (= 타임딜 상품 목록)  /timedeals/:timedealId
// 타임딜 상품 상세  /alloffProducts/:alloffProductId
// 모아보기 /brands
// 상품상세 /products/:productId

func SendNotification(noti *domain.NotificationDAO) error {
	client := &http.Client{}

	payload := MakePayload(noti)
	resp, err := CallWithJson(client, config.NotificationUrl, POST, payload)
	if err != nil {
		log.Println(err, "err occured in send notification")
		return err
	}

	var notiResult NotificationResponse
	err = json.Unmarshal(resp, &notiResult)
	if err != nil {
		log.Println("Err occured in result unmarshaling", err)
		return err
	}

	failedCounts := 0
	failedDeviceIds := []string{}
	for _, logResult := range notiResult.Logs {
		if logResult.Type == "failed-push" {
			failedCounts += 1
			failedDeviceIds = append(failedDeviceIds, logResult.Token)
		}
	}

	for _, failedDeviceID := range failedDeviceIds {
		err := ioc.Repo.Devices.MakeRemoved(failedDeviceID)
		if err != nil {
			log.Println("err on failed make removed", err)
		}
	}

	noti.Sended = time.Now()
	noti.Updated = time.Now()
	noti.Status = domain.NOTIFICATION_SUCCEEDED
	noti.NumUsersFailed = failedCounts
	noti.NumUsersPushed = len(noti.DeviceIDs) - failedCounts

	_, err = ioc.Repo.Notifications.Update(noti)
	if err != nil {
		return err
	}
	return nil
}

func MakePayload(noti *domain.NotificationDAO) []byte {
	payload := NotificationPayload{}
	payload.Notifications = append(payload.Notifications, NotificationStruct{
		Tokens:   noti.DeviceIDs,
		Platform: 2,
		Message:  noti.Message,
		Title:    noti.Title,
		Data: NavigationStruct{
			Type:       "navigation",
			NavigateTo: config.NavigateUrl + noti.NavigateTo + noti.ReferenceID,
		},
	})

	doc, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error in json marshalling in make payload")
		return nil
	}
	return doc
}

func CallWithJson(client *http.Client, url string, method Method, param []byte) ([]byte, error) {
	req, err := http.NewRequest(string(method), url, bytes.NewBuffer([]byte(param)))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(HeaderContentType, HeaderContentTypeJson)

	res, err := call(client, req)
	if err != nil {
		return []byte{}, err
	}

	return res, nil
}

func call(client *http.Client, req *http.Request) ([]byte, error) {
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	err = errorHandler(res)
	if err != nil {
		return []byte{}, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return resBody, nil
}

func errorHandler(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return errors.New(ErrStatusUnauthorized)
	case http.StatusNotFound:
		return errors.New(ErrStatusNotFound)
	default:
		return errors.New(ErrUnknown)
	}
}
