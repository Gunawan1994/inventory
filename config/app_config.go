package config

import (
	"github.com/spf13/viper"
)

type App struct {
	CurrentEnv  string
	HttpPort    string `validate:"required,number" name:"HTTP_PORT"`
	LogFilePath string `validate:"required" name:"LOG_PATH"`
}

func AppConfig() *App {
	return &App{
		CurrentEnv:  viper.GetString("CURRENT_ENV"),
		HttpPort:    viper.GetString("HTTP_PORT"),
		LogFilePath: viper.GetString("LOG_PATH"),
	}
}
