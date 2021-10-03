package manalyzer

func Fibonacci(top, low float64) map[float64]float64 {
	out := map[float64]float64{
		1:     low,
		0.786: calculateFibo(top, low, 0.786),
		0.618: calculateFibo(top, low, 0.618),
		0.5:   calculateFibo(top, low, 0.5),
		0.382: calculateFibo(top, low, 0.382),
		0.236: calculateFibo(top, low, 0.236),
		0:     top,
	}

	return out
}

func calculateFibo(top, low, number float64) float64 {
	return top - (top-low)*number
}

func PriceOnFibo(top, low, price float64) float64 {
	fib := Fibonacci(top, low)

	if price < fib[0.786] && price >= fib[1] {
		return 1
	}

	if price < fib[0.618] && price >= fib[0.786] {
		return 0.786
	}

	if price < fib[0.5] && price >= fib[0.618] {
		return 0.618
	}

	if price < fib[0.382] && price >= fib[0.5] {
		return 0.5
	}

	if price < fib[0.236] && price >= fib[0.382] {
		return 0.382
	}

	if price < fib[0] {
		return 0.236
	}

	return -1

}
