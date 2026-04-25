package handlers

import (
	"net/http"
	"strconv"
	"time"
)

func DelayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	secondsStr := r.PathValue("seconds")

	seconds, err := strconv.Atoi(secondsStr)
	if err != nil || seconds < 0 {
		http.Error(w, "Invalid seconds amount", http.StatusBadRequest)
		return
	}

	time.Sleep(time.Duration(seconds) * time.Second)
	w.WriteHeader(200)
}
