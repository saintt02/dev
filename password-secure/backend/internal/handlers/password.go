package handlers

import (
	"net/http"
	config "pwd-safe/internal/database"
	"pwd-safe/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type PasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

func AddPassword(c *gin.Context) {

	var req PasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hasheo de la password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
    return
	}

	// Guardar la password dentro de MongoDB
	password := models.Password{HashedPassword: string(hashedPassword)}
	if err := config.SavePassword(password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password stored successfully"})

}