package main

import (
	"log"
	di "services/internal"
	"services/internal/config"
	"services/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	_, userController := di.SetupDependencies()
	
	router := gin.Default()

	routes.SetupRoutes(router, userController)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}