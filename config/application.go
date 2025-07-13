package config

import (
	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	JWT_TOKEN      string
	REFERESH_TOKEN string
}

var applicationConfig ApplicationConfig

func GetApplicationConfig() *ApplicationConfig {

	return &applicationConfig

}

func LoadApplicationConfig() {

	applicationConfig = ApplicationConfig{
		JWT_TOKEN:      viper.GetString("JWT_TOKEN"),
		REFERESH_TOKEN: viper.GetString("REFERESH_TOKEN"),
	}

}
