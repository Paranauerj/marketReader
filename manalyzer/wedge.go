package manalyzer

import (
	"time"

	"github.com/local/api"
)

type Wedge struct {
	Exists       bool       `json:"exists"`
	ValuesTop    []DayValue `json:"top"`
	ValuesBottom []DayValue `json:"bottom"`
	Type         string     `json:"type"`
}

type DayValue struct {
	Time  time.Time `json:"time"`
	Value float64   `json:"value"`
}

func GetWedgeExtremes() Wedge {
	var valuesHigh, valuesLow []DayValue
	var peaks, bottoms []DayValue

	for _, val := range api.Series.Candles {
		valuesHigh = append(valuesHigh, DayValue{
			Time:  val.Period.Start,
			Value: val.MaxPrice.Float(),
		})

		valuesLow = append(valuesLow, DayValue{
			Time:  val.Period.Start,
			Value: val.MinPrice.Float(),
		})
	}

	findWedgePeak(valuesHigh, &peaks, 0)
	findWedgeBottom(valuesLow, &bottoms, 0)

	return calculateWedge(peaks, bottoms)

}

func findWedgePeak(values []DayValue, peaks *[]DayValue, index int) {
	if index >= 2 && index < len(values)-2 {
		if values[index].Value > values[index-1].Value && values[index].Value > values[index-2].Value && values[index].Value > values[index+1].Value && values[index].Value > values[index+2].Value {
			*peaks = append(*peaks, values[index])
		}
	}

	if index == 1 {
		if values[index].Value > values[index-1].Value && values[index].Value > values[index+1].Value && values[index].Value > values[index+2].Value {
			*peaks = append(*peaks, values[index])
		}
	}

	if index == 0 {
		if values[index].Value > values[index+1].Value && values[index].Value > values[index+2].Value {
			*peaks = append(*peaks, values[index])
		}
	}

	if index == len(values)-1 {
		if values[index].Value > values[index-1].Value && values[index].Value > values[index-2].Value {
			*peaks = append(*peaks, values[index])
		}
	}

	if index == len(values)-2 {
		if values[index].Value > values[index-1].Value && values[index].Value > values[index-2].Value && values[index].Value > values[index+1].Value {
			*peaks = append(*peaks, values[index])
		}
	}

	if index < len(values)-1 {
		findWedgePeak(values, peaks, index+1)
	}
}

func findWedgeBottom(values []DayValue, bottoms *[]DayValue, index int) {
	if index >= 2 && index < len(values)-2 {
		if values[index].Value < values[index-1].Value && values[index].Value < values[index-2].Value && values[index].Value < values[index+1].Value && values[index].Value < values[index+2].Value {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index == 1 {
		if values[index].Value < values[index-1].Value && values[index].Value < values[index+1].Value && values[index].Value < values[index+2].Value {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index == 0 {
		if values[index].Value < values[index+1].Value && values[index].Value < values[index+2].Value {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index == len(values)-1 {
		if values[index].Value < values[index-1].Value && values[index].Value < values[index-2].Value {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index == len(values)-2 {
		if values[index].Value < values[index-1].Value && values[index].Value < values[index-2].Value && values[index].Value < values[index+1].Value {
			*bottoms = append(*bottoms, values[index])
		}
	}

	if index < len(values)-1 {
		findWedgeBottom(values, bottoms, index+1)
	}
}

// Testar depois em outros tempos de mercado
func calculateWedge(tops, lows []DayValue) Wedge {
	const numOfTops = 4
	const numOfBottoms = 6

	response := new(Wedge)
	response.Exists = false

	var sumTops, averageTop, sumBottoms, averageBottom float64
	sumTops = 0
	for i := 1; i <= numOfTops; i++ {
		response.ValuesTop = append(response.ValuesTop, tops[len(tops)-i])
		sumTops += tops[len(tops)-i].Value
	}
	averageTop = sumTops / numOfTops

	sumBottoms = 0
	for i := 1; i <= numOfBottoms; i++ {
		response.ValuesBottom = append(response.ValuesBottom, lows[len(lows)-i])
		sumBottoms += lows[len(lows)-i].Value
	}
	averageBottom = sumBottoms / numOfBottoms

	//fmt.Println(averageTop, tops[len(tops)-1], averageBottom, lows[len(lows)-1])

	// Topos subindo e fundos subindo
	if averageTop < tops[len(tops)-1].Value && averageBottom < lows[len(lows)-1].Value {
		// Fundos alcançando topos
		if (tops[len(tops)-1].Value - averageTop) < (lows[len(lows)-1].Value - averageBottom) {
			// Razão entre a aproximação
			if (lows[len(lows)-1].Value-averageBottom)/(tops[len(tops)-1].Value-averageTop) > 0.3 {
				// Cunha Ascendente 1 sendo formada
				response.Exists = true
				response.Type = "ascending"
			}
		}
	}

	// Topos descendo e fundos subindo
	if averageTop > tops[len(tops)-1].Value && averageBottom < lows[len(lows)-1].Value {
		// Fundos alcançando topos
		if (tops[len(tops)-1].Value - averageTop) < (lows[len(lows)-1].Value - averageBottom) {
			if (lows[len(lows)-1].Value-averageBottom)/(tops[len(tops)-1].Value-averageTop) > 0.3 {
				// Cunha Ascendente 2 sendo formada
				response.Exists = true
				response.Type = "ascending"
			}
		} else { // Topos alcançando fundos
			if (tops[len(tops)-1].Value-averageTop)/(lows[len(lows)-1].Value-averageBottom) > 0.3 {
				// Cunha Descendente 2 sendo formada
				response.Exists = true
				response.Type = "descending"
			}
		}
	}

	// Topos descendo e fundos descendo
	if averageTop > tops[len(tops)-1].Value && averageBottom > lows[len(lows)-1].Value {
		// Fundos alcançando topos
		if (tops[len(tops)-1].Value - averageTop) < (lows[len(lows)-1].Value - averageBottom) {
			if (tops[len(tops)-1].Value-averageTop)/(lows[len(lows)-1].Value-averageBottom) > 0.3 {
				// Cunha Descendente 1 sendo formada
				response.Exists = true
				response.Type = "descending"
			}
		}
	}

	return *response

}
