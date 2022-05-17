package elasticsearch

// import (
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func TestConnElasticSearch(t *testing.T) {
// 	t.Run("Test Connection", func(t *testing.T) {
// 		conn := NewElasticSearch()
// 		res, err := conn.Client.Info()
// 		if err != nil {
// 			t.Errorf("Error getting response : %v", err)
// 		}
// 		defer res.Body.Close()

// 		require.Equal(t, 200, res.StatusCode)
// 	})
// }
