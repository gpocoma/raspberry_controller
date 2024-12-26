package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "raspberry-controller/system_monitor/services"
)

func TestService(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "message": "Service is running",
    })
}

func RamUsageService(context *gin.Context) {
    data, err := services.GetRamUsage()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, data)
}

func CpuUsageService(context *gin.Context) {
    data, err := services.GetCpuUsage()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, data)
}

func SystemInfoService(context *gin.Context) {
    data, err := services.GetSystemInfo()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, data)
}

func CpuTemperatureService(context *gin.Context) {
    data, err := services.GetCpuTemperature()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, data)
}

func ShutdownSystemService(context *gin.Context) {
    err := services.ShutdownSystem()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "message": "Shutdown command executed successfully",
    })
}
