package alimtalk

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func MakePostRequest(url string, bodyStr []byte) (*http.Response, error) {
	var client http.Client
	XSecretKey := "oE07IvyX"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyStr))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("X-Secret-Key", XSecretKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return resp, nil
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(b))
		resp.Body.Close()
		return resp, nil
	}
}

func MakeDeleteRequest(url string) (*http.Response, error) {
	var client http.Client
	XSecretKey := "oE07IvyX"

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	req.Header.Add("X-Secret-Key", XSecretKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		return resp, nil
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(b))
		resp.Body.Close()
		return resp, nil
	}
}
