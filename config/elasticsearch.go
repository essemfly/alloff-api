package config

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"strings"
)

var EsClient *elasticsearch.Client

func InitElasticSearch(conf Configuration) {
	defaultIndices := []string{"product_view", "product_order"}

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

	alreadyExist, err := checkIndexExists(defaultIndices)
	if err != nil {
		log.Printf("Error on checking index exists %s \n", err)
	}

	if !alreadyExist {
		err = createDefaultIndexMapping(defaultIndices)
		if err != nil {
			log.Printf("Error on creating default index %s \n", err)
		}
		log.Println("default index created")
	}
}

// checkIndexExists : 입력된 이름의 인덱스가 있는지 확인
func checkIndexExists(indices []string) (bool, error) {
	req := esapi.IndicesExistsRequest{
		Index: indices,
	}
	res, err := req.Do(context.Background(), EsClient)
	defer res.Body.Close()
	if err != nil {
		log.Printf("Error getting response: %s\n", err)
		return false, err
	}

	if res.StatusCode == 404 {
		return false, nil
	}
	return true, nil

}

// createDefaultIndexMapping : 기본 인덱스의 구조를 생성
// 다른건 필드는 동적으로 생성하고 created_at 만 date 타입의 필드로 사전에 생성해둔다.
func createDefaultIndexMapping(indices []string) error {
	bodyStr := `{
		"mappings": {
			"properties": {
				"created_at": {
	   			"type": "date",
	   			"format": "yyyy-MM-dd HH:mm:ss"
	 			}
			}
		}
	}`

	for _, index := range indices {
		log.Printf("creating index mapping for %s\n", index)
		req := esapi.IndicesCreateRequest{
			Index: index,
			Body:  strings.NewReader(bodyStr),
		}
		res, err := req.Do(context.Background(), EsClient)
		if err != nil {
			log.Printf("Error getting response: %s\n", err)
			res.Body.Close()
			return err
		}
		res.Body.Close()
	}
	return nil
}
