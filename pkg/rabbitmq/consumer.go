package rabbitmq

import "log"

func StartConsumer(queueName string, handler func(msg string)) {
	msgs, err := Channel.Consume(
		queueName, "", true, false, false, false, nil,
	)
	if err != nil {
		log.Fatal("❌ Failed to register consumer:", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("📥 Received: %s", d.Body)
			handler(string(d.Body))
		}
	}()
}
