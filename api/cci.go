package api

import "github.com/sdcoffey/techan"

func CCI(days int) float64 {

	c := techan.NewCCIIndicator(Series, days)

	return c.Calculate(Series.LastIndex()).Float()

}
