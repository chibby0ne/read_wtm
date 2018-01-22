package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"
)

const (
	REGEXP = `([G|g]obyte|[E|e]thereum|[T|t]rezarcoin|[Z|z]cash|[Z|z]classic|[Z|z]encash)`
)

func init() {
	// array that contains names of gpus
	var GPU_Names [NUMBER_OF_GPUS]string
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

}

func calculateHashRateAndPowerForRig(totalGPUsDevices map[string]uint64) GPU {

	// total GPU characteristics
	partialGPUsCharacteristics := make(map[string]GPU)
	for k, _ := range totalGPUsDevices {
		gpu := GPUs[k]

		// Multiply each algorithm explicilty per the number of GPUs
		// Another way of doing it is using reflection to iterate over the fields of the structure

		// // Ethash
		gpu.Ethash.HashRate *= float64(totalGPUsDevices[k])
		gpu.Ethash.Power *= float64(totalGPUsDevices[k])

		// // Groestl
		gpu.Groestl.HashRate *= float64(totalGPUsDevices[k])
		gpu.Groestl.Power *= float64(totalGPUsDevices[k])

		// // X12Gost
		gpu.X11Gost.HashRate *= float64(totalGPUsDevices[k])
		gpu.X11Gost.Power *= float64(totalGPUsDevices[k])

		// // CryptoNight
		gpu.CryptoNight.HashRate *= float64(totalGPUsDevices[k])
		gpu.CryptoNight.Power *= float64(totalGPUsDevices[k])

		// // Equihash
		gpu.Equihash.HashRate *= float64(totalGPUsDevices[k])
		gpu.Equihash.Power *= float64(totalGPUsDevices[k])

		// // Lyra2REv2
		gpu.Lyra2REv2.HashRate *= float64(totalGPUsDevices[k])
		gpu.Lyra2REv2.Power *= float64(totalGPUsDevices[k])

		// // NeoScrypt
		gpu.NeoScrypt.HashRate *= float64(totalGPUsDevices[k])
		gpu.NeoScrypt.Power *= float64(totalGPUsDevices[k])

		// // LBRY
		gpu.LBRY.HashRate *= float64(totalGPUsDevices[k])
		gpu.LBRY.Power *= float64(totalGPUsDevices[k])

		// // Blake2b
		gpu.Blake2b.HashRate *= float64(totalGPUsDevices[k])
		gpu.Blake2b.Power *= float64(totalGPUsDevices[k])

		// // Blake14r
		gpu.Blake14r.HashRate *= float64(totalGPUsDevices[k])
		gpu.Blake14r.Power *= float64(totalGPUsDevices[k])

		// // Pascal
		gpu.Pascal.HashRate *= float64(totalGPUsDevices[k])
		gpu.Pascal.Power *= float64(totalGPUsDevices[k])

		// // Skunkhash
		gpu.Skunkhash.HashRate *= float64(totalGPUsDevices[k])
		gpu.Skunkhash.Power *= float64(totalGPUsDevices[k])

		// store back the total GPU Characteristics
		partialGPUsCharacteristics[k] = gpu
	}

	// instance GPU that contains the total hashing rate and power for all the GPUS listed in conf.json
	var totalGPUsCharacteristics GPU
	for _, v := range partialGPUsCharacteristics {

		// // Ethash
		totalGPUsCharacteristics.Ethash.HashRate += v.Ethash.HashRate
		totalGPUsCharacteristics.Ethash.Power += v.Ethash.Power

		// // Groestl
		totalGPUsCharacteristics.Groestl.HashRate += v.Groestl.HashRate
		totalGPUsCharacteristics.Groestl.Power += v.Groestl.Power

		// // X11Gost
		totalGPUsCharacteristics.X11Gost.HashRate += v.X11Gost.HashRate
		totalGPUsCharacteristics.X11Gost.Power += v.X11Gost.Power

		// // CryptoNight
		totalGPUsCharacteristics.CryptoNight.HashRate += v.CryptoNight.HashRate
		totalGPUsCharacteristics.CryptoNight.Power += v.CryptoNight.Power

		// // Equihash
		totalGPUsCharacteristics.Equihash.HashRate += v.Equihash.HashRate
		totalGPUsCharacteristics.Equihash.Power += v.Equihash.Power

		// // Lyra2REv2
		totalGPUsCharacteristics.Lyra2REv2.HashRate += v.Lyra2REv2.HashRate
		totalGPUsCharacteristics.Lyra2REv2.Power += v.Lyra2REv2.Power

		// // NeoScrypt
		totalGPUsCharacteristics.NeoScrypt.HashRate += v.NeoScrypt.HashRate
		totalGPUsCharacteristics.NeoScrypt.Power += v.NeoScrypt.Power

		// // LBRY
		totalGPUsCharacteristics.LBRY.HashRate += v.LBRY.HashRate
		totalGPUsCharacteristics.LBRY.Power += v.LBRY.Power

		// // Blake2b
		totalGPUsCharacteristics.Blake2b.HashRate += v.Blake2b.HashRate
		totalGPUsCharacteristics.Blake2b.Power += v.Blake2b.Power

		// // Blake14r
		totalGPUsCharacteristics.Blake14r.HashRate += v.Blake14r.HashRate
		totalGPUsCharacteristics.Blake14r.Power += v.Blake14r.Power

		// // Pascal
		totalGPUsCharacteristics.Pascal.HashRate += v.Pascal.HashRate
		totalGPUsCharacteristics.Pascal.Power += v.Pascal.Power

		// // Skunkhash
		totalGPUsCharacteristics.Skunkhash.HashRate += v.Skunkhash.HashRate
		totalGPUsCharacteristics.Skunkhash.Power += v.Skunkhash.Power
	}
	return totalGPUsCharacteristics
}

