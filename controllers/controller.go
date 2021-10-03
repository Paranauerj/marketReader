package controllers

import (
	"errors"
	"strconv"

	"github.com/local/api"
	"github.com/local/manalyzer"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "everything OK",
	})
}

func LoadPair(c *gin.Context) {
	pair := c.Param("pair")
	api.ReadMarket(pair, "coinbase-pro")
	c.JSON(200, api.Series.Candles)
}

func Backtrack(c *gin.Context) {
	days := c.Param("days")
	daysInt, err := strconv.Atoi(days)

	if err != nil {
		c.Error(errors.New("Invalid Input"))
	}

	api.Backtrack(daysInt)
}

func GetCandles(c *gin.Context) {
	c.JSON(200, api.Series.Candles)
}

func GetAverages(c *gin.Context) {
	c.JSON(200, Averages{
		Ema56:  api.EMALine(56),
		Ema200: api.EMALine(200),
		Ema500: api.EMALine(500),
	})
}

func GetSupport(c *gin.Context) {
	days := c.Param("days")
	daysInt, err := strconv.Atoi(days)

	if err != nil {
		c.Error(errors.New("Invalid Input"))
	}

	c.JSON(200, CalculateSupport(daysInt))
}

func GetResistance(c *gin.Context) {
	days := c.Param("days")
	daysInt, err := strconv.Atoi(days)

	if err != nil {
		c.Error(errors.New("Invalid Input"))
	}

	c.JSON(200, CalculateResistance(daysInt))
}

func GetRSI(c *gin.Context) {
	c.JSON(200, RSI{Rsi: api.RSILine(14)})
}

func GetWedge(c *gin.Context) {
	c.JSON(200, manalyzer.GetWedgeExtremes())
}

func GetTrends(c *gin.Context) {
	c.JSON(200, manalyzer.CalculateTrends())
}

func GetTargets(c *gin.Context) {
	c.JSON(200, gin.H{
		"tops": manalyzer.TargetsTop(),
		"lows": manalyzer.TargetsBottom(),
	})
}
