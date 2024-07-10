// File: config/config.go

package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Backend struct {
		Host string
		Port string
	}
	BFFServer struct {
		Port string
	}
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config

	// Set default values
	viper.SetDefault("backend.host", "localhost")
	viper.SetDefault("backend.port", "9000")
	viper.SetDefault("bffserver.port", "8090")

	// Read the config file
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal the config file values
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Override with environment variables if they are set
	if envHost := os.Getenv("DOMAIN_SERVER_HOST"); envHost != "" {
		config.Backend.Host = envHost
	}
	if envPort := os.Getenv("DOMAIN_SERVER_PORT"); envPort != "" {
		config.Backend.Port = envPort
	}
	if envBFFPort := os.Getenv("BFF_SERVER_PORT"); envBFFPort != "" {
		config.BFFServer.Port = envBFFPort
	}

	return &config, nil
}
