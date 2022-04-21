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
	"net/http"
	"strings"
	"time"
)

func logRequest(request *LogFields) (int, error) {
	index := "access_log"
	bodyStr := request.jsonEncoder()
	header := http.Header{}
	header.Add("Authorization", "Basic "+config.EsAPIKEY)
	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(bodyStr),
		Refresh: "true",
		Header:  header,
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

func ProductLogRequest(request *domain.ProductDAO) (int, error) {
	index := "product_log"
	header := http.Header{}
	header.Add("Authorization", "Basic "+config.EsAPIKEY)

	now := time.Now().Format("2006-01-02 15:04:05")
	bd := &customLogging{
		Data: request,
		Ts:   now,
	}

	bodyBuffer := new(bytes.Buffer)
	encoder := json.NewEncoder(bodyBuffer)
	encoder.SetEscapeHTML(false)
	encoder.Encode(&bd)
	bodyStr := bodyBuffer.String()

	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(bodyStr),
		Header:  header,
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

type customLogging struct {
	Data *domain.ProductDAO `json:"data"`
	Ts   string             `json:"ts"`
}
