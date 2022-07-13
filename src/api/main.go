package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"OKAY": "I'M HEALTHY",
	})
}

func randomNumber(c *gin.Context) {
	c.JSON(200, gin.H{
		"randomNumber": rand.Intn(10 - 0),
	})
}

const (
	apiVersion = "/api/v1/"
)

func main() {

	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	m := ginmetrics.GetMonitor()
	m.SetMetricPath(apiVersion + "metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	m.Use(r)

	r.GET("/api/v1/health", healthCheck)
	r.Run()
}
