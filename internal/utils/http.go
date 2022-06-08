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
		"Cookie": "X-Origin-Cookie=2; T_ch=M; T_ss=AL; _abck=16110E3B9B12B0A801B9597FE599234E~-1~YAAQi3pGaBh6bTuBAQAAdyC1PAh1mrSzhBes4GGKtXlrGk3UppX4XUQvdqHZmjXJ9HZRHbWks/dobFRuLL+tVKkkUN02VV6nRPcFFe49AY4IWRCEm4ZAKq0VJRnlW9WGoHaUR0uNtrSpjlKa0qa1r524Mh0QW2uZ2Y7Z81lGLIJemC+obGGJzWiOywwy8rqo0/U9IKIj5+UijtogHopA9ZOlG8JLoyM1jxkwcWprN2ZPvQU0LvzCwhUzx4l63j7ElmJ1cGJ47ngnk+Mq2HeTm4Lq0agIP71f7jFtZkx/xIQNkNMEOw7jK7KlQRc/Ve9Z96RMsn+YzvxVY0gKbjx9YDBhESDEJ3KU9ga08Ce0uh1fAeLmOWNbROI9HKAUH2ieiF9WxAjM35AC~-1~-1~-1; ak_bmsc=8389B5833048E2A893CDEB94BBA2C791~000000000000000000000000000000~YAAQZpc7F7jMwTeBAQAApZAyPRD0DDT6zmufhdKsU3HfkpOTWlBi7P33gN2MX1MsT82jRyx5kzJkxBgVqu76GBCIRcR7wooAZ1LFsqaoFOk9vyTOjcJ7WuyWeUibLfIISQn6rcFRBJBGfHuLKMZ4WczO5M4J4v19bihy/9+AnJHp097iuPXLnSk9sju1ceR3G0MXhcYF3JCn9HxWcYugLAemmnI/uloy2N+AYjFTBYB2E0zIcytCIm5sTrClnZj4RMAOt66GW8QIRkRepEXsy6Yqwrzwu8jCdGlx0wqJMhITcR5rUatpSYxV2jOTZttds7sDF7VHxoe3NYu6wJuewpB2p6si08pkAoptFsQ2I7ysOcVAVfPMu7G2qNEVhkg=; bm_sv=AAA1AE3CE598B48CF2CA51E6341E9F8C~YAAQZpc7Fw/uwTeBAQAAwRU5PRAaDv345PBXvL3af9qtZxxnC+q3s1l8MWtJ5VmjmyT19+0HczkALnKXGMIZQzGbAnPYrNwLpcERT2NEV6ZoYKlhTf9jVEIMroHZCEJMEodYxCiS27T56/Qz01RRHi0HJjQ4aQpVNrSeOTAx6X3bYp5ng88Xe+VVoyIsYtQW84wqu75qoRCAQrG3g5EKbKOP8gsZmU4xvlp4TXDJYeL/nsw9ycXqTC0AQXKRWc8SO1U=~1; bm_sz=9EBABD5FF3BCD9F7E91EE11D3237AF5D~YAAQi3pGaBp6bTuBAQAAdyC1PBC0+pkbbw7/+51UG5SA2QHRZstE6T5Ui0gYHn69lft9/X2h2uIbOs6EthcUT9YEwTzwGX0ww4pymbQgtlffzdV2c7CHCTwvx38Hlp6K3BIVRu+ltOc4APmCULhjoVWPWc8s9otokd7VZI2lVzfqUOyEv2q0cnLeLOGZaOcpaMm+RzTZJbZNJGKc/AhfCRQVNaF8Ib44AvwhgRTbZn+U4i2+ya7RWTq8XO9VmUBEKQaoy4cWBUe/Tl1C2OMKofb9PU2xFb98Bz08DVOyBKxzvUWC0A==~3360056~3490615; FlannelsFashion_AuthenticationCookie=18e297ea-9bd6-41b3-9ac1-1f8f7744abb7; TS01a19d95=01e4dc9a76f4679e1c8ab331454675101983505bf233f95d4164cc9c4022affbbeb2dd1a3216dedf40b79a9cd8c112fe1c95ecea08; X-SD-URep=e20cf74f-7ea3-488e-9853-d85aff50a30c",
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
