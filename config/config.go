package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {

	fmt.Println("Yaha Tak Aa aa raha hain")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	//LoadDatabaseConfig()
	LoadApplicationConfig()

}
