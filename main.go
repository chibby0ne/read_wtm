package main

import (
	// "encoding/json"
	// "flag"
	"fmt"
	// "io/ioutil"
	// "log"
	// "strconv"
	// "bytes"
	// "strings"
	// "reflect"
)

// array that contains names of gpus
var GPU_Names [NUMBER_OF_GPUS]string

// instance of json representation of config file (named conf.json by default)
var config ConfigFileJson

// func init() {
//     // initialize array that contains names of gpus
// 	GPU_Names[0] = "GPU280x"
// 	GPU_Names[1] = "GPU380"
// 	GPU_Names[2] = "GPUFury"
// 	GPU_Names[3] = "GPU470"
// 	GPU_Names[4] = "GPU480"
// 	GPU_Names[5] = "GPU570"
// 	GPU_Names[6] = "GPU580"
// 	GPU_Names[7] = "GPUVega56"
// 	GPU_Names[8] = "GPUVega64"
// 	GPU_Names[9] = "GPU750Ti"
// 	GPU_Names[10] = "GPU1050Ti"
// 	GPU_Names[11] = "GPU1060"
// 	GPU_Names[12] = "GPU1070"
// 	GPU_Names[13] = "GPU1080"
// 	GPU_Names[14] = "GPU1080Ti"

//     // initilize array that contains GPU characteristics in an order corresponding the name written in the GPU_Names array
// 	var GPU_HashRates [NUMBER_OF_GPUS]GPU
// 	GPU_HashRates[0] = GPU280x
// 	GPU_HashRates[1] = GPU380
// 	GPU_HashRates[2] = GPUFury
// 	GPU_HashRates[3] = GPU470
// 	GPU_HashRates[4] = GPU480
// 	GPU_HashRates[5] = GPU570
// 	GPU_HashRates[6] = GPU580
// 	GPU_HashRates[7] = GPUVega56
// 	GPU_HashRates[8] = GPUVega64
// 	GPU_HashRates[9] = GPU750Ti
// 	GPU_HashRates[10] = GPU1050Ti
// 	GPU_HashRates[11] = GPU1060
// 	GPU_HashRates[12] = GPU1070
// 	GPU_HashRates[13] = GPU1080
// 	GPU_HashRates[14] = GPU1080Ti

//     // initialize map that contains GPUs characteristics scraped from whattomine
// 	GPUs = make(map[string]GPU)
// 	for i := range GPU_Names {
// 		GPUs[GPU_Names[i]] = GPU_HashRates[i]
// 	}

//     // read config file passed as command line argument or conf.json if none was passed
//     configFilePathPtr := flag.String("config", "./conf.json", "Config file with mining rig specs")
//     flag.Parse()
// 	readConfig(*configFilePathPtr, &config)

//     // store relevant GPUs (the ones that are being used)
//     r := reflect.ValueOf(config.GPU)
//     totalGPUsDevices := make(map[string]uint64)
//     for k, _ := range GPUs {
//         numOfGPUs := r.FieldByName(k)
//         totalGPUsDevices[k] = numOfGPUs.Uint()
//     }

//     // Calculate
//     totalGPUsCharacteristics := make(map[string]GPU)

//     for k, v := range GPUs {
//         s := reflect.ValueOf(&v).Elem()
//         // totalGPUsCharacteristics[k] = v
//         totalV := totalGPUsCharacteristics[k]
//         totalS := reflect.ValueOf(&totalV).Elem()
//         for i := 0; i < s.NumField(); i++ {
//             alg := s.Field(i)
//             algInterface := alg
//             HashRate := algInterface.FieldByName("HashRate").Float()
//             Power := algInterface.FieldByName("Power").Float()

//             totalHashRate := HashRate * float64(totalGPUsDevices[k])
//             totalPower := Power * float64(totalGPUsDevices[k])

//             // newAlg := new(Algorithm)
//             // newAlgValue := reflect.ValueOf(newAlg)
//             // newAlg := Algorithm{totalHashRate, totalPower}

//             // fmt.Printf("k = %v, v = %v\n", k, v)
//             // fmt.Printf("%v\n", totalS.Field(i))
//             // fmt.Printf("%v\n", s.Field(i))

//             // fmt.Printf("%v\n", totalS.Field(i).FieldByName("HashRate"))
//             totalS.Field(i).FieldByName("HashRate").SetFloat(totalHashRate)
//             totalS.Field(i).FieldByName("Power").SetFloat(totalPower)
//             // totalGPUsCharacteristics[k] = HashRate float64(totalGPUsDevices[k])
//         }
//     }
//     fmt.Printf("\n%v\n", totalGPUsCharacteristics)
// }

