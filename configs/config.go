// File: configs/config.go
// Loads application configuration using Viper

package configs

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}
}
