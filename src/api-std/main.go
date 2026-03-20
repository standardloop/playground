package main

import (
	"api-std/integrations/postgres"
	"api-std/logging"
	"log/slog"
)

//	func handler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Hello, World!")
//	}

func main() {
	logging.Init()
	slog.Error("Starting up!")

	postgres.JoshTest2()

	//server.Init()
}
