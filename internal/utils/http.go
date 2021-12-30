package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
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

	if resp.StatusCode == http.StatusOK {
		return resp, nil
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()

		errObject := HttpError{RespBody: string(b), StatusCode: strconv.Itoa(resp.StatusCode)}
		return nil, &errObject
	}
}
