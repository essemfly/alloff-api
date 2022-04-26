package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type brandLogRepo struct {
	client *elasticsearch.Client
}

func (repo *brandLogRepo) Index(request *domain.BrandDAO) (int, error) {
	index := "brand_log"

	now := time.Now().Format("2006-01-02 15:04:05")
	bd := &domain.BrandLogDAO{
		Brand: request,
		Ts:    now,
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
	res, err := req.Do(context.Background(), repo.client)
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

func EsBrandLogRepo(conn *ESClient) repository.BrandLogRepository {
	return &brandLogRepo{
		client: conn.Client,
	}
}
