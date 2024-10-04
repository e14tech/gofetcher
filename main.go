package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxRetries    = 2
	retryDelay    = 60 * time.Second
	fetchInterval = 300 * time.Second
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
	var priceData PriceData
	urlLink := "https://api.coingecko.com/api/v3/coins/bitcoin-cash"
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		CatchSig()
		os.Exit(1)
	}()
	for {
		GetData(urlLink, &priceData)

		fmt.Printf("Bitcoin Cash price in USD: $%.2f\n", priceData.MarketData.CurrentPrice.USD)
		fmt.Printf("Bitcoin Cash price in BTC: ₿%f\n", priceData.MarketData.CurrentPrice.BTC)
		fmt.Printf("Bitcoin Cash price in ETH: %f Ether\n\n", priceData.MarketData.CurrentPrice.ETH)

		time.Sleep(fetchInterval)
	}
}

func CatchSig() {
	fmt.Println("Quitting.")
}

func GetData(url string, coinData *PriceData) {
	for i := 0; i < maxRetries; i++ {
		httpResp, httpErr := http.Get(url)
		if httpErr != nil {
			PrintRetry(i, httpErr)
			continue
		}
		if httpResp.StatusCode != 200 {
			log.Println("HTTP error code: ", httpResp.StatusCode)
			htmlData, htmlErr := ioutil.ReadAll(httpResp.Body)
			if htmlErr != nil {
				PrintRetry(i, htmlErr)
			}
			htmlErr = errors.New(string(htmlData))
			PrintRetry(i, htmlErr)
			continue
		}
		defer httpResp.Body.Close()

		htmlData, htmlErr := ioutil.ReadAll(httpResp.Body)
		if htmlErr != nil {
			PrintRetry(i, htmlErr)
		}

		jsonErr := json.Unmarshal(htmlData, &coinData)
		if jsonErr != nil {
			PrintRetry(i, jsonErr)
		}
	}
}

func PrintRetry(tries int, err error) {
	if tries == 0 {
		log.Println(err)
		fmt.Printf("Will try again in one minute.\n")
	} else {
		log.Fatal("No more tries. ", err)
	}
	time.Sleep(retryDelay)
}
