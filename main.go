package main

import (
	"github.com/local/controllers"

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
	r.GET("/targets", controllers.GetTargets)

	r.Static("/frontend", "./frontend")

	r.Run()

}
