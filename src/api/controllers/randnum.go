package controllers

import (
	"api/database/dbmongo"
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
	dbmysql.GormDB.First(&randNumModel, "id = ?", strconv.Itoa(rand.Intn(99-0)))

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

// wip
func (r RandNumController) RandomNumberFromMongo(c *gin.Context) {

	randNumList, err := dbmongo.GetOne()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"randomNumberFromDB": "NULL",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"randomNumberFromDB": randNumList[0].RandNum,
	})
}
