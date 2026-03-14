package server

import (
	"api-std/config"
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

func envHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, config.Env.String())
}

func Init() {
	mux := http.NewServeMux()
	mux.Handle("GET /env", loggingMiddleware(http.HandlerFunc(envHandler)))

	http.ListenAndServe(fmt.Sprintf(":%s", config.Env.AppPort), mux)
}
