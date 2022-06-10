package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

var UserCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

var ErrInvalidToken = errors.New("ERR000:invalid token")
var ErrDeviceChanged = errors.New("ERR001:device changed")
var ErrNoTokenGiven = errors.New("ERR002:no token given")

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if header == "" {
				ctx := context.WithValue(r.Context(), UserCtxKey, ErrNoTokenGiven)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}
			//validate jwt token
			tokenStr := header
			mobile, uuid, err := ParseToken(tokenStr)
			// Expired된 경우 여기로 통과해서 들어옴
			if err != nil {
				//http.Error(w, "Invalid token", http.StatusForbidden)
				ctx := context.WithValue(r.Context(), UserCtxKey, ErrInvalidToken)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			// create user and check if user exists in db
			user, err := ioc.Repo.Users.GetByMobile(mobile)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			// 다른 기기에서 API를 사용하는 경우 Forbidden 사용
			if user.Uuid != uuid {
				//http.Error(w, "Device Changed", http.StatusForbidden)
				ctx := context.WithValue(r.Context(), UserCtxKey, ErrDeviceChanged)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), UserCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) (*domain.UserDAO, error) {
	raw, ok := ctx.Value(UserCtxKey).(*domain.UserDAO)
	if ok {
		return raw, nil
	}

	errMsg, _ := ctx.Value(UserCtxKey).(error)
	switch errMsg {
	case ErrInvalidToken:
		return nil, ErrInvalidToken
	case ErrDeviceChanged:
		return nil, ErrDeviceChanged
	case ErrNoTokenGiven:
		return nil, ErrNoTokenGiven
	default:
		return nil, fmt.Errorf("ERR_UNKNOWN")
	}
}
