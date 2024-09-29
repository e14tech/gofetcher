package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PriceData struct {
	MarketData struct {
		CurrentPrice struct {
			USD float64 `json:"usd"`
			BTC float64 `json:"btc"`
			ETH float64 `json:"eth"`
		} `json:"current_price"`
	} `json:"market_data"`
}

func main() {
	urlLinks := []string{"https://api.coingecko.com/api/v3/coins/bitcoin-cash"}
	var priceData PriceData

	GetData(urlLinks, &priceData)

	fmt.Printf("Bitcoin Cash price in USD: $%.2f\n", priceData.MarketData.CurrentPrice.USD)
	fmt.Printf("Bitcoin Cash price in BTC: ₿%f\n", priceData.MarketData.CurrentPrice.BTC)
	fmt.Printf("Bitcoin Cash price in ETH: %f Ether\n", priceData.MarketData.CurrentPrice.ETH)
}

func GetData(urls []string, coinData *PriceData) {
	for _, u := range urls {
		httpResp, httpErr := http.Get(u)
		if httpErr != nil {
			log.Fatal(httpErr)
		}
		if httpResp.StatusCode != 200 {
			log.Println("HTTP error code: ", httpResp.StatusCode)
			htmlData, htmlErr := ioutil.ReadAll(httpResp.Body)
			if htmlErr != nil {
				log.Println(htmlErr)
			}
			log.Println(string(htmlData))
			continue
		}
		defer httpResp.Body.Close()

		htmlData, htmlErr := ioutil.ReadAll(httpResp.Body)
		if htmlErr != nil {
			fmt.Println(htmlErr)
			os.Exit(1)
		}

		jsonErr := json.Unmarshal(htmlData, &coinData)
		if jsonErr != nil {
			fmt.Println(jsonErr)
			os.Exit(1)
		}
	}
}
