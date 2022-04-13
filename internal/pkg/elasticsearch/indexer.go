package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type Status string

const (
	ORDER = Status("ORDER")
	VIEW  = Status("VIEW")
)

type Response struct {
	Index string `json:"_index"`
}

// logProduct : ElasticSearch에 기록할 상품정보를 받아 알맞은 인덱스에 상품 ID와 Brand를 기록한다.
func logProduct(product *domain.ProductDAO, status Status) (int, string, error) {
	var resp Response

	index := ""
	switch status {
	case VIEW:
		index = "product_view"
	case ORDER:
		index = "product_order"
	}

	// TODO 어떤 정보들 인덱싱 할지 ?
	bodyStr := fmt.Sprintf(`{
				"productId": "%s",
				"brand": "%s",
				"created_at": "%v"
				}`,
		product.ID.Hex(),
		product.ProductInfo.Brand.KeyName,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(bodyStr),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), config.EsClient)
	defer res.Body.Close()
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return 400, "", err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 400, "", err
	}

	if err = json.Unmarshal(resBody, &resp); err != nil {
		log.Println(err)
	}
	return res.StatusCode, resp.Index, err
}
