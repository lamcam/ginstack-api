package config

import (
	"fmt"
	"github.com/spf13/viper"
)
var (
	AppHost    string
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func init() {
	loadConfig()

	// server configuration
	AppHost = viper.GetString("APP_HOST")
	AppPort = viper.GetString("APP_PORT")

	// database configuration
	DBHost = viper.GetString("DB_HOST")
	DBPort = viper.GetString("DB_PORT")
	DBUser = viper.GetString("DB_USER")
	DBPassword = viper.GetString("DB_PASSWORD")
	DBName = viper.GetString("DB_NAME")
}

func loadConfig() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Config file loaded from ENV")
		return
	}

	fmt.Println("Failed to load anyt config file")
}
