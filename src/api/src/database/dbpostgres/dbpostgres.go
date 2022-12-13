package dbpostgres

import (
	"api/src/config"
	"api/src/util"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GormDB = dbInit()
var globalID uint = 0

func DBSeed() {
	GormDB.Migrator().CreateTable(&util.RandNum{})

	for i := 1; i < 100; i++ {
		globalID += 1
		randNum := &util.RandNum{
			ID:      globalID,
			RandNum: rand.Intn(100 - 0),
		}
		GormDB.Create(randNum)
	}
}

func DBHealthCheck(c *gin.Context) {
	// cleanup
	realDB, err := GormDB.DB()
	if err != nil || realDB.Ping() != nil {
		c.JSON(500, gin.H{
			"FAIL": "DB UNPINGABLE",
		})
	}

	c.JSON(200, gin.H{
		"OKAY": "DB HEALTHY",
	})
}

func dbInit() *gorm.DB {

	if !config.Env.PostgresEnabled {
		return nil
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=America/Denver", config.Env.PostgresHost, config.Env.PostgresUser,
		config.Env.PostgresPass, config.Env.PostgresPort)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("postgres initial init rip")
	}

	// do not do this in production
	dbc := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", config.Env.PostgresDBName))
	if dbc.Error != nil {
		log.Fatal("postgres cleanup db rip")
	}

	dbc = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", config.Env.PostgresDBName))
	if dbc.Error != nil {
		log.Fatal("postgres create db rip")
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Denver", config.Env.PostgresHost, config.Env.PostgresUser,
			config.Env.PostgresPass, config.Env.PostgresDBName, config.Env.PostgresPort),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("postgres connect to db rip")
	}

	return db
}
