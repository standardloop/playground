package server

import (
	"api-std/config"
	"api-std/server/handlers"
	v1 "api-std/server/handlers/api/v1"
	"api-std/server/handlers/api/v1/health"
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
	apiMux.Handle("GET /health/simple", loggingMiddleware(http.HandlerFunc(health.BasicHealthHandler)))
	apiMux.Handle("GET /health/postgres", loggingMiddleware(http.HandlerFunc(health.PostgresHealthHandler)))
	apiMux.Handle("GET /health/mysql", loggingMiddleware(http.HandlerFunc(health.MYSQLHealthHandler)))
	apiMux.Handle("GET /health/mongo", loggingMiddleware(http.HandlerFunc(health.MongoHealthHandler)))
	apiMux.Handle("GET /health/redis", loggingMiddleware(http.HandlerFunc(health.RedisHealthHandler)))

	// todo
	apiMux.Handle("GET /health/intergrations", loggingMiddleware(http.HandlerFunc(health.BasicHealthHandler)))

	mainMux.Handle(config.ApiVersion+"/", http.StripPrefix(config.ApiVersion, apiMux))
	mainMux.HandleFunc("/", http.HandlerFunc(handlers.GenericNotFoundHandler))

	http.ListenAndServe(fmt.Sprintf(":%s", config.Env.AppPort), mainMux)
}
