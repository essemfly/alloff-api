package elasticsearch

import (
	"context"
	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"log"
	"net/http"
	"strings"
)

type ESClient struct {
	Client *es8.Client
}

func NewElasticSearch(conf config.Configuration) *ESClient {
	defaultIndexName := []string{"access_log", "product_log", "search_log"}

	header := http.Header{}
	header.Add("Authorization", "Basic "+conf.ELASTICSEARCH_APIKEY)

	cfg := es8.Config{
		Addresses: []string{
			conf.ELASTICSEARCH_URL,
		},
		Header: header,
	}

	esClient, err := es8.NewClient(cfg)
	if err != nil {
		log.Println("Error creating Client : ", err)
	}

	alreadyExist, err := checkIndexExists(defaultIndexName, esClient)
	if err != nil {
		log.Printf("Error on checking index exists %s \n", err)
	}

	if !alreadyExist {
		err = createDefaultIndexMapping(defaultIndexName, esClient)
		if err != nil {
			log.Printf("Error on creating default index %s \n", err)
		} else {
			log.Println("default index created")
		}
	}

	return &ESClient{
		Client: esClient,
	}
}

func (conn *ESClient) RegisterRepos() {
	ioc.Repo.AccessLog = EsAccessLogRepo(conn)
	ioc.Repo.ProductLog = EsProductLogRepo(conn)
	ioc.Repo.SearchLog = EsSearchLogRepo(conn)
}

// checkIndexExists : 입력된 이름의 인덱스가 있는지 확인
func checkIndexExists(index []string, esClient *es8.Client) (bool, error) {
	req := esapi.IndicesExistsRequest{
		Index: index,
	}
	res, err := req.Do(context.Background(), esClient)
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
func createDefaultIndexMapping(index []string, esClient *es8.Client) error {
	bodyStr := `{
		"mappings": {
			"properties": {
				"ts": {
	   				"type": "date",
	   				"format": "yyyy-MM-dd HH:mm:ss"
	 			},
				"uri": {
					"type": "keyword"
				}
			}
		}
	}`

	for _, index := range index {
		req := esapi.IndicesCreateRequest{
			Index: index,
			Body:  strings.NewReader(bodyStr),
		}
		res, err := req.Do(context.Background(), esClient)
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
