package utils

import (
	"data-collector/domain"
	"log"
)

// LogData logs data entries
func LogData(data domain.Bess) {
	log.Printf("Received data: %+v\n", data)
}
