package elasticsearch

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

func JsonEncoder(v interface{}) string {
	if v == nil || reflect.ValueOf(v).IsNil() {
		return "{}"
	} else {
		bytes, err := json.Marshal(v)
		if err != nil {
			log.Println("error on marshaling : ", err)
			return "{}"
		}
		return string(bytes)
	}
}

func RequestIndex(index, reqBody string, client *elasticsearch.Client) (int, error) {
	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(reqBody),
		Refresh: "true",
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Println("Error getting response : ", err)
		return 400, err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response : ", err)
		return 400, err
	}
	return res.StatusCode, nil
}

func RequestQuery(index, reqBody string, client *elasticsearch.Client) ([]byte, error) {
	indices := []string{index}
	req := esapi.SearchRequest{
		Index: indices,
		Body:  strings.NewReader(reqBody),
	}
	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Println("Error getting response : ", err)
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading response : ", err)
		return nil, err
	}
	return resBody, nil

}
