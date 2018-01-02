package main

import (
	"encoding/json"
	"flag"
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

/* Any algorithm has a specific hash rate and a specific power consumption*/
type Algorithm struct {
	HashRate float64
	Power    float64
}

/* Each GPU has a specifc Algorithm performance */
type GPU struct {
	Ethash      Algorithm
	Groestl     Algorithm
	X11Gost     Algorithm
	CryptoNight Algorithm
	Equihash    Algorithm
	Lyra2REv2   Algorithm
	NeoScrypt   Algorithm
	LBRY        Algorithm
	Blake2b     Algorithm
	Blake14r    Algorithm
	Pascal      Algorithm
	Skunkhash   Algorithm
}

/* Data scraped from www.whattomine.com  01.01.2018 */
var GPU280x GPU = GPU{Algorithm{11, 220}, Algorithm{23.8, 250}, Algorithm{2.9, 200}, Algorithm{490, 220}, Algorithm{290, 230}, Algorithm{14050, 220}, Algorithm{490, 250}, Algorithm{60, 200}, Algorithm{960, 250}, Algorithm{1450, 220}, Algorithm{580, 250}, Algorithm{0.0, 0.0}}
var GPU380 GPU = GPU{Algorithm{20.2, 145}, Algorithm{15.5, 130}, Algorithm{2.5, 120}, Algorithm{530, 120}, Algorithm{205, 130}, Algorithm{6400, 125}, Algorithm{350, 145}, Algorithm{44, 135}, Algorithm{760, 150}, Algorithm{1140, 155}, Algorithm{480, 145}, Algorithm{9, 120}}
var GPUFury GPU = GPU{Algorithm{28.2, 180}, Algorithm{17.4, 180}, Algorithm{4.5, 140}, Algorithm{800, 120}, Algorithm{455, 200}, Algorithm{14200, 190}, Algorithm{500, 160}, Algorithm{83, 200}, Algorithm{1400, 260}, Algorithm{1900, 270}, Algorithm{950, 270}, Algorithm{0.0, 0.0}}
var GPU470 GPU = GPU{Algorithm{26, 120}, Algorithm{14.5, 120}, Algorithm{5.3, 125}, Algorithm{660, 100}, Algorithm{260, 110}, Algorithm{4400, 120}, Algorithm{600, 140}, Algorithm{80, 120}, Algorithm{800, 120}, Algorithm{1100, 120}, Algorithm{510, 120}, Algorithm{15.0, 105}}
var GPU480 GPU = GPU{Algorithm{29.5, 135}, Algorithm{18, 130}, Algorithm{6.7, 140}, Algorithm{730, 110}, Algorithm{290, 120}, Algorithm{4900, 130}, Algorithm{650, 150}, Algorithm{95, 140}, Algorithm{990, 150}, Algorithm{1400, 150}, Algorithm{690, 135}, Algorithm{18, 115}}
var GPU570 GPU = GPU{Algorithm{27.9, 120}, Algorithm{15.5, 110}, Algorithm{5.6, 110}, Algorithm{700, 110}, Algorithm{260, 110}, Algorithm{5500, 110}, Algorithm{630, 140}, Algorithm{115, 115}, Algorithm{840, 115}, Algorithm{1140, 115}, Algorithm{580, 135}, Algorithm{16.3, 110}}
var GPU580 GPU = GPU{Algorithm{30.2, 135}, Algorithm{18.5, 135}, Algorithm{6.9, 110}, Algorithm{690, 115}, Algorithm{290, 120}, Algorithm{5700, 120}, Algorithm{650, 150}, Algorithm{135, 145}, Algorithm{990, 150}, Algorithm{1350, 130}, Algorithm{690, 145}, Algorithm{18.5, 115}}
var GPUVega56 GPU = GPU{Algorithm{36.5, 210}, Algorithm{38, 190}, Algorithm{10.5, 230}, Algorithm{1850, 190}, Algorithm{440, 190}, Algorithm{13000, 190}, Algorithm{290, 160}, Algorithm{260, 210}, Algorithm{1900, 230}, Algorithm{2600, 210}, Algorithm{1350, 230}, Algorithm{36, 210}}
var GPUVega64 GPU = GPU{Algorithm{40, 230}, Algorithm{44, 200}, Algorithm{12, 250}, Algorithm{1850, 200}, Algorithm{450, 200}, Algorithm{13000, 200}, Algorithm{290, 170}, Algorithm{280, 230}, Algorithm{2200, 250}, Algorithm{2900, 230}, Algorithm{1550, 250}, Algorithm{40, 230}}
var GPU750Ti GPU = GPU{Algorithm{0.5, 45}, Algorithm{8.3, 80}, Algorithm{2.0, 55}, Algorithm{250, 55}, Algorithm{75, 55}, Algorithm{6640, 70}, Algorithm{220, 75}, Algorithm{51, 75}, Algorithm{350, 75}, Algorithm{610, 75}, Algorithm{200, 55}, Algorithm{0.0, 0.0}}
var GPU1050Ti GPU = GPU{Algorithm{13.9, 70}, Algorithm{14.5, 75}, Algorithm{4.9, 75}, Algorithm{300, 50}, Algorithm{180, 75}, Algorithm{14500, 75}, Algorithm{420, 75}, Algorithm{110, 75}, Algorithm{700, 75}, Algorithm{1050, 75}, Algorithm{380, 75}, Algorithm{11.5, 75}}
var GPU1060 GPU = GPU{Algorithm{22.5, 90}, Algorithm{20.5, 90}, Algorithm{7.2, 90}, Algorithm{430, 70}, Algorithm{270, 90}, Algorithm{20300, 90}, Algorithm{500, 90}, Algorithm{170, 90}, Algorithm{990, 80}, Algorithm{1550, 90}, Algorithm{580, 90}, Algorithm{18, 90}}
var GPU1070 GPU = GPU{Algorithm{30, 120}, Algorithm{35.5, 130}, Algorithm{11.5, 120}, Algorithm{500, 100}, Algorithm{430, 120}, Algorithm{35500, 130}, Algorithm{1000, 155}, Algorithm{270, 120}, Algorithm{1600, 120}, Algorithm{2500, 125}, Algorithm{940, 120}, Algorithm{26.5, 120}}
var GPU1080 GPU = GPU{Algorithm{23.3, 140}, Algorithm{44.5, 150}, Algorithm{13.5, 145}, Algorithm{580, 100}, Algorithm{550, 130}, Algorithm{46500, 150}, Algorithm{1060, 150}, Algorithm{360, 150}, Algorithm{2150, 150}, Algorithm{3300, 150}, Algorithm{1250, 150}, Algorithm{36.5, 150}}
var GPU1080Ti GPU = GPU{Algorithm{35, 140}, Algorithm{58.5, 210}, Algorithm{19.5, 170}, Algorithm{830, 140}, Algorithm{685, 190}, Algorithm{64000, 190}, Algorithm{1400, 190}, Algorithm{460, 190}, Algorithm{2800, 190}, Algorithm{4350, 210}, Algorithm{1700, 210}, Algorithm{47.5, 190}}

var GPUs map[string]GPU

func init() {
	var GPU_Names [15]string
	GPU_Names[0] = "280x"
	GPU_Names[1] = "380"
	GPU_Names[2] = "Fury"
	GPU_Names[3] = "470"
	GPU_Names[4] = "480"
	GPU_Names[5] = "570"
	GPU_Names[6] = "580"
	GPU_Names[7] = "Vega56"
	GPU_Names[8] = "Vega64"
	GPU_Names[9] = "750Ti"
	GPU_Names[10] = "1050Ti"
	GPU_Names[11] = "1060"
	GPU_Names[12] = "1070"
	GPU_Names[13] = "1080"
	GPU_Names[14] = "1080Ti"

	var GPU_HashRates [15]GPU
	GPU_HashRates[0] = GPU280x
	GPU_HashRates[1] = GPU380
	GPU_HashRates[2] = GPUFury
	GPU_HashRates[3] = GPU470
	GPU_HashRates[4] = GPU480
	GPU_HashRates[5] = GPU570
	GPU_HashRates[6] = GPU580
	GPU_HashRates[7] = GPUVega56
	GPU_HashRates[8] = GPUVega64
	GPU_HashRates[9] = GPU750Ti
	GPU_HashRates[10] = GPU1050Ti
	GPU_HashRates[11] = GPU1060
	GPU_HashRates[12] = GPU1070
	GPU_HashRates[13] = GPU1080
	GPU_HashRates[14] = GPU1080Ti

	GPUs = make(map[string]GPU)
	for i := range GPU_Names {
		GPUs[GPU_Names[i]] = GPU_HashRates[i]
	}
}

func constructUrlQuery() {
	// "https://whattomine.com/coins.json?utf8=%E2%9C%93&
	// adapt_q_280x=0&
	// adapt_q_380=0&
	// adapt_q_fury=0&
	// adapt_q_470=0&
	// adapt_q_480=3&
	// adapt_q_570=0&
	// adapt_q_580=0&
	// adapt_q_vega56=0&
	// adapt_q_vega64=0&
	// adapt_q_750Ti=0&
	// adapt_q_1050Ti=0&
	// adapt_q_10606=0&
	// adapt_q_1070=4&
	// adapt_1070=true&
	// adapt_q_1080=0&
	// adapt_q_1080Ti=0&
	// eth=true&
	// factor%5Beth_hr%5D=120.0&
	// factor%5Beth_p%5D=590.0&
	// grof=true&
	// factor%5Bgro_hr%5D=142.0&
	// factor%5Bgro_p%5D=590.0&
	// x11gf=true&
	// factor%5Bx11g_hr%5D=46.0&
	// factor%5Bx11g_p%5D=590.0&
	// cn=true&
	// factor%5Bcn_hr%5D=2000.0&
	// factor%5Bcn_p%5D=590.0&
	// eq=true&
	// factor%5Beq_hr%5D=1720.0&
	// factor%5Beq_p%5D=590.0&
	// lre=true&
	// factor%5Blrev2_hr%5D=142000.0&
	// factor%5Blrev2_p%5D=590.0&
	// ns=true&
	// factor%5Bns_hr%5D=4000.0&
	// factor%5Bns_p%5D=590.0&
	// lbry=true&
	// factor%5Blbry_hr%5D=1080.0&
	// factor%5Blbry_p%5D=590.0&
	// bk2bf=true&
	// factor%5Bbk2b_hr%5D=6400.0&
	// factor%5Bbk2b_p%5D=590.0&
	// bk14=true&
	// factor%5Bbk14_hr%5D=10000.0&
	// factor%5Bbk14_p%5D=590.0&
	// pas=true&
	// factor%5Bpas_hr%5D=3760.0&
	// factor%5Bpas_p%5D=590.0&
	// skh=true&
	// factor%5Bskh_hr%5D=106.0&
	// factor%5Bskh_p%5D=590.0&
	// factor%5Bl2z_hr%5D=420.0&
	// factor%5Bl2z_p%5D=300.0&
	// factor%5Bcost%5D=0.3&
	// sort=Profitability24&
	// volume=0&
	// revenue=24h&
	// factor%5Bexchanges%5D%5B%5D=&
	// factor%5Bexchanges%5D%5B%5D=abucoins&
	// factor%5Bexchanges%5D%5B%5D=bitfinex&
	// factor%5Bexchanges%5D%5B%5D=bittrex&
	// factor%5Bexchanges%5D%5B%5D=bleutrade&
	// factor%5Bexchanges%5D%5B%5D=cryptopia&
	// factor%5Bexchanges%5D%5B%5D=hitbtc&
	// factor%5Bexchanges%5D%5B%5D=poloniex&
	// factor%5Bexchanges%5D%5B%5D=yobit&
	// dataset=Main&
	// commit=Calculate"
    fmt.Println("Created correct url")
}

type ConfigFileJson struct {
	GPU       []GPUConfig `json:"gpu"`
	Power     float64     `json:"power_consumption"`
	CostPerKw float64     `json:"cost_per_kw"`
}

type GPUConfig struct {
	GPU280x   uint64 `json:"280x"`
	GPU380    uint64 `json:"380"`
	GPUFury   uint64 `json:"Fury"`
	GPU470    uint64 `json:"470"`
	GPU480    uint64 `json:"480"`
	GPU570    uint64 `json:"570"`
	GPU580    uint64 `json:580"`
	GPUVega56 uint64 `json:"Vega56"`
	GPUVega64 uint64 `json:"Vega64"`
	GPU750Ti  uint64 `json:"750Ti"`
	GPU1050Ti uint64 `json:"1050Ti"`
	GPU1060   uint64 `json:"1060"`
	GPU1070   uint64 `json:"1070"`
	GPU1080   uint64 `json:"1080"`
	GPU1080Ti uint64 `json:"1080Ti"`
}

func readConfig(configFile string) error {
	configFileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	configFilePathPtr := flag.String("config", "./config.json", "Config file with mining rig specs")
	flag.Parse()
	readConfig(*configFilePathPtr)

	// url := "https://whattomine.com/coins.json"
	// url := "https://whattomine.com/coins.json?utf8=%E2%9C%93&adapt_q_280x=0&adapt_q_380=0&adapt_q_fury=0&adapt_q_470=0&adapt_q_480=3&adapt_q_570=0&adapt_q_580=0&adapt_q_vega56=0&adapt_q_vega64=0&adapt_q_750Ti=0&adapt_q_1050Ti=0&adapt_q_10606=0&adapt_q_1070=4&adapt_1070=true&adapt_q_1080=0&adapt_q_1080Ti=0&eth=true&factor%5Beth_hr%5D=120.0&factor%5Beth_p%5D=590.0&grof=true&factor%5Bgro_hr%5D=142.0&factor%5Bgro_p%5D=590.0&x11gf=true&factor%5Bx11g_hr%5D=46.0&factor%5Bx11g_p%5D=590.0&cn=true&factor%5Bcn_hr%5D=2000.0&factor%5Bcn_p%5D=590.0&eq=true&factor%5Beq_hr%5D=1720.0&factor%5Beq_p%5D=590.0&lre=true&factor%5Blrev2_hr%5D=142000.0&factor%5Blrev2_p%5D=590.0&ns=true&factor%5Bns_hr%5D=4000.0&factor%5Bns_p%5D=590.0&lbry=true&factor%5Blbry_hr%5D=1080.0&factor%5Blbry_p%5D=590.0&bk2bf=true&factor%5Bbk2b_hr%5D=6400.0&factor%5Bbk2b_p%5D=590.0&bk14=true&factor%5Bbk14_hr%5D=10000.0&factor%5Bbk14_p%5D=590.0&pas=true&factor%5Bpas_hr%5D=3760.0&factor%5Bpas_p%5D=590.0&skh=true&factor%5Bskh_hr%5D=106.0&factor%5Bskh_p%5D=590.0&factor%5Bl2z_hr%5D=420.0&factor%5Bl2z_p%5D=300.0&factor%5Bcost%5D=0.3&sort=Profitability24&volume=0&revenue=24h&factor%5Bexchanges%5D%5B%5D=&factor%5Bexchanges%5D%5B%5D=abucoins&factor%5Bexchanges%5D%5B%5D=bitfinex&factor%5Bexchanges%5D%5B%5D=bittrex&factor%5Bexchanges%5D%5B%5D=bleutrade&factor%5Bexchanges%5D%5B%5D=cryptopia&factor%5Bexchanges%5D%5B%5D=hitbtc&factor%5Bexchanges%5D%5B%5D=poloniex&factor%5Bexchanges%5D%5B%5D=yobit&dataset=Main&commit=Calculate"
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
