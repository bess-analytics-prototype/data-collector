package utils

import (
	"data-collector/internal/model"
	"log"
)

// LogData logs data entries
func LogData(data model.BESSData) {
	log.Printf("Received data: %+v\n", data)
}
