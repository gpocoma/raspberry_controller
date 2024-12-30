package routes

import (
	"github.com/gin-gonic/gin"
	"raspberry-controller/postgresql_app/controllers"
)

// SetupPostgreSQLRoutes configura las rutas para la aplicación de PostgreSQL
func SetupPostgreSQLRoutes(r *gin.Engine) {
	postgreSQLGroup := r.Group("/postgresql")
	{
		postgreSQLGroup.GET("/", controllers.PostgreSQLStatusService)
		postgreSQLGroup.POST("/start", controllers.StartPostgreSQLService)
		postgreSQLGroup.POST("/stop", controllers.StopPostgreSQLService)
	}

	// Aquí puedes definir otros grupos de rutas para otras aplicaciones en el futuro
	// por ejemplo:
	// anotherAppGroup := r.Group("/another-app")
	// {
	//     anotherAppGroup.GET("/some-endpoint", anotherAppController.SomeEndpoint)
	// }
}