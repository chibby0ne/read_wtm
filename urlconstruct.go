package main


import (
    "bytes"
    "strings"
    "strconv"
)

const (
    urlBase string = "https://whattomine.com/coins.json?utf8=%E2%9C%93&"
    trueString string = "=true&"
)

// appends the strings adaptString and numberOfGPUs
// additionally appends 
func writeOneParameterQuery(buffer *bytes.Buffer, adaptString, numberOfGPUs string) {
	buffer.WriteString(adaptString)
	buffer.WriteString(numberOfGPUs)
	buffer.WriteString("&")
	// Add adapt_MODEL=true& whenever there's > 0 cards for that model
	if numberOfGPUs != "0" {
		parts := strings.Split(adaptString, "_")
		adapt_true := parts[0] + "_" + parts[2] + trueString
		buffer.WriteString(adapt_true)
	}
}

func constructUrlQuery() string {
	var buffer bytes.Buffer
	buffer.WriteString(urlBase)
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


