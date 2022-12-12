package dbpostgres

import (
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

	// secret management later
	host := util.GetEnv("POSTGRES_HOST", "localhost")
	user := util.GetEnv("POSTGRES_USER", "root")
	password := util.GetEnv("POSTGRES_PASSWORD", "mypassword")
	port := util.GetEnv("POSTGRES_PORT", "5432")
	dbname := util.GetEnv("POSTGRES_DBNAME", "playground")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable TimeZone=America/Denver", host, user, password, port)
	log.Error("WHATTUP SUCKAS")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("postgres initial init rip")
	}

	// do not do this in production
	dbc := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbname))
	if dbc.Error != nil {
		log.Fatal("postgres cleanup db rip")
	}

	dbc = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbname))
	if dbc.Error != nil {
		log.Fatal("postgres create db rip")
	}

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=America/Denver", host, user, password, dbname, port),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("postgres connect to db rip")
	}

	return db
}
