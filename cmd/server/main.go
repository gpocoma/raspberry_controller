package main

import (
	"log"
	"raspberry-controller/stats_app/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Crear el enrutador Gin
	r := gin.Default()

	// Configurar las rutas
	routes.SetupRoutes(r)

	// Iniciar el servidor en el puerto 3000
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
