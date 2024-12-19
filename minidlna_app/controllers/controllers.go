package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "raspberry-controller/minidlna_app/services"
)

func MinidlnaStatusService(context *gin.Context) {
    active, err := services.GetStatus()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{"active": active})
}

func StartMinidlnaService(context *gin.Context) {
    err := services.StartMinidlna()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Minidlna started successfully"})
}

func StopMinidlnaService(context *gin.Context) {
    err := services.StopMinidlna()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "Minidlna stopped successfully"})
}