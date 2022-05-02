package elasticsearch

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccessLogDAO_JsonEncoder(t *testing.T) {
	t.Run("nil pointer exception on AccessLogDAO", func(t *testing.T) {
		var v *domain.AccessLogDAO // pass nil value
		res := JsonEncoder(v)
		require.Equal(t, "{}", res)
	})
}
