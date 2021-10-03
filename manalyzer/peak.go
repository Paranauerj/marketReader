package manalyzer

func FindPeak(values []float64, peaks *[]float64, index int) {
	if index >= 2 && index < len(values)-2 {
		if values[index] > values[index-1] && values[index] > values[index-2] && values[index] > values[index+1] && values[index] > values[index+2] {
			*peaks = append(*peaks, values[index])
		}
	}

	if index == 1 {
		if values[index] > values[index-1] && values[index] > values[index+1] && values[index] > values[index+2] {
			*peaks = append(*peaks, values[index])
		}
	}

	if index == 0 {
		if values[index] > values[index+1] && values[index] > values[index+2] {
			*peaks = append(*peaks, values[index])
		}
	}

	if index == len(values)-1 {
		if values[index] > values[index-1] && values[index] > values[index-2] {
			*peaks = append(*peaks, values[index])
		}
	}

	if index == len(values)-2 {
		if values[index] > values[index-1] && values[index] > values[index-2] && values[index] > values[index+1] {
			*peaks = append(*peaks, values[index])
		}
	}

	if index < len(values)-1 {
		FindPeak(values, peaks, index+1)
	}
}
