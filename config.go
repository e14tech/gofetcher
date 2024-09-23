package main

type Config struct {
	apiKey   string `mapstructure:"API_KEY"`
	dbName   string `mapstructure:"DATABASE_NAME"`
	dbServer string `mapstructure:"DATABASE_SERVER"`
}
