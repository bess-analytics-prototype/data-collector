package api

import (
	//"data-collector/internal/kafka"
	"data-collector/internal/model"
	"net/http"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
)

// Pending: Producer and topic needs to be initialized globally or passed via context
// Here, we're using a global variable for simplicity
var producer *confluentKafka.Producer
var kafkaTopic string

// Temp removal of kafka
// func init() {
// 	var err error
// 	//  @TODO Initialize these values from config system or environment variables
// 	config := &kafka.ProducerConfig{BootstrapServers: "localhost:9092"}
// 	producer, err = kafka.NewProducer(config)
// 	if err != nil {
// 		panic("Failed to create Kafka producer")
// 	}
// 	kafkaTopic = "data-topic"
// }

// PostData handles incoming POST requests and sends data to Kafka
func PostData(c *gin.Context) {
	var data model.BESSData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Send data to Kafka
	// Temp removal of kafka
	// if err := kafka.SendToKafka(data, producer, kafkaTopic); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{"status": "data sent to Kafka"})
}
