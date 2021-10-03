package api

func Backtrack(days int) {
	for i := 1; i < days; i++ {
		Series.Candles = remove(Series.Candles, len(Series.Candles)-1)
	}
	CurrentValue = Series.LastCandle().ClosePrice.Float()
}
