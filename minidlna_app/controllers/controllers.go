package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "raspberry-controller/minidlna_app/services"
)

func MinidlnaStatusService(context *gin.Context) {
    status, err := services.GetStatus()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, status)
}

func StartMinidlnaService(context *gin.Context) {
    message, err := services.StartMinidlna()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, message)
}

func StopMinidlnaService(context *gin.Context) {
    message, err := services.StopMinidlna()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, message)
}