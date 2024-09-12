package main

import (
	"log"
	"pwd-safe/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Route to add password
    router.POST("/pwd", handlers.AddPassword)

    // Start server
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}