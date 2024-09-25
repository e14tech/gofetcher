package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type CoinGecko struct {
	BitcoinCash struct {
		USD float64 `json:"usd"`
		BTC float64 `json:"btc"`
	} `json:"bitcoin-cash"`
}

func main() {
	bchUSD := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin-cash&vs_currencies=USD"
	bchBTC := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin-cash&vs_currencies=BTC"
	httpResp, httpErr := http.Get(bchUSD)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
	defer httpResp.Body.Close()

	htmlData, htmlErr := ioutil.ReadAll(httpResp.Body)
	if htmlErr != nil {
		fmt.Println(htmlErr)
		os.Exit(1)
	}

	fmt.Println(string(htmlData))

	var coinGecko CoinGecko
	jsonErr := json.Unmarshal(htmlData, &coinGecko)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		os.Exit(1)
	}

	httpResp, httpErr = http.Get(bchBTC)
	if httpErr != nil {
		log.Fatal(httpErr)
	}
	defer httpResp.Body.Close()

	htmlData, htmlErr = ioutil.ReadAll(httpResp.Body)
	if htmlErr != nil {
		fmt.Println(htmlErr)
		os.Exit(1)
	}

	fmt.Println(string(htmlData))

	jsonErr = json.Unmarshal(htmlData, &coinGecko)
	if jsonErr != nil {
		fmt.Println(jsonErr)
		os.Exit(1)
	}

	fmt.Printf("coinGecko: %v\n", coinGecko)
	fmt.Printf("parsedJSON: %v\n", coinGecko.BitcoinCash.USD)
	fmt.Printf("parsedJSON: %v\n", coinGecko.BitcoinCash.BTC)
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
