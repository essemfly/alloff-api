package middleware

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// secret key being used to sign tokens
var (
	SecretKey = []byte("lessbutteroutletapi")
)

func GenerateToken(mobile, uuid string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["mobile"] = mobile
	claims["uuid"] = uuid
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Middleware에서 parsing해서 valid한지 체크 + Refreshtoken요청할때 ParseToken으로 사용됨
func ParseToken(tokenStr string) (string, string, error) {
	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		mobile := claims["mobile"].(string)
		uuid := claims["uuid"].(string)
		return mobile, uuid, nil
	} else {
		// Expired 된경우 -> 그냥 Refresh에서 새거 받아올수있도록 한다.
		// Return으로 mobile과 uuid가 필요한 이유는 resolver refresh token 함수에서 mobile로 체크하고, 해결해주기 위함
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			if claims["mobile"] != nil {
				mobile := claims["mobile"].(string)
				uuid := claims["uuid"].(string)
				return mobile, uuid, errors.New("Invalid token")
			}
			// 시간이 지났는데 안에가 변조된 경우
			return "", "", errors.New("Invalid token")
		} else {
			// 시간 안지났는데 안에가 변조된 경우
			return "", "", errors.New("Invalid token")
		}
	}
}
