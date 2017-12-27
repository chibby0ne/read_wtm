package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

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

type Coins struct {
	Coins map[string]Coin `json:"coins"`
}

/* Structure that contains key and value of map */
type Pair struct {
	key   string
	value float64
}

/* Slice of Pair that implements sort.Interface to sort by value */
type PairList []Pair

func (p PairList) Len() int {
	return len(p)
}
func (p PairList) Less(i, j int) bool {
	return p[i].value < p[j].value
}
func (p PairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

/* Returns a slice which is sorted by value */
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

var myClient = &http.Client{Timeout: 10 * time.Second}

/* Read a url an if successful writes returns the body of the page as a byte slice */
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

func convertToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func main() {
	// url := "https://whattomine.com/coins.json"
    url := "https://whattomine.com/coins.json?utf8=%E2%9C%93&adapt_q_280x=0&adapt_q_380=0&adapt_q_fury=0&adapt_q_470=0&adapt_q_480=3&adapt_q_570=0&adapt_q_580=0&adapt_q_vega56=0&adapt_q_vega64=0&adapt_q_750Ti=0&adapt_q_1050Ti=0&adapt_q_10606=0&adapt_q_1070=4&adapt_1070=true&adapt_q_1080=0&adapt_q_1080Ti=0&eth=true&factor%5Beth_hr%5D=120.0&factor%5Beth_p%5D=590.0&grof=true&factor%5Bgro_hr%5D=142.0&factor%5Bgro_p%5D=590.0&x11gf=true&factor%5Bx11g_hr%5D=46.0&factor%5Bx11g_p%5D=590.0&cn=true&factor%5Bcn_hr%5D=2000.0&factor%5Bcn_p%5D=590.0&eq=true&factor%5Beq_hr%5D=1720.0&factor%5Beq_p%5D=590.0&lre=true&factor%5Blrev2_hr%5D=142000.0&factor%5Blrev2_p%5D=590.0&ns=true&factor%5Bns_hr%5D=4000.0&factor%5Bns_p%5D=590.0&lbry=true&factor%5Blbry_hr%5D=1080.0&factor%5Blbry_p%5D=590.0&bk2bf=true&factor%5Bbk2b_hr%5D=6400.0&factor%5Bbk2b_p%5D=590.0&bk14=true&factor%5Bbk14_hr%5D=10000.0&factor%5Bbk14_p%5D=590.0&pas=true&factor%5Bpas_hr%5D=3760.0&factor%5Bpas_p%5D=590.0&skh=true&factor%5Bskh_hr%5D=106.0&factor%5Bskh_p%5D=590.0&factor%5Bl2z_hr%5D=420.0&factor%5Bl2z_p%5D=300.0&factor%5Bcost%5D=0.3&sort=Profitability24&volume=0&revenue=24h&factor%5Bexchanges%5D%5B%5D=&factor%5Bexchanges%5D%5B%5D=abucoins&factor%5Bexchanges%5D%5B%5D=bitfinex&factor%5Bexchanges%5D%5B%5D=bittrex&factor%5Bexchanges%5D%5B%5D=bleutrade&factor%5Bexchanges%5D%5B%5D=cryptopia&factor%5Bexchanges%5D%5B%5D=hitbtc&factor%5Bexchanges%5D%5B%5D=poloniex&factor%5Bexchanges%5D%5B%5D=yobit&dataset=Main&commit=Calculate"
	var coins Coins
	readJsonFromUrl(url, &coins)


	url = "https://api.coinmarketcap.com/v1/ticker/bitcoin/"
	bitcoin := make([]CoinMarketCapCoin, 0)
	readJsonFromUrl(url, &bitcoin)


	dailyBtcRevenueTable := make(map[string]float64)
	for coinName, coinContent := range coins.Coins {
		dailyBtcRevenueTable[coinName] = convertToFloat64(coinContent.Btc_revenue24)
	}

	fmt.Println("\nDaily BTC revenue")
	sortedDailyBtcRevenue := SortMapByValue(dailyBtcRevenueTable)
	for i := 0; i < len(sortedDailyBtcRevenue); i++ {
		// fmt.Printf("%s = %f\n", sortedDailyBtcRevenue[i].key, sortedDailyBtcRevenue[i].value)
	}

	// Calculate profits for each cryptocurrency
	bitcoinPrice := convertToFloat64(bitcoin[0].Price_USD)
    dailyDollarRevenue := make(map[string]float64)
    fmt.Println("\nDaily $ revenue (BTC price: " + bitcoin[0].Price_USD + ")")
	for i := 0; i < len(sortedDailyBtcRevenue); i++ {
        dailyDollarRevenue[sortedDailyBtcRevenue[i].key] = sortedDailyBtcRevenue[i].value * bitcoinPrice
	}

    sortedDailyDollarRevenue := SortMapByValue(dailyDollarRevenue)

    for i := 0; i < len(sortedDailyDollarRevenue); i++ {
		fmt.Printf("%s = %f\n", sortedDailyDollarRevenue[i].key, sortedDailyDollarRevenue[i].value)
    }



}
