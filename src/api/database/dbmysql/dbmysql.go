package dbmysql

import (
	"api/config"
	"api/models"
	"fmt"
	"math/rand"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var globalID uint = 0

func GetDB() *gorm.DB {
	return gormDB
}

func dbSeed() {
	log.Debug().Msg("seed MYSQL")

	if gormDB == nil {
		log.Fatal().Msg("MYSQL gormDB is nil?")
	}

	gormDB.Migrator().CreateTable(&models.RandNum{})

	for i := 1; i < 100; i++ {
		globalID += 1
		randNum := &models.RandNum{
			ID:      globalID,
			RandNum: rand.Intn(100 - 0),
		}
		gormDB.Create(randNum)
	}
}

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.Env.MySQLUser, config.Env.MySQLPass, config.Env.MySQLHost, config.Env.MySQLPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("initial init fail")
	}
	dbc := db.Exec("SET global general_log = 1;")
	if dbc.Error != nil {
		log.Fatal().Msg("set log fail")
	}

	dbc = db.Exec("CREATE DATABASE IF NOT EXISTS playground")
	if dbc.Error != nil {
		log.Fatal().Msg("create db fail")
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Env.MySQLUser, config.Env.MySQLPass,
		config.Env.MySQLHost, config.Env.MySQLPort, config.Env.MySQLDBName)
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Msg("full connect fail")
	}
	dbSeed()
}
