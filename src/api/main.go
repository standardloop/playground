package main

import (
	"api/src/config"
	"api/src/database/dbmysql"
	"api/src/database/dbpostgres"
	"api/src/server"
)

func main() {

	if config.Env.MySQLEnabled {
		dbmysql.DBSeed()
	}
	if config.Env.PostgresEnabled {
		dbpostgres.DBSeed()
	}
	server.Init()
}
