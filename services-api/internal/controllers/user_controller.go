package controllers

import (
	"net/http"
	"services/internal/models"
	"services/internal/services"

	"github.com/gin-gonic/gin"
)

// UserController maneja las operaciones relacionadas con usuarios
type UserController struct {
	UserService services.UserService
}

// NewUserController crea una nueva instancia de UserController
func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// RegisterUser maneja el registro de nuevos usuarios
func (uc *UserController) RegisterUser(c *gin.Context) {
	var user models.User

	// Bind JSON input to user model
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to register the user
	if err := uc.UserService.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// LoginUser maneja el inicio de sesión de usuarios
func (uc *UserController) LoginUser(c *gin.Context) {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON input to login request model
	if err := c.BindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to login the user
	user, err := uc.UserService.LoginUser(loginReq.Email, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "user": user})
}
