package elasticsearch

import (
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/dto"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	alloffEs "github.com/lessbutter/alloff-api/internal/pkg/elasticsearch"
	"log"
	"time"
)

type brandLogRepo struct {
	client *elasticsearch.Client
}

const index = "brand_log"

func (repo *brandLogRepo) Index(request *domain.BrandDAO) (int, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	bd := &domain.BrandLogDAO{
		Brand: request,
		Ts:    now,
	}

	bodyStr := alloffEs.JsonEncoder(bd)
	statusCode, err := alloffEs.RequestIndex(index, bodyStr, repo.client)
	if err != nil {
		return 400, err
	}
	return statusCode, nil
}

func (repo *brandLogRepo) GetRank(limit int, from time.Time, to time.Time) (*dto.DocumentCountDTO, error) {
	var documentCountDTO dto.DocumentCountDTO

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
					"field": "brand.ID.keyword",
					"size": %v
				}
			}
		}
	}
	`, fromStr, toStr, limit)

	resBody, err := alloffEs.RequestQuery(index, bodyStr, repo.client)
	if err != nil {
		return nil, err
	}
	log.Println(string(resBody))
	json.Unmarshal(resBody, &documentCountDTO)
	return &documentCountDTO, nil
}

func EsBrandLogRepo(conn *ESClient) repository.BrandLogRepository {
	return &brandLogRepo{
		client: conn.Client,
	}
}
