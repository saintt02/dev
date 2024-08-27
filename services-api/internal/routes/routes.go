package routes

import (
	"services/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	api := router.Group("/api")
	{
		api.POST("/register", userController.RegisterUser)
		api.POST("/login", userController.LoginUser)

		// rutas protegidas
		authorized := api.Group("/")
		// middleware autenticacion
			{
					// reservas
					authorized.POST("/reservas",)
					authorized.GET("/reservas",)
					authorized.GET("/reservas/:id",)
					authorized.PUT("/reservas/:id",)
					authorized.DELETE("/reservas/:id",)
				  // talleres
					authorized.POST("/talleres",)
					authorized.GET("/talleres",)
					authorized.GET("/talleres/:id",)
					authorized.PUT("/talleres/:id",)
					authorized.DELETE("/talleres/:id",)
					// servicios
					authorized.POST("/servicios",)
					authorized.GET("/servicios",)
					authorized.GET("/servicios/:id",)
					authorized.PUT("/servicios/:id",)
					authorized.DELETE("/servicios/:id",)
			}
	}
}