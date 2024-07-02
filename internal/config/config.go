package config

import (
	"os"
)

// Config stores the configuration settings for the application
type Config struct {
	KafkaBrokerURL string
	ServerPort     string
}

// LoadConfig loads environment variables into Config
func LoadConfig() *Config {
	return &Config{
		KafkaBrokerURL: os.Getenv("KAFKA_BROKER_URL"),
		ServerPort:     os.Getenv("SERVER_PORT"),
	}
}
