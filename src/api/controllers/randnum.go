package controllers

import (
	"api/database/dbmysql"
	"api/database/dbpostgres"
	"api/models"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RandNumController struct{}

var randNumModel = new(models.RandNum)

func (r RandNumController) RandomNumber(c *gin.Context) {
	c.JSON(200, gin.H{
		"randomNumber": rand.Intn(10 - 0),
	})
}

func (r RandNumController) RandomNumberFromMySQL(c *gin.Context) {
	randID := strconv.Itoa(rand.Intn(99 - 0))
	dbmysql.GormDB.First(&randNumModel, "id = ?", randID)

	c.JSON(http.StatusOK, gin.H{
		"randomNumberFromDB": randNumModel.RandNum,
	})
}

func (r RandNumController) RandomNumberFromPostgres(c *gin.Context) {
	randID := strconv.Itoa(rand.Intn(99 - 0))
	dbpostgres.GormDB.First(&randNumModel, "id = ?", randID)

	c.JSON(http.StatusOK, gin.H{
		"randomNumberFromDB": randNumModel.RandNum,
	})
}
