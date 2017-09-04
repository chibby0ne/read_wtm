package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
);

type Coin struct {
    Id uint64 `json:"id"`
    Tag string `json:"tag"`
    Algorithm string `json:"algorithm"`
    Block_time float64 `json:"block_time"`
    Block_reward float64 `json:"block_reward"`
    Block_reward24 float64 `json:"block_reward24"`
    Last_block uint64 `json:"last_block"`
    Difficulty float64 `json:"difficulty"`
    Difficulty24 float64 `json:"difficulty24"`
    Nethash float64 `json:"nethash"`
    Exchange_rate float64 `json:"exchange_rate"`
    Exchange_rate24 float64 `json:"exchange_rate24"`
    Exchange_rage_vol float64 `json:"exchange_rage_vol"`
    Exchange_rage_curr string `json:"exchange_rage_curr"`
    Market_cap string `json:"market_cap"`
    Estimated_rewards float64 `json:"estimated_rewards"`
    Estimated_rewards24 float64 `json:"estimated_rewards24"`
    Btc_revenue float64 `json:"btc_revenue"`
    Btc_revenue24 float64 `json:"btc_revenue24"`
    Profitability uint64 `json:"profitability"`
    Profitability24 uint64 `json:"profitability24"`
    Lagging bool `json:"lagging"`
    Timestamp uint64 `json:"timestamp"`
}

type Coins struct {
    Coins map[string]Coin `json:"coins"`
}

func main() {
    url := "https://whattomine.com/coins.json"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

    var coins Coins
    json.Unmarshal(responseData, &coins)

    var maxProfitability uint64
    var maxCoin string
    for coinName, coinContent := range coins.Coins {
        if coinContent.Profitability > maxProfitability {
            maxProfitability = coinContent.Profitability
            maxCoin = coinName
        }
    }
    fmt.Printf("The most profitable coin is: %s with %v of Profitability\n", maxCoin, maxProfitability)
}
