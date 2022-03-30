package utils

import (
	"flag"

	"github.com/spf13/viper"
)

type Config struct {
	Port                  string `mapstructure:"PORT"`
	SignKey               string `mapstructure:"SIGN_KEY"`
	DisableFraudDetection bool   `mapstructure:"DISABLE_FRAUD_DETECTION"`
	DisableLimit          bool   `mapstructure:"DISABLE_LIMIT"`
	DisableDailyLimit     bool   `mapstructure:"DISABLE_DAILY_LIMIT"`
	set                   bool
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
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	setAfterFlags()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	c.set = true

	return c
}

func setAfterFlags() {
	fdFlag := flag.Bool("disable_fraud", false, "disable fraud detection")
	lFlag := flag.Bool("disable_limit", false, "disable limit")
	dlFlag := flag.Bool("disable_daily_limit", false, "disable daily limit")
	flag.Parse()

	if *fdFlag {
		viper.Set("DISABLE_FRAUD_DETECTION", true)
	}
	if *lFlag {
		viper.Set("DISABLE_LIMIT", true)
	}
	if *dlFlag {
		viper.Set("DISABLE_DAILY_LIMIT", true)
	}
}
