package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port                string `mapstructure:"PORT"`
	SignKey             string `mapstructure:"SIGN_KEY"`
	AllowFraudDetection bool   `mapstructure:"ALLOW_FRAUD_DETECTION"`
	AllowLimit          bool   `mapstructure:"ALLOW_LIMIT_DETECTION"`
	AllowDailyLimit     bool   `mapstructure:"ALLOW_DAILY_LIMIT"`
	set                 bool
}

var config Config

func GetConfig() Config {
	if !config.set {
		config = setup()
	}
	return config
}

func setup() Config {

	c := Config{}
	setDefaults()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	c.set = true

	return c
}

func setDefaults() {
	viper.SetDefault("PORT", "8000")
	viper.SetDefault("ALLOW_FRAUD_DETECTION", true)
	viper.SetDefault("ALLOW_LIMIT", true)
	viper.SetDefault("ALLOW_DAILY_LIMIT", true)
}
