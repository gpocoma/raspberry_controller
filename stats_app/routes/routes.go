package routes

import (
	"github.com/gin-gonic/gin"
	"raspberry-controller/stats_app/controllers"
)

func SetupRoutes(r *gin.Engine) {
	// Definir una ruta GET
	r.GET("/", controllers.TestService)
}