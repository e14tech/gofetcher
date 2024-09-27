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
	urlLinks := []string{"https://api.coingecko.com/api/v3/simple/price?ids=bitcoin-cash&vs_currencies=USD", "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin-cash&vs_currencies=BTC"}

	var coinGecko CoinGecko

	GetData(urlLinks, &coinGecko)

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

func GetData(urls []string, coinData *CoinGecko) {
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
