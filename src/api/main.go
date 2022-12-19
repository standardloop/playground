package main

import (
	"api/config"
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
		log.Debug().Msg("seeding mysql")
	}
	if config.Env.PostgresEnabled {
		dbpostgres.DBSeed()
		log.Debug().Msg("seeding postgres")
	}
	log.Debug().Msg("initializing server")
	server.Init()
}
