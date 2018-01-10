package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Json representation of config file
type ConfigFileJson struct {
	GPU       GPUConfig `json:"gpu"`
	CostPerKw float64   `json:"cost_per_kw"`
}

// Each GPU type possible
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

/*
Example configuration file:
{
	"gpu": {
		"1070": 4,
		"1080": 5
	},
	"cost_per_kw": 0.3
}
*/
func readConfig(configFile string, target interface{}) {
	configFileContent, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configFileContent, target)
	if err != nil {
		log.Fatal(err)
	}
}
