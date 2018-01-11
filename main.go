package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

// array that contains names of gpus
var GPU_Names [NUMBER_OF_GPUS]string

// instance of json representation of config file (named conf.json by default)
var config ConfigFileJson

// instance GPU that contains the total hashing rate and power for all the GPUS listed in conf.json
var totalGPUsCharacteristics GPU

// directory containing the miners
var minersDirectory string

func init() {
	// initialize array that contains names of gpus
	GPU_Names[0] = "GPU280x"
	GPU_Names[1] = "GPU380"
	GPU_Names[2] = "GPUFury"
	GPU_Names[3] = "GPU470"
	GPU_Names[4] = "GPU480"
	GPU_Names[5] = "GPU570"
	GPU_Names[6] = "GPU580"
	GPU_Names[7] = "GPUVega56"
	GPU_Names[8] = "GPUVega64"
	GPU_Names[9] = "GPU750Ti"
	GPU_Names[10] = "GPU1050Ti"
	GPU_Names[11] = "GPU1060"
	GPU_Names[12] = "GPU1070"
	GPU_Names[13] = "GPU1080"
	GPU_Names[14] = "GPU1080Ti"

	// initilize array that contains GPU characteristics in an order corresponding the name written in the GPU_Names array
	var GPU_HashRates [NUMBER_OF_GPUS]GPU
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

	// initialize map that contains GPUs characteristics scraped from whattomine
	GPUs = make(map[string]GPU)
	for i := range GPU_Names {
		GPUs[GPU_Names[i]] = GPU_HashRates[i]
	}

	// read config file passed as command line argument or conf.json if none was passed
	configFilePathPtr := flag.String("config", "./conf.json", "Config file with mining rig specs")
	flag.Parse()
	readConfig(*configFilePathPtr, &config)
	minersDirectory = config.MinerDirectory

	// store the gpus and quantities used (taken from conf.json)
	r := reflect.ValueOf(config.GPU)
	totalGPUsDevices := make(map[string]uint64)
	for k, _ := range GPUs {
		numOfGPUs := r.FieldByName(k).Uint()
		if numOfGPUs != 0 {
			totalGPUsDevices[k] = numOfGPUs
		}
	}

	// total GPU characteristics
	partialGPUsCharacteristics := make(map[string]GPU)

	for k, _ := range totalGPUsDevices {
		partialGPUsCharacteristics[k] = GPUs[k]
		gpu := partialGPUsCharacteristics[k]

		// Multiply each algorithm explicilty per the number of GPUs
		// Another way of doing it is using reflection to iterate over the fields of the structure

		// Ethash
		gpu.Ethash.HashRate *= float64(totalGPUsDevices[k])
		gpu.Ethash.Power *= float64(totalGPUsDevices[k])

		// Groestl
		gpu.Groestl.HashRate *= float64(totalGPUsDevices[k])
		gpu.Groestl.Power *= float64(totalGPUsDevices[k])

		// X12Gost
		gpu.X11Gost.HashRate *= float64(totalGPUsDevices[k])
		gpu.X11Gost.Power *= float64(totalGPUsDevices[k])

		// CryptoNight
		gpu.CryptoNight.HashRate *= float64(totalGPUsDevices[k])
		gpu.CryptoNight.Power *= float64(totalGPUsDevices[k])

		// Equihash
		gpu.Equihash.HashRate *= float64(totalGPUsDevices[k])
		gpu.Equihash.Power *= float64(totalGPUsDevices[k])

		// Lyra2REv2
		gpu.Lyra2REv2.HashRate *= float64(totalGPUsDevices[k])
		gpu.Lyra2REv2.Power *= float64(totalGPUsDevices[k])

		// NeoScrypt
		gpu.NeoScrypt.HashRate *= float64(totalGPUsDevices[k])
		gpu.NeoScrypt.Power *= float64(totalGPUsDevices[k])

		// LBRY
		gpu.LBRY.HashRate *= float64(totalGPUsDevices[k])
		gpu.LBRY.Power *= float64(totalGPUsDevices[k])

		// Blake2b
		gpu.Blake2b.HashRate *= float64(totalGPUsDevices[k])
		gpu.Blake2b.Power *= float64(totalGPUsDevices[k])

		// Blake14r
		gpu.Blake14r.HashRate *= float64(totalGPUsDevices[k])
		gpu.Blake14r.Power *= float64(totalGPUsDevices[k])

		// Pascal
		gpu.Pascal.HashRate *= float64(totalGPUsDevices[k])
		gpu.Pascal.Power *= float64(totalGPUsDevices[k])

		// Skunkhash
		gpu.Skunkhash.HashRate *= float64(totalGPUsDevices[k])
		gpu.Skunkhash.Power *= float64(totalGPUsDevices[k])

		// store back the total GPU Characteristics
		partialGPUsCharacteristics[k] = gpu
	}

	for _, v := range partialGPUsCharacteristics {

		// Ethash
		totalGPUsCharacteristics.Ethash.HashRate += v.Ethash.HashRate
		totalGPUsCharacteristics.Ethash.Power += v.Ethash.Power

		// Groestl
		totalGPUsCharacteristics.Groestl.HashRate += v.Groestl.HashRate
		totalGPUsCharacteristics.Groestl.Power += v.Groestl.Power

		// X11Gost
		totalGPUsCharacteristics.X11Gost.HashRate += v.X11Gost.HashRate
		totalGPUsCharacteristics.X11Gost.Power += v.X11Gost.Power

		// CryptoNight
		totalGPUsCharacteristics.CryptoNight.HashRate += v.CryptoNight.HashRate
		totalGPUsCharacteristics.CryptoNight.Power += v.CryptoNight.Power

		// Equihash
		totalGPUsCharacteristics.Equihash.HashRate += v.Equihash.HashRate
		totalGPUsCharacteristics.Equihash.Power += v.Equihash.Power

		// Lyra2REv2
		totalGPUsCharacteristics.Lyra2REv2.HashRate += v.Lyra2REv2.HashRate
		totalGPUsCharacteristics.Lyra2REv2.Power += v.Lyra2REv2.Power

		// NeoScrypt
		totalGPUsCharacteristics.NeoScrypt.HashRate += v.NeoScrypt.HashRate
		totalGPUsCharacteristics.NeoScrypt.Power += v.NeoScrypt.Power

		// LBRY
		totalGPUsCharacteristics.LBRY.HashRate += v.LBRY.HashRate
		totalGPUsCharacteristics.LBRY.Power += v.LBRY.Power

		// Blake2b
		totalGPUsCharacteristics.Blake2b.HashRate += v.Blake2b.HashRate
		totalGPUsCharacteristics.Blake2b.Power += v.Blake2b.Power

		// Blake14r
		totalGPUsCharacteristics.Blake14r.HashRate += v.Blake14r.HashRate
		totalGPUsCharacteristics.Blake14r.Power += v.Blake14r.Power

		// Pascal
		totalGPUsCharacteristics.Pascal.HashRate += v.Pascal.HashRate
		totalGPUsCharacteristics.Pascal.Power += v.Pascal.Power

		// Skunkhash
		totalGPUsCharacteristics.Skunkhash.HashRate += v.Skunkhash.HashRate
		totalGPUsCharacteristics.Skunkhash.Power += v.Skunkhash.Power
	}
	fmt.Printf("\ntotalGPU Charactheristics\n%v\n", totalGPUsCharacteristics)

}

