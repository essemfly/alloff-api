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

	t.Run("Test Indexing Data For Product View", func(t *testing.T) {
		product, _ := ioc.Repo.Products.Get("62429cadc45891a3a2bef8b8")
		statusCode, index, err := logProduct(product, VIEW)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
		require.Equal(t, "product_view", index)
	})

	t.Run("Test Indexing Data For Product Order", func(t *testing.T) {
		product, _ := ioc.Repo.Products.Get("62429cadc45891a3a2bef8b8")
		statusCode, index, err := logProduct(product, ORDER)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
		require.Equal(t, "product_order", index)
	})
}
