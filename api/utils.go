package api

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/sdcoffey/techan"
)

func round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func reverseArray(arr []float64) []float64 {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func float64ToString(number float64) string {
	return fmt.Sprintf("%f", number)
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func remove(slice []*techan.Candle, s int) []*techan.Candle {
	return append(slice[:s], slice[s+1:]...)
}

func strToFloat64(s string) float64 {
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	} else {
		panic(err)
	}
}
