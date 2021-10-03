package manalyzer

import "math"

func Support(values []float64, days int) float64 {
	support := math.MaxFloat64
	for i := 1; i <= days; i++ {
		if values[len(values)-i] < support {
			support = values[len(values)-i]
		}
	}

	return support
}

func Resistance(values []float64, days int) float64 {
	resistance := -1.0
	for i := 1; i <= days; i++ {
		if values[len(values)-i] > resistance {
			resistance = values[len(values)-i]
		}
	}

	return resistance
}
