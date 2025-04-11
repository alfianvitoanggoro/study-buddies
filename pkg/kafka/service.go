package kafka

import (
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaTopic         = "class-registration"
	kafkaScheduleTopic = "schedule-registration"
	kafkaBrokerAddress = "study-buddies-kafka:9092"
)

func NewWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaBrokerAddress),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

func NewReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBrokerAddress},
		Topic:    kafkaTopic,
		GroupID:  "study-buddies-group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Second,
	})
}

func NewScheduleWriter() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaBrokerAddress),
		Topic:    kafkaScheduleTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

func NewScheduleReader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBrokerAddress},
		Topic:    kafkaScheduleTopic,
		GroupID:  "study-buddies-group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  1 * time.Second,
	})
}
