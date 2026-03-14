package main

import (
	"api-std/logging"
	"api-std/server"
	"log/slog"
)

//	func handler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Hello, World!")
//	}

func main() {
	logging.Init()
	slog.Error("Starting up!")

	server.Init()
}
