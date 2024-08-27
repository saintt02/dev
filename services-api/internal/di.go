package di

import (
	"services/internal/config"
	"services/internal/controllers"
	"services/internal/repositories"
	"services/internal/services"

	"go.mongodb.org/mongo-driver/mongo"
)

// SetupDependencies inicializa todas las dependencias necesarias
func SetupDependencies() (*mongo.Database, *controllers.UserController) {
	// Conectar a la base de datos
	db := config.ConnectDB()

	// Inicializar los repositorios
	userRepo := repositories.NewMongoUserRepository(config.Client.Database("maintenance-service"))
	// reservationRepo := repositories.NewMongoReservationRepository(config.Client.Database("maintenance-service"))
	// workshopRepo := repositories.NewMongoWorkshopRepository(config.Client.Database("maintenance-service"))

	// Inicializar los servicios
	userService := services.NewUserService(userRepo)
	// reservationService := services.NewReservationService(reservationRepo)
	// workshopService := services.NewWorkshopService(workshopRepo)

	// Inicializar los controladores
	userController := controllers.NewUserController(userService)
	// reservationController := controllers.NewReservationController(reservationService)
	// workshopController := controllers.NewWorkshopController(workshopService)

	return db, userController //reservationController, workshopController
}
