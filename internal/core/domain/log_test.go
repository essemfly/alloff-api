package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccessLogDAO_JsonEncoder(t *testing.T) {
	t.Run("nil pointer exception", func(t *testing.T) {
		var f *AccessLogDAO // pass nil value
		res := f.JsonEncoder()
		require.Equal(t, "{}", res)
	})
}
