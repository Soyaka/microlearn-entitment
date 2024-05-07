package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	Consumer *kafka.Consumer
}

var (
	KafkaServer  = "localhost:9092"
	KafkaTopics  = []string{"bookmarks", "interests", "progress", "subscriptions"}
	KafkaGroupId = "rdkafka-f9a5f53b-cc4e-45b403cee"
)

func NewConsumer() *Consumer {

	return &Consumer{
		Consumer: NewBookmarkConsumer(),
	}
}

func NewBookmarkConsumer() *kafka.Consumer {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{

		"bootstrap.servers": KafkaServer,
		"group.id":          KafkaGroupId,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	c.SubscribeTopics(KafkaTopics, nil)
	return c
}
