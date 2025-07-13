package config

import (
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Database string
	User     string
	Password string
	Port     string
}

var databaseConfig DatabaseConfig

func getDatabaseConfig() *DatabaseConfig {
	return &databaseConfig
}

func LoadDatabaseConfig() {

	databaseConfig = DatabaseConfig{

		Host:     viper.GetString("DATABASE_HOST"),
		Database: viper.GetString("DATABASE_NAME"),
		User:     viper.GetString("DATABASE_USER"),
		Password: viper.GetString("DATABASE_PASSWORD"),
		Port:     viper.GetString("DATABASE_PORT"),
	}

}
