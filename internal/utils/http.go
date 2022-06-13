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

func GetFlannelsHeader() map[string]string {
	return map[string]string{
		"Cookie": "X-Origin-Cookie=2; " +
			"T_ch=M; " +
			"T_ss=AL; " +
			"_abck=74284572184F517B929A478DA748D818~0~YAAQmplMF8ZxtjyBAQAACF0jWwgMaEL4Il04KRlMNcZiNlrWQkhzFJ4Zd+PhdXHMGyYwL3Iu246pmwUaH9fvnlZBVzL/zlxUYlUxnA9xZ1U3JH9jCJiGw/0lvs4kwHvp94pqk1byZOaS4bTFUBy620Prn5QM/rkILRQBqx/wlMFHXk6kVvkjBmCADShyck2x1embnG+BGc9GxMoW/m6CzEnPvrP9dMU5TDkhpfCY62330rVROp2YCv+50ekyBWtm2EMI07ltoIygDpgdMVEjaIuHieq67X60FDdtUVow/ZRhKbQbrvdS0ZGTMW8P1B5AacmO53ts6aimq6K567plpIRJnGPKNEwlFybCIxHpRlrkOByIYTOEGoz3qgFm8IAj~-1~-1~-1; " +
			"ak_bmsc=16403ACF226BD15DFC1E86FBFA126D2D~000000000000000000000000000000~YAAQdHpGaCbQSimBAQAAHZkgWxDhhgfkv/1r5XX+TzBVkZ8RxyYzTpAfC267J7LWDHJiYNacnZGPtbkOXVMWnFi237npmNNDa71Gk8UDmTDeAGzg3V5hXp3JZvFshlxWiALWVFqkz+gw1KFSrTF+RWez6EshB/mrwXdy+4YVrT+jzC+L0Mlu1X9PHQi5F2gWbJeyE5MQUCipwpa/S5Kml43xNg9Fb3yGFpYgZlM4q/Urw6m9Xqy4olNLPzTAoVLi9vboAoYa81nJgW3I5QY4j++lTWlx7x4H4Ev5VdDVwfuUAMP75vwk5hJe6l1NnIO8a65kyEN6Y4juj+A5DYGVMhbdhpAnIxR92ibLvn27dGc/VaS/EfqelrzElwR0fzJQ0K+LN/z525WPEwbwurPICCXYagHabh9xI0QwL1f75NqallhbAc3VbyOkTIOXc+w6KSFXh2hWOg==; " +
			"bm_sv=A15C4C3342DC2E101F267EC252E9064F~YAAQdHpGaP/QSimBAQAA8KogWxBiMlq2fpwiWmqBSYM1LeDr90HDE/neSObJWFRazO96ev6ZeYd10FhJLO4YFWd6dCHPJ5YIiBXTRlASP7vWV7v2spcxiJOUkvTjtOk452NMerVgbm1HuVXPieAQw7h4JRZmhZ4Wrtw4oAjYn4ITyfYgAkwpjLw6cAqLKsnTLZaNn39QuhrQmL1/nkXNEa00qEtCk22P/oI5Qutkisb/D8TJPgZSTdd70Q41Ss2L/r0=~1; " +
			"bm_sz=6E1C04F9786C4CBDE1F6A1CDE6DE916A~YAAQF60sF4Jr2DuBAQAALl3NWhAaZSaZERbKHWU+9OiIHuDl33SVGKrOIjAZNzY9hSnJJUzxtuvuMaGPm1jGGjzG5IfrqjrFfDjrGKj7kPnVMzaGXROWEcAt0lnA9bREEGhAzRedRLkn2WhxIUxPuYnxv7mfGMMB7mqY/eOR1lvi6i7NO2oHI6FFPIf60AB4rXxAVY4R2+NHVr3Tc6Ae33F0kic/g2LBUP4KZEs2/sLVxJu1q6zMeCctC7vSOYsvq3AE6fbMbrtMyQi2JWc/gNYOWxRS1WBxk93XQ2s0X9XuTnbM4Q==~3225922~3490096; " +
			"FlannelsFashion_AuthenticationCookie=4f84c1d5-d5c0-483f-8f46-a7dcd90634a0; " +
			"TS01a19d95=01e4dc9a76eb659592266abd9bd0795f98781ca2ef362fefd9dc6774ce22193c5a1ae0be8742cf6c4e4bf0f5943b285898514fedc7bdeb4637db785bcb685a4f08f53125de; " +
			"X-SD-URep=00ec5f40-ae6f-4da2-ac7d-26695f2bb348"}
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