func writeOneParameterQuery(buffer *bytes.Buffer, s, t string) {
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

func constructUrlQuery() string {
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

	// Ethash //

	// "eth=true&"
	buffer.WriteString("eth=true&")
	// factor%5Beth_hr%5D=120.0&
	buffer.WriteString("factor%5Beth_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Ethash.HashRate, 'f', -1, 64) + "&")
	// factor%5Beth_p%5D=590.0&
	buffer.WriteString("factor%5Beth_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Ethash.Power, 'f', -1, 64) + "&")

	// Groestl //

	// grof=true&
	buffer.WriteString("grof=true&")
	// factor%5Bgro_hr%5D=142.0&
	buffer.WriteString("factor%5Bgro_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Groestl.HashRate, 'f', -1, 64) + "&")
	// factor%5Bgro_p%5D=590.0&
	buffer.WriteString("factor%5Bgro_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Groestl.Power, 'f', -1, 64) + "&")

	// X11Gost //

	// x11gf=true&
	buffer.WriteString("x11gf=true&")
	// factor%5Bx11g_hr%5D=46.0&
	buffer.WriteString("factor%5Bx11g_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.X11Gost.HashRate, 'f', -1, 64) + "&")
	// factor%5Bx11g_p%5D=590.0&
	buffer.WriteString("factor%5Bx11g_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.X11Gost.Power, 'f', -1, 64) + "&")

	// CryptoNight //

	// cn=true&
	buffer.WriteString("cn=true&")
	// factor%5Bcn_hr%5D=2000.0&
	buffer.WriteString("factor%5Bcn_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.CryptoNight.HashRate, 'f', -1, 64) + "&")
	// factor%5Bcn_p%5D=590.0&
	buffer.WriteString("factor%5Bcn_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.CryptoNight.Power, 'f', -1, 64) + "&")

	// Equihash //

	// eq=true&
	buffer.WriteString("eq=true&")
	// factor%5Beq_hr%5D=1720.0&
	buffer.WriteString("factor%5Bcn_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.CryptoNight.HashRate, 'f', -1, 64) + "&")
	// factor%5Beq_p%5D=590.0&
	buffer.WriteString("factor%5Bcn_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.CryptoNight.Power, 'f', -1, 64) + "&")

	// Lyra2REv2 //

	// lre=true&
	buffer.WriteString("lre=true&")
	// factor%5Blrev2_hr%5D=142000.0&
	buffer.WriteString("factor%5Blrev2_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Lyra2REv2.HashRate, 'f', -1, 64) + "&")
	// factor%5Blrev2_p%5D=590.0&
	buffer.WriteString("factor%5Blrev2_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Lyra2REv2.Power, 'f', -1, 64) + "&")

	// NeoScrypt //

	// ns=true&
	buffer.WriteString("ns=true&")
	// factor%5Bns_hr%5D=4000.0&
	buffer.WriteString("factor%5Bns_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.NeoScrypt.HashRate, 'f', -1, 64) + "&")
	// factor%5Bns_p%5D=590.0&
	buffer.WriteString("factor%5Bns_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.NeoScrypt.Power, 'f', -1, 64) + "&")

	// LBRY //

	// lbry=true&
	buffer.WriteString("lbry=true&")
	// factor%5Blbry_hr%5D=1080.0&
	buffer.WriteString("factor%5Blrev2_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.LBRY.HashRate, 'f', -1, 64) + "&")
	// factor%5Blbry_p%5D=590.0&
	buffer.WriteString("factor%5Blrev2_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.LBRY.Power, 'f', -1, 64) + "&")

	// Blake2b //

	// bk2bf=true&
	buffer.WriteString("bk2bf=true&")
	// factor%5Bbk2b_hr%5D=6400.0&
	buffer.WriteString("factor%5Bbk2b_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Blake2b.HashRate, 'f', -1, 64) + "&")
	// factor%5Bbk2b_p%5D=590.0&
	buffer.WriteString("factor%5Bbk2b_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Blake2b.Power, 'f', -1, 64) + "&")

	// Blake14r //

	// bk14=true&
	buffer.WriteString("bk14=true&")
	// factor%5Bbk14_hr%5D=10000.0&
	buffer.WriteString("factor%5Bbk14_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Blake14r.HashRate, 'f', -1, 64) + "&")
	// factor%5Bbk14_p%5D=590.0&
	buffer.WriteString("factor%5Bbk14_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Blake14r.Power, 'f', -1, 64) + "&")

	// Pascal //

	// pas=true&
	buffer.WriteString("pas=true&")
	// factor%5Bpas_hr%5D=3760.0&
	buffer.WriteString("factor%5Bpas_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Pascal.HashRate, 'f', -1, 64) + "&")
	// factor%5Bpas_p%5D=590.0&
	buffer.WriteString("factor%5Bpas_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Pascal.Power, 'f', -1, 64) + "&")

	// Skunkhash //

	// skh=true&
	buffer.WriteString("skh=true&")
	// factor%5Bskh_hr%5D=106.0&
	buffer.WriteString("factor%5Bskh_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Skunkhash.HashRate, 'f', -1, 64) + "&")
	// factor%5Bskh_p%5D=590.0&
	buffer.WriteString("factor%5Bskh_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Skunkhash.Power, 'f', -1, 64) + "&")
	// factor%5Bl2z_hr%5D=420.0&
	buffer.WriteString("factor%5B12z_hr%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Skunkhash.HashRate, 'f', -1, 64) + "&")
	// factor%5Bl2z_p%5D=300.0&
	buffer.WriteString("factor%5B12z_p%5D=" + strconv.FormatFloat(totalGPUsCharacteristics.Skunkhash.Power, 'f', -1, 64) + "&")

	// Cost and rest of url//

	buffer.WriteString("factor%5Bcost%5D=" + strconv.FormatFloat(config.CostPerKw, 'f', 1, 64) + "&")
	buffer.WriteString("sort=Profitability24&volume=0&revenue=24h&factor%5Bexchanges%5D%5B%5D=&factor%5Bexchanges%5D%5B%5D=abucoins&")
	buffer.WriteString("factor%5Bexchanges%5D%5B%5D=bitfinex&factor%5Bexchanges%5D%5B%5D=bittrex&factor%5Bexchanges%5D%5B%5D=bleutrade&")
	buffer.WriteString("factor%5Bexchanges%5D%5B%5D=cryptopia&factor%5Bexchanges%5D%5B%5D=hitbtc&factor%5Bexchanges%5D%5B%5D=poloniex&")
	buffer.WriteString("factor%5Bexchanges%5D%5B%5D=yobit&dataset=Main&commit=Calculate")
	return buffer.String()
}

func main() {
	// read current values from www.whattomine.com
	url := constructUrlQuery()
	// url := "https://whattomine.com/coins.json"
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
		fmt.Println(coinName)
		dailyDollarRevenue[coinName] = convertToFloat64(coinContent.Btc_revenue24) * bitcoinPrice
	}

	// sort the map into a sorted pairlist
	sortedDailyDollarRevenue := SortMapByValue(dailyDollarRevenue)

	fmt.Println("\nDaily $ revenue (BTC price: " + bitcoin[0].Price_USD + ")")
	for i := 0; i < len(sortedDailyDollarRevenue); i++ {
		fmt.Printf("%s = %f\n", sortedDailyDollarRevenue[i].key, sortedDailyDollarRevenue[i].value)
	}

	minersDirectory = filepath.Clean(minersDirectory)
	files, err := ioutil.ReadDir(minersDirectory)
	for err != nil {
		log.Fatal(err)
	}

	minersScripts := make([]string, 10)
	for i, file := range files {
		fmt.Printf("minersScripts[%v] =  %v\n", i, file.Name())
		minersScripts[i] = file.Name()
	}

	bestRevenue := -1000.0
	var bestCoin string
	for _, minerScriptName := range minersScripts {
		coinName := strings.Split(minerScriptName, ".")[0]
		if dailyDollarRevenue[coinName] > bestRevenue {
			bestCoin = coinName
			bestRevenue = dailyDollarRevenue[coinName]
		}
	}
	fmt.Println("The best coin is: " + bestCoin + " with $ " + strconv.FormatFloat(bestRevenue, 'f', 6, 64) + " of revenue")

}
