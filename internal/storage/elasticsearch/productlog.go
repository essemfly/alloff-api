package elasticsearch

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/dto"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	alloffEs "github.com/lessbutter/alloff-api/internal/pkg/elasticsearch"
	"time"
)

type productLogRepo struct {
	client *elasticsearch.Client
}

func (repo *productLogRepo) Index(request *domain.ProductDAO, t domain.LogType) (int, error) {
	index := "product_log"

	now := time.Now().Format("2006-01-02 15:04:05")
	bd := &domain.ProductLogDAO{
		Product: request,
		Ts:      now,
		Type:    t,
	}

	bodyStr := alloffEs.JsonEncoder(bd)
	statusCode, err := alloffEs.RequestIndex(index, bodyStr, repo.client)
	if err != nil {
		return 400, err
	}
	return statusCode, nil
}

func (repo *productLogRepo) GetRank(limit int, from time.Time, to time.Time) (*dto.DocumentCountDTO, error) {
	var documentCountDTO dto.DocumentCountDTO

	index := "product_log"
	fromStr := from.Format("2006-01-02 15:04:05")
	toStr := to.Format("2006-01-02 15:04:05")

	bodyStr := fmt.Sprintf(`{
		"size": 0,
		"query": {
			"range": {
				"ts": {
					"gt": "%s",
					"lt": "%s"
				}
			}
		},
		"aggs": {
			"group_by_state": {
				"terms": {
					"field": "product.ID.keyword",
					"size": %v
				}
			}
		}
	}`, fromStr, toStr, limit)

	resBody, err := alloffEs.RequestQuery(index, bodyStr, repo.client)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resBody, &documentCountDTO)
	return &documentCountDTO, nil
}

func (repo *productLogRepo) GetRankByCatId(limit int, from time.Time, to time.Time, catId string) (*dto.DocumentCountDTO, error) {
	var documentCountDTO dto.DocumentCountDTO

	index := "product_log"
	fromStr := from.Format("2006-01-02 15:04:05")
	toStr := to.Format("2006-01-02 15:04:05")

	bodyStr := fmt.Sprintf(`
	{
	   "size": 0,
	   "query": {
		   "bool": {
			   "must": [
				   {
					   "range": {
						   "ts": {
							   "gt": "%s",
							   "lt": "%s"
							   }
					   }
				   },
				   {
					   "match": {
						   "product.AlloffCategories.First.ID": "%s"
					   }
				   }
			   ]
		   }
		},
	   "aggs": {
		  "group_by_state": {
			 "terms": {
				"field": "product.ID.keyword",
				"size": %v
			 }
		  }
	   }
	}
	`, fromStr, toStr, catId, limit)

	resBody, err := alloffEs.RequestQuery(index, bodyStr, repo.client)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resBody, &documentCountDTO)
	return &documentCountDTO, nil
}

func EsProductLogRepo(conn *ESClient) repository.ProductLogRepository {
	return &productLogRepo{
		client: conn.Client,
	}
}
