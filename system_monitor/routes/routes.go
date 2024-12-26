package routes

import (
    "github.com/gin-gonic/gin"
    "raspberry-controller/system_monitor/controllers"
)

func SetupSystemMonitorRoutes(r *gin.Engine) {
    // Grupo de rutas para system_monitor
    systemMonitorGroup := r.Group("/system-monitor")
    {
        systemMonitorGroup.GET("/", controllers.TestService)
        systemMonitorGroup.GET("/ram-usage", controllers.RamUsageService)
        systemMonitorGroup.GET("/cpu-usage", controllers.CpuUsageService)
		systemMonitorGroup.GET("/system-info", controllers.SystemInfoService)
		systemMonitorGroup.GET("/cpu-temperature", controllers.CpuTemperatureService)
        systemMonitorGroup.GET("/shutdown", controllers.ShutdownSystemService)
    }

    // Aquí puedes definir otros grupos de rutas para otras aplicaciones en el futuro
    // por ejemplo:
    // anotherAppGroup := r.Group("/another-app")
    // {
    //     anotherAppGroup.GET("/some-endpoint", anotherAppController.SomeEndpoint)
    // }
}