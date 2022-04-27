package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	alloffEs "github.com/lessbutter/alloff-api/internal/pkg/elasticsearch"
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

func EsAccessLogRepo(conn *ESClient) repository.AccessLogRepository {
	return &accessLogRepo{
		client: conn.Client,
	}
}
