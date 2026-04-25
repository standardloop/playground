package handlers

import (
	"net/http"
	"os"
	"strconv"
)

func CrashHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(255)
}

func CrashCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	codeStr := r.PathValue("code")

	code, err := strconv.Atoi(codeStr)
	if err != nil {
		os.Exit(255)
	}
	os.Exit(code)
}
