package util

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DB_DRIVER           string        `mapstructure:"DB_DRIVER"`
	DB_SOURCE           string        `mapstructure:"DB_SOURCE"`
	SERVER_ADDRESS      string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	REDIRECT_URL        string        `mapstructure:"REDIRECT_URL"`
	CLIENT_ID           string        `mapstructure:"CLIENT_ID"`
	CLIENT_SECRET       string        `mapstructure:"CLIENT_SECRET"`
	SCOPES              string        `mapstructure:"SCOPES"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}
	return
}
