package api

import (
	"github.com/sdcoffey/techan"
)

func EMA(days int) float64 {

	closePrices := techan.NewClosePriceIndicator(Series)
	movingAverage := techan.NewEMAIndicator(closePrices, days) // Create an exponential moving average with a window of 56

	return movingAverage.Calculate(Series.LastIndex()).Float()

}

func EMALine(days int) []float64 {

	var emas []float64
	closePrices := techan.NewClosePriceIndicator(Series)
	movingAverage := techan.NewEMAIndicator(closePrices, days) // Create an exponential moving average with a window of 56

	for i := 0; i < len(Series.Candles); i++ {
		emas = append(emas, movingAverage.Calculate(i).Float())
	}

	return emas

}
