package controllers

import (
	"github.com/local/api"
	"github.com/local/manalyzer"
)

type Averages struct {
	Ema56  []float64 `json:"ema56"`
	Ema200 []float64 `json:"ema200"`
	Ema500 []float64 `json:"ema500"`
}

type RSI struct {
	Rsi []float64 `json:"rsi"`
}

func CalculateSupport(days int) float64 {
	var valuesLow []float64
	candlesLength := len(api.Series.Candles)

	for key, val := range api.Series.Candles {
		if key != candlesLength-1 {
			valuesLow = append(valuesLow, val.MinPrice.Float())
		}
	}

	return manalyzer.Support(valuesLow, days)
}

func CalculateResistance(days int) float64 {
	var valuesHigh []float64
	candlesLength := len(api.Series.Candles)

	for key, val := range api.Series.Candles {
		if key != candlesLength-1 {
			valuesHigh = append(valuesHigh, val.MaxPrice.Float())
		}
	}

	return manalyzer.Resistance(valuesHigh, days)
}
