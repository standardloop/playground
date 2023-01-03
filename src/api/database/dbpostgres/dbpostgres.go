package dbpostgres

import (
	"api/config"
	"api/models"
	"fmt"
	"math/rand"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var gormDB *gorm.DB
var globalID uint = 0

func GetDB() *gorm.DB {
	return gormDB
}

func dbSeed() {
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

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=America/Denver", config.Env.PostgresHost, config.Env.PostgresUser,
		config.Env.PostgresPass, config.Env.PostgresPort)

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal().Msg("postgres initial init rip")
	}

	// do not do this in production
	dbc := gormDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", config.Env.PostgresDBName))
	if dbc.Error != nil {
		log.Fatal().Msg("postgres cleanup db rip")
	}

	dbc = gormDB.Exec(fmt.Sprintf("CREATE DATABASE %s;", config.Env.PostgresDBName))
	if dbc.Error != nil {
		log.Fatal().Msg("postgres create db rip")
	}

	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Denver", config.Env.PostgresHost, config.Env.PostgresUser,
			config.Env.PostgresPass, config.Env.PostgresDBName, config.Env.PostgresPort),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal().Msg("postgres connect to db rip")
	}
	dbSeed()
}
