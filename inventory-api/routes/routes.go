package routes

import (
	"inventory/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api") 
	{
		// CRUD Routes
		api.GET("/parts", controllers.GetParts)
		api.GET("/parts/:id", controllers.GetPartByID)
		api.POST("/parts", controllers.PostParts)
		api.PUT("/parts/:id", controllers.UpdatePart)
		api.DELETE("/parts/:id", controllers.DeletePart)
	}
}