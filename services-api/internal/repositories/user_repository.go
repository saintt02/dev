package repositories

import (
	"context"
	"services/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// definimos la interfaz para operaciones relacionadas con usuarios
type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Save(user *models.User) error
}

// implementamos el repositorio usando MongoDB
type MongoUserRepository struct {
	collection *mongo.Collection
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository {
		collection: db.Collection("users"),
	}
}

// FindByEmail busca al user por email
func (ur *MongoUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := ur.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No se encontró el documento, lo que significa que no existe el usuario
		}
		return nil, err
	}

	return &user, nil
}

func (ur *MongoUserRepository) Save(user *models.User) error {
	_, err := ur.collection.InsertOne(context.TODO(), user)
	return err
}