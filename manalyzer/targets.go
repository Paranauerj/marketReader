package manalyzer

import (
	"github.com/local/api"
)

type Target struct {
	Value float64 `json:"value"`
}

func newTarget(value float64) Target {
	return Target{
		Value: value,
	}
}

func TargetsTop() []Target {
	const daysBeforeCalc = 3
	const dayInverval = 10
	targets := []Target{}
	cont := 0
	oldVal := 0.00

	ema56Values := api.EMALine(56)

	for key, val := range ema56Values {
		if key < daysBeforeCalc {
			targets = append(targets, newTarget(calcTopTarget(val, api.Series.Candles[key].ClosePrice.Float())))
		} else {
			if cont%dayInverval == 0 {
				oldVal = calcTopTarget(val, api.Series.Candles[key-daysBeforeCalc].ClosePrice.Float())
				/*aux := calcTopTarget(val, api.Series.Candles[key-daysBeforeCalc].ClosePrice.Float())
				if !(oldVal > 0.95*aux && oldVal < 1.05*aux) {
					oldVal = aux
				}*/
			}
			targets = append(targets, newTarget(oldVal))
		}
		cont++
	}

	return targets
}

func TargetsBottom() []Target {
	const daysBeforeCalc = 3
	const dayInverval = 10
	targets := []Target{}
	cont := 0
	oldVal := 0.00

	ema56Values := api.EMALine(56)

	for key, val := range ema56Values {
		if key < daysBeforeCalc {
			targets = append(targets, newTarget(calcBottomTarget(val, api.Series.Candles[key].ClosePrice.Float())))
		} else {
			if cont%dayInverval == 0 {
				oldVal = calcBottomTarget(val, api.Series.Candles[key-daysBeforeCalc].ClosePrice.Float())
				/*aux := calcBottomTarget(val, api.Series.Candles[key-daysBeforeCalc].ClosePrice.Float())
				if !(oldVal > 0.95*aux && oldVal < 1.05*aux) {
					oldVal = aux
				}*/
			}
			targets = append(targets, newTarget(oldVal))
		}
		cont++
	}

	return targets
}

func calcTopTarget(emaVal, dayVal float64) float64 {
	response := 0.00
	if dayVal >= emaVal*1.1 {
		response = ((dayVal - emaVal) / 2) + dayVal
	} else if dayVal <= emaVal*0.9 {
		response = ((emaVal - dayVal) / 2) + dayVal
	} else if dayVal > emaVal {
		response = ((dayVal - emaVal) * 2) + dayVal
	} else {
		response = ((emaVal - dayVal) * 2) + dayVal
	}

	return response
}

func calcBottomTarget(emaVal, dayVal float64) float64 {
	response := 0.00
	if dayVal >= emaVal*1.1 {
		response = dayVal - ((dayVal - emaVal) / 2)
	} else if dayVal <= emaVal*0.9 {
		response = dayVal - ((emaVal - dayVal) / 2)
	} else if dayVal > emaVal {
		response = dayVal - ((dayVal - emaVal) * 2)
	} else {
		response = dayVal - ((emaVal - dayVal) * 2)
	}

	return response
}

func resistanceByIndex(index, days int) DayValue {
	biggest := DayValue{Value: 0}
	for key, val := range api.Series.Candles {
		if key < index && key >= index-days {
			if val.MaxPrice.Float() > biggest.Value {
				biggest.Value = val.MaxPrice.Float()
				biggest.Time = val.Period.End
			}
		}
	}
	return biggest
}

func supportByIndex(index, days int) DayValue {
	lowest := DayValue{Value: 100000000}
	for key, val := range api.Series.Candles {
		if key < index && key >= index-days {
			if val.MinPrice.Float() < lowest.Value {
				lowest.Value = val.MaxPrice.Float()
				lowest.Time = val.Period.End
			}
		}
	}
	return lowest
}
