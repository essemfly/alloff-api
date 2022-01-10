package iamport

import (
	"bytes"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
)

// Status 결제 상태
type Status string

const (
	StatusAll      = "all"      // 전체
	StatusReady    = "ready"    // 미결제
	StatusPaid     = "paid"     // 결제완료
	StatusCanceled = "canceled" // 결제취소
	StatusFailed   = "failed"   // 결제실패

	SortDESCStarted = "-started" // 결제시작시각(결제창오픈시각) 기준 내림차순(DESC) 정렬
	SortASCStarted  = "started"  // 결제시작시각(결제창오픈시각) 기준 오름차순(ASC) 정렬
	SortDESCPaid    = "-paid"    // 결제완료시각 기준 내림차순(DESC) 정렬
	SortASCPaid     = "paid"     // 결제완료시각 기준 오름차순(ASC) 정렬
	SortDESCUpdated = "-updated" // 최종수정시각(결제건 상태변화마다 수정시각 변경됨) 기준 내림차순(DESC) 정렬
	SortASCUpdated  = "updated"  // 최종수정시각(결제건 상태변화마다 수정시각 변경됨) 기준 오름차순(ASC) 정렬
)

const (
	CodeOK = 0

	ErrStatusUnauthorized = "iamport: unauthorized"
	ErrStatusNotFound     = "iamport: invalid imp_uid"
	ErrUnknown            = "iamport: unknown error"

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

func Call(client *http.Client, token string, url string, method Method) ([]byte, error) {
	req, err := http.NewRequest(string(method), url, nil)

	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(HeaderAuthorization, token)

	res, err := call(client, req)
	if err != nil {
		return []byte{}, err
	}

	return res, nil
}

func CallWithForm(client *http.Client, token string, url string, method Method, param []byte) ([]byte, error) {

	// json 형식을 form 형태에 맞게 변환
	jsonStr := string(param)
	jsonStr = strings.Replace(jsonStr, `{`, "", -1)
	jsonStr = strings.Replace(jsonStr, `}`, "", -1)
	jsonStr = strings.Replace(jsonStr, `"`, "", -1)
	jsonStr = strings.Replace(jsonStr, `:`, "=", -1)
	jsonStr = strings.Replace(jsonStr, `,`, "&", -1)

	req, err := http.NewRequest(string(method), url, bytes.NewBufferString(jsonStr))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(HeaderAuthorization, token)
	req.Header.Set(HeaderContentType, HeaderContentTypeForm)

	res, err := call(client, req)
	if err != nil {
		return []byte{}, err
	}

	return res, nil
}

func CallWithJson(client *http.Client, token string, url string, method Method, param []byte) ([]byte, error) {
	req, err := http.NewRequest(string(method), url, bytes.NewReader(param))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(HeaderAuthorization, token)
	req.Header.Set(HeaderContentType, HeaderContentTypeJson)

	res, err := call(client, req)
	if err != nil {
		return []byte{}, err
	}

	return res, nil
}

func call(client *http.Client, req *http.Request) ([]byte, error) {
	res, err := client.Do(req)
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

func GetQueryPrefix(isFirst *bool) string {
	if *isFirst {
		*isFirst = false
		return "?"
	} else {
		return "&"
	}
}

func ValidateStatusParameter(src string) bool {
	if src == "" || src == StatusReady || src == StatusAll || src == StatusPaid || src == StatusFailed || src == StatusCanceled {
		return true
	}

	return false
}

func ValidateSortParameter(src string) bool {
	if src == "" || src == SortDESCStarted || src == SortASCStarted || src == SortDESCPaid || src == SortASCPaid || src == SortASCUpdated || src == SortDESCUpdated {
		return true
	}

	return false
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

func GetRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVEWXYZabcdefghijklmnopqrstuvewxyz0123456789")
	var bytes strings.Builder
	for i := 0; i < length; i++ {
		bytes.WriteRune(chars[rand.Intn(len(chars))])
	}

	return bytes.String()
}

func GetJoinString(values ...string) string {
	len := len(values)
	urls := make([]string, len)

	for _, s := range values {
		urls = append(urls, s)
	}

	url := strings.Join(urls, "")
	return url
}

var Unmarshaler = protojson.UnmarshalOptions{
	DiscardUnknown: true,
}
