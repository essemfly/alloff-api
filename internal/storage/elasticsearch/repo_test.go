package elasticsearch

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvent(t *testing.T) {
	testConf := config.GetConfiguration("dev")
	mongoConn := mongo.NewMongoDB(testConf)
	mongoConn.RegisterRepos()
	esConn := NewElasticSearch(testConf)
	esConn.RegisterRepos()

	t.Run("Test Indexing Product For Product View", func(t *testing.T) {
		product, _ := ioc.Repo.Products.Get("62429cb0c45891a3a2bef8c2")
		statusCode, err := ioc.Repo.ProductLog.Index(product, domain.PRODUCT_VIEW)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
	})

	t.Run("Test Indexing Product For Product Order", func(t *testing.T) {
		product, _ := ioc.Repo.Products.Get("62429cb0c45891a3a2bef8c2")
		statusCode, err := ioc.Repo.ProductLog.Index(product, domain.ORDERED_ITEM)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
	})

	t.Run("Test Indexing Search Log", func(t *testing.T) {
		keyword := "test_keyword"
		statusCode, err := ioc.Repo.SearchLog.Index(keyword)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
	})

	t.Run("Test Indexing Brand Log", func(t *testing.T) {
		brand, _ := ioc.Repo.Brands.Get("61d699eb74b2b71fe80ff4b6")
		statusCode, err := ioc.Repo.BrandLog.Index(brand)
		require.Nil(t, err)
		require.Equal(t, 201, statusCode)
	})

}
