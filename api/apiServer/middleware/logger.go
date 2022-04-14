package middleware

import (
	"github.com/go-chi/chi/middleware"
	"github.com/lessbutter/alloff-api/internal/pkg/elasticsearch"
	"log"
	"net/http"
)

// ElasticSearchLogger : LogFormatter 의 구현체인 StructuredLogger 를 미들웨어에 등록한다.
func ElasticSearchLogger(logger *log.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&elasticsearch.StructuredLogger{Logger: logger})
}
