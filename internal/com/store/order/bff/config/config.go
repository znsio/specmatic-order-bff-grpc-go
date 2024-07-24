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

// func LoadConfig(configPath string) (*Config, error) {
// 	var config Config

// 	// Set default values
// 	viper.SetDefault("backend.host", "localhost")
// 	viper.SetDefault("backend.port", "9000")
// 	viper.SetDefault("bffserver.port", "8080")

// 	// Read the config file
// 	viper.SetConfigFile(configPath)
// 	if err := viper.ReadInConfig(); err != nil {
// 		return nil, fmt.Errorf("failed to read config file: %w", err)
// 	}

// 	// Unmarshal the config file values
// 	if err := viper.Unmarshal(&config); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
// 	}

// 	// Override with environment variables if they are set
// 	if envHost := os.Getenv("DOMAIN_SERVER_HOST"); envHost != "" {
// 		config.Backend.Host = envHost
// 	}
// 	if envPort := os.Getenv("DOMAIN_SERVER_PORT"); envPort != "" {
// 		config.Backend.Port = envPort
// 	}
// 	if envBFFPort := os.Getenv("BFF_SERVER_PORT"); envBFFPort != "" {
// 		config.BFFServer.Port = envBFFPort
// 	}

// 	return &config, nil
// }
