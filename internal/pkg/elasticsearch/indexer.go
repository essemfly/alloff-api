package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type logType string

const (
	PRODUCT_VIEW = logType("PRODUCT_VIEW")
	ORDERED_ITEM = logType("ORDERED_ITEM")
)

type productLogging struct {
	Product *domain.ProductDAO `json:"product"`
	Ts      string             `json:"ts"`
	Type    logType            `json:"type"`
}

type searchLogging struct {
	Keyword string `json:"keyword"`
	Ts      string `json:"ts"`
}

func logRequest(request *LogFields) (int, error) {
	index := "access_log"
	bodyStr := request.jsonEncoder()
	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(bodyStr),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), config.EsClient)
	if err != nil {
		log.Println("Error getting response : ", err)
		return 400, err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return 400, err
	}
	return res.StatusCode, nil
}

func ProductLogRequest(request *domain.ProductDAO, t logType) (int, error) {
	index := "product_log"

	now := time.Now().Format("2006-01-02 15:04:05")
	bd := &productLogging{
		Product: request,
		Ts:      now,
		Type:    t,
	}

	bodyBuffer := new(bytes.Buffer)
	encoder := json.NewEncoder(bodyBuffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(&bd)
	bodyStr := bodyBuffer.String()

	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(bodyStr),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), config.EsClient)
	if err != nil {
		log.Println("err getting response : ", err)
		return 400, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	log.Println(string(b))
	if err != nil {
		log.Println("err reading response : ", err)
		return 400, err
	}
	return res.StatusCode, nil
}

func SearchLogRequest(keyword string) (int, error) {
	index := "search_log"

	now := time.Now().Format("2006-01-02 15:04:05")
	bd := &searchLogging{
		Keyword: keyword,
		Ts:      now,
	}

	bodyBuffer := new(bytes.Buffer)
	encoder := json.NewEncoder(bodyBuffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(&bd)
	bodyStr := bodyBuffer.String()

	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(bodyStr),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), config.EsClient)
	if err != nil {
		log.Println("err getting response : ", err)
		return 400, err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("err reading response : ", err)
		return 400, err
	}
	return res.StatusCode, nil
}
