package config

import (
	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	JWT_TOKEN            string
	REFERESH_TOKEN       string
	GOOGLE_CLIENT_ID     string
	GOOGLE_CLIENT_SECRET string
	GOOGLE_REDIRECT_URL  string
}

var applicationConfig ApplicationConfig

func GetApplicationConfig() *ApplicationConfig {

	return &applicationConfig

}

func LoadApplicationConfig() {

	applicationConfig = ApplicationConfig{
		JWT_TOKEN:            viper.GetString("JWT_TOKEN"),
		REFERESH_TOKEN:       viper.GetString("REFERESH_TOKEN"),
		GOOGLE_CLIENT_ID:     viper.GetString("GOOGLE_CLIENT_ID"),
		GOOGLE_CLIENT_SECRET: viper.GetString("GOOGLE_CLIENT_SECRET"),
		GOOGLE_REDIRECT_URL:  viper.GetString("GOOGLE_REDIRECT_URL"),
	}

}
