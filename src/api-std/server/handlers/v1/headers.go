package v1

import (
	"encoding/json"
	"net/http"
)

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{}
	for name, values := range r.Header {
		for _, value := range values {
			response[name] = value
		}
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}
}
