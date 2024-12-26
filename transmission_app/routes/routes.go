package routes

import (
    "github.com/gin-gonic/gin"
    "raspberry-controller/transmission_app/controllers"
)

func SetupTransmissionRoutes(r *gin.Engine) {
    transmissionGroup := r.Group("/transmission")
    {
        transmissionGroup.GET("/", controllers.TransmissionStatusService)
        transmissionGroup.POST("/start", controllers.StartTransmissionService)
        transmissionGroup.POST("/stop", controllers.StopTransmissionService)
    }

    // Aquí puedes definir otros grupos de rutas para otras aplicaciones en el futuro
    // por ejemplo:
    // anotherAppGroup := r.Group("/another-app")
    // {
    //     anotherAppGroup.GET("/some-endpoint", anotherAppController.SomeEndpoint)
    // }
}