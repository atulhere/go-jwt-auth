package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	LoadDatabaseConfig()
	LoadApplicationConfig()

}
