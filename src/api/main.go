package main

import (
	"api/config"
	"api/database/dbmysql"
	"api/database/dbpostgres"
	"api/logging"
	"api/server"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logging.Trace(logrus.Fields{
		"foo": "bar",
	}, "hi")
	os.Exit(1)
	if config.Env.MySQLEnabled {
		dbmysql.DBSeed()
		//logging.Trace("seeding mysql")
	}
	if config.Env.PostgresEnabled {
		dbpostgres.DBSeed()
		//logging.Trace("seeding postgres")
	}
	//logging.Trace("init server")
	server.Init()
}
