package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type ClassRegistrationMessage struct {
	StudentID string `json:"student_id"`
	ClassID   string `json:"class_id"`
	Timestamp int64  `json:"timestamp"`
}

type ScheduleRegistrationMessage struct {
	ScheduleID string `json:"schedule_id"`
	ClassID    string `json:"class_id"`
	MaterialID string `json:"material_id"`
	Timestamp  int64  `json:"timestamp"`
}

func PublishClassRegistration(msg ClassRegistrationMessage) error {
	writer := NewWriter()
	defer writer.Close()

	value, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(msg.StudentID),
			Value: value,
		},
	)

	if err != nil {
		log.Printf("❌ Failed to write message: %v", err)
		return err
	}

	log.Println("✅ Message published to Kafka")
	return nil
}

func PublishScheduleRegistration(msg ScheduleRegistrationMessage) error {
	writer := NewScheduleWriter()
	defer writer.Close()

	value, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(msg.ScheduleID),
			Value: value,
		},
	)

	if err != nil {
		log.Printf("❌ Failed to write message: %v", err)
		return err
	}

	log.Println("✅ Message published to Kafka")
	return nil
}
