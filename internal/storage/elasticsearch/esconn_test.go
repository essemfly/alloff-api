package elasticsearch

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConnElasticSearch(t *testing.T) {
	testConf := config.GetConfiguration("local")
	t.Run("Test Connection", func(t *testing.T) {
		conn := NewElasticSearch(testConf)
		res, err := conn.Client.Info()
		if err != nil {
			t.Errorf("Error getting response : %v", err)
		}
		defer res.Body.Close()

		require.Equal(t, 200, res.StatusCode)
	})

}
