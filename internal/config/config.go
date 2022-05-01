package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	AllowedOrigin string `mapstructure:"ALLOWED_ORIGIN"`
	PostgresUri   string `mapstructure:"POSTGRES_URI"`
	JwtSecret     string `mapstructure:"JWT_SECRET"`
	Debug         bool   `mapstructure:"DEBUG"`
	SentryDSN     string `mapstructure:"SENTRY_DSN"`
	Port          int    `mapstructure:"PORT"`
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
