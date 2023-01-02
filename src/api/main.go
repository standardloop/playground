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
		log.Debug().Msg("seeding mysql")
		dbmysql.DBSeed()
	}
	if config.Env.PostgresEnabled {
		log.Debug().Msg("seeding postgres")
		dbpostgres.DBSeed()
	}
	if config.Env.MongoEnabled {
		log.Debug().Msg("seeding mongo")
		dbmongo.DBSeed()
	}

	log.Debug().Msg("initializing server")
	server.Init()
}
