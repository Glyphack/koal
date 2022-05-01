package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	allowedOrigin string `mapstructure:"ALLOWED_ORIGIN"`
	postgresUri   string `mapstructure:"POSTGRES_URI"`
	jwtSecret     string `mapstructure:"JWT_SECRET"`
	debug         bool   `mapstructure:"DEBUG"`
	sentryDSN     string `mapstructure:"SENTRY_DSN"`
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal parsing config file: %w \n", err))
	}
}
