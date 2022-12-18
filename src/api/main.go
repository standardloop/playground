package main

import (
	"fmt"
	"strconv"

	"math/rand"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	"api/src/config"
	"api/src/database/dbmysql"
	"api/src/database/dbpostgres"
	"api/src/util"
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

func randomNumberFromMySQLDB(c *gin.Context) {
	//test := database.GormDB.Exec("SELECT rand_num FROM rand_nums ORDER BY RAND() LIMIT 1;")
	var randNum util.RandNum
	randID := strconv.Itoa(rand.Intn(99 - 0))
	dbmysql.GormDB.First(&randNum, "id = ?", randID)

	c.JSON(200, gin.H{
		"randomNumberFromDB": randNum.RandNum,
	})
}

func randomNumberFromPostgresDB(c *gin.Context) {
	var randNum util.RandNum
	randID := strconv.Itoa(rand.Intn(99 - 0))
	dbpostgres.GormDB.First(&randNum, "id = ?", randID)

	c.JSON(200, gin.H{
		"randomNumberFromDB": randNum.RandNum,
	})
}

func main() {

	if config.Env.MySQLEnabled {
		dbmysql.DBSeed()
	}
	if config.Env.PostgresEnabled {
		dbpostgres.DBSeed()
	}

	gin.SetMode(config.Env.GinMode)
	r := gin.New()
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, fmt.Sprintf("%s/metrics", config.ApiVersion)),
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
	m.SetMetricPath(config.ApiVersion + "/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	m.Use(r)

	r.GET(fmt.Sprintf("%s/health", config.ApiVersion), healthCheck)
	r.GET(fmt.Sprintf("%s/rand", config.ApiVersion), randomNumber)

	if config.Env.MySQLEnabled {
		r.GET(fmt.Sprintf("%s/randMySQLDB", config.ApiVersion), randomNumberFromMySQLDB)
		r.GET(fmt.Sprintf("%s/mysql-health", config.ApiVersion), dbmysql.DBHealthCheck)
	}
	if config.Env.PostgresEnabled {
		r.GET(fmt.Sprintf("%s/randPostgresDB", config.ApiVersion), randomNumberFromPostgresDB)
		r.GET(fmt.Sprintf("%s/postgres-health", config.ApiVersion), dbpostgres.DBHealthCheck)
	}

	r.Run(config.Env.AppPort)
}
