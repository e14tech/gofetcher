package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/wormhole-foundation/wormhole-explorer/common/coingecko"
)

func main() {
	// var apiKey is located within the apikey.go file under the same directory as main.go.
	// apikey.go structure is just:
	// package main
	// var apiKey string = "YOUR_API_KEY"
	apiConnect := coingecko.NewCoinGeckoAPI("https://api.coingecko.com", "x_cg_demo_api_key", apiKey)

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
