package main

import "github.com/spf13/viper"

type Config struct {
	apiKey   string `mapstructure:"API_KEY"`
	dbName   string `mapstructure:"DATABASE_NAME"`
	dbServer string `mapstructure:"DATABASE_SERVER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
