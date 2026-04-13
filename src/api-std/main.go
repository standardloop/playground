package main

import (
	"api-std/config"
	"api-std/integrations/mongo"
	"api-std/integrations/mysql"
	"api-std/integrations/postgres"
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

	if config.Env.PostgresEnabled {
		postgres.PostgresPoolInit()
	}

	if config.Env.MySQLEnabled {
		mysql.MySQLPoolInit()
	}

	if config.Env.MongoEnabled {
		mongo.MongoPoolInit()
	}

	server.Init()
}
