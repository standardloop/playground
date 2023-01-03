package main

import (
	"api/config"
	"api/database/dbmongo"
	"api/database/dbmysql"
	"api/database/dbpostgres"
	"api/logging"
	"api/server"

	"github.com/rs/zerolog/log"
)

func main() {

	logging.Init()
	log.Trace().Msg("Starting main()")
	if config.Env.MySQLEnabled {
		dbmysql.DBSeed()
	}
	if config.Env.PostgresEnabled {
		dbpostgres.DBSeed()
	}
	if config.Env.MongoEnabled {
		dbmongo.Init()
	}

	log.Debug().Msg("initializing server")
	server.Init()
}
