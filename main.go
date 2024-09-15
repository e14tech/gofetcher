package main

import (
	"fmt"
	"strings"

	"github.com/wormhole-foundation/wormhole-explorer/common/coingecko"
)

func main() {
	// var apiKey is located within the apikey.go file under the same directory as main.go.
	// apikey.go structure is just:
	// package main
	// var apiKey string = "YOUR_API_KEY"
	apiConnect := coingecko.NewCoinGeckoAPI("https://api.coingecko.com", "x_cg_demo_api_key", apiKey)

	bchMarketData, err := apiConnect.GetMarketData("bitcoin-cash")

	if err != nil {
		panic(err)
	}

	bchPrice := fmt.Sprintf("%s", bchMarketData.MarketData.CurrentPrice)

	bchPrice = strings.ReplaceAll(strings.ReplaceAll(bchPrice, "{", ""), "}", "")

	fmt.Printf("Current Bitcoin Cash price: $%s\n", bchPrice)
}
