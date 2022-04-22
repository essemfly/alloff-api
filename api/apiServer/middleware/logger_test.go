package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/utils"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"os"
	"testing"
)

type TestResponse struct {
	Hits struct {
		Hits []struct {
			Source struct {
				URI string `json:"uri"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func TestElasticSearchLogger(t *testing.T) {
	cmd.SetBaseConfig("local")
	sv := chi.NewRouter()
	sv.Use(chimiddleware.RequestID)
	sv.Use(ElasticSearchLogger(log.New(os.Stdout, "", 0)))

	randEndpoint := utils.CreateShortUUID()
	sv.Get("/"+randEndpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to test endpoint"))
	})

	ts := httptest.NewServer(sv)
	defer ts.Close()

	t.Run("test es middleware", func(t *testing.T) {
		// request to test endpoint
		req, _ := http.NewRequest("GET", ts.URL+"/"+randEndpoint, nil)
		http.DefaultClient.Do(req)

		// waiting for elastic search indexing on outside of main goroutine
		time.Sleep(time.Second * 1)

		// esapi body
		var buf bytes.Buffer
		query := fmt.Sprintf(`
		{
			"query": {
				"term": {
					"uri": "%s/%s"
				}
			}
		}
		`, ts.URL, randEndpoint)
		body := strings.NewReader(query)
		json.NewEncoder(&buf).Encode(query)

		// esapi req
		searchReq := esapi.SearchRequest{
			Index: []string{"access_log"},
			Body:  body,
		}
		res, err := searchReq.Do(context.Background(), config.EsClient)
		if err != nil {
			log.Println("err getting response : ", err)
		}
		defer res.Body.Close()

		// decode response
		var respBody TestResponse
		err = json.NewDecoder(res.Body).Decode(&respBody)

		// test
		require.NoError(t, err)
		require.Equal(t, ts.URL+"/"+randEndpoint, respBody.Hits.Hits[0].Source.URI)
	})
}
