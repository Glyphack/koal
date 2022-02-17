package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("koal")
	viper.BindEnv("jwt_secret")
	viper.BindEnv("postgres_uri")
	viper.BindEnv("debug")
	viper.SetDefault("debug", true)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
