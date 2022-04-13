package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInitElasticSearch(t *testing.T) {
	testConf := GetConfiguration("local")
	t.Run("Test Connection", func(t *testing.T) {
		InitElasticSearch(testConf)
		res, err := EsClient.Info()
		if err != nil {
			t.Errorf("Error getting response: %s", err)

		}
		defer res.Body.Close()

		require.Equal(t, 200, res.StatusCode)
	})
}
