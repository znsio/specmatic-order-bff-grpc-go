// File: config/config.go

package config

import (
	"fmt"
	"os"
)

type Config struct {
	Backend struct {
		Host string
		Port string
	}
	BFFServer struct {
		Port string
	}
	KafkaService struct {
		Port  string
		Host  string
		Topic string
	}
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Backend: struct {
			Host string
			Port string
		}{
			Port: getEnvOrDefault("DOMAIN_SERVER_PORT", "9000"),
			Host: getEnvOrDefault("DOMAIN_SERVER_HOST", "order-api-mock"),
		},
		BFFServer: struct {
			Port string
		}{
			Port: getEnvOrDefault("SERVER_PORT", "8080"),
		},
		KafkaService: struct {
			Port  string
			Host  string
			Topic string
		}{
			Port:  getEnvOrDefault("KAFKA_PORT", "9093"),
			Host:  getEnvOrDefault("KAFKA_HOST", "specmatic-kafka"),
			Topic: getEnvOrDefault("KAFKA_TOPIC", "product-queries"),
		},
	}

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		fmt.Printf("%s received via env var: %s\n", key, value)
		return value
	}
	fmt.Printf("%s using default value: %s\n", key, defaultValue)
	return defaultValue
}
