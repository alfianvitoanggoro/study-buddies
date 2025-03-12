package elasticsearch

import (
	"fmt"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

func Init() (*elasticsearch.Client, error) {
	elasticsearchUrl := os.Getenv("ELASTICSEARCH_URL")

	cfg := elasticsearch.Config{
		Addresses: []string{elasticsearchUrl}, // Ganti dengan alamat Elasticsearch
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	// Mengecek koneksi
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error connecting to Elasticsearch: %v", err)
	}
	defer res.Body.Close()

	fmt.Println("Connected to Elasticsearch")

	return es, nil
}
