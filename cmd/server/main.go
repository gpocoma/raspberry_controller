package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"raspberry-controller/server"
)

func main() {
	// Crear el enrutador Gin
	r := gin.Default()

	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
        c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method Not Allowed"})
    })

	// Configurar las rutas
	server.SetupRoutes(r)

	// Iniciar el servidor en el puerto 3000
	if err := r.Run(":3000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
