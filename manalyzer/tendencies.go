package manalyzer

import (
	"strings"

	"github.com/local/api"
)

type Tendencies struct {
	ShortTerm []string `json:"short_term"`
	MidTerm   []string `json:"mid_term"`
	LongTerm  []string `json:"long_term"`
}

func CalculateTrends() Tendencies {
	var last10DaysVolume []float64
	response := new(Tendencies)
	rsi := api.RSI(14)
	cci := api.CCI(20)
	ema56 := api.EMA(56)
	ema200 := api.EMA(200)

	// Fundos de ciclo
	if ema56 < ema200 {
		if rsi >= 70 {
			response.LongTerm = append(response.LongTerm, "Fundo com boa recuperação")
		}
		if rsi >= 60 && rsi < 70 {
			response.LongTerm = append(response.LongTerm, "Fundo com moderada recuperação")
		}
		if rsi >= 40 && rsi < 60 {
			response.LongTerm = append(response.LongTerm, "Fundo lateralizado (fundo dos deuses)")
		}
		if rsi > 30 && rsi < 40 {
			response.LongTerm = append(response.LongTerm, "Fundo com moderada tendência de baixa")
		}
		if rsi <= 30 {
			response.LongTerm = append(response.LongTerm, "Fundo com tendência de baixa")
		}
	}

	// Troca de tendências longo prazo
	if ema56 > ema200*0.9 && ema56 < ema200*1.1 {
		changingTrend := "Troca de Tendência"
		if rsi > 60 {
			changingTrend += " para Alta"
		} else if rsi < 40 {
			changingTrend += " para Baixa"
		} else {
			changingTrend += " para Lateral"
		}
		response.LongTerm = append(response.LongTerm, changingTrend)
	}

	// Topos de verão próximos
	if ema56 > ema200*1.26 && rsi > 60 {
		response.LongTerm = append(response.LongTerm, "Alta de Verão - EMA200/EMA56")
	}

	// Médias de 8 do augusto backes
	if api.CurrentValue < 0.9*ema56 && rsi < 60 {
		response.MidTerm = append(response.MidTerm, "Tendência de baixa - Média de 8")
	}

	if api.CurrentValue > 1.15*ema56 && rsi > 40 {
		response.MidTerm = append(response.MidTerm, "Tendência de alta - Média de 8")
	}

	// Técnica cci e rsi
	if cci > 100 && rsi >= 60 {
		response.ShortTerm = append(response.ShortTerm, "Alta próxima - RSI/CCI")
	} else if cci < 0 && rsi <= 40 {
		response.ShortTerm = append(response.ShortTerm, "Baixa próxima - RSI/CCI")
	} else {
		response.ShortTerm = append(response.ShortTerm, "Lateralidade - RSI/CCI")
	}

	// Técnica volume
	for key, val := range api.Series.Candles {
		if key >= api.Series.LastIndex()-10 {
			last10DaysVolume = append(last10DaysVolume, val.Volume.Float())
		}
	}

	if getAvg(last10DaysVolume[:5]) < getAvg(last10DaysVolume[5:10]) && strings.Split(response.ShortTerm[0], " ")[0] == "Lateralidade" {
		response.ShortTerm = append(response.ShortTerm, "Aumento da Acumulação Em Lateralidade")
	}

	return *response
}

func getAvg(arr []float64) float64 {
	sum := 0.00

	for _, val := range arr {
		sum += val
	}
	return sum / float64(len(arr))
}
