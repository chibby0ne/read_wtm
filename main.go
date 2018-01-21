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
		r := reflect.ValueOf(&gpu)
		e := r.Elem()
		for i := 0; i < e.NumField(); i++ {
			castedAlgo, ok := e.Field(i).Interface().(Algorithm)
			checkFatalError(ok)
			castedAlgo.HashRate *= float64(totalGPUsDevices[k])
			castedAlgo.Power *= float64(totalGPUsDevices[k])
			castedAlgoAsValue := reflect.ValueOf(castedAlgo)
			e.Field(i).Set(castedAlgoAsValue)
		}

		// store back the total GPU Characteristics
		partialGPUsCharacteristics[k] = gpu
	}

	// instance GPU that contains the total hashing rate and power for all the GPUS listed in conf.json
	var totalGPUsCharacteristics GPU
	totalReflect := reflect.ValueOf(&totalGPUsCharacteristics)
	totalReflectElem := totalReflect.Elem()
	for _, v := range partialGPUsCharacteristics {
		partialReflect := reflect.ValueOf(v)

		for i := 0; i < totalReflectElem.NumField(); i++ {
			castedPartialAlgo, ok := partialReflect.Field(i).Interface().(Algorithm)
			checkFatalError(err)
			castedTotalAlgo, ok := totalReflectElem.Field(i).Interface().(Algorithm)
			checkFatalError(err)

			castedTotalAlgo.HashRate += castedPartialAlgo.HashRate
			castedTotalAlgo.Power += castedPartialAlgo.Power
			castedTotalAlgoAsValue := reflect.ValueOf(castedTotalAlgo)
			totalReflectElem.Field(i).Set(castedTotalAlgoAsValue)
		}
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
	// and loop forever
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

// Returns the most profitable script filename
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
