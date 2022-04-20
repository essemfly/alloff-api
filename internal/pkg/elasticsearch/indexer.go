package elasticsearch

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/lessbutter/alloff-api/config"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func logRequest(request *LogFields) (int, error) {
	index := "access_log"
	bodyStr := request.jsonEncoder()
	header := http.Header{}
	header.Add("Authorization", "Basic "+config.EsAPIKEY)
	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(bodyStr),
		Refresh: "true",
		Header:  header,
	}
	res, err := req.Do(context.Background(), config.EsClient)
	defer res.Body.Close()
	if err != nil {
		log.Println("Error getting response: %s", err)
		return 400, err
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return 400, err
	}
	return res.StatusCode, nil
}
