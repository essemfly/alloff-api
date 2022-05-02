package elasticsearch

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
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

	t.Run("Test Query Brand Count", func(t *testing.T) {
		from := time.Now().Add(-24 * 365 * time.Hour)
		log.Println(from.String())
		to := time.Now()
		limit := 100
		res, _ := ioc.Repo.BrandLog.GetRank(limit, from, to)
		log.Println(res)
	})

	t.Run("Test Query Product Count", func(t *testing.T) {
		from := time.Now().Add(-24 * 365 * time.Hour)
		to := time.Now()
		limit := 100
		res, _ := ioc.Repo.ProductLog.GetRank(limit, from, to)
		log.Println(res)
	})

	t.Run("Test Query AccessLog", func(t *testing.T) {
		from := time.Now().Add(-24 * 365 * time.Hour)
		to := time.Now()
		limit := 100
		order := "desc"

		res, _ := ioc.Repo.AccessLog.List(limit, from, to, order)
		log.Println(res)
	})

	t.Run("Test Get Latest Log", func(t *testing.T) {
		res, _ := ioc.Repo.AccessLog.GetLatest(100)
		log.Println(res)
	})

	t.Run("Test Get rank by cat id", func(t *testing.T) {
		from := time.Now().Add(-24 * 365 * time.Hour)
		to := time.Now()
		alloffCatIds := []string{""}
		alloffLev1Cats, _ := ioc.Repo.AlloffCategories.List(&alloffCatIds[0])

		res, _ := ioc.Repo.ProductLog.GetRankByCatId(100, from, to, alloffLev1Cats[0].ID.Hex())
		for _, pd := range res.Aggregations.GroupByState.Buckets {
			pd, _ := ioc.Repo.Products.Get(pd.Key)
			require.Equal(t, pd.AlloffCategories.First.ID.Hex(), alloffLev1Cats[0].ID.Hex())
		}
	})
}
