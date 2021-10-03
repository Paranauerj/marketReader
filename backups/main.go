package main

import (
	"fmt"

	"github.com/local/api"
	"github.com/local/controllers"
	"github.com/local/manalyzer"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", controllers.Ping)
	r.GET("/load/:pair", controllers.LoadPair)
	r.GET("/get", controllers.GetCandles)
	r.GET("/emas", controllers.GetAverages)
	r.GET("/support/:days", controllers.GetSupport)
	r.GET("/resistance/:days", controllers.GetResistance)
	r.GET("/backtrack/:days", controllers.Backtrack)
	r.GET("/rsi", controllers.GetRSI)
	r.GET("/wedge", controllers.GetWedge)
	r.GET("/trends", controllers.GetTrends)

	r.Static("/frontend", "./frontend")

	r.Run(":8090")
	// Fazer análise de volumes ascendente e descendente

	var ema56, ema200, rsi, cci, support, resistance, priceOnFibo float64
	var valuesHigh, valuesLow, volumes []float64
	var peaks, bottoms []float64

	api.ReadMarket("btcusd", "coinbase-pro")
	// api.Backtrack(0)

	for _, val := range api.Series.Candles {
		valuesHigh = append(valuesHigh, val.MaxPrice.Float())
		valuesLow = append(valuesLow, val.MinPrice.Float())
		volumes = append(volumes, val.Volume.Float())
	}

	manalyzer.FindPeak(valuesHigh, &peaks, 0)
	manalyzer.FindBottom(valuesLow, &bottoms, 0)

	ema56 = api.EMA(56)
	ema200 = api.EMA(200)
	rsi = api.RSI(14)
	cci = api.CCI(20)
	support = manalyzer.Support(valuesLow, 30)
	resistance = manalyzer.Resistance(valuesHigh, 30)
	fibo := manalyzer.Fibonacci(resistance, support)
	priceOnFibo = manalyzer.PriceOnFibo(resistance, support, api.CurrentValue)

	/*fmt.Println(`
	Indicadores:
		Price:	` + strconv.FormatFloat(api.CurrentValue, 'f', 6, 64) + `
		RSI:	` + strconv.FormatFloat(rsi, 'f', 6, 64) + `
		CCI:	` + strconv.FormatFloat(cci, 'f', 6, 64) + `
		ema56:	` + strconv.FormatFloat(ema56, 'f', 6, 64) + `
		ema200: ` + strconv.FormatFloat(ema200, 'f', 6, 64) + `
		30d sp:	` + strconv.FormatFloat(manalyzer.Support(valuesLow, 30), 'f', 6, 64) + `
		30d rs:	` + strconv.FormatFloat(manalyzer.Resistance(valuesHigh, 30), 'f', 6, 64) + `
	`)*/

	fmt.Printf("%#v\n%f\n", fibo, priceOnFibo)

	fmt.Println("Hints:")

	// Gera informação sobre cunhas
	// manalyzer.Wedge(peaks, bottoms)

	// fmt.Println(ema56, ema200)
	// Fundos de ciclo
	if ema56 < ema200 {
		if rsi >= 70 {
			fmt.Println("Fundo com boa recuperação")
		}
		if rsi >= 60 && rsi < 70 {
			fmt.Println("Fundo com moderada recuperação")
		}
		if rsi >= 40 && rsi < 60 {
			fmt.Println("Fundo lateralizado (fundo dos deuses)")
		}
		if rsi > 30 && rsi < 40 {
			fmt.Println("Fundo com moderada tendência de baixa")
		}
		if rsi <= 30 {
			fmt.Println("Fundo com tendência de baixa")
		}
	}

	// Troca de tendências longo prazo
	if ema56 > ema200*0.9 && ema56 < ema200*1.1 {
		fmt.Print("Troca de Tendência")
		if rsi > 60 {
			fmt.Print(" para Alta\n")
		} else if rsi < 40 {
			fmt.Print(" para Baixa\n")
		} else {
			fmt.Print(" para Lateral\n")
		}
	}

	// Topos de verão próximos
	if ema56 > ema200*1.26 && rsi > 60 {
		fmt.Println("Alta de Verão - EMA200/EMA56")
	}

	// Médias de 8 do augusto backes
	if api.CurrentValue < 0.9*ema56 && rsi < 60 {
		fmt.Println("Tendência de baixa - Média de 8")
	}

	if api.CurrentValue > 1.15*ema56 && rsi > 40 {
		fmt.Println("Tendência de alta - Média de 8")
	}

	// Técnica cci e rsi
	if cci > 100 && rsi >= 60 {
		fmt.Println("Alta próxima - RSI/CCI")
	} else if cci < 0 && rsi <= 40 {
		fmt.Println("Baixa próxima - RSI/CCI")
	} else {
		fmt.Println("Lateralidade - RSI/CCI")
	}

}
