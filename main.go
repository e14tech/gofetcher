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

	client := &http.Client{}
	req, reqErr := http.NewRequest("GET", urlLink, nil)
	if reqErr != nil {
		log.Fatal(reqErr)
	}
	req.Header.Set("User-Agent", "gofetcher")

	for {
		UnmarshalJSON(GetData(req, client), &priceData)

		fmt.Printf("Bitcoin Cash price in USD: $%.2f\n", priceData.MarketData.CurrentPrice.USD)
		fmt.Printf("Bitcoin Cash price in BTC: ₿%f\n", priceData.MarketData.CurrentPrice.BTC)
		fmt.Printf("Bitcoin Cash price in ETH: %f Ether\n\n", priceData.MarketData.CurrentPrice.ETH)

		time.Sleep(fetchInterval)
	}
}

func CatchSig() {
	fmt.Println("Quitting.")
}

// Grabs HTML data from CoinGecko.
// Also retries in case of an error.
func GetData(req *http.Request, client *http.Client) []byte {
	var htmlData []byte
	for i := 0; i < maxRetries; i++ {
		httpResp, httpErr := client.Do(req)

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
		return htmlData
	}
	return htmlData
}

// Unmarshals the JSON data from the HTML data and puts in into the PriceData struct.
// If we've gotten this far with a 200 Status code AND the JSON unmarshaling errors out, something else is wrong.
func UnmarshalJSON(htmlData []byte, coinData *PriceData) {
	jsonErr := json.Unmarshal(htmlData, &coinData)
	if jsonErr != nil {
		log.Fatal(jsonErr)
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
