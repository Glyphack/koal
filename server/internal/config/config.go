package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
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
	viper.AddConfigPath("..")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warningf("No config file found in paths: %v", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("Fatal parsing config file: %w \n", err))
	}
}
