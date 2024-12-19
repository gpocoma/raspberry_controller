package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"raspberry-controller/server"
)

func main() {
	// Crear el enrutador Gin
	r := gin.Default()

	// Configurar las rutas
	server.SetupRoutes(r)

	// Iniciar el servidor en el puerto 3000
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
