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

type accessLogRepo struct {
	client *elasticsearch.Client
}

func (repo *accessLogRepo) Index(request *domain.AccessLogDAO) (int, error) {
	index := "access_log"

	bodyStr := alloffEs.JsonEncoder(request)
	statusCode, err := alloffEs.RequestIndex(index, bodyStr, repo.client)
	if err != nil {
		return 400, err
	}
	return statusCode, nil
}

func (repo *accessLogRepo) List(limit int, from, to time.Time, order string) (*dto.AccessLogDTO, error) {
	var accessLogDTO dto.AccessLogDTO

	index := "access_log"
	fromStr := from.Format("2006-01-02 15:04:05")
	toStr := to.Format("2006-01-02 15:04:05")

	bodyStr := fmt.Sprintf(`{
		"size": %v,
		"query": {
			"range": {
				"ts": {
					"gt": "%s",
					"lt": "%s"
				}
			}
		},
		"sort": {"ts": "%s"}
	}`, limit, fromStr, toStr, order)

	resBody, err := alloffEs.RequestQuery(index, bodyStr, repo.client)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resBody, &accessLogDTO)
	return &accessLogDTO, nil
}

func (repo *accessLogRepo) GetLatest(limit int) (*dto.AccessLogDTO, error) {
	var accessLogDTO dto.AccessLogDTO

	now := time.Now().Add(time.Minute).Format("2006-01-02 15:04:05")
	index := "access_log"

	bodyStr := fmt.Sprintf(`{
		"size": %v,
		"sort": {"ts": "desc"},
		"query": {
			"range": {
				"ts": {
					"lt": "%s"
				}
			}
		}
	}`, limit, now)

	resBody, err := alloffEs.RequestQuery(index, bodyStr, repo.client)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(resBody, &accessLogDTO)
	return &accessLogDTO, nil
}

func EsAccessLogRepo(conn *ESClient) repository.AccessLogRepository {
	return &accessLogRepo{
		client: conn.Client,
	}
}
