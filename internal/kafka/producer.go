// internal/kafka/producer.go

package kafka

import (
	"data-collector/domain"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// ProducerConfig struct to hold configuration for Kafka producer
type ProducerConfig struct {
	BootstrapServers string
}

// NewProducer creates and returns a new Kafka Producer based on the provided configuration
func NewProducer(config *ProducerConfig) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.BootstrapServers,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %v", err)
	}
	return producer, nil
}

// SendToKafka sends BESS data to the specified Kafka topic
func SendToKafka(data domain.Bess, producer *kafka.Producer, topic string) error {
	// Marshal data into bytes. Assume data is JSON-encoded
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling data to JSON: %v", err)
	}

	// Produce a message
	deliveryChan := make(chan kafka.Event, 1)
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          dataBytes,
	}, deliveryChan)

	if err != nil {
		return fmt.Errorf("produce failed: %v", err)
	}

	// Wait for message to be delivered and check for errors
	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %v", m.TopicPartition.Error)
	}

	// Close the channel
	close(deliveryChan)

	return nil
}
