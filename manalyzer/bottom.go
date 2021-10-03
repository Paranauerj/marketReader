package manalyzer

func FindBottom(values []float64, bottoms *[]float64, index int) {
	if index >= 2 && index < len(values)-2 {
		if values[index] < values[index-1] && values[index] < values[index-2] && values[index] < values[index+1] && values[index] < values[index+2] {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index == 1 {
		if values[index] < values[index-1] && values[index] < values[index+1] && values[index] < values[index+2] {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index == 0 {
		if values[index] < values[index+1] && values[index] < values[index+2] {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index == len(values)-1 {
		if values[index] < values[index-1] && values[index] < values[index-2] {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index == len(values)-2 {
		if values[index] < values[index-1] && values[index] < values[index-2] && values[index] < values[index+1] {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index < len(values)-1 {
		FindBottom(values, bottoms, index+1)
	}
}
