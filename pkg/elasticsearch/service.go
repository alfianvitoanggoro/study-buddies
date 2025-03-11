package elasticsearch

import (
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func Init() (*elasticsearch.Client, error) {
	elasticsearchUrl := os.Getenv("ELASTICSEARCH_URL")

	cfg := elasticsearch.Config{
		Addresses: []string{elasticsearchUrl}, // Ganti dengan alamat Elasticsearch
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	return client, nil
}
