package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"raspberry-controller/postgresql_app/services"
)

func PostgreSQLStatusService(context *gin.Context) {
	status, err := services.GetStatus()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, status)
}

func StartPostgreSQLService(context *gin.Context) {
	message, err := services.StartPostgreSQL()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, message)
}

func StopPostgreSQLService(context *gin.Context) {
	message, err := services.StopPostgreSQL()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, message)
}