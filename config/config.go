package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"inventory-service/helpers/xvalidator"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	AppEnv   *App
	Database *Database
}

func (c Config) IsStaging() bool {
	return c.AppEnv.CurrentEnv != "production"
}

func (c Config) IsProd() bool {
	return c.AppEnv.CurrentEnv == "production"
}

func InitConfig(validate *xvalidator.Validator) *Config {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		slog.Error(fmt.Sprintf("Failed to read config file: %s", err))
		slog.Info("Using default environment variables...")
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()

	// Iterate through all the settings
	for key, value := range viper.AllSettings() {
		envKey := strings.ToUpper(key)
		err := os.Setenv(strings.ToUpper(key), fmt.Sprintf("%v", value))
		if err != nil {
			slog.Error(fmt.Sprintf("Error setting environment variable %s: %v", envKey, err))
			os.Exit(1)
		}
	}
	c := Config{
		AppEnv:   AppConfig(),
		Database: DatabaseConfig(),
	}
	errs := validate.Struct(c)
	if errs != nil {
		for k, e := range errs {
			slog.Error(fmt.Sprintf("Failed to load env: %s, msg: %s", k, strings.ToLower(e)))
		}
		os.Exit(1)
	}
	slog.Info("Config loaded")
	return &c
}
