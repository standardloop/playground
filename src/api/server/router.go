package server

import (
	"api/config"
	"api/controllers"
	"api/middleware"
	"encoding/json"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

func jsonLoggerMiddleware() gin.HandlerFunc {
	return gin.LoggerWithFormatter(
		func(params gin.LogFormatterParams) string {
			log := make(map[string]interface{})

			log["status_code"] = params.StatusCode
			log["path"] = params.Path
			log["method"] = params.Method
			log["start_time"] = params.TimeStamp.Format("2006/01/02 - 15:04:05")
			log["remote_addr"] = params.ClientIP
			log["response_time"] = params.Latency.String()

			s, _ := json.Marshal(log)
			return string(s) + "\n"
		},
	)
}

func NewRouter() *gin.Engine {
	gin.SetMode(config.Env.GinMode)
	r := gin.New()

	health := new(controllers.HealthController)
	randNum := new(controllers.RandNumController)

	r.Use(
		// gin.LoggerWithWriter(gin.DefaultWriter, fmt.Sprintf("%s/metrics", config.ApiVersion)),
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
		middleware.GinLogger(),
	)
	m := ginmetrics.GetMonitor()
	m.SetMetricPath(config.ApiVersion + "/metrics")
	m.SetSlowTime(10)
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	m.Use(r)

	apiVersion := r.Group(config.ApiVersion)
	{

		apiVersion.GET("/health", health.Status)
		apiVersion.GET("/rand", randNum.RandomNumber)

		if config.Env.MySQLEnabled {
			apiVersion.GET("/randMySQLDB", randNum.RandomNumberFromMySQL)
			apiVersion.GET("/mysql-health", health.MySQLStatus)
		}
		if config.Env.PostgresEnabled {
			apiVersion.GET("/randPostgresDB", randNum.RandomNumberFromPostgres)
			apiVersion.GET("/postgres-health", health.PostgresStatus)
		}
	}
	return r
}
