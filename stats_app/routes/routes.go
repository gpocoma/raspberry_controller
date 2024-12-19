package routes

import (
    "github.com/gin-gonic/gin"
    "raspberry-controller/stats_app/controllers"
)

func SetupStatsRoutes(r *gin.Engine) {
    // Grupo de rutas para stats_app
    statsGroup := r.Group("/stats")
    {
        statsGroup.GET("/", controllers.TestService)
        statsGroup.GET("/ram-usage", controllers.RamUsageService)
        statsGroup.GET("/cpu-usage", controllers.CpuUsageService)
		statsGroup.GET("/system-info", controllers.SystemInfoService)
		statsGroup.GET("/cpu-temperature", controllers.CpuTemperatureService)
    }

    // Aquí puedes definir otros grupos de rutas para otras aplicaciones en el futuro
    // por ejemplo:
    // anotherAppGroup := r.Group("/another-app")
    // {
    //     anotherAppGroup.GET("/some-endpoint", anotherAppController.SomeEndpoint)
    // }
}