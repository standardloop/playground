package v1

import (
	"api-std/config"
	"fmt"
	"net/http"
)

func EnvHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := config.Env.JSON()
	if response == "" {
		http.Error(w, "env is empty!?", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, response)
}