/*

func writeOneParameterQuery(buffer *bytes.Buffer,  s, t string) {
   buffer.WriteString(s)
   buffer.WriteString(t)
   buffer.WriteString("&")
   // Add adapt_MODEL=true& whenever there's > 0 cards for that model
   if t != "0" {
       parts := strings.Split(s, "_")
       adapt_true := parts[0] + "_" + parts[2] + "=true&"
       buffer.WriteString(adapt_true)
   }
}

var activeGPUs []GPU

func constructUrlQuery(config ConfigFileJson) {
    var buffer bytes.Buffer
    buffer.WriteString("https://whattomine.com/coins.json?utf8=%E2%9C%93&")
	writeOneParameterQuery(&buffer, "adapt_q_280x=", strconv.FormatUint(config.GPU.GPU280x, 10))
	writeOneParameterQuery(&buffer, "adapt_q_380=", strconv.FormatUint(config.GPU.GPU380, 10))
	writeOneParameterQuery(&buffer, "adapt_q_fury=", strconv.FormatUint(config.GPU.GPUFury, 10))
	writeOneParameterQuery(&buffer, "adapt_q_470=", strconv.FormatUint(config.GPU.GPU470, 10))
	writeOneParameterQuery(&buffer, "adapt_q_480=", strconv.FormatUint(config.GPU.GPU480, 10))
	writeOneParameterQuery(&buffer, "adapt_q_570=", strconv.FormatUint(config.GPU.GPU570, 10))
	writeOneParameterQuery(&buffer, "adapt_q_580=", strconv.FormatUint(config.GPU.GPU580, 10))
	writeOneParameterQuery(&buffer, "adapt_q_vega56=", strconv.FormatUint(config.GPU.GPUVega56, 10))
	writeOneParameterQuery(&buffer, "adapt_q_vega64=", strconv.FormatUint(config.GPU.GPUVega64, 10))
	writeOneParameterQuery(&buffer, "adapt_q_750Ti=", strconv.FormatUint(config.GPU.GPU750Ti, 10))
	writeOneParameterQuery(&buffer, "adapt_q_1050Ti=", strconv.FormatUint(config.GPU.GPU1050Ti, 10))
    // this "10606" seems like a typo but that's the way the parameter is written
	writeOneParameterQuery(&buffer, "adapt_q_10606=", strconv.FormatUint(config.GPU.GPU1060, 10))
	writeOneParameterQuery(&buffer, "adapt_q_1070=", strconv.FormatUint(config.GPU.GPU1070, 10))
	writeOneParameterQuery(&buffer, "adapt_q_1080=", strconv.FormatUint(config.GPU.GPU1080, 10))
	writeOneParameterQuery(&buffer, "adapt_q_1080Ti=", strconv.FormatUint(config.GPU.GPU1080Ti, 10))
	// adapt_1070=true&

    buffer.WriteString("eth=true&")
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
    buffer.WriteString("factor%5Bcost%5D=")
    buffer.WriteString(strconv.FormatFloat(config.CostPerKw, 'f', 1, 64))
    buffer.WriteString("&")
    buffer.WriteString("sort=Profitability24&volume=0&revenue=24h&factor%5Bexchanges%5D%5B%5D=&factor%5Bexchanges%5D%5B%5D=abucoins&factor%5Bexchanges%5D%5B%5D=bitfinex&factor%5Bexchanges%5D%5B%5D=bittrex&factor%5Bexchanges%5D%5B%5D=bleutrade&factor%5Bexchanges%5D%5B%5D=cryptopia&factor%5Bexchanges%5D%5B%5D=hitbtc&factor%5Bexchanges%5D%5B%5D=poloniex&factor%5Bexchanges%5D%5B%5D=yobit&dataset=Main&commit=Calculate")
    fmt.Println("Created correct url")
}

*/

func main() {
	// read current values from www.whattomine.com
	url := "https://whattomine.com/coins.json"
	var coins Coins
	readJsonFromUrl(url, &coins)

	// read current value of bitcoin
	url = "https://api.coinmarketcap.com/v1/ticker/bitcoin/"
	bitcoin := make([]CoinMarketCapCoin, 0)
	readJsonFromUrl(url, &bitcoin)

	// Create map 'coinName' -> USD revenue 24 hr
	dailyDollarRevenue := make(map[string]float64)
	// Convert bitcoin price to float64
	bitcoinPrice := convertToFloat64(bitcoin[0].Price_USD)
	for coinName, coinContent := range coins.Coins {
		dailyDollarRevenue[coinName] = convertToFloat64(coinContent.Btc_revenue24) * bitcoinPrice
	}

	sortedDailyDollarRevenue := SortMapByValue(dailyDollarRevenue)

	fmt.Println("\nDaily $ revenue (BTC price: " + bitcoin[0].Price_USD + ")")
	for i := 0; i < len(sortedDailyDollarRevenue); i++ {
		fmt.Printf("%s = %f\n", sortedDailyDollarRevenue[i].key, sortedDailyDollarRevenue[i].value)
	}

}
