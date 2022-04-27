package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/lessbutter/alloff-api/internal/core/domain"
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

func EsProductLogRepo(conn *ESClient) repository.ProductLogRepository {
	return &productLogRepo{
		client: conn.Client,
	}
}
