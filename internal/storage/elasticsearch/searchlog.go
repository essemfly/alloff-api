package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	alloffEs "github.com/lessbutter/alloff-api/internal/pkg/elasticsearch"
	"time"
)

type searchLogRepo struct {
	client *elasticsearch.Client
}

func (repo *searchLogRepo) Index(keyword string) (int, error) {
	index := "search_log"

	now := time.Now().Format("2006-01-02 15:04:05")
	bd := &domain.SearchLogDAO{
		Keyword: keyword,
		Ts:      now,
	}

	bodyStr := alloffEs.JsonEncoder(bd)
	statusCode, err := alloffEs.RequestIndex(index, bodyStr, repo.client)
	if err != nil {
		return 400, err
	}
	return statusCode, nil
}

func EsSearchLogRepo(conn *ESClient) repository.SearchLogRepository {
	return &searchLogRepo{
		client: conn.Client,
	}
}
