package api

import (
	"strconv"
	"strings"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

func getOHLC(coin string, exchange string) [][]float64 {

	// url := "https://api.cryptowat.ch/markets/" + exchange + "/" + coin + "/ohlc?periods=86400"

	url := "https://api.binance.com/api/v3/klines?symbol=" + strings.ToUpper(coin) + "&interval=1d&limit=5000&startTime=" + strconv.FormatInt(loadStartTime().Unix(), 10) + "000"

	resp := new(response)
	getJson(url, resp)

	var ret [][]float64

	for _, val := range *resp {
		openTime := val[0].(float64)
		open := strToFloat64(val[1].(string))
		high := strToFloat64(val[2].(string))
		low := strToFloat64(val[3].(string))
		close := strToFloat64(val[4].(string))
		volume := strToFloat64(val[5].(string))
		ret = append(ret, []float64{openTime, open, high, low, close, volume})
	}

	return ret

}

func loadStartTime() time.Time {
	startTime, _ := time.Parse(layoutISO, "2019-01-05")
	today := time.Now()

	for {
		if today.Sub(startTime).Hours()/24 > 1000 {
			startTime = startTime.Add(time.Hour * 24 * 30)
		} else {
			break
		}
	}

	return startTime
}
