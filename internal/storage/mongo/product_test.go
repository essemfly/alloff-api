package mongo

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestMongoProductsRepo(t *testing.T) {
	conf := config.GetConfiguration("dev")
	conn := NewMongoDB(conf)
	conn.RegisterRepos()

	t.Run("test get products by list of id", func(t *testing.T) {
		ids := []string{}
		samplePds, _, _ := ioc.Repo.Products.List(0, 100, bson.M{}, bson.M{})
		for _, pd := range samplePds {
			ids = append(ids, pd.ID.Hex())
		}

		pds, _ := ioc.Repo.Products.ListByIds(ids)
		require.Equal(t, samplePds, pds)
	})
}
