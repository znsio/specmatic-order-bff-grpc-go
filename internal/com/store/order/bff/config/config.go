// File: config/config.go

package config

import (
	"fmt"

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
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config

	// Set default values
	viper.SetDefault("backend.host", "localhost")
	viper.SetDefault("backend.port", "9000")
	viper.SetDefault("bff-server.port", "8090")

	// Map environment variables
	viper.BindEnv("backend.host", "DOMAIN_SERVER_HOST")
	viper.BindEnv("backend.port", "DOMAIN_SERVER_PORT")
	viper.BindEnv("bff-server.port", "BFF_SERVER_PORT")

	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
