package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/wormhole-foundation/wormhole-explorer/common/coingecko"
)

func main() {
	// Load config file
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	apiConnect := coingecko.NewCoinGeckoAPI("https://api.coingecko.com", "x_cg_demo_api_key", config.apiKey)

	for i := 0; i < 3; i++ {
		bchMarketData, err := apiConnect.GetMarketData("bitcoin-cash")
		if err == nil {
			bchPrice := fmt.Sprintf("%s", bchMarketData.MarketData.CurrentPrice)
			bchPrice = strings.ReplaceAll(strings.ReplaceAll(bchPrice, "{", ""), "}", "")
			fmt.Printf("Current Bitcoin Cash price: $%s\n", bchPrice)
			break
		} else {
			fmt.Printf("%s\n", err)
			if i == 0 {
				fmt.Println("Will try again in a minute for two more tries.")
			}
			if i == 1 {
				fmt.Println("Will try again in a minute for one more try.")
			}
			time.Sleep(60 * time.Second)

			if i >= 2 {
				panic(err)
			}
		}
	}
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
