package main

import (
	"inventory/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	routes.SetupRoutes(router)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}