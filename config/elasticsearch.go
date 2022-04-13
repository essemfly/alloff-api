package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

var EsClient *elasticsearch.Client

func InitElasticSearch(conf Configuration) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.ELASTICSEARCH_URL,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Printf("Error creating the client: %s", err)
	}
	EsClient = es
}
