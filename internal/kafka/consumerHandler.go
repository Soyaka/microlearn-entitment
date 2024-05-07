package kafka

import (
	"fmt"

	"github.com/IBM/sarama"
)

type CustomConsumerGroupHandler struct {
}

func NewConsumerGroupHandler() *CustomConsumerGroupHandler {
	return &CustomConsumerGroupHandler{}
}

func (h *CustomConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	// Called when a new consumer group session is started
	return nil
}

func (h *CustomConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	// Called when a consumer group session is finished
	return nil
}

func (h *CustomConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	// Called when the consumer group claims partitions for consumption
	for msg := range claim.Messages() {
		// Process the consumed message
		fmt.Printf("Message claimed: Topic - %s, Partition - %d, Offset - %d, Value - %s\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		// Mark the message as processed
		session.MarkMessage(msg, "")
	}
	return nil
}
