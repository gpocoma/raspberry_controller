package routes

import (
    "github.com/gin-gonic/gin"
    "raspberry-controller/minidlna_app/controllers"
)

func SetupMinidlnaRoutes(r *gin.Engine) {
    // Grupo de rutas para stats_app
    minidlnaGroup := r.Group("/minidlna")
    {
        minidlnaGroup.GET("/", controllers.MinidlnaStatusService)
        minidlnaGroup.POST("/start", controllers.StartMinidlnaService)
        minidlnaGroup.POST("/stop", controllers.StopMinidlnaService)
    }

    // Aquí puedes definir otros grupos de rutas para otras aplicaciones en el futuro
    // por ejemplo:
    // anotherAppGroup := r.Group("/another-app")
    // {
    //     anotherAppGroup.GET("/some-endpoint", anotherAppController.SomeEndpoint)
    // }
}