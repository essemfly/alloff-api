package elasticsearch

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvent(t *testing.T) {
	testConf := config.GetConfiguration("local")
	conn := mongo.NewMongoDB(testConf)
	conn.RegisterRepos()
	config.InitElasticSearch(testConf)

	t.Run("Test Indexing Product For Product View", func(t *testing.T) {
		product, _ := ioc.Repo.Products.Get("62429cb0c45891a3a2bef8c2")
		statusCode, err := ProductLogRequest(product, PRODUCT_VIEW)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
	})

	t.Run("Test Indexing Product For Product Order", func(t *testing.T) {
		product, _ := ioc.Repo.Products.Get("62429cb0c45891a3a2bef8c2")
		statusCode, err := ProductLogRequest(product, ORDERED_ITEM)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
	})

	t.Run("Test Indexing Search Log", func(t *testing.T) {
		keyword := "test_keyword"
		statusCode, err := SearchLogRequest(keyword)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
	})
}
