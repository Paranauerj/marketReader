package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

type response [][]interface{}

var myClient = &http.Client{Timeout: 10 * time.Second}
var Series = techan.NewTimeSeries()
var COIN string
var CurrentValue float64

/*
	Coin - pair of values. E.g.: btcusd, ethusd
	Exchange - coinbase-pro, bitstamp, okex, kraken, binance
*/
func ReadMarket(coin string, exchange string) {
	myClient = &http.Client{Timeout: 10 * time.Second}
	Series = techan.NewTimeSeries()
	COIN = ""
	CurrentValue = 0

	COIN = coin
	var dataset [][]string
	for _, val := range getOHLC(coin, exchange) {

		// Timestamp, Open, Close, High, Low, volume
		var tmstmp int = int(val[0])

		dataset = append(dataset, []string{
			strconv.Itoa(tmstmp),
			float64ToString(val[1]),
			float64ToString(val[4]),
			float64ToString(val[2]),
			float64ToString(val[3]),
			float64ToString(val[5]),
		})

	}

	for _, datum := range dataset {
		start, _ := strconv.ParseInt(datum[0][:10], 10, 64)
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(datum[1])
		candle.ClosePrice = big.NewFromString(datum[2])
		candle.MaxPrice = big.NewFromString(datum[3])
		candle.MinPrice = big.NewFromString(datum[4])
		candle.Volume = big.NewFromString(datum[5])

		Series.AddCandle(candle)
	}

	CurrentValue = Series.LastCandle().ClosePrice.Float()
}
