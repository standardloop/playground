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
		slog.Info(r.URL.Path)
	})
}

func Init() {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

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

	mainMux.Handle(config.ApiVersion+"/", http.StripPrefix(config.ApiVersion, apiMux))
	mainMux.HandleFunc("GET /crash", http.HandlerFunc(handlers.CrashHandler))
	mainMux.HandleFunc("GET /crash/{code}", http.HandlerFunc(handlers.CrashCodeHandler))
	mainMux.HandleFunc("GET /status/{code}", http.HandlerFunc(handlers.StatusHandler))
	mainMux.HandleFunc("/", http.HandlerFunc(handlers.GenericNotFoundHandler))

	http.ListenAndServe(fmt.Sprintf(":%s", config.Env.AppPort), mainMux)
}
