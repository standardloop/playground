package dbmysql

import (
	"api/src/config"
	"api/src/logging"
	"api/src/util"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
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
	if !config.Env.MySQLEnabled {
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", config.Env.MySQLUser, config.Env.MySQLPass, config.Env.MySQLHost, config.Env.MySQLPort)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Fatal("initial init fail")
	}
	dbc := db.Exec("SET global general_log = 1;")
	if dbc.Error != nil {
		logging.Fatal("set log fail")
	}

	dbc = db.Exec("CREATE DATABASE IF NOT EXISTS playground")
	if dbc.Error != nil {
		logging.Fatal("create db fail")
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Env.MySQLUser, config.Env.MySQLPass,
		config.Env.MySQLHost, config.Env.MySQLPort, config.Env.MySQLDBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Fatal("full connect fail")
	}
	return db
}
