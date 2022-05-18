package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type RequestType string

const (
	REQUEST_GET  = RequestType("GET")
	REQUEST_POST = RequestType("POST")
)

type HttpError struct {
	RespBody   string
	StatusCode string
}

func (m *HttpError) Error() string {
	return m.StatusCode + " " + m.RespBody
}

func RequestRetryer(url string, method RequestType, headers map[string]string, bodyJsonStr string, errorMessage string) (*http.Response, error) {
	// err := sentry.Init(sentry.ClientOptions{
	// 	Dsn: "https://3bb7aa9c71b44397928e0101ebfecef2@o306501.ingest.sentry.io/6011225",
	// })
	// if err != nil {
	// 	log.Fatalf("sentry.Init: %s", err)
	// }

	resp, err := MakeRequest(url, method, headers, bodyJsonStr)
	retryCount := 0
	for err != nil && retryCount < 3 {
		retryCount += 1
		if retryCount == 3 {
			// sentry.CaptureMessage(errorMessage + " " + err.Error())
			return nil, err
		}
		resp, err = MakeRequest(url, method, headers, bodyJsonStr)
	}
	return resp, nil
}

func MakeRequest(url string, method RequestType, headers map[string]string, bodyJsonStr string) (*http.Response, error) {
	var client http.Client
	jsonStrByte := []byte(bodyJsonStr)
	req, err := http.NewRequest(string(method), url, bytes.NewBuffer(jsonStrByte))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		return resp, nil
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		errObject := HttpError{RespBody: string(b), StatusCode: strconv.Itoa(resp.StatusCode)}
		return nil, &errObject
	}
}

func GetHeader() map[string]string {
	return map[string]string{
		"accept":          "*/*",
		"content-type":    "application/x-www-form-urlencoded;charset=UTF-8",
		"connection":      "keep-alive",
		"user-agent":      "Crawler",
		"accept-language": "ko-KR",
	}
}

func GetGeneralHeader() map[string]string {
	return map[string]string{
		"accept":       "*/*",
		"content-type": "application/json",
		"connection":   "keep-alive",
		"user-agent":   "Crawler",
	}
}

func GetJsonHeader() map[string]string {
	return map[string]string{
		"accept": "application/json",
	}
}

func GetSSFHeaders() map[string]string {
	coockie := time.Now().UnixNano()
	s := strconv.FormatInt(coockie, 10)
	return map[string]string{
		"accept": "application/json",
		"Cookie": "PCID=" + s,
	}
}

func GetTheoutnetHeaders() map[string]string {
	return map[string]string{
		"x-ibm-client-id": "19c36e19-5bc7-4de4-a4a9-65ffb9dcb727",
		"accept":          "*/*",
		"accept-encoding": "gzip, deflate, br",
		"connection":      "keep-alive",
		"user-agent":      "PostmanRuntime/7.29.0",
		"content-type":    "application/x-www-form-urlencoded",
	}
}

func GetOmniousHeader(omniousKey string) map[string]string {
	return map[string]string{
		"x-api-key":       omniousKey,
		"accept-language": "ko",
	}
}
