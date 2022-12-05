package main

import (
	_ "fmt"
	"strconv"

	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	"api/src/constants"
	"api/src/database"
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

func randomNumberFromDB(c *gin.Context) {

	//test := database.GormDB.Exec("SELECT rand_num FROM rand_nums ORDER BY RAND() LIMIT 1;")
	var randNum database.RandNum
	randID := strconv.Itoa(rand.Intn(99 - 0))
	database.GormDB.First(&randNum, "id = ?", randID)

	c.JSON(200, gin.H{
		"randomNumberFromDB": randNum.RandNum,
	})
}

func main() {
	database.DBSeed()
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/metrics"),
		gin.Recovery(),
		cors.New(cors.Config{
			//AllowOrigins:     []string{"http://localhost:3000", "http://localhost:80", "http://ui.local:80"},
			AllowOrigins:     []string{"*"},
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
	m.SetMetricPath(constants.ApiVersion + "metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	m.Use(r)

	r.GET("/api/v1/health/", healthCheck)
	r.GET("/api/v1/db-health/", database.DBHealthCheck)
	r.GET("/api/v1/rand/", randomNumber)
	r.GET("/api/v1/randDB/", randomNumberFromDB)

	r.Run()
}
