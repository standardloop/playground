package main

import (
	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
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
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/metrics"),
		gin.Recovery(),
		cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000", "http://localhost:80", "http://ui.local:80"},
			AllowMethods:     []string{"GET", "PUT", "PATCH"},
			AllowHeaders:     []string{"Accepts"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "http://localhost"
			},
			MaxAge: 12 * time.Hour,
		}),
	)

	m := ginmetrics.GetMonitor()
	m.SetMetricPath(apiVersion + "metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	m.Use(r)

	r.GET("/api/v1/health/", healthCheck)
	r.GET("/api/v1/rand/", randomNumber)
	r.Run()
}
