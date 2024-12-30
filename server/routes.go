package server

import (
    "github.com/gin-gonic/gin"
    systemMonitorRoutes "raspberry-controller/system_monitor/routes"
    transmissionRoutes "raspberry-controller/transmission_app/routes"
    minidlnaRoutes "raspberry-controller/minidlna_app/routes"
    postgresqlRoutes "raspberry-controller/postgresql_app/routes"
)

func SetupRoutes(r *gin.Engine) {
    // Configurar las rutas para system_monitor
    systemMonitorRoutes.SetupSystemMonitorRoutes(r)

    // Configurar las rutas para transmission_app
    transmissionRoutes.SetupTransmissionRoutes(r)

    // Configurar las rutas para minidlna_app
    minidlnaRoutes.SetupMinidlnaRoutes(r)

    // Configurar las rutas para postgresql_app
    postgresqlRoutes.SetupPostgreSQLRoutes(r)
}