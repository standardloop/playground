package server

import (
	"api-std/config"
	"api-std/server/handlers"
	v1 "api-std/server/handlers/v1"
	"api-std/server/handlers/v1/health"
	"fmt"
	"log/slog"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		next.ServeHTTP(w, r)
		slog.Info(fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	})
}

func Init() {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	// API
	apiMux.HandleFunc("/", http.HandlerFunc(handlers.GenericNotFoundHandler))
	apiMux.Handle("GET /env", loggingMiddleware(http.HandlerFunc(v1.EnvHandler)))
	apiMux.Handle("GET /healthz", loggingMiddleware(http.HandlerFunc(health.BasicHealthHandler)))

	if config.Env.MySQLEnabled {
		apiMux.Handle("GET /health/mysql", loggingMiddleware(http.HandlerFunc(health.MYSQLHealthHandler)))
	}

	if config.Env.PostgresEnabled {
		apiMux.Handle("GET /health/postgres", loggingMiddleware(http.HandlerFunc(health.PostgresHealthHandler)))
	}

	if config.Env.MongoEnabled {
		apiMux.Handle("GET /health/mongo", loggingMiddleware(http.HandlerFunc(health.MongoHealthHandler)))
	}

	if config.Env.RedisEnabled {
		apiMux.Handle("GET /health/redis", loggingMiddleware(http.HandlerFunc(health.RedisHealthHandler)))
	}

	// if config.Env.ElasticEnabled {
	// 	apiMux.Handle("GET /health/es", loggingMiddleware(http.HandlerFunc(health.ESHealthHandler)))
	// }

	// todo
	apiMux.Handle("GET /health/intergrations", loggingMiddleware(http.HandlerFunc(health.BasicHealthHandler)))

	apiMux.Handle("GET /headers", loggingMiddleware(http.HandlerFunc(v1.HeadersHandler)))
	apiMux.Handle("GET /crash", loggingMiddleware(http.HandlerFunc(v1.CrashHandler)))
	apiMux.Handle("GET /crash/{code}", loggingMiddleware(http.HandlerFunc(v1.CrashCodeHandler)))
	apiMux.Handle("GET /status/{code}", loggingMiddleware(http.HandlerFunc(v1.StatusHandler)))
	apiMux.Handle("GET /delay/{seconds}", loggingMiddleware(http.HandlerFunc(v1.DelayHandler)))
	apiMux.Handle("POST /echo", loggingMiddleware(http.HandlerFunc(v1.EchoHandler)))

	mainMux.Handle(config.ApiVersion+"/", http.StripPrefix(config.ApiVersion, apiMux))

	// GLOBAL
	mainMux.Handle("/", loggingMiddleware(http.HandlerFunc(handlers.GenericNotFoundHandler)))

	http.ListenAndServe(fmt.Sprintf(":%s", config.Env.AppPort), mainMux)
}
