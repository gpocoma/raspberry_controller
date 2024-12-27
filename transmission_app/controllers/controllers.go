package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "raspberry-controller/transmission_app/services"
)

func TransmissionStatusService(context *gin.Context) {
    status, err := services.GetStatus()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, status)
}

func StartTransmissionService(context *gin.Context) {
    message, err := services.StartTransmission()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, message)
}

func StopTransmissionService(context *gin.Context) {
    message, err := services.StopTransmission()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, message)
}