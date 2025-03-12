package repository

import (
	"bytes"
	"context"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
)

type ElasticsearchRepository interface {
	Insert(ctx context.Context, index string, doc interface{}) error
}

type elasticsearchRepository struct {
	es *elasticsearch.Client
}

func NewElasticsearch(es *elasticsearch.Client) *elasticsearchRepository {
	return &elasticsearchRepository{
		es,
	}
}

var (
	appName = "study-buddies"
)

func (er *elasticsearchRepository) Insert(ctx context.Context, index string, doc interface{}) error {
	jsonBody, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to encode document: %w", err)
	}

	req := esapi.IndexRequest{
		Index:   fmt.Sprintf("%s.%s", appName, index),
		Body:    bytes.NewReader(jsonBody),
		Refresh: "true",
	}

	res, err := req.Do(ctx, er.es)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"app_name": appName,
			"index":    fmt.Sprintf("%s.%s", appName, index),
			"data":     string(jsonBody),
		}).Error("Failed to insert document")
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("insert failed, status: %s", res.Status())
	}

	logrus.WithFields(logrus.Fields{
		"app_name": appName,
		"index":    fmt.Sprintf("%s.%s", appName, index),
		"data":     string(jsonBody),
	}).Info("Document inserted successfully")

	return nil
}
