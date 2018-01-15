package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

const (
	BITCOINURL string = "https://api.coinmarketcap.com/v1/ticker/bitcoin/"
)

// Struct type that represents one coin from www.whattomine.com
type Coin struct {
	Id                  uint64  `json:"id"`
	Tag                 string  `json:"tag"`
	Algorithm           string  `json:"algorithm"`
	Block_time          float64 `json:"block_time"`
	Block_reward        float64 `json:"block_reward"`
	Block_reward24      float64 `json:"block_reward24"`
	Last_block          uint64  `json:"last_block"`
	Difficulty          float64 `json:"difficulty"`
	Difficulty24        float64 `json:"difficulty24"`
	Nethash             float64 `json:"nethash"`
	Exchange_rate       float64 `json:"exchange_rate"`
	Exchange_rate24     float64 `json:"exchange_rate24"`
	Exchange_rage_vol   float64 `json:"exchange_rage_vol"`
	Exchange_rage_curr  string  `json:"exchange_rage_curr"`
	Market_cap          string  `json:"market_cap"`
	Estimated_rewards   string  `json:"estimated_rewards"`
	Estimated_rewards24 string  `json:"estimated_rewards24"`
	Btc_revenue         string  `json:"btc_revenue"`
	Btc_revenue24       string  `json:"btc_revenue24"`
	Profitability       uint64  `json:"profitability"`
	Profitability24     uint64  `json:"profitability24"`
	Lagging             bool    `json:"lagging"`
	Timestamp           uint64  `json:"timestamp"`
}

// Struct type that represents one coin from www.coinmarketcap.com/api
type CoinMarketCapCoin struct {
	Id                 string `json:"id"`                 // "bitcoin"
	Name               string `json:"name"`               // "Bitcoin"
	Symbol             string `json:"symbol"`             //: "BTC",
	Rank               uint64 `json:"rank"`               // "1",
	Price_USD          string `json:"price_usd"`          //: "13887.8",
	Print_BTC          string `json:"price_btc"`          //: "1.0",
	Day_Volume_USD     string `json:"24h_volume_usd"`     //: "10438500000.0",
	Market_Cap_USD     string `json:"market_cap_usd"`     //: "232799274884",
	Available_Supply   string `json:"available_supply"`   //: "16762862.0",
	Total_Supply       string `json:"total_supply"`       //: "16762862.0",
	Max_Supply         string `json:"max_supply"`         //: "21000000.0",
	Percent_Change_1h  string `json:"percent_change_1h"`  //: "-0.55",
	Percent_Change_24h string `json:"percent_change_24h"` //: "1.9",
	Percent_Change_7d  string `json:"percent_change_7d"`  //: "-26.6",
	Last_Updated       string `json:"last_updated"`       //: "1514242461"
}

// Struct type that represents the top level type in www.whattomine.com, which contains the data from all coins
type Coins struct {
	Coins map[string]Coin `json:"coins"`
}

// Structure that contains key and value of map
// Used for sorting the keys inside the map according to value
type Pair struct {
	key   string
	value float64
}

// Slice of Pair that implements sort.Interface to sort by value
type PairList []Pair

// Method that returns the length for a PairList
func (p PairList) Len() int {
	return len(p)
}

// Method that returns true if the ith element is < jth element in a PairList slice
func (p PairList) Less(i, j int) bool {
	return p[i].value < p[j].value
}

// Method that swaps the ith element  with the jth element in a PairList slice
// Used for sorting
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//  Returns a PairList slice which is sorted by value
func SortMapByValue(m map[string]float64) PairList {
	// Create a PairList containing all the pairs
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	// Sort the PairList
	sort.Sort(sort.Reverse(p))
	return p
}

// Global httpClient with custom Timeout of 10 sec
var myClient = &http.Client{Timeout: 10 * time.Second}

// Read a url an if successful writes returns the body of the page as a byte slice
func readJsonFromUrl(url string, target interface{}) {
	// Get request to the site, and defer closing body
	response, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// read all body of page
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(responseData, target)
}

// Converts string to float64 and returns float64. Includes error handling
func convertToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
