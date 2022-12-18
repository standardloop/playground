package main

import (
	"api/config"
	"api/database/dbmysql"
	"api/database/dbpostgres"
	"api/server"
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
