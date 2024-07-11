// File: scripts/read_port.go

package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	port := viper.GetString("bffserver.port")
	if port == "" {
		log.Fatalf("BFF server port not found in config")
	}

	// Output in a format that GitHub Actions can use to set an environment variable
	fmt.Printf("BFF_SERVER_PORT=%s\n", port)
}