func getNumberOfGPUs(config ConfigFileJson) map[string]uint64 {
	// store the gpus and quantities used (taken from conf.json)
	r := reflect.ValueOf(config.GPU)
	// map of type  [ GPU_Name ]  -> Number of  gpus
	totalGPUsDevices := make(map[string]uint64)
	for k, _ := range GPUs {
		numOfGPUs := r.FieldByName(k).Uint()
		if numOfGPUs != 0 {
			totalGPUsDevices[k] = numOfGPUs
		}
	}
	return totalGPUsDevices
}

func main() {

	// Parse command
	config := parseConfig()

	// get number of GPUs
	numOfGPUs := getNumberOfGPUs(config)

	// Calculate final hashrate and power
	totalGPUsCharacteristics := calculateHashRateAndPowerForRig(numOfGPUs)

	// read current values from www.whattomine.com
	url := constructUrlQuery(config, totalGPUsCharacteristics)

	regexp := compileRegex()
	getMostProfitableCoin(url, regexp, config)

	log.Println("Starting to mine: " + bestCoin)
	cmd := exec.Command(bestCoin)
	cmd.Start()

	// now we need to start that script if it is not started
	// and loop for
	ticker := time.NewTicker(time.Minute * 5)
	// go checkAndRun(ticker, url, bestCoin)
	for t := range ticker.C {
		log.Println("Checking for new best coin at: ", t)
		// checked new bestCoin
		newBestCoin := getMostProfitableCoin(url, regexp, config)
		if bestCoin != newBestCoin {
			// start new bestCoin
			log.Println("Starting to mine: " + bestCoin)
			cmd := exec.Command(bestCoin)
			cmd.Start()
		}
	}
}

// Returns the compiled regexp
func compileRegex() *Regexp {
	re, err := regexp.Compile(REGEXP)
	checkFatalError(err)
	return re
}

// checks for err and return log fatal if any error
func checkFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Returns the most
func getMostProfitableCoin(url string, regexp *Regexp, config ConfigFileJson) string {
	// read json from url
	var coins Coins
	readJsonFromUrl(url, &coins)

	// read current value of bitcoin
	bitcoin := make([]CoinMarketCapCoin, 0)
	readJsonFromUrl(BITCOINURL, &bitcoin)

	// Create map 'coinName' -> USD revenue 24 hr
	dailyDollarRevenue := make(map[string]float64)

	// Convert bitcoin price to float64
	bitcoinPrice := convertToFloat64(bitcoin[0].Price_USD)
	for coinName, coinContent := range coins.Coins {
		dailyDollarRevenue[coinName] = convertToFloat64(coinContent.Btc_revenue24) * bitcoinPrice
	}

	// sort the map into a sorted pairlist
	sortedDailyDollarRevenue := SortMapByValue(dailyDollarRevenue)

	// Print the coins and their revenue
	log.Println("Daily $ revenue (BTC price: " + bitcoin[0].Price_USD + ")")
	for i := 0; i < len(sortedDailyDollarRevenue); i++ {
		log.Println(sortedDailyDollarRevenue[i].key + " = " + sortedDailyDollarRevenue[i].value)
	}

	// Get the miners scripts in the minerDirectory
	minerDirectory := filepath.Clean(config.MinerDirectory)
	files, err := ioutil.ReadDir(minerDirectory)
	checkFatalError(err)

	// Create a map of type: map[coin name] = script name
	minersScripts := make(map[string]string, len(files))
	for _, file := range files {
		if result := regexp.FindString(file.Name()); result != "" {
			minersScripts[strings.ToLower(result)] = file.Name()
		}
	}

	// Select the most profitable coin from the corresponding mining scripts available
	var bestCoin string
	for i := 0; i < len(sortedDailyDollarRevenue); i++ {
		bestCoin = minersScripts[strings.ToLower(sortedDailyDollarRevenue[i].key)]
		if bestCoin != "" {
			log.Println("Most profitable is: " + bestCoin)
			break
		}
	}
	return bestCoin
}
