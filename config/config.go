package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig() // fallback to system env variables if not in file

	// Read the .env file
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	LoadApplicationConfig()
	LoadDatabaseConfig()
}
