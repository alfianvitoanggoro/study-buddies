package kafka

import (
	"context"
	"log"
)

func ConsumeClassRegistration() {
	reader := NewReader()
	defer reader.Close()

	log.Println("👂 Listening for class registration events...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("❌ Error reading message: %v", err)
			continue
		}

		log.Printf("📥 Received: %s", string(m.Value))
	}
}

func ConsumeScheduleRegistration() {
	reader := NewScheduleReader()
	defer reader.Close()

	log.Println("👂 Listening for schedule registration events...")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("❌ Error reading message: %v", err)
			continue
		}

		log.Printf("📥 Received Schedule: %s", string(m.Value))
	}
}
