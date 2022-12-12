package dbmysql

import (
	"api/src/util"
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
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

	// secret management later
	host := util.GetEnv("MYSQL_HOST", "localhost")
	port := util.GetEnv("MYSQL_PORT", "3306")
	user := util.GetEnv("MYSQL_USER", "root")
	password := util.GetEnv("MYSQL_PASSWORD", "mypassword")
	dbname := util.GetEnv("MYSQL_DBNAME", "playground")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, password, host, port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("initial init fail")
	}
	dbc := db.Exec("SET global general_log = 1;")
	if dbc.Error != nil {
		log.Fatal("set log fail")
	}

	dbc = db.Exec("CREATE DATABASE IF NOT EXISTS playground")
	if dbc.Error != nil {
		log.Fatal("create db fail")
	}

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("full connect fail")
	}
	return db
}
