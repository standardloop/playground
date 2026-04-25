package v1

import (
	"net/http"
	"strconv"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	codeStr := r.PathValue("code")

	code, err := strconv.Atoi(codeStr)
	if err != nil || code < 100 || code > 599 {
		http.Error(w, "Invalid status code", http.StatusBadRequest)
		return
	}
	w.WriteHeader(code)
}
