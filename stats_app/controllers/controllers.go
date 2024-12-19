package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/host"
    "time"
)

func TestService(context *gin.Context) {
    context.JSON(http.StatusOK, gin.H{
        "message": "Service is running",
    })
}

func RamUsageService(context *gin.Context) {
    v, err := mem.VirtualMemory()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "totalMB":     v.Total / 1024 / 1024,
        "usedMB":      v.Used / 1024 / 1024,
        "freeMB":      v.Free / 1024 / 1024,
        "usedPercent": v.UsedPercent,
    })
}

func CpuUsageService(context *gin.Context) {
    percentages, err := cpu.Percent(1*time.Second, false)
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "cpuUsagePercent": percentages[0],
    })
}

func SystemInfoService(context *gin.Context) {
    info, err := host.Info()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    context.JSON(http.StatusOK, gin.H{
        "hostname":       info.Hostname,
        "uptime":         info.Uptime,
        "bootTime":       info.BootTime,
        "os":             info.OS,
        "platform":       info.Platform,
        "platformFamily": info.PlatformFamily,
        "platformVersion": info.PlatformVersion,
        "kernelVersion":  info.KernelVersion,
        "virtualizationSystem": info.VirtualizationSystem,
        "virtualizationRole":   info.VirtualizationRole,
    })
}

func CpuTemperatureService(context *gin.Context) {
    sensors, err := host.SensorsTemperatures()
    if err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    for _, sensor := range sensors {
        if sensor.SensorKey == "cpu_thermal" {
            context.JSON(http.StatusOK, gin.H{
                "cpuTemperature": sensor.Temperature,
            })
            return
        }
    }

    context.JSON(http.StatusNotFound, gin.H{
        "error": "CPU temperature sensor not found",
    })
}