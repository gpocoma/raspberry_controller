package server

import (
    "github.com/gin-gonic/gin"
    statsRoutes "raspberry-controller/stats_app/routes"
    transmissionRoutes "raspberry-controller/transmission_app/routes"
)

func SetupRoutes(r *gin.Engine) {
    // Configurar las rutas para stats_app
    statsRoutes.SetupStatsRoutes(r)

    // Configurar las rutas para transmission_app
    transmissionRoutes.SetupTransmissionRoutes(r)
}