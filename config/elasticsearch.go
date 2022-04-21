package config

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"net/http"
	"strings"
)

var EsClient *elasticsearch.Client
var EsAPIKEY string

func InitElasticSearch(conf Configuration) {
	defaultIndexName := []string{"access_log", "product_log"}

	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.ELASTICSEARCH_URL,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s \n", err)
	}

	EsClient = es
	EsAPIKEY = conf.ELASTICSEARCH_APIKEY

	alreadyExist, err := checkIndexExists(defaultIndexName)
	if err != nil {
		log.Printf("Error on checking index exists %s \n", err)
	}

	if !alreadyExist {
		err = createDefaultIndexMapping(defaultIndexName)
		if err != nil {
			log.Printf("Error on creating default index %s \n", err)
		} else {
			log.Println("default index created")
		}
	}
}

// checkIndexExists : 입력된 이름의 인덱스가 있는지 확인
func checkIndexExists(index []string) (bool, error) {
	header := http.Header{}
	header.Add("Authorization", "Basic "+EsAPIKEY)
	req := esapi.IndicesExistsRequest{
		Index:  index,
		Header: header,
	}
	res, err := req.Do(context.Background(), EsClient)
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return false, nil
	}
	return true, nil

}

// createDefaultIndexMapping : 기본 인덱스의 구조를 생성
// 다른건 필드는 동적으로 생성하고 created_at 만 date 타입의 필드로 사전에 생성해둔다.
func createDefaultIndexMapping(index []string) error {
	bodyStr := `{
		"mappings": {
			"properties": {
				"ts": {
	   			"type": "date",
	   			"format": "yyyy-MM-dd HH:mm:ss"
	 			}
			}
		}
	}`
	header := http.Header{}
	header.Add("Authorization", "Basic "+EsAPIKEY)

	for _, index := range index {
		req := esapi.IndicesCreateRequest{
			Index:  index,
			Body:   strings.NewReader(bodyStr),
			Header: header,
		}
		res, err := req.Do(context.Background(), EsClient)
		if err != nil {
			log.Printf("Error getting response on creating default index mapping: %s\n", err)
			return err
		} else {
			log.Printf("creating index mapping for %s\n", index)
		}
		res.Body.Close()
	}
	return nil
}
