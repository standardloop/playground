package handlers

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func GenericNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	response := errorResponse{
		Message: "The requested resource could not be found.",
		Status:  http.StatusNotFound,
	}
	json.NewEncoder(w).Encode(response)
}
