package middleware

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJwt(t *testing.T) {
	token := ""
	testUserMobile := "01073881067"
	testUserUUID := "aeb06898-5183-4fca-9e37-851999f26f5a"

	t.Run("Test GenerateToken", func(t *testing.T) {
		generatedToken, err := GenerateToken(testUserMobile, testUserUUID)
		require.NoError(t, err)
		token = generatedToken
	})

	t.Run("Test ParseToken", func(t *testing.T) {
		// TODO Expired 됐는데 정상일때 / Expired 됐고 비정상일 때 정상응답오는지도 테스트 필요
		if token == "" {
			generatedToken, err := GenerateToken(testUserMobile, testUserUUID)
			require.NoError(t, err)
			token = generatedToken
		}
		m, u, err := ParseToken(token)
		require.NoError(t, err)
		require.Equal(t, testUserMobile, m)
		require.Equal(t, testUserUUID, u)
	})
}
