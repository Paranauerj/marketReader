package api

import "github.com/sdcoffey/techan"

func RSI(days int) float64 {

	closePrices := techan.NewClosePriceIndicator(Series)
	r := techan.NewRelativeStrengthIndexIndicator(closePrices, days)

	return r.Calculate(Series.LastIndex()).Float()

}

func RSILine(days int) []float64 {

	var rsis []float64
	closePrices := techan.NewClosePriceIndicator(Series)
	r := techan.NewRelativeStrengthIndexIndicator(closePrices, days)

	for i := 0; i < len(Series.Candles); i++ {
		rsis = append(rsis, r.Calculate(i).Float())
	}

	return rsis

}
