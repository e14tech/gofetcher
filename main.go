package main

import (
	"fmt"
	"strings"

	"github.com/wormhole-foundation/wormhole-explorer/common/coingecko"
)

func main() {
	apiConnect := coingecko.NewCoinGeckoAPI("https://api.coingecko.com", "x_cg_demo_api_key", "CG-gr2qVv94bSy3XJjuv29GqyZy")

	bchMarketData, err := apiConnect.GetMarketData("bitcoin-cash")

	if err != nil {
		panic(err)
	}

	bchPrice := fmt.Sprintf("%s", bchMarketData.MarketData.CurrentPrice)

	bchPrice = strings.ReplaceAll(strings.ReplaceAll(bchPrice, "{", ""), "}", "")

	fmt.Printf("Current Bitcoin Cash price: $%s\n", bchPrice)
}
