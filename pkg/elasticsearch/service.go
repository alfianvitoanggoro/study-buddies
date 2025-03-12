package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/sirupsen/logrus"
)

var (
	appName = os.Getenv("APP")
	Client  *elasticsearch.Client
)

// CustomLogger implements elasticsearch Logger
type CustomLogger struct{}

// LogRoundTrip logs request and response details
func (l *CustomLogger) LogRoundTrip(req *http.Request, res *http.Response, err error, start time.Time, dur time.Duration) error {
	entry := logrus.WithFields(logrus.Fields{
		"app_name": appName,
		"method":   req.Method,
		"url":      req.URL.String(),
		"duration": dur,
	})

	if err != nil {
		entry.WithError(err).Error("Request failed")
		return err
	}

	entry = entry.WithFields(logrus.Fields{
		"status_code": res.StatusCode,
		"response":    readResponseBody(res),
	})

	if res.StatusCode >= 500 {
		entry.Error("Server error")
	} else if res.StatusCode >= 400 {
		entry.Warn("Client error")
	} else {
		entry.Info("Request successful")
	}

	return nil
}

// RequestBodyEnabled enables request logging
func (l *CustomLogger) RequestBodyEnabled() bool { return true }

// ResponseBodyEnabled enables response logging
func (l *CustomLogger) ResponseBodyEnabled() bool { return true }

// readResponseBody reads and returns response body
func readResponseBody(res *http.Response) string {
	if res.Body == nil {
		return ""
	}
	bodyBytes, _ := io.ReadAll(res.Body)
	res.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Reset response body
	return string(bodyBytes)
}

// Init initializes Elasticsearch connection
func Init() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			os.Getenv("ELASTIC_SEARCH_URL"),
		},
		Transport: &http.Transport{
			// MaxIdleConnsPerHost:   10,
			// ResponseHeaderTimeout: time.Millisecond,
			// DialContext:           (&net.Dialer{Timeout: time.Nanosecond}).DialContext,
			// TLSClientConfig: &tls.Config{
			// MinVersion: tls.VersionTLS12,
			// ...
			// },
		},

		Logger: &CustomLogger{},
	}

	var err error
	Client, err = elasticsearch.NewClient(cfg)
	if err != nil {
		logrus.Fatalf("Error creating Elasticsearch client: %s", err)
	}
	logrus.Info("✅ Elasticsearch successfully connected")
}

// Insert adds a document to Elasticsearch
func Insert(ctx context.Context, index string, doc interface{}) error {
	jsonBody, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to encode document: %w", err)
	}

	req := esapi.IndexRequest{
		Index:   fmt.Sprintf("%s.%s", appName, index),
		Body:    bytes.NewReader(jsonBody),
		Refresh: "true",
	}

	res, err := req.Do(ctx, Client)
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

// Update modifies an existing document
func Update(ctx context.Context, index, ID string, update map[string]interface{}) error {
	jsonBody, err := json.Marshal(map[string]interface{}{"doc": update})
	if err != nil {
		return fmt.Errorf("failed to encode update document: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: ID,
		Body:       bytes.NewReader(jsonBody),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, Client)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"index": fmt.Sprintf("%s.%s", appName, index),
			"id":    ID,
			"data":  update,
		}).Error("Failed to update document")
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("update failed, status: %s", res.Status())
	}

	logrus.WithFields(logrus.Fields{
		"index": fmt.Sprintf("%s.%s", appName, index),
		"id":    ID,
	}).Info("Document updated successfully")

	return nil
}

// Search queries Elasticsearch
func Search(ctx context.Context, index string, query map[string]interface{}) (map[string]interface{}, error) {
	jsonBody, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("failed to encode search query: %w", err)
	}

	req := esapi.SearchRequest{
		Index: []string{index},
		Body:  bytes.NewReader(jsonBody),
	}

	res, err := req.Do(ctx, Client)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"index": fmt.Sprintf("%s.%s", appName, index),
			"query": string(jsonBody),
		}).Error("Failed to execute search")
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("search failed, status: %s", res.Status())
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode search response: %w", err)
	}

	logrus.WithFields(logrus.Fields{
		"index": fmt.Sprintf("%s.%s", appName, index),
		"query": string(jsonBody),
	}).Info("Search executed successfully")

	return result, nil
}
