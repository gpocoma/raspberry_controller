package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "raspberry-controller/transmission_app/services"
)

func TransmissionStatusService(context *gin.Context) {
    active, err := services.GetStatus()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{"active": active})
}

func StartTransmissionService(context *gin.Context) {
    err := services.StartTransmission()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Transmission started successfully"})
}

func StopTransmissionService(context *gin.Context) {
    err := services.StopTransmission()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Transmission stopped successfully"})
}