package elasticsearch

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"io/ioutil"
	"log"
	"strings"
)

type accessLogRepo struct {
	client *elasticsearch.Client
}

func (repo *accessLogRepo) Index(request *domain.AccessLogDAO) (int, error) {
	index := "access_log"
	bodyStr := request.JsonEncoder()
	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(bodyStr),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), repo.client)
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

func EsAccessLogRepo(conn *ESClient) repository.AccessLogRepository {
	return &accessLogRepo{
		client: conn.Client,
	}
}
