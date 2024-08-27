package services

import (
	"errors"
	"services/internal/models"
	"services/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

// UserService describe los métodos del servicio de usuario
type UserService interface {
	RegisterUser(user *models.User) error
	LoginUser(email, password string) (*models.User, error)
}

// userService implementa la interfaz UserService
type userService struct {
	repo repositories.UserRepository
}

// NewUserService crea una nueva instancia del servicio de usuario
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

// RegisterUser registra un nuevo usuario
func (us *userService) RegisterUser(user *models.User) error {
	// Verificar si el usuario ya existe
	existingUser, _ := us.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Hashear la contraseña antes de guardarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return us.repo.Save(user)
}

// LoginUser maneja el inicio de sesión del usuario
func (us *userService) LoginUser(email, password string) (*models.User, error) {
	// Buscar el usuario por email
	existingUser, err := us.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Comparar la contraseña proporcionada con el hash almacenado
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	return existingUser, nil
}
