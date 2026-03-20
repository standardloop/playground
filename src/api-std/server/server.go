package server

import (
	"api-std/config"
	v1 "api-std/server/handlers/api/v1"
	"api-std/server/handlers/health"
	"fmt"
	"log/slog"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(r.URL.Path)
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
		// log.Printf("Finished request: %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func Init() {
	mainMux := http.NewServeMux()
	apiMux := http.NewServeMux()

	apiMux.HandleFunc("/", http.HandlerFunc(v1.NotFoundHandler))
	apiMux.Handle("GET /env", loggingMiddleware(http.HandlerFunc(v1.EnvHandler)))

	mainMux.Handle(config.ApiVersion+"/", http.StripPrefix(config.ApiVersion, apiMux))

	mainMux.Handle("/health", loggingMiddleware(http.HandlerFunc(health.HealthHandler)))

	//mainMux.NotFound = loggingMiddleware(http.HandlerFunc(handlers.NotFoundHandler))

	http.ListenAndServe(fmt.Sprintf(":%s", config.Env.AppPort), mainMux)
}
