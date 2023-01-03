package controllers

import (
	"api/database/dbmongo"
	"api/database/dbmysql"
	"api/database/dbpostgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"OKAY": "I'M HEALTHY",
	})
}

func (h HealthController) MySQLStatus(c *gin.Context) {
	realDB, err := dbmysql.GormDB.DB()
	if err != nil || realDB.Ping() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"FAIL": "DB UNPINGABLE",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"OKAY": "DB HEALTHY",
	})
}

func (h HealthController) PostgresStatus(c *gin.Context) {
	realDB, err := dbpostgres.GormDB.DB()
	if err != nil || realDB.Ping() != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"FAIL": "DB UNPINGABLE",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"OKAY": "DB HEALTHY",
	})
}

func (h HealthController) MongoStatus(c *gin.Context) {
	err := dbmongo.HealthCheck()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"FAIL": "DB UNPINGABLE",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"OKAY": "DB HEALTHY",
	})
}
